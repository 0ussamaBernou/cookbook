from yaml import load

# try to import the c version of yaml SafeLoader
try:
    from yaml import CSafeLoader as SafeLoader
except ImportError:
    from yaml import SafeLoader

with open("config.yaml", "r") as f:
    config_file = f.read()
# config_file = "./config.yaml"
config = load(config_file, SafeLoader)
print(config)
