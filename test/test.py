import yaml

with open('aikitfile-unsloth.yaml', 'r') as config_file:
    try:
        data = yaml.safe_load(config_file)
        # print(data)
    except yaml.YAMLError as exc:
        print(exc)

output = data.get('output')

if output.get('token') != "":
    print('hi')


# cfg = data.get('config').get('unsloth')
# ml = cfg.get('maxSeqLength')
# print(ml)


# ds = data.get('datasets')[0]['source']
# print(ds)
