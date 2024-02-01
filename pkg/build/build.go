package build

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/containerd/containerd/platforms"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/moby/buildkit/frontend/dockerui"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/pkg/errors"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/aikit2llb/finetune"
	"github.com/sozercan/aikit/pkg/aikit2llb/inference"
	"github.com/sozercan/aikit/pkg/utils"
)

const (
	LocalNameDockerfile   = "dockerfile"
	keyFilename           = "filename"
	defaultDockerfileName = "aikitfile.yaml"
)

func Build(ctx context.Context, c client.Client) (*client.Result, error) {
	inferenceCfg, finetuneCfg, err := getAikitfileConfig(ctx, c)
	if err != nil {
		return nil, errors.Wrap(err, "getting aikitfile")
	}

	if finetuneCfg != nil {
		return buildFineTune(ctx, c, finetuneCfg)
	} else if inferenceCfg != nil {
		return buildInference(ctx, c, inferenceCfg)
	}

	return nil, nil
}

func buildFineTune(ctx context.Context, c client.Client, cfg *config.FineTuneConfig) (*client.Result, error) {
	err := validateFinetuneConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "validating aikitfile")
	}

	st, img := finetune.Aikit2LLB(cfg)

	def, err := st.Marshal(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	config, err := json.Marshal(img)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal image config")
	}
	k := platforms.Format(platforms.DefaultSpec())

	res.AddMeta(fmt.Sprintf("%s/%s", exptypes.ExporterImageConfigKey, k), config)
	res.SetRef(ref)

	return res, nil
}

func buildInference(ctx context.Context, c client.Client, cfg *config.InferenceConfig) (*client.Result, error) {
	err := validateConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "validating aikitfile")
	}

	st, img := inference.Aikit2LLB(cfg)

	def, err := st.Marshal(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	config, err := json.Marshal(img)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal image config")
	}
	k := platforms.Format(platforms.DefaultSpec())

	res.AddMeta(fmt.Sprintf("%s/%s", exptypes.ExporterImageConfigKey, k), config)
	res.SetRef(ref)

	return res, nil
}

func getAikitfileConfig(ctx context.Context, c client.Client) (*config.InferenceConfig, *config.FineTuneConfig, error) {
	opts := c.BuildOpts().Opts
	filename := opts[keyFilename]
	if filename == "" {
		filename = defaultDockerfileName
	}

	name := "load aikitfile"
	if filename != "aikitfile" {
		name += " from " + filename
	}

	src := llb.Local(LocalNameDockerfile,
		llb.IncludePatterns([]string{filename}),
		llb.SessionID(c.BuildOpts().SessionID),
		llb.SharedKeyHint(defaultDockerfileName),
		dockerui.WithInternalName(name),
	)

	def, err := src.Marshal(ctx)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to marshal local source")
	}

	var dtDockerfile []byte
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}

	ref, err := res.SingleRef()
	if err != nil {
		return nil, nil, err
	}

	dtDockerfile, err = ref.ReadFile(ctx, client.ReadRequest{
		Filename: filename,
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to read dockerfile")
	}

	inferenceCfg, finetuneCfg, err := config.NewFromBytes(dtDockerfile)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting config")
	}

	return inferenceCfg, finetuneCfg, nil
}

func validateFinetuneConfig(c *config.FineTuneConfig) error {
	if c.APIVersion == "" {
		return errors.New("apiVersion is not defined")
	}
	return nil
}

func validateConfig(c *config.InferenceConfig) error {
	if c.APIVersion == "" {
		return errors.New("apiVersion is not defined")
	}

	if c.APIVersion != utils.APIv1alpha1 {
		return errors.Errorf("apiVersion %s is not supported", c.APIVersion)
	}

	if len(c.Models) == 0 {
		return errors.New("no models defined")
	}

	if len(c.Backends) > 1 {
		return errors.New("only one backend is supported at this time")
	}

	if slices.Contains(c.Backends, utils.BackendStableDiffusion) && (slices.Contains(c.Backends, utils.BackendExllama) || slices.Contains(c.Backends, utils.BackendExllamaV2)) {
		return errors.New("cannot specify both stablediffusion with exllama or exllama2 at this time")
	}

	if (slices.Contains(c.Backends, utils.BackendExllama) || slices.Contains(c.Backends, utils.BackendExllamaV2) || slices.Contains(c.Backends, utils.BackendMamba)) && c.Runtime != utils.RuntimeNVIDIA {
		return errors.New("exllama and mamba only supports nvidia cuda runtime. please add 'runtime: cuda' to your aikitfile.yaml")
	}

	backends := []string{utils.BackendExllama, utils.BackendExllamaV2, utils.BackendStableDiffusion, utils.BackendMamba}
	for _, b := range c.Backends {
		if !slices.Contains(backends, b) {
			return errors.Errorf("backend %s is not supported", b)
		}
	}

	runtimes := []string{"", utils.RuntimeNVIDIA, utils.RuntimeCPUAVX, utils.RuntimeCPUAVX2, utils.RuntimeCPUAVX512}
	if !slices.Contains(runtimes, c.Runtime) {
		return errors.Errorf("runtime %s is not supported", c.Runtime)
	}

	return nil
}
