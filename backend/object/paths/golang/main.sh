test=(-tests-)
export TERM=xterm  # TERM değişkeni ayarlandı

for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    
    
    go install golang.org/x/tools/cmd/goimports@latest > /dev/null 2>&1
    goimports -w ../main.go > /dev/null 2>&1

    compile_output=$(go build -o main ../main.go 2>&1)

    if [ $? -ne 0 ]; then
        echo "_|||_|||_|||$compile_output"
        exit 1
    fi

    result=$(go run ../main.go)  

    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|||$result|||_|||_"
    else
        echo "_|||$result|||$expected_result|||_"
        exit 2
    fi
done
