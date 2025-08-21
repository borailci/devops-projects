#!/bin/bash

# Nginx Log Analyzer Script
# This script analyzes nginx access logs and provides statistics
# Author: DevOps Engineer
# Date: $(date +%Y-%m-%d)

# Colors for output formatting
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if log file is provided
if [ $# -eq 0 ]; then
    echo -e "${RED}Usage: $0 <nginx-log-file>${NC}"
    echo -e "${YELLOW}Example: $0 nginx-access.log${NC}"
    exit 1
fi

LOG_FILE="$1"

# Check if log file exists
if [ ! -f "$LOG_FILE" ]; then
    echo -e "${RED}Error: Log file '$LOG_FILE' not found!${NC}"
    exit 1
fi

# Display header
echo -e "${BLUE}=== NGINX LOG ANALYSIS REPORT ===${NC}"
echo -e "${GREEN}Log file: $LOG_FILE${NC}"
echo -e "${GREEN}Analysis date: $(date)${NC}"
echo -e "${GREEN}Total log entries: $(wc -l < "$LOG_FILE")${NC}"
echo -e "${BLUE}==================================${NC}"

# Function to analyze top 5 IP addresses
analyze_top_ips() {
    echo -e "\n${YELLOW}Top 5 IP addresses with the most requests:${NC}"
    awk '{print $1}' "$LOG_FILE" | sort | uniq -c | sort -nr | head -5 | while read count ip; do
        echo -e "${GREEN}$ip - $count requests${NC}"
    done
}

# Function to analyze top 5 requested paths
analyze_top_paths() {
    echo -e "\n${YELLOW}Top 5 most requested paths:${NC}"
    awk '{print $7}' "$LOG_FILE" | sort | uniq -c | sort -nr | head -5 | while read count path; do
        echo -e "${GREEN}$path - $count requests${NC}"
    done
}

# Function to analyze top 5 status codes
analyze_top_status_codes() {
    echo -e "\n${YELLOW}Top 5 response status codes:${NC}"
    awk '{print $9}' "$LOG_FILE" | sort | uniq -c | sort -nr | head -5 | while read count status; do
        echo -e "${GREEN}$status - $count requests${NC}"
    done
}

# Function to analyze top 5 user agents
analyze_top_user_agents() {
    echo -e "\n${YELLOW}Top 5 user agents:${NC}"
    awk -F'"' '{print $6}' "$LOG_FILE" | sort | uniq -c | sort -nr | head -5 | while read count agent; do
        echo -e "${GREEN}\"$agent\" - $count requests${NC}"
    done
}

# Function to show additional statistics
show_additional_stats() {
    echo -e "\n${BLUE}=== ADDITIONAL STATISTICS ===${NC}"
    
    echo -e "\n${YELLOW}Request methods:${NC}"
    awk '{print $6}' "$LOG_FILE" | tr -d '"' | sort | uniq -c | sort -nr | head -5 | while read count method; do
        echo -e "${GREEN}$method - $count requests${NC}"
    done
    
    echo -e "\n${YELLOW}Response size distribution:${NC}"
    awk '{
        size = $10
        if (size == "-") size = 0
        if (size < 1024) print "< 1KB"
        else if (size < 10240) print "1KB-10KB"
        else if (size < 102400) print "10KB-100KB"
        else if (size < 1048576) print "100KB-1MB"
        else print "> 1MB"
    }' "$LOG_FILE" | sort | uniq -c | sort -nr | while read count range; do
        echo -e "${GREEN}$range - $count requests${NC}"
    done
    
    echo -e "\n${YELLOW}Hourly request distribution (last 24 hours):${NC}"
    awk '{
        # Extract hour from timestamp like [04/Oct/2024:00:00:18 +0000]
        match($4, /[0-9]{2}:[0-9]{2}:[0-9]{2}/)
        hour = substr($4, RSTART, 2)
        print hour
    }' "$LOG_FILE" | sort | uniq -c | sort -k2 | while read count hour; do
        echo -e "${GREEN}${hour}:00 - $count requests${NC}"
    done
}

# Function to generate summary report
generate_summary() {
    echo -e "\n${BLUE}=== SUMMARY ===${NC}"
    
    # Get unique IPs
    unique_ips=$(awk '{print $1}' "$LOG_FILE" | sort | uniq | wc -l)
    echo -e "${GREEN}Unique IP addresses: $unique_ips${NC}"
    
    # Get error rate (4xx and 5xx status codes)
    total_requests=$(wc -l < "$LOG_FILE")
    error_requests=$(awk '$9 ~ /^[45][0-9]{2}$/ {count++} END {print count+0}' "$LOG_FILE")
    error_rate=$(echo "scale=2; $error_requests * 100 / $total_requests" | bc -l 2>/dev/null || echo "0")
    echo -e "${GREEN}Error rate: ${error_rate}% ($error_requests out of $total_requests requests)${NC}"
    
    # Get most active IP
    most_active_ip=$(awk '{print $1}' "$LOG_FILE" | sort | uniq -c | sort -nr | head -1)
    echo -e "${GREEN}Most active IP: $most_active_ip${NC}"
}

# Main execution
main() {
    analyze_top_ips
    analyze_top_paths
    analyze_top_status_codes
    analyze_top_user_agents
    show_additional_stats
    generate_summary
    
    echo -e "\n${BLUE}=== END OF REPORT ===${NC}"
    echo -e "${YELLOW}Report generated successfully!${NC}"
}

# Alternative solution using grep and sed (commented out)
# This is the stretch goal - alternative implementation
alternative_solution() {
    echo -e "\n${BLUE}=== ALTERNATIVE SOLUTION USING GREP AND SED ===${NC}"
    
    # Top IPs using grep and sed
    echo -e "\n${YELLOW}Top 5 IPs (alternative method):${NC}"
    grep -oE '^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+' "$LOG_FILE" | sort | uniq -c | sort -nr | head -5
    
    # Top paths using grep and sed
    echo -e "\n${YELLOW}Top 5 paths (alternative method):${NC}"
    grep -oE '"[A-Z]+ [^ ]+ HTTP/[0-9.]+' "$LOG_FILE" | sed 's/"[A-Z]* \([^ ]*\) HTTP.*/\1/' | sort | uniq -c | sort -nr | head -5
    
    # Top status codes using grep and sed
    echo -e "\n${YELLOW}Top 5 status codes (alternative method):${NC}"
    grep -oE '" [0-9]{3} ' "$LOG_FILE" | sed 's/[^0-9]//g' | sort | uniq -c | sort -nr | head -5
}

# Check if alternative flag is provided
if [ "$2" = "--alternative" ]; then
    alternative_solution
else
    main
fi

# Option to save report to file
if [ "$2" = "--save" ] || [ "$3" = "--save" ]; then
    report_file="nginx_analysis_report_$(date +%Y%m%d_%H%M%S).txt"
    echo -e "${YELLOW}Saving report to $report_file...${NC}"
    $0 "$LOG_FILE" > "$report_file" 2>&1
    echo -e "${GREEN}Report saved to $report_file${NC}"
fi
