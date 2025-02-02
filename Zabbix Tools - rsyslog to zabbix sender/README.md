# Introduction

This python script can be set as a listener for rsyslog messages containing Zabbix streaming metrics. Any metrics
received will be passed along using the Zabbix Sender binary to the localhost Zabbix Server. 

## Requirements

* Python >= 3

## Zabbix server configuration

Items sent need to match a corresponding host, with items as type Zabbix Trapepr.
