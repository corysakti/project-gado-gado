#!/bin/bash

# Source file path
source_file="standalone-rita-12.xml"

# Destination details
destination_user="rita2"
destination_ip="10.64.17.12"
destination_path="/app/EAP-7.0.0/standalone/configuration/standalone.xml1"

# SCP command to copy the file
scp "$source_file" "${destination_user}@${destination_ip}:${destination_path}"

# Check if SCP command was successful
if [ $? -eq 0 ]; then
    echo "File copied successfully."
else
    echo "Error: File copy failed."
fi
