#!/bin/bash

set -e

# # Extract scrambled from output2
unscrambled2=$(python3 -c "import json; data=json.load(open('output.txt')); print(data['unscrambled_raw'])" 2>/dev/null)n first command
echo "Running first command..."
go run cmd/main.go -f input.txt > ignore_output.txt

# Extract scrambled
scrambled=$(python3 -c "import json; data=json.load(open('output.txt')); print(data['scrambled_raw'])" 2>/dev/null)

if [ -z "$scrambled" ]; then
    echo "Failed to extract scrambled from output.txt"
    exit 1
fi

# Write to input2.txt
echo "$scrambled" > input2.txt  

# Run second command
echo "Running second command..."
go run cmd/main.go -f input2.txt > ignore_output.txt

# Extract scrambled from output2
unscrambled2=$(python3 -c "import json; data=json.load(open('output.txt')); print(data['unscrambled'])" 2>/dev/null)

if [ -z "$unscrambled2" ]; then
    echo "Failed to extract unscrambled  from output.txt"
    exit 1
fi

# Read original input
original=$(cat input.txt)

# Compare
if [ "$unscrambled2" == "$original" ]; then
    echo "Test passed: unscrambled2 matches original input"
else
    echo "Test failed"
    echo "Original: $original"
    echo "Unscrambled2: $unscrambled2"
fi