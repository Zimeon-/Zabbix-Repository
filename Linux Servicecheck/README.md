#### Zabbix Service checking ####
Add the service to visudo to allow zabbix to query the status. 

Example:
#Zabbix User needs privileges to chek service status
zabbix ALL = NOPASSWD: /etc/init.d/rsyslog

If you are using rhel/centos, also remove the requirement of a tty in /etc/sudoers by adding # before requiretty
Example:
# Defaults    requiretty


You can test result of query by using the following command:
su -s /bin/bash zabbix -c "zabbix_agentd -t service.status[rsyslog]"