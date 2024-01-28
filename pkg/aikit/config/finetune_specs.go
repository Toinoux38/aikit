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
	Common  FineTuneConfigCommonSpec  `yaml:"common,omitempty"`
	Axolotl FineTuneConfigAxolotlSpec `yaml:"axolotl,omitempty"`
	Unsloth FineTuneConfigUnslothSpec `yaml:"unsloth,omitempty"`
}

type FineTuneConfigCommonSpec struct {
	Packing                   bool    `yaml:"packing,omitempty"`
	LoadIn4bit                bool    `yaml:"loadIn4bit,omitempty"`
	BatchSize                 int     `yaml:"batchSize,omitempty"`
	GradientAccumulationSteps int     `yaml:"gradientAccumulationSteps,omitempty"`
	WarmupSteps               int     `yaml:"warmupSteps,omitempty"`
	MaxSteps                  int     `yaml:"maxSteps,omitempty"`
	LearningRate              float64 `yaml:"learningRate,omitempty"`
	Fp16                      bool    `yaml:"fp16,omitempty"`
	Bf16                      bool    `yaml:"bf16,omitempty"`
	Tf32                      bool    `yaml:"tf32,omitempty"`
	LoggingSteps              int     `yaml:"loggingSteps,omitempty"`
	Optimizer                 string  `yaml:"optimizer,omitempty"`
	WeightDecay               float64 `yaml:"weightDecay,omitempty"`
	LrSchedular               string  `yaml:"lrSchedular,omitempty"`
	Seed                      int     `yaml:"seed,omitempty"`
}

type FineTuneConfigAxolotlSpec struct {
	FlashAttention bool `yaml:"flashAttention,omitempty"`
}

type FineTuneConfigUnslothSpec struct {}

type FineTuneOutputSpec struct {
	Path     string `yaml:"path,omitempty"`
	Format   string `yaml:"format,omitempty"`
	Quantize string `yaml:"quantize,omitempty"`
}

type Dataset struct {
	Source string `yaml:"source,omitempty"`
	Type   string `yaml:"type,omitempty"`
}
