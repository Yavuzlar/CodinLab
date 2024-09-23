test=(-tests-)
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    
    compile_output=$(g++ -o main main.cpp 2>&1) 
    
    if [ $? -ne 0 ]; then 
        echo "Compilation failed:"
        echo "$compile_output" 
        exit 1 
    fi
    
    result=$(./main)

    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed"
    else
        echo "Test Failed:"
        echo " Expected: $expected_result,"
        echo "but got $result"
    fi
done
