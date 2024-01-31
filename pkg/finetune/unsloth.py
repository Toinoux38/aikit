#!/usr/bin/env python3
# https://colab.research.google.com/drive/1Dyauq4kTZoLewQ1cApceUQVNcnnNTzg_?usp=sharing#scrollTo=FqfebeAdT073

from unsloth import FastLanguageModel
import torch
from trl import SFTTrainer
from transformers import TrainingArguments
from datasets import load_dataset
import yaml

with open('config.yaml', 'r') as config_file:
    try:
        data = yaml.safe_load(config_file)
        print(data)
    except yaml.YAMLError as exc:
        print(exc)

max_seq_length = 2048  # Supports RoPE Scaling interally, so choose any!
# Get LAION dataset
url = "https://huggingface.co/datasets/laion/OIG/resolve/main/unified_chip2.jsonl"
dataset = load_dataset("json", data_files={"train": url}, split="train")

# Load Llama model
model, tokenizer = FastLanguageModel.from_pretrained(
    # Supports Llama, Mistral - replace this!
    model_name="unsloth/llama-2-7b-bnb-4bit",
    max_seq_length=max_seq_length,
    dtype=None,
    load_in_4bit=True,
)

# Do model patching and add fast LoRA weights
model = FastLanguageModel.get_peft_model(
    model,
    r=16,
    target_modules=["q_proj", "k_proj", "v_proj", "o_proj",
                    "gate_proj", "up_proj", "down_proj",],
    lora_alpha=16,
    lora_dropout=0,  # Supports any, but = 0 is optimized
    bias="none",    # Supports any, but = "none" is optimized
    use_gradient_checkpointing=True,
    random_state=3407,
    max_seq_length=max_seq_length,
)

trainer = SFTTrainer(
    model=model,
    train_dataset=dataset,
    dataset_text_field="text",
    max_seq_length=max_seq_length,
    tokenizer=tokenizer,
    args=TrainingArguments(
        per_device_train_batch_size=2,
        gradient_accumulation_steps=4,
        warmup_steps=10,
        max_steps=60,
        fp16=not torch.cuda.is_bf16_supported(),
        bf16=torch.cuda.is_bf16_supported(),
        logging_steps=1,
        output_dir="outputs",
        optim="adamw_8bit",
        seed=3407,
    ),
)

trainer.train()

# model.save_pretrained("lora_model")  # Local saving

# model.save_pretrained_gguf("model_gguf", tokenizer, quantization_method="q4_k_m")

model.push_to_hub_gguf("model_gguf", tokenizer,
                       quantization_method="q4_k_m", token="")
