test=(-tests-)
export TERM=xterm  

if [ ${#test[@]} -eq 0 ]; then
    result=$(python3 ../main.py)
    echo "Test Passed|||$result|||_|||_"
    exit 0
fi

for i in "${!test[@]}"; do
    expected_result="${test[$i]}"

    result=$(python3 ../main.py)  

    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|||$result|||_|||_"
    else
        echo "_|||$result|||$expected_result|||_"
        exit 2
    fi
done
