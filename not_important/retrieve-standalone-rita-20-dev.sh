#!/bin/bash

# Source file path
destination_file="standalone-rita-dev-20.xml"

# Destination details
source_user="rita2"
source_ip="10.64.17.20"
source_path="/app/EAP-7.0.0-DEV-8080/standalone/configuration/standalone.xml"

# SCP command to copy the file
scp "${source_user}@${source_ip}:${source_path}" "$destination_file"

# Check if SCP command was successful
if [ $? -eq 0 ]; then
    echo "File copied successfully."
else
    echo "Error: File copy failed."
fi
