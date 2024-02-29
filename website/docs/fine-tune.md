---
title: Fine Tuning
---

:::note
Due to current BuildKit and NVIDIA limitations, your host GPU driver version must match the driver that AIKit will install into the container during build. 

To find your host GPU driver, you can run `nvidia-smi` or `cat /proc/driver/nvidia/version`

For a list of supported drivers for AIKit, please refer to https://download.nvidia.com/XFree86/Linux-x86_64/

If you don't see your host GPU driver in that list, you'll need to install one from the list.
:::


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
    # TODO: add validations
    packing: False
    maxSeqLength: 2048
    loadIn4bit: true
    batchSize: 2
    gradientAccumulationSteps: 4
    warmupSteps: 10
    maxSteps: 60
    learningRate: 0.0002
    loggingSteps: 1
    optimizer: adamw_8bit
    weightDecay: 0.01
    lrSchedulerType: linear
    seed: 42
output:
  quantize: q8_0
  name: model

```

```bash
docker build --allow security.insecure -f test/aikitfile-unsloth.yaml --output _output --target unsloth .
```