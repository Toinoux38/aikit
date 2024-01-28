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

	APIv1alpha1 = "v1alpha1"

	DebianSlim     = "docker.io/library/debian:12-slim"
)
