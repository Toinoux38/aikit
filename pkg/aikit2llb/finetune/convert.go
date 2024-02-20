package finetune

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
	"gopkg.in/yaml.v2"
)

func Aikit2LLB(c *config.FineTuneConfig) (llb.State, *specs.Image) {
	imageCfg := NewImageConfig(c)

	// TODO: not sure if these are necessary
	env := map[string]string{
		"PATH":                       system.DefaultPathEnv("linux") + ":/usr/local/cuda/bin",
		"NVIDIA_REQUIRE_CUDA":        "cuda>=12.0",
		"NVIDIA_DRIVER_CAPABILITIES": "compute,utility",
		"NVIDIA_VISIBLE_DEVICES":     "all",
		"LD_LIBRARY_PATH":            "/usr/local/cuda/lib64",
	}

	// TODO: need input from the user or set up requirements with nvidia-container-cli
	state := llb.Image("docker.io/sozercan/nvidia:545-23-06")
	for k, v := range env {
		state = state.AddEnv(k, v)
	}

	// var runOpts []llb.RunOption
	// runOpts = append(runOpts, llb.Security(llb.SecurityModeInsecure))

	// state = state.Run(utils.Sh("mknod --mode 666 /dev/nvidiactl c 195 255 && mknod --mode 666 /dev/nvidia-modeset c 195 254 && mknod --mode 666 /dev/nvidia-uvm c 235 0 && mknod --mode 666 /dev/nvidia-uvm-tools c 235 1 && mknod --mode 666 /dev/nvidia0 c 195 0 && chmod 0666 /dev/nvidia* && ls -al /dev && /root/NVIDIA-Linux-x86_64-545.23.06/nvidia-smi"), llb.Security(llb.SecurityModeInsecure)).Root()

	state = state.Run(utils.Sh("apt-get update && apt-get install -y python3 python3-pip python-is-python3 git"), llb.IgnoreCache).Root()

	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	state = state.File(
		llb.Copy(cudaKeyring, utils.FileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(cudaKeyringURL)),
	)
	state = state.Run(utils.Sh("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()

	// TODO: use required packages instead of cuda-toolkit if possible
	state = state.Run(utils.Sh("apt-get update && apt-get install -y --no-install-recommends cuda-toolkit cuda-nvcc-12-3 && apt-get clean")).Root()
	state = state.Run(utils.Sh("pip install --upgrade pip && pip install packaging torch==2.1.0 ipython")).Root()
	state = state.Run(utils.Sh("pip install 'unsloth[cu121_ampere] @ git+https://github.com/unslothai/unsloth.git'")).Root()

	// TODO: replace the branch with a release tag and have it update with a release
	unslothScriptURL := "https://raw.githubusercontent.com/sozercan/aikit/finetune/pkg/finetune/provider_unsloth.py"
	var opts []llb.HTTPOption
	opts = append(opts, llb.Chmod(0o755))
	unslothScript := llb.HTTP(unslothScriptURL, opts...)
	state = state.File(
		llb.Copy(unslothScript, utils.FileNameFromURL(unslothScriptURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(unslothScriptURL)),
	)

	cfg, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	state = state.Run(utils.Shf("echo -n \"%s\" > /config.yaml", string(cfg))).Root()

	// TODO: remove ls /dev and nvidia-smi
	state = state.Run(utils.Sh("mknod --mode 666 /dev/nvidiactl c 195 255 && mknod --mode 666 /dev/nvidia-modeset c 195 254 && mknod --mode 666 /dev/nvidia-uvm c 235 0 && mknod --mode 666 /dev/nvidia-uvm-tools c 235 1 && mknod --mode 666 /dev/nvidia0 c 195 0 && chmod 0666 /dev/nvidia* && ls -al /dev && /root/NVIDIA-Linux-x86_64-545.23.06/nvidia-smi && /provider_unsloth.py"), llb.Security(llb.SecurityModeInsecure)).Root()

	scratch := llb.Scratch().File(llb.Copy(state, "model_gguf-unsloth.Q4_K_M.gguf", "model_gguf-unsloth.Q4_K_M.gguf"))

	return scratch, imageCfg
}
