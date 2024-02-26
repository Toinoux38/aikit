package finetune

import (
	"fmt"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
	"gopkg.in/yaml.v2"
)

const (
	// https://github.com/unslothai/unsloth/releases/tag/February-2024
	unslothCommitSHA = "3e4c5a323c16bbda2c92212b790073c4e99c2a55"
)

func Aikit2LLB(c *config.FineTuneConfig) (llb.State, *specs.Image) {
	imageCfg := NewImageConfig(c)

	env := map[string]string{
		"PATH":                       system.DefaultPathEnv("linux") + ":/usr/local/cuda/bin",
		"NVIDIA_REQUIRE_CUDA":        "cuda>=12.0",
		"NVIDIA_DRIVER_CAPABILITIES": "compute,utility",
		"NVIDIA_VISIBLE_DEVICES":     "all",
		"LD_LIBRARY_PATH":            "/usr/local/cuda/lib64",
	}

	state := llb.Image(utils.DebianSlim)
	for k, v := range env {
		state = state.AddEnv(k, v)
	}

	// installing dependencies
	// due to buildkit run limitations, we need to install nvidia drivers and driver version must match the host
	state = state.Run(utils.Sh("apt-get update && apt-get install -y --no-install-recommends python3-dev python3 python3-pip python-is-python3 git wget kmod && cd /root && VERSION=$(cat /proc/driver/nvidia/version | sed -n 's/.*NVIDIA UNIX x86_64 Kernel Module  \\([0-9]\\+\\.[0-9]\\+\\.[0-9]\\+\\).*/\\1/p') && wget -q https://us.download.nvidia.com/XFree86/Linux-x86_64/$VERSION/NVIDIA-Linux-x86_64-$VERSION.run && chmod +x NVIDIA-Linux-x86_64-$VERSION.run && ./NVIDIA-Linux-x86_64-$VERSION.run -x && rm NVIDIA-Linux-x86_64-$VERSION.run && /root/NVIDIA-Linux-x86_64-$VERSION/nvidia-installer -a -s --skip-depmod --no-dkms --no-nvidia-modprobe --no-questions --no-systemd --no-x-check --no-kernel-modules --no-kernel-module-source && rm -rf /root/NVIDIA-Linux-x86_64-$VERSION"), llb.IgnoreCache).Root()

	// installing cuda
	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	state = state.File(
		llb.Copy(cudaKeyring, utils.FileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(cudaKeyringURL)),
	)
	state = state.Run(utils.Sh("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()
	state = state.Run(utils.Sh("apt-get update && apt-get install -y --no-install-recommends cuda-toolkit cuda-nvcc-12-3 && apt-get clean")).Root()

	// installing unsloth
	state = state.Run(utils.Sh("pip install --upgrade pip --break-system-packages && pip install packaging torch==2.1.0 ipython --break-system-packages")).Root()
	state = state.Run(utils.Shf("pip install 'unsloth[cu121_ampere] @ git+https://github.com/unslothai/unsloth.git@%[1]s' --break-system-packages", unslothCommitSHA)).Root()

	// TODO: replace the branch with a release tag and have it update with a release
	unslothScriptURL := "https://raw.githubusercontent.com/sozercan/aikit/finetune/pkg/finetune/provider_unsloth.py"
	var opts []llb.HTTPOption
	opts = append(opts, llb.Chmod(0o755))
	unslothScript := llb.HTTP(unslothScriptURL, opts...)
	state = state.File(
		llb.Copy(unslothScript, utils.FileNameFromURL(unslothScriptURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(unslothScriptURL)),
	)

	// write config to /config.yaml
	cfg, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	state = state.Run(utils.Shf("echo -n \"%s\" > /config.yaml", string(cfg))).Root()

	// setup nvidia devices and run unsloth
	// due to buildkit run limitations, we need to create the devices manually and run unsloth in the same command
	state = state.Run(utils.Sh("mknod --mode 666 /dev/nvidiactl c 195 255 && mknod --mode 666 /dev/nvidia-uvm c 235 0 && mknod --mode 666 /dev/nvidia-uvm-tools c 235 1 && mknod --mode 666 /dev/nvidia0 c 195 0 && nvidia-smi && /provider_unsloth.py"), llb.Security(llb.SecurityModeInsecure)).Root()

	// copy gguf to scratch which will be the output
	const inputFile = "*.gguf"
	copyOpts := []llb.CopyOption{}
	copyOpts = append(copyOpts, &llb.CopyInfo{AllowWildcard: true})
	outputFile := fmt.Sprintf("%s-%s.gguf", c.Output.Name, c.Output.Quantize)
	scratch := llb.Scratch().File(llb.Copy(state, inputFile, outputFile, copyOpts...))

	return scratch, imageCfg
}
