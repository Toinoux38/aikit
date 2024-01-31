import yaml

with open('config.yaml', 'r') as config_file:
    try:
        data = yaml.safe_load(config_file)
        # print(data)
    except yaml.YAMLError as exc:
        print(exc)

cfg = data.get('config').get('unsloth')
ml = cfg.get('maxSeqLength')
print(ml)


# ds = data.get('datasets')[0]['source']
# print(ds)
