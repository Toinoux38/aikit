package config

type Config struct {
	APIVersion string   `yaml:"apiVersion"`
	Debug      bool     `yaml:"debug,omitempty"`
	Runtime    string   `yaml:"runtime,omitempty"`
	Backends   []string `yaml:"backends,omitempty"`
	Models     []Model  `yaml:"models"`
	Config     string   `yaml:"config,omitempty"`
}

type Model struct {
	Name            string           `yaml:"name"`
	Source          string           `yaml:"source"`
	SHA256          string           `yaml:"sha256,omitempty"`
	PromptTemplates []PromptTemplate `yaml:"promptTemplates,omitempty"`
}

type PromptTemplate struct {
	Name     string `yaml:"name,omitempty"`
	Template string `yaml:"template,omitempty"`
}
