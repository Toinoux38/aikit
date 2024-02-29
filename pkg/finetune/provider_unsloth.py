#!/usr/bin/env python3
# https://colab.research.google.com/drive/1Dyauq4kTZoLewQ1cApceUQVNcnnNTzg_?usp=sharing

from unsloth import FastLanguageModel
import torch
from trl import SFTTrainer
from transformers import TrainingArguments
from datasets import load_dataset
import yaml

with open('config.yaml', 'r') as config_file:
    try:
        data = yaml.safe_load(config_file)
        print(data)   # TODO: remove this
    except yaml.YAMLError as exc:
        print(exc)

cfg = data.get('config').get('unsloth')

max_seq_length = cfg.get('maxSeqLength')
url = data.get('datasets')[0]['source']
dataset = load_dataset("json", data_files={"train": url}, split="train")

model, tokenizer = FastLanguageModel.from_pretrained(
    model_name=data.get('baseModel'),
    max_seq_length=max_seq_length,
    dtype=None,
    load_in_4bit=True,
)

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
    dataset_num_proc = 2,
    packing = cfg.get('packing'), # Can make training 5x faster for short sequences.
    args=TrainingArguments(
        per_device_train_batch_size=cfg.get('batchSize'),
        gradient_accumulation_steps=cfg.get('gradientAccumulationSteps'),
        warmup_steps=cfg.get('warmupSteps'),
        max_steps=cfg.get('maxSteps'),
        learning_rate = cfg.get('learningRate'),
        fp16=not torch.cuda.is_bf16_supported(),
        bf16=torch.cuda.is_bf16_supported(),
        logging_steps=cfg.get('loggingSteps'),
        optim=cfg.get('optimizer'),
        weight_decay = cfg.get('weightDecay'),
        lr_scheduler_type = cfg.get('lrSchedulerType'),
        seed=cfg.get('seed'),
        output_dir="outputs",
    ),
)
trainer.train()

output = data.get('output')
model.save_pretrained_gguf(output.get('name'), tokenizer,
                           quantization_method=output.get('quantize'))
