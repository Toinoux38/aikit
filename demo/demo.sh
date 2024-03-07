#!/bin/bash

. third_party/demo-magic/demo-magic.sh

clear
DEMO_PROMPT="${GREEN}➜  ${COLOR_RESET}"

echo "✨ In this demo, we are going to start by fine tuning a model and then deploy the model as a minimal container!"

echo ""

echo "👷‍♀️ First, we are going to create a new builder"

echo ""

pei "docker buildx create --name aikit-builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'"

echo ""

echo "🗃️ Create a configuration for the fine tuning. We are going to be using a LLama 2 model and fine tune using yahma/alpaca-cleaned dataset."

cat > aikit-finetune.yaml << EOF
#syntax=sozercan/aikit:latest
apiVersion: v1alpha1
baseModel: unsloth/llama-2-7b-bnb-4bit
datasets:
  - source: "yahma/alpaca-cleaned"
    type: alpaca
config:
  unsloth:
EOF

echo ""

pei "bat aikit-finetune.yaml"

echo ""

echo "🎵 Starting the fine tuning process using the above configuration file, and output fine tuned model will be saved in _output folder."

echo ""

pei "docker buildx build --builder aikit-builder --allow security.insecure --file 'aikit-finetune.yaml' --output '_output' --target unsloth --progress plain ."

echo ""

echo "✅ We have finished fine tuning the model. Let's look at the output..."

echo ""

pei "ls -al _output"

echo ""

echo "📦 Now that we have a fine tuned model. We can deploy it as a minimal container."

echo ""

echo "📃 We'll start by creating a basic inference configuration file for the deployment."

cat > aikit-inference.yaml << EOF
#syntax=sozercan/aikit:latest
apiVersion: v1alpha1
debug: true
runtime: cuda
models:
  - name: llama-2-finetuned
    source: aikit-model-q4_k_m.gguf
config: |
  - name: llama-2-finetuned
    parameters:
      model: aikit-model-q4_k_m.gguf
    context_size: 4096
    gpu_layers: 35
    f16: true
    batch: 512
    mmap: true
EOF

pei "bat aikit-inference.yaml"

echo ""

echo "🏗️ We can now build a minimal container for the model using the configuration file."

echo ""

pei "docker buildx build -t llama-finetuned -f aikit-inference.yaml --load --progress plain _output"

echo ""

echo "🏃 We have finished building the minimal container. Let's start the container and test it."

echo ""

pei "docker run --name llama-2-finetuned -d --rm -p 8080:8080 llama-finetuned"

echo ""

echo "🧪 We can now test the container using a sample query. Since this is OpenAI API compatible, you can use it as a drop-in replacement for any application that uses OpenAI API."

echo ""

pei "curl http://localhost:8080/v1/chat/completions -H \"Content-Type: application/json\" -d '{\"model\": \"llama-2-finetuned\", \"messages\": [{\"role\": \"user\", \"content\": \"Compose a haiku about cats\"}]}'"

echo ""

echo "🙌 We have successfully deployed the fine tuned model as a minimal container and successfully verified it! We can now stop the container if we wish."

echo ""

pei "docker stop llama-2-finetuned"

echo ""

echo "❤️ In this demo, we have shown how to fine tune a model and deploy it as a minimal container using AIKit. Thank you for watching!"

pei "docker buildx rm aikit-builder"
