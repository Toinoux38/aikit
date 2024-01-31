package finetune

import (
	"github.com/moby/buildkit/client/llb"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
)

func Aikit2LLB(c *config.FineTuneConfig) (llb.State, *specs.Image) {
	imageCfg := NewImageConfig(c)

	state := llb.Image(utils.DebianSlim)
	state = state.Run(sh("apt-get update && apt-get install -y python3 python3-pip python-is-python3 git"), llb.IgnoreCache).Root()

	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	state = state.File(
		llb.Copy(cudaKeyring, utils.FileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(cudaKeyringURL)),
	)
	state = state.Run(sh("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()

	// TODO: use required packages instead of cuda-toolkit if possible
	state = state.Run(sh("apt-get update && apt-get install -y --no-install-recommends cuda-toolkit cuda-nvcc-12-3 && apt-get clean")).Root()
	state = state.Run(sh("pip install packaging torch==2.1.0 ipython --break-system-packages")).Root()
	state = state.Run(sh("pip install 'unsloth[cu121_ampere] @ git+https://github.com/unslothai/unsloth.git' --break-system-packages")).Root()

	// replace
	unslothScriptURL := "https://gist.githubusercontent.com/sozercan/f00e6e307511718771bb386553ef408f/raw/67fd2fa0d951b3cfb4f04adbbb1d2af42f68d3e7/script.py"
	var opts []llb.HTTPOption
	opts = append(opts, llb.Chmod(0o755))
	unslothScript := llb.HTTP(unslothScriptURL, opts...)
	state = state.File(
		llb.Copy(unslothScript, utils.FileNameFromURL(unslothScriptURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(unslothScriptURL)),
	)

	// replace
	configURL := "https://gist.githubusercontent.com/sozercan/a4b4d902b91fa7c0880807dac8c29ec6/raw/6adbea7895415200bf24539c2b9d59cb07d5951e/config.yaml"
	config := llb.HTTP(configURL)
	state = state.File(
		llb.Copy(config, utils.FileNameFromURL(configURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(configURL)),
	)

	return state, imageCfg
}

func sh(cmd string) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", cmd})
}
