#!/bin/bash

. third_party/demo-magic/demo-magic.sh

clear
DEMO_PROMPT="${GREEN}➜  ${COLOR_RESET}"

echo "In this demo, we are going to start by fine tuning a model and then deploy the model as a minimal container!"

echo "First, we are going to create a new builder"

pei "docker buildx create --name aikit-builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'"

echo "Create a configuration for the fine tuning. We are going to be using a LLama 2 model and fine tune using yahma/alpaca-cleaned dataset."

pei cat <<EOF >> aikit-finetune.yaml
#syntax=sozercan/aikit:latest
apiVersion: v1alpha1
baseModel: unsloth/llama-2-7b-bnb-4bit
datasets:
  - source: "yahma/alpaca-cleaned"
    type: alpaca
EOF

# echo "Starting the process using the above configuration file, and output fine tuned model will be saved in _output folder."

# pei "docker buildx build --builder aikit-builder --allow security.insecure --file 'aikit-finetune.yaml' --output '_output' --target unsloth --progress plain ."

# echo "We have finished fine tuning the model. Let's look at the output."

# pei "ls -al _output"

# echo "Now that we have a fine tuned model. We can deploy it as a minimal container."

# echo "We'll start by creating a basic inference configuration file for the deployment."

# pei "cat <<EOF >> aikit-inference.yaml
# #syntax=sozercan/aikit:latest
# apiVersion: v1alpha1
# models:
#   - name: llama-2-finetuned
#     source: aikit-model-q4_k_m.gguf
# EOF"

# echo "We can now build a minimal container for the model using the configuration file."

# pei "docker buildx build -t llama-finetuned -f aikit-inference.yaml --load _output"

# echo "We have finished building the minimal container. Let's start the container and test it."

# pei "docker run --name llama-2-finetuned -d --rm -p 8080:8080 llama-finetuned"

# echo "We can now test the container using a sample query"

# pei "curl http://localhost:8080/v1/chat/completions -H 'Content-Type: application/json' -d '{'model': 'llama-2-finetuned', 'messages': [{'role': 'user', 'content': 'Compose a haiku about cats'}]}'"

# echo "We have successfully deployed the fine tuned model as a minimal container and tested it. We can now stop the container."

# pei "docker stop llama-2-finetuned"

# echo "In this demo, we have shown how to fine tune a model and deploy it as a minimal container using AIKit. Thank you for watching!"
