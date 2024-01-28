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
	state = state.Run(llb.Shlex("apt-get update && apt-get install -y python3 python3-pip"), llb.IgnoreCache).Root()

	return state, imageCfg
}

