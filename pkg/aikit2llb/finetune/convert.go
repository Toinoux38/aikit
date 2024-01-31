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
	state = state.Run(sh("apt-get update && apt-get install -y python3 python3-pip git python3-is-python"), llb.IgnoreCache).Root()

	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	state = state.File(
		llb.Copy(cudaKeyring, utils.FileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(cudaKeyringURL)),
	)
	state = state.Run(sh("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()

	// TODO: use required packages instead of cuda-toolkit if possible
	state = state.Run(sh("apt-get update && apt-get install -y --no-install-recommends cuda-toolkit cuda-nvcc-12-3 && apt-get clean")).Root()
	state = state.Run(sh("pip install packaging torch==2.1.0")).Root()
	state = state.Run(sh("pip install 'unsloth[cu121_ampere] @ git+https://github.com/unslothai/unsloth.git' --break-system-packages")).Root()

	unslothScriptURL := "https://gist.githubusercontent.com/sozercan/f00e6e307511718771bb386553ef408f/raw/8d52de7cdd3a14c715c15770ef8eac06bc3b7ff5/script.py"
	unslothScript := llb.HTTP(unslothScriptURL)
	state = state.File(
		llb.Copy(unslothScript, utils.FileNameFromURL(unslothScriptURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(unslothScriptURL)),
	)

	return state, imageCfg
}

func sh(cmd string) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", cmd})
}
