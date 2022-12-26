import json, re
import os, time

def load_ruuvitags():
   f = open("/home/ruuvi/Ruuvitag/ruuvitags.json")
   data = json.load(f)
   f.close()
   return data

config = load_ruuvitags()
ruuvitags = []

for tag in config['config']:
   ruuvitags.append({"{#NAME}": tag["name"], "{#DISPLAYNAME}": tag["displayname"]})

print(json.dumps(ruuvitags))
