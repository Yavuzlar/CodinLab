test=(-tests-)
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    
    compile_output=$(g++ -o main main.cpp 2>&1) 
    
    if [ $? -ne 0 ]; then 
        echo "Compilation failed:"
        echo "_|*_|*_|*$compile_output" 
        exit 1 
    fi
    
    result=$(./main)

    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|*$result|*_|*_"
    else
        echo "_|*$expected_result|*$result|*_"
    fi
done
