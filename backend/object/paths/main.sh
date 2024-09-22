test=(-tests-) 
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    result=$(-cmd-) 
    echo "$result+1"
    if [[ "$result" == "$expected_result" ]]; then
        echo "Test passed"
    else
        echo "Test failed "
        echo "Expected: $expected_result"
        echo " Got: $result"
    fi
done