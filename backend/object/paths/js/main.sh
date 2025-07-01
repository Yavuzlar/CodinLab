#!/bin/bash
test=(-tests-)  # test array is defined

export TERM=xterm  # TERM variable is set

# If test array is empty, it runs once. This is a path used for learning and doesn't require validation.
if [ ${#test[@]} -eq 0 ]; then
    result=$(node main.js)
    echo "Test Passed|||$result|||_|||_"
    exit 0
fi

# Test loop
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    
    # Runs the JavaScript file
    result=$(node main.js 2>&1)
    
    # Compares the actual result with the expected result
    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|||$result|||_|||_"
    else
        echo "_|||$result|||$expected_result|||_"
        exit 2
    fi
done
