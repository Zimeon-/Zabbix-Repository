# Introduction

This python script can be set as a listener for rsyslog messages containing Zabbix streaming metrics. Any metrics
received will be passed along using the Zabbix Sender binary to the localhost Zabbix Server. 

1. chmod the file with "chmod u+x zbx_logreader.py"

2. Run the listener using "python3 zbx_logreader.py /path/to/logfile.log"


## Requirements

* Python >= 3

## Zabbix server configuration

Items sent need to match a corresponding host, with items as type Zabbix Trapepr.

# 