---
title: Fine Tuning
---


```bash
docker buildx create --name builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'
docker build --allow security.insecure -f test/aikitfile-unsloth.yaml -output out .
```