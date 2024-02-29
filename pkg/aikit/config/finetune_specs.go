package config

type FineTuneConfig struct {
	APIVersion string             `yaml:"apiVersion"`
	Target     string             `yaml:"target,omitempty"`
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
	Quantize string `yaml:"quantize,omitempty"`
	Name     string `yaml:"name,omitempty"`
}
