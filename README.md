# gocamole
Go tool for monitor php-fpm status

## Usage ##
This tool is designed for monitoring status of several PHP-FPM pools generally.
For example, if you have 3 pools in one server serving the same site 
(people say this was has more performance), you'd like to know the state of all pools in general.

So, for this you need to do next:

* compile binary: `go build -o gocamole`
* put this gocamole on some path in the monitoring server
* put config to the server as well 
(you can find [config.example.json](https://github.com/scukonick/gocamole/blob/master/config.example.json) in the root of the project
* run gocamole with some metric:
```
[root@server ~]# ./gocamole --config /etc/gocamole.json active_processes
3
[root@server ~]# ./gocamole --config /etc/gocamole.json total_processes
300
[root@server ~]# ./gocamole --config /etc/gocamole.json available
1
[root@server ~]# ./gocamole --config /etc/gocamole.json accepted_conn
121745657
```
### Zabbix ###
You can find config for zabbix user parameters here:
[https://github.com/scukonick/gocamole/blob/master/zabbix_userparameter.conf](https://github.com/scukonick/gocamole/blob/master/zabbix_userparameter.conf)
