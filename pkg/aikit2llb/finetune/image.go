package finetune

import (
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
)

func NewImageConfig(c *config.FineTuneConfig) *specs.Image {
	img := emptyImage(c)

	img.Config.Entrypoint = []string{"python3"}
	return img
}

func emptyImage(c *config.FineTuneConfig) *specs.Image {
	img := &specs.Image{
		Platform: specs.Platform{
			Architecture: "amd64",
			OS:           "linux",
		},
	}
	img.RootFS.Type = "layers"
	img.Config.WorkingDir = "/"

	img.Config.Env = []string{
		"PATH=" + system.DefaultPathEnv("linux"),
	}

	return img
}
