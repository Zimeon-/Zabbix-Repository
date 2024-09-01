# Zabbix Template for Unifi API

## How to use this template from this repository?
- Download and [import a template](https://www.zabbix.com/documentation/current/manual/xml_export_import/templates#importing) into the supported Zabbix version.
- Generate a new API key at https://unifi.ui.com/api and save it to your favorite password manager
- Create a new [host](https://www.zabbix.com/documentation/current/en/manual/config/hosts/host) in Zabbix for Unifi API monitoring
- Link template "Ubiquiti UniFi API"
- Configure API key {$APIKEY} in host inherited template [macros](https://www.zabbix.com/documentation/current/en/manual/config/macros)

- Note: Template "Ubiquiti UniFi API Host" is used by any discovered Unifi Hosts (Gateways). Discovered Hosts are created in "Unifi Site Hosts" Host Group automatically.

## Template Zabbix results
![msedge_avPwOZ23Wv](https://github.com/user-attachments/assets/ff59e3fc-e427-4ca8-9f8d-aa9ef08faa42)
![msedge_l3VxjwhMnn](https://github.com/user-attachments/assets/730b7c75-8aac-46cd-9f48-2d25ed5f5e75)

Resources:
- https://ubntwiki.com/products/software/unifi-controller/api
- https://developer.ui.com/unifi-api/
