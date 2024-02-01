package config

type FineTuneConfig struct {
	APIVersion string             `yaml:"apiVersion"`
	Debug      bool               `yaml:"debug,omitempty"`
	Runtime    string             `yaml:"runtime,omitempty"`
	Provider   string             `yaml:"provider,omitempty"`
	BaseModel  string             `yaml:"baseModel,omitempty"`
	Datasets   []Dataset          `yaml:"datasets"`
	Config     FineTuneConfigSpec `yaml:"config,omitempty"`
	Output     FineTuneOutputSpec `yaml:"output,omitempty"`
}

type FineTuneConfigSpec struct {
	Axolotl FineTuneConfigAxolotlSpec `yaml:"axolotl,omitempty"`
	Unsloth FineTuneConfigUnslothSpec `yaml:"unsloth,omitempty"`
}

type Dataset struct {
	Source string `yaml:"source,omitempty"`
	Type   string `yaml:"type,omitempty"`
}

type FineTuneConfigAxolotlSpec struct{}

type FineTuneConfigUnslothSpec struct {
	MaxSeqLength              int     `yaml:"maxSeqLength,omitempty"`
	LoadIn4bit                bool    `yaml:"loadIn4bit,omitempty"`
	BatchSize                 int     `yaml:"batchSize,omitempty"`
	GradientAccumulationSteps int     `yaml:"gradientAccumulationSteps,omitempty"`
	WarmupSteps               int     `yaml:"warmupSteps,omitempty"`
	MaxSteps                  int     `yaml:"maxSteps,omitempty"`
	LoggingSteps              int     `yaml:"loggingSteps,omitempty"`
	Optimizer                 string  `yaml:"optimizer,omitempty"`
	WeightDecay               float64 `yaml:"weightDecay,omitempty"`
	LrSchedular               string  `yaml:"lrSchedular,omitempty"`
	Seed                      int     `yaml:"seed,omitempty"`
}

type FineTuneOutputSpec struct {
	Path       string `yaml:"path,omitempty"`
	Format     string `yaml:"format,omitempty"`
	Quantize   string `yaml:"quantize,omitempty"`
	UploadToHF bool   `yaml:"uploadToHF,omitempty"`
	Token      string `yaml:"token,omitempty"`
}
