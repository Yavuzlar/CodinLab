#!/bin/bash
test=(-tests-) 

export TERM=xterm  


# If the test array is empty, this will run. It means that it is a path for learning, not for getting an answer.
if [ ${#test[@]} -eq 0 ]; then
    result=$(node main.js)
    echo "Test Passed|||$result|||_|||_"
    exit 0
fi

# Test loop
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"

    compile_output=$(go build -o main ../main.go 2>&1)

        if [ $? -ne 0 ]; then 
        echo "_|||_|||_|||$compile_output" 
        exit 1 
        fi


    result=$(go run ../main.go)  

    # the result will be compared with the expected result
    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|||$result|||_|||_"
    else
        echo "_|||$result|||$expected_result|||_"
        exit 2
    fi
done
