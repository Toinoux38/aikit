package finetune

import (
	"github.com/moby/buildkit/client/llb"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
	"gopkg.in/yaml.v2"
)

func Aikit2LLB(c *config.FineTuneConfig) (llb.State, *specs.Image) {
	imageCfg := NewImageConfig(c)

	state := llb.Image(utils.DebianSlim)
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
	state = state.Run(utils.Sh("pip install packaging torch==2.1.0 ipython --break-system-packages")).Root()
	state = state.Run(utils.Sh("pip install 'unsloth[cu121_ampere] @ git+https://github.com/unslothai/unsloth.git' --break-system-packages")).Root()

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

	state = state.Run(utils.Sh("/provider_unsloth.py")).Root()

	scratch := llb.Scratch().File(llb.Copy(state, "/provider_unsloth.py", "/provider_unsloth.py"))

	return scratch, imageCfg
}
