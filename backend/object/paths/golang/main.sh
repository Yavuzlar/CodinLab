test=(-tests-)
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    
   $(go install golang.org/x/tools/cmd/goimports@latest && goimports -w ../main.go 2>&1)
    
    compile_output=$(go build -o main ../main.go 2>&1)
    
    if [ $? -ne 0 ]; then
        echo "Compilation failed:"
        echo "$compile_output" 
        exit 1  
    fi
    
    result=$(go run ../main.go)

    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed"
    else
        echo "Test Failed:"
        echo " Expected: $expected_result,"
        echo "but got $result"
    fi
done