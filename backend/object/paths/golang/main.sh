#!/bin/bash
test=(-tests-) # test dizisi tanımlandı

export TERM=xterm  # TERM değişkeni ayarlandı

# Eğer test dizisi boşsa, bir kere çalıştır. Cevap gerekmeyen öğrenmek için olan bir pathdir.
if [ ${#test[@]} -eq 0 ]; then
    result=$(node main.js)
    echo "Test Passed|||$result|||_|||_"
    exit 0
fi

# Test döngüsü
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"

    # GO dosyasını çalıştır 
    compile_output=$(go build -o main ../main.go 2>&1)
    result=$(go run ../main.go)  

    # Sonucu beklenen sonuç ile karşılaştır
    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|||$result|||_|||_"
    else
        echo "_|||$result|||$expected_result|||_"
        exit 2
    fi
done
