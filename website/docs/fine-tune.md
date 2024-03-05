---
title: Fine Tuning
---

Fine tuning process allows the adaptation of pre-trained models to domain-specific data. At this time, AIKit fine tuning process is only supported with NVIDIA GPUs.

:::note
Due to current BuildKit and NVIDIA limitations, your host GPU driver version must match the driver that AIKit will install into the container during build. 

To find your host GPU driver version, you can run `nvidia-smi` or `cat /proc/driver/nvidia/version`

For a list of supported driver versions for AIKit, please refer to https://download.nvidia.com/XFree86/Linux-x86_64/

If you don't see your host GPU driver version in that list, you'll need to install one that's matching the version in that list. You don't need to install drivers from that location, only the versions need to match.

This might be further optimizated in the future to remove this requirement, if possible.
:::

## Getting Started

To get started, you need to create a builder to be able to access host GPU devices.

Create a builder with the following configuration:

```bash
docker buildx create --name aikit-builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'
```

## Targets and Configuration

AIKit is capable of supporting multiple fine tuning implementation targets. At this time, [Unsloth](https://github.com/unslothai/unsloth) is the only supported target, but can be extended for other fine tuning implementations in the future.

### Unsloth

Create a YAML file with your configuration. For example, minimum config looks like:

```yaml
#syntax=ghcr.io/sozercan/aikit:latest
apiVersion: v1alpha1
baseModel: unsloth/llama-2-7b-bnb-4bit # base model to be fine tuned. can be any model from huggingface. for unsloth optimized base models, see https://huggingface.co/unsloth
datasets:
  - source: "yahma/alpaca-cleaned" # data set to be used for fine tuning.
    type: alpaca # type of dataset. only alpaca is supported at this time.
```

For full configuration, please refer to [Fine Tune API Specifications](./specs-finetune.md)

:::note
Please refer to [Unsloth documentation](https://github.com/unslothai/unsloth) for more information about Unsloth configuration.
:::

## Build

Build using following command and make sure to replace `--target` with the fine-tuning implementation of your choice (`unsloth` is the only option supported at this time), `--file` with the path to your configuration YAML and `--output` with the output directory of the finetuned model.

```bash
docker buildx build --builder aikit-builder --allow security.insecure --file "/path/to/config.yaml" --output "/path/to/output" --target unsloth --progress plain .
```

Depending on your setup and configuration, build process may take some time. At the end of the build, the fine-tuned model will automatically be quantized with the specified format and output to the path specified in the `--output`.

Output will be a `GGUF` model file with the name and quanization format from the configuration. For example:

```bash
$ ls -al _output
-rw-r--r--  1 sozercan sozercan 7161089856 Mar  3 00:19 aikit-model-q4_k_m.gguf
```

## What's next?

👉 Now that you have a fine-tuned model output as a GGUF file, you can refer to [Creating Model Images](./create-images.md) on how to create an image with AIKit to serve your fine-tuned model!
