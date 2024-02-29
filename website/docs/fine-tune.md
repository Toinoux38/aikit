---
title: Fine Tuning
---

```bash
docker buildx create --name builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'
```

```yaml
#syntax=sozercan/aikit:test
apiVersion: v1alpha1
baseModel: unsloth/llama-2-7b-bnb-4bit
datasets:
  - source: "https://huggingface.co/datasets/laion/OIG/resolve/main/unified_chip2.jsonl"
    type: alpaca
config:
  unsloth:
    maxSeqLength: 2048
    loadIn4bit: true
    batchSize: 2
    gradientAccumulationSteps: 4
    warmupSteps: 10
    maxSteps: 60
    loggingSteps: 1
    optimizer: adamw_8bit
    weightDecay: 0.01
    lrSchedular: linear
    seed: 42
output:
  quantize: q8_0
  name: model
```

```bash
docker build --allow security.insecure -f test/aikitfile-unsloth.yaml --output _output --target unsloth .
```