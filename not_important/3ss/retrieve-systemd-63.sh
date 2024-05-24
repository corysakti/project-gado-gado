#!/bin/bash
# Source file path
destination_file="kpk-voucher-tools.service.bak"

# Destination details
source_user="odp"
source_ip="10.64.1.63"
source_path="/etc/systemd/system/kpk-voucher-tools.service"

# SCP command to copy the file
scp "${source_user}@${source_ip}:${source_path}" "$destination_file"

# Check if SCP command was successful
if [ $? -eq 0 ]; then
    echo "File copied successfully."
else
    echo "Error: File copy failed."
fi
