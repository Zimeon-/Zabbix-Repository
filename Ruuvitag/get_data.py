from ruuvitag_sensor.ruuvitag import RuuviTag
import json, re
import os, time, socket, subprocess

def load_ruuvitags():
   f = open("/home/ruuvi/Ruuvitag/ruuvitags.json")
   data = json.load(f)
   f.close()
   return data

def get_ruuvitag_data(mac):
   from ruuvitag_sensor.ruuvitag import RuuviTag
   # Ruuvitag Bluetooth MAC
   sensor = RuuviTag(mac)
   # update state from the device
   state = sensor.update()
   # get latest state (does not get it from the device)
   state = sensor.state
   return state

def write_zbx_data(zbxhostname, tag, state):
   for key in state:
      value = state[key]
      f.write("{} ruuvitag.{}[{}] {}\n".format(zbxhostname, key, tag["name"], value))

#Load Ruuvitags from config
ruuvitags = load_ruuvitags()
#Open file for writing
epoch_time = int(time.time())
zbxfile = "/tmp/ruuvisender-{}.data".format(epoch_time)
f = open(zbxfile, "w")
#Get Hostname for ZBX File
zbxhostname = socket.gethostname()
#Get Ruuvitag data from sensors in config file
for tag in ruuvitags['config']:
   sensor_data = get_ruuvitag_data(tag["mac"])
   write_zbx_data(zbxhostname, tag, sensor_data)
f.close()
#Catch Zabbix Sender output
proc = subprocess.Popen("/usr/bin/zabbix_sender -c /etc/zabbix/zabbix_agentd.conf -i /tmp/ruuvisender-{}.data".format(epoch_time), shell=True, stdout=subprocess.PIPE)
(out, err) = proc.communicate()
zbxr = re.findall(r"failed: (\d+)",out.decode('utf-8'))
zbxt = re.findall(r"total: (\d+)",out.decode('utf-8'))
#Check if sender failed any items
if zbxr and int(zbxr[0]) == 0:
   print("No Errors Detected, removing temporary file {}".format(zbxfile))
   os.remove(zbxfile)
elif int(zbxt[0]) == 0:
   print("Nothing send to Zabbix Server, use {} to debug Zabbix Sender issues.".format(zbxfile))
   f=open(zbxfile, "a")
   f.write("ZBX Sender:{}".format(out.decode('utf-8')))
   #Close file
   f.close()
else:
   print("Errors found in zabbix sender process. Use {} to debug Zabbix Sender issues.".format(zbxfile))
   f=open(zbxfile, "a")
   f.write("ZBX Sender:{}".format(out.decode('utf-8')))
   #Close file
   f.close()
