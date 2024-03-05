---
title: Fine Tuning
---

Fine tuning process allows the adaptation of pre-trained models to domain-specific data. At this time, AIKit fine tuning process is only supported with NVIDIA GPUs.

:::note
Due to current BuildKit and NVIDIA limitations, your host GPU driver version must match the driver that AIKit will install into the container during build. 

To find your host GPU driver, you can run `nvidia-smi` or `cat /proc/driver/nvidia/version`

For a list of supported drivers for AIKit, please refer to https://download.nvidia.com/XFree86/Linux-x86_64/

If you don't see your host GPU driver version in that list, you'll need to install one from the list.

This might be further optimizated in the future to remove this requirement, if possible.
:::

## Getting Started

To get started, you need to create a builder to be able to access host GPU devices. There are two ways you can get started:

### `docker-container` driver

Create a builder with the following configuration:

```bash
docker buildx create --name aikit-builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'
```

### `docker` driver with containerd image store

:::note
Containerd image store requires Docker v24 and later.
:::

You can enable [containerd image store](https://docs.docker.com/storage/containerd/) and required security entitlements for GPU device access by editing `/etc/docker/daemon.json` to add the following configuration:

```json
{
  ...
    "features": {
        "containerd-snapshotter": true
    },
    "builder": {
        "entitlements": {
            "security-insecure": true
        }
    }
}
```

After editing `/etc/docker/daemon.json`, restart the service with `sudo systemctl restart docker`. 

## Targets and Configuration

AIKit is capable of supporting multiple fine tuning implementation targets. 

At this time, [unsloth](https://github.com/unslothai/unsloth) is the only supported target, but future support for [axolotl](https://github.com/OpenAccess-AI-Collective/axolotl) is planned.

Please refer to [Fine Tuning API Specifications](./specs-finetune.md) for more information.

### Unsloth

Create a YAML file with your configuration. For example, minimum config looks like:

```yaml
#syntax=sozercan/aikit:test
apiVersion: v1alpha1
baseModel: unsloth/llama-2-7b-bnb-4bit # base model to be fine tuned. can be any model from huggingface. for unsloth optimized base models, see https://huggingface.co/unsloth
datasets:
  - source: "yahma/alpaca-cleaned" # data set to be used for fine tuning.
    type: alpaca # type of dataset. only alpaca is supported at this time.
```

For full configuration, please refer to [Fine Tune Specifications](./specs-finetune.md)

:::note
Please refer to [unsloth documentation](https://github.com/unslothai/unsloth) for more information about the configuration.
:::

## Build

Build using following command and make sure to replace `--target` with the fine-tuning implementation of your choice (`unsloth` is the only option supported at this time), `--file` with the path to your configuration YAML and `--output` with the output directory of the finetuned model.

```bash
docker buildx build --builder aikit-builder --allow security.insecure --file "/path/to/config.yaml" --output "/path/to/output" --target unsloth --progress plain .
```

:::tip
If you are using containerd image store option, you can build with `docker build` which will default to the `default` `docker` driver builder.
:::

Depending on your setup and configuration, build process may take some time. At the end of the build, the fine-tuned model will automatically be quantized with the specified format and output to the path specified in the `--output`.

Output will be a `GGUF` model file with the name and quanization format from the configuration. For example:

```bash
$ ls -al _output
-rw-r--r--  1 sozercan sozercan 7161089856 Mar  3 00:19 aikit-model-q4_k_m.gguf
```

## What's next?

Now that you have a fine-tuned model output as a GGUF file, you can refer to [Creating Model Images](./create-images.md) on how to create an image with AIKit to serve your fine-tuned model!
