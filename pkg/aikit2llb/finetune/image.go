package finetune

import (
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func NewImageConfig() *specs.Image {
	return emptyImage()
}

func emptyImage() *specs.Image {
	img := &specs.Image{
		Platform: specs.Platform{
			Architecture: "amd64",
			OS:           "linux",
		},
	}
	img.RootFS.Type = "layers"
	img.Config.WorkingDir = "/"

	return img
}
