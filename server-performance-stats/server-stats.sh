#!/bin/bash

# Script to display server performance statistics with cross-platform compatibility.

# --- Style Definitions ---
# Using tput for portability
BOLD=$(tput bold)
UNDERLINE=$(tput smul)
RESET=$(tput sgr0)
RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)

# --- Helper Functions ---
print_header() {
    echo "${BOLD}${GREEN}$1${RESET}"
    echo "${YELLOW}-----------------------------${RESET}"
}

# --- OS Detection ---
OS_TYPE=$(uname)

echo "${BOLD}Server Performance Statistics${RESET}"
echo "==========================="
echo "OS Detected: ${BOLD}$OS_TYPE${RESET}"
echo ""


# --- CPU Usage ---
print_header "Total CPU Usage"
if [[ "$OS_TYPE" == "Linux" ]]; then
    # Linux: Use top to get idle percentage, then subtract from 100
    CPU_IDLE=$(top -b -n 1 | grep "Cpu(s)" | awk '{print $8}' | cut -d'%' -f1)
    CPU_USAGE=$(echo "100 - $CPU_IDLE" | bc)
    echo "CPU Usage: ${CPU_USAGE}%"
elif [[ "$OS_TYPE" == "Darwin" ]]; then
    # macOS: Grep CPU usage from top
    top -l 1 | grep "CPU usage"
fi
echo ""


# --- Memory Usage ---
print_header "Total Memory Usage"
if [[ "$OS_TYPE" == "Linux" ]]; then
    # Linux: Use free for detailed memory stats
    free -h
elif [[ "$OS_TYPE" == "Darwin" ]]; then
    # macOS: Parse memory info from top
    MEM_INFO=$(top -l 1 | grep "PhysMem")
    USED_MEM=$(echo "$MEM_INFO" | awk '{print $2}')
    UNUSED_MEM=$(echo "$MEM_INFO" | awk '{print $6}')
    echo "Used: ${USED_MEM} | Unused: ${UNUSED_MEM}"
fi
echo ""


# --- Disk Usage ---
print_header "Total Disk Usage"
# df is fairly consistent, but we'll format it nicely.
# We'll show the main filesystem usage.
df -h | grep -E '^/dev/|Filesystem'
echo ""


# --- Top 5 CPU Processes ---
print_header "Top 5 Processes by CPU"
if [[ "$OS_TYPE" == "Linux" ]]; then
    ps -eo pid,user,%cpu,comm --sort=-%cpu | head -n 6
elif [[ "$OS_TYPE" == "Darwin" ]]; then
    ps -arcx -o "pid,%cpu,comm" | head -n 6
fi
echo ""


# --- Top 5 Memory Processes ---
print_header "Top 5 Processes by Memory"
if [[ "$OS_TYPE" == "Linux" ]]; then
    ps -eo pid,user,%mem,comm --sort=-%mem | head -n 6
elif [[ "$OS_TYPE" == "Darwin" ]]; then
    ps -arcx -o "pid,%mem,comm" | head -n 6
fi
echo ""


# --- Stretch Goals ---
print_header "System Information"
if [[ "$OS_TYPE" == "Linux" ]]; then
    # Show OS version from /etc/os-release
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        echo "OS Version: ${PRETTY_NAME}"
    fi
    # Show logged in users
    echo "Logged-in Users: $(who | wc -l)"
elif [[ "$OS_TYPE" == "Darwin" ]]; then
    # Show macOS version
    echo "OS Version: $(sw_vers -productName) $(sw_vers -productVersion)"
fi
# Uptime and Load Average (consistent across both)
echo "Uptime & Load: $(uptime)"
echo ""
