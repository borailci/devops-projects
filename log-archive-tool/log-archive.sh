#!/bin/bash

# Check if a log directory is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <log-directory>"
  exit 1
fi

LOG_DIR=$1
ARCHIVE_DIR="archives"
DATE=$(date +"%Y%m%d_%H%M%S")
ARCHIVE_FILE="logs_archive_${DATE}.tar.gz"
LOG_FILE="log-archive.log"

# Check if the log directory exists
if [ ! -d "$LOG_DIR" ]; then
  echo "Error: Directory '$LOG_DIR' not found."
  exit 1
fi

# Create the archive directory if it doesn't exist
mkdir -p $ARCHIVE_DIR

# Compress the logs
tar -czf "${ARCHIVE_DIR}/${ARCHIVE_FILE}" -C "$LOG_DIR" .

# Log the archive event
echo "${DATE}: Archived logs from '${LOG_DIR}' to '${ARCHIVE_DIR}/${ARCHIVE_FILE}'" >> $LOG_FILE

echo "Log archive complete. Archive stored in '${ARCHIVE_DIR}/${ARCHIVE_FILE}'"
