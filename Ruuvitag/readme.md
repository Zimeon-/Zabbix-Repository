Zabbix Ruuvitag

1. Follow installation instructions for Python Ruuvitag (https://github.com/ttu/ruuvitag-sensor)
2. Create directory structure for Ruuvitag scripts #mkdir -p /etc/zabbix/scripts/ruuvitag
3. Copy python and json files to created folder /etc/zabbix/scripts/ruuvitag
4. Use find_ruuvitags.py to locate your ruuvitags to find out their bluetooth macaddr # python /etc/zabbix/scripts/ruuvitag/find_tags.py
5. Update configuration file with correct MAC, Name and Displayname values # vi /etc/zabbix/scripts/ruuvitag/ruuvitags.json
6. Add Userparameters to the end of the zabbix agent configuration # vi /etc/zabbix/zabbix_agentd.conf

#Ruuvitag Key

UserParameter=ruuvitag.get,python /etc/zabbix/scripts/ruuvitag/get_data.py

UserParameter=ruuvitag.discover,python /etc/zabbix/scripts/ruuvitag/discover_tags.py

7. Restart the Zabbix Agent for user parameters to work.
8. Test the python script to get date # python /etc/zabbix/scripts/ruuvitag/get_data.py
9. Add a cron entry for data sender # crontab -e (You can use the zabbix agent item instead of crontab, test functionality!)

#Send Ruuvitag Updates every 1min

*/1 * * * * python /etc/zabbix/scripts/ruuvitag/get_data.py
10. Import Zabbix Template .yaml
11. Add Zabbix Template to your host of choise
