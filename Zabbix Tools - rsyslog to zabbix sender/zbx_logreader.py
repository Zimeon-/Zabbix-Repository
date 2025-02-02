import sys
import json
import re
import subprocess
import argparse

# Zabbix server details
ZABBIX_SERVER = "127.0.0.1"  # Change this to your Zabbix server
ZABBIX_SENDER = "/usr/bin/zabbix_sender"  # Path to zabbix_sender binary

# Regex to extract JSON parts from the log line (supports multiple JSON objects)
LOG_JSON_REGEX = re.compile(r'(\{.*?\})(?=#012|$)')

def send_to_zabbix(host, key, value):
    """Send data to Zabbix using zabbix_sender"""
    cmd = [
        ZABBIX_SENDER,
        "-z", ZABBIX_SERVER,
        "-s", host,
        "-k", key,
        "-o", str(value)
    ]
    
    try:
        result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        if result.returncode != 0:
            print(f"Error sending data to Zabbix: {result.stderr}", file=sys.stderr)
        else:
            print(f"Sent to Zabbix: {host}, {key}, {value}")
    except Exception as e:
        print(f"Exception while sending data: {e}", file=sys.stderr)

def extract_streaming_key(item_tags):
    """Extract the value of the 'streaming_key' tag"""
    for tag in item_tags:
        if tag["tag"] == "streaming_key":
            return tag["value"]
    return None

def process_log_line(line):
    """Check for ZbxStream, extract JSON objects, and send to Zabbix"""
    if "ZbxStream" not in line:
        return  # Skip lines that don't contain ZbxStream

    json_matches = LOG_JSON_REGEX.findall(line)
    for json_str in json_matches:
        try:
            log_data = json.loads(json_str)
            host = log_data["host"]["host"]
            item_tags = log_data.get("item_tags", [])
            streaming_key = extract_streaming_key(item_tags)
            value = log_data["value"]

            if streaming_key:
                send_to_zabbix(host, streaming_key, value)
            else:
                print("Warning: 'streaming_key' tag not found, skipping.", file=sys.stderr)
        except (json.JSONDecodeError, KeyError) as e:
            print(f"Error parsing log JSON: {e}\nFaulty JSON: {json_str}", file=sys.stderr)

def follow_log_file(log_file):
    """Follow a log file in real-time using tail -F"""
    with subprocess.Popen(["tail", "-F", log_file], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True) as process:
        for line in process.stdout:
            process_log_line(line)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Zabbix log stream reader")
    parser.add_argument("logfile", help="Path to the log file to follow")
    args = parser.parse_args()

    try:
        follow_log_file(args.logfile)
    except KeyboardInterrupt:
        print("\nExiting...")
        sys.exit(0)
