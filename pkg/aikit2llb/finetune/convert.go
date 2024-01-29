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
	state = state.Run(sh("apt-get update && apt-get install -y python3 python3-pip git"), llb.IgnoreCache).Root()
	state = state.Run(sh("pip install 'unsloth[cu121_ampere] @ git+https://github.com/unslothai/unsloth.git'")).Root()

	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	state = state.File(
		llb.Copy(cudaKeyring, utils.FileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(cudaKeyringURL)),
	)
	state = state.Run(sh("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()

	state = state.Run(sh("apt-get install -y --no-install-recommends libcublas-12-3 cuda-cudart-12-3 && apt-get clean")).Root()


	return state, imageCfg
}

func sh(cmd string) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", cmd})
}
