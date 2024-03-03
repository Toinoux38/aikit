package finetune

import (
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
)

func NewImageConfig(c *config.FineTuneConfig) *specs.Image {
	img := emptyImage(c)

	img.Config.Entrypoint = []string{"/target_unsloth.py"}
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
		"PATH=" + system.DefaultPathEnv("linux") + ":/usr/local/cuda/bin",
		"NVIDIA_REQUIRE_CUDA=cuda>=12.0",
		"NVIDIA_DRIVER_CAPABILITIES=compute,utility",
		"NVIDIA_VISIBLE_DEVICES=all",
		"LD_LIBRARY_PATH=/usr/local/cuda/lib64",
	}

	return img
}
