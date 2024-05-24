#!/bin/bash

# Get the directory of the script
script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Define the source file and destination directory
source_file="$script_dir/standalone.xml1"
destination_file="$script_dir/standalone.xml"

# Backup directory
backup_dir="$script_dir/backup"
backup_file="$backup_dir/standalone_backup_$(date +'%Y-%m-%d_%H-%M-%S').xml"

# Check if the source file exists
if [ ! -f "$source_file" ]; then
    echo "Error: Source file '$source_file' not found."
    exit 1
fi

# Create the backup directory if it doesn't exist
mkdir -p "$backup_dir"

# Check if the destination file exists
if [ -f "$destination_file" ]; then
    # Backup the destination file to the backup directory
    cp "$destination_file" "$backup_file"
    echo "Previous version of standalone.xml backed up to $backup_file"
fi

# Copy the source file to the destination directory
cp "$source_file" "$destination_file"

echo "File copied successfully."
