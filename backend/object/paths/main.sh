test=(-tests-)
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    result=$(-cmd-)
    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed"
    else
        echo " Test Failed:"
        echo " Expected: $expected_result,"
        echo " but got $result"
    fi
done

