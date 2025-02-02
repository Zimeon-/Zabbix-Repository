# Introduction

A basic http server written on Golang to receive Zabbix server data stream and export it into file.
Currently supports streaming of events and history data.
Supports both http and https.

## Requirements

* Go >= 1.18

## Setup

Make sure you have the correct go version and run `go build`. This will create an executable called `Example_interface`,
in project directory.

## Configuration

The interface can be configured by providing an json configuration flag with the `-c` command. If no value for the flag
is provided, then the default config file `config.json` in the application root directory will be used.

**port** — The port for the http server
*Default value:* 80

**data_path** — Path for files for the received data.
*Default value:* 'path_to_executable_directory'/data

**log_path** — Path for log file.
*Default value:* 'path_to_executable_directory'/Example_interface.log

**enable_tls** — If true uses https instead of http
*Default value:* false

**cert_file** — Full path to cert file for https.
*Default value:*

**key_file** — Full path to key file for https.
*Default value:*

## Zabbix server configuration

Receiver listens for history data stream on path v1/history and events on v1/events. So on server the connector URLs
must be configured as http://<receiver address>:<receiver port>/v1/history and
http://<receiver address>:<receiver port>/v1/events. For encrypted connection specify cert_file and key_file
configuration options and replace http with https in URL.
