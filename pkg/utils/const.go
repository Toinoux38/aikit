package utils

const (
	RuntimeNVIDIA    = "cuda"
	RuntimeCPUAVX    = "avx"
	RuntimeCPUAVX2   = "avx2"
	RuntimeCPUAVX512 = "avx512"

	BackendStableDiffusion = "stablediffusion"
	BackendExllama         = "exllama"
	BackendExllamaV2       = "exllama2"
	BackendMamba           = "mamba"

	TargetUnsloth = "unsloth"

	DatasetAlpaca = "alpaca"

	APIv1alpha1 = "v1alpha1"

	DebianSlim = "docker.io/library/debian:12-slim"
	PythonBase = "docker.io/library/python:3.11"
)
