# CodinLab Project Documentation

Bu doküman, **CodinLab** projesine içerik eklemeye yardımcı olması için hazırlanmıştır.

Proje içerisinde varsayılan olarak dil dosyası oluşturulmuştur. Bu dosyasının içerisinde mevcut diller (Go, Javascript, Python, C++) yer almaktadır. Ekleyeceğiniz içeriğin dili mevcut değil ise ilk önce o dili eklemeniz gerekmektedir. Eğer içerik ekleyeceğiniz dil mevcut ise bu adımı geçebilirsiniz.

Dillere eklenebilecek içerik olarak iki ayrı kategori bulunmaktadır:

**Path** : Yazılım öğrenmek isteyenlere, öğrenmek istediği dildeki konularda öğretici ve uygulamalı alıştırmalar sunarak o dilde kendini geliştirmesini sağlayan **yol haritasının adımlarıdır**.

**Lab** : Path'de öğrenilen bilgiler ışığında kullanıcıların konu üzerinde daha fazla pratik yapmasını amaçlayan bir **laboratuvar ortamıdır**.

Aşağıdaki başlıklar ile daha detaylı bilgilere sahip olabilirsiniz.

## Dil Ekleme

### Dizin Yapısı

Proje içerisinde dil dosyası object klasörü altında yer almaktadır. Buradaki `inventory.json` dosyası içerisinde varsayılan olarak ekli diller bulunmaktadır. Yeni bir dil eklenmek istendiği zaman bu dosya içerisine ekleme yapılması gerekmektedir.

### Nasıl Yeni Dil Eklenir?

Bunun için öncelikle `inventory.json` dosyasının yapısını incelememiz gerekmektedir:

```json
[
	{
		"id": 1,
		"name": "C++",
		"dockerImage": "gcc:latest",
		"labDir": "object/labs",
		"pathDir": "object/paths/c++",
		"iconPath": "images/c++.png",
		"cmd": [
			"sh",
			"-c",
			"g++ -o main main.cpp && ./main"
		],
		"bashCmd": [
			"bash",
			"-c",
			"./main.sh"
		],
		"fileExtension": "cpp",
		"monacoEditor": "cpp",
		"languages": [
			{
				"lang": "en",
				"title": "What is C++",
				"description": "C is a general-purpose programming language. It was created in the 1970s by Dennis Ritchie and remains very widely used and influential."
			},
			{
				"lang": "tr",
				"title": "C++ nedir?",
				"description": "C genel amaçlı bir programlama dilidir. 1970'lerde Dennis Ritchie tarafından yaratılmıştır ve hala çok yaygın olarak kullanılmakta ve etkili olmaya devam etmektedir."
			}
		]
	}
]
```

#### Açıklama

- **`id`**: Eklenen dilin id'si, onu diğerlerinden ayıran numarasıdır. En son eklenen dilin id'sinin bir fazlası eklenmelidir.

- **`name`**: Buraya eklenen dilin adı büyük harflerle eklenmelidir. Ön yüz tarafında başlık olarak kullanılmaktadır.

- **`dockerImage`**: Buraya eklemek istediğiniz dilin **docker image** adını ve versiyonunu yazmanız gerekmektedir. Yazılması gereken image'i öğrenmek için [dockerhub](https://hub.docker.com/)'a bakmanız yeterlidir. Eğer image'in sağına `latest` yazarsanız bu en son sürümü indir demektir. Dilerseniz belirli bir versiyon da belirtebilirsiniz. Kullanıcıların yazdıkları kodların bu dilde derlenebilmesi için gereklidir.

- **`labDir`**: Burası sabittir. Her zaman **object/labs** yazmanız gereklidir. Lablar dillere göre yalnızca template olarak ayrılması sebebiyle burası hep aynıdır.

- **`pathDir`**: Burası dile göre değişmektedir. Eklediğiniz dile göre **object/paths/`eklenecek dil`** burası değişmektedir.

- **`iconPath`**: Ekleyeceğiniz dil için **object/icons** altına dilin iconunu eklemeniz gerekmektedir.

- **`cmd`**: Cmd içersine, bu dili derlemek ve çalıştırmak için gereken komutu yazmanız gerekmektedir. Bu şekilde kullanıcının yazmış olduğu kod Docker'da çalışacaktır. Komutlar, önce "sh", ardından "-c" ile başlar.

   - **sh**: Shell (kabuk) çalıştırır.
   - **-c**: Gelen komutun shell içerisinde çalıştırılacağını belirtir.

Bunların ardından derlemek & çalıştırmak için gerekli komut yazılır. Bu komut C++ dili için **g++ -o main main.cpp && ./main**'dir. Eklenecek dile göre düzenlenmelidir.

- **`bashCmd`**: Cmd'den farklı olarak burada **main.sh** dosyasını çalıştırmak için gerekli sh komutu yazılmalıdır. Bu dosyanın ne işe yaradığına [buradan](#mainsh-dosyası-yazımı) ulaşabilirsiniz.

- **`fileExtension`**: Eklenecek olan dilin dosya uzantısını yazmanız gerekmektedir. Kullanıcının yazacağı dilin dosyasını belirlenmesi için gerekmektedir.

- **`monacoEditor`**: Bu kısma ön yüzde, kullanıcının yazacağı kod dosyasının dile göre özelleştirilebilmesi için monacoEditor'ün istediği uzantı yazılmalıdır.

- **`languages`**: Json dosyası içerisine yazmanız gereken son özellik ise Türkçe ve İngilizce olacak şekilde; eklenecek programlama dilinin Türkçe/İngilizce olarak ayrı ayrı başlık ve açıklamasıdır.

   - **lang**: Eklenecek dil (tr veya en)
   - **title**: Eklenen programlama dilinin başlığı 
   - **description**: Eklenen programlama dilinin açıklaması

Belirtilen özelliklerin hepsinin yazılması sonucunda yeni dil için lab veya path eklemesi yapılmaya başlanabilmektedir. Yeni bir dil eklediğiniz zaman hata vermemesi için bu dilde en az 1 adet path eklemeniz gerekmektedir.

## Path Ekleme

### 1- Roadmap Dizin Yapısının Anlaşılması

Projenin `objects\paths` dizininde, desteklenen programlama dillerine ait klasörler bulunur. Eğer eklemek istediğiniz dile ait klasör burada bulunmuyorsa:

1. `inventory` dosyasına gerekli formatta yeni [dil eklenmelidir](#dil-ekleme).
2. Daha sonra, bu dizinde eklemek istediğiniz dilin adıyla bir klasör oluşturularak, o dile ait yeni pathler eklenebilir.

Örnek dosya yapısı aşağıda yer almaktadır.

```
object/ 
└── paths/ 
	├── c++/ 
	├── go/ 
	├── js/ 
	└── python/
```


Her dilin klasörü, o dile ait pathleri sıralı bir şekilde numaralandırılmış klasörlerde tutar. Bu klasörler hem path sırasını hem de `id` numarasını temsil eder.

Örnek dosya yapısı aşağıda yer almaktadır.

```
c++/ 
├── 1/ 
├── 2/ 
├── 3/ 
└── 4/
```

### 2- Roadmap'e Yeni Path Ekleme

1. Eklemek istenen dilin klasörüne gidin.(örn  `c++`)
2. En son path klasörünü kontrol edin (örn. `4`). Bir sonraki path için sıradaki numara ile yeni bir klasör oluşturun (örn. `5`).
3. Yeni oluşturulan klasöre şu iki dosyayı ekleyin:
	- `quest.json`: Path'in görev tanımları ve testlerini içerir.
	- `template.txt`: Kullanıcının kod yazma ortamı, Docker ile çalıştırma ve testlerin tanımlarını içerir.

## Lab Ekleme

### 1- Labs Dizin Yapısının Anlaşılması

Projenin `objects\labs` dizininde, her soru için bir klasör bulunur. Bütün lablar bu klasör içinde yer alınır. Sorunun hangi programlama dili içerisinde görüleceği 'code template' lere bağlı olarak değişmektedir.

1. `inventory` dosyasına gerekli formatta yeni dil eklenmelidir.

Örnek dosya yapısı aşağıda yer almaktadır.
```
object/ 
└── labs/ 
	├── 1
	    └──  quest.json
	    └──  c++.txt
	    └──  js.txt
	├── 2
```
Burada labs klasöründe 2 soru eklenmiştir.

### 2- Yeni Lab Ekleme

1. Labs klasörüne istediğiniz isimde bir klasör açıp bir 'quest.json' oluşturunuz.
2. Bu eklenen lab için 'code templates' yazınız. İleride detaylı açıklanacaktır.

## Soru Yapısı - Path & Lab

Bu dosya, sorular hakkında bilgiler, görev tanımı ve testleri içerir.

**Örnek Yapı:**

```json
{
    "id": 4, // Path veya Lab ID (Yeni path veya lab için sıra numarasıyla eşleşmeli)
    "languages": [
        {
            "lang": "en",
            "title": "Conditional Statements",
            "description": "Understanding conditional statements in C++",
            "content": "Learn how to use if-else statements to make decisions...", // Yalnızca path'de vardır
            "note": "Detailed instructions about the task in English...",
            "hint":"Hint for solving the lab" // Yalnızca lablarda eklenir pathde yoktur
        },
        {
            "lang": "tr",
            "title": "Koşullu İfadeler",
            "description": "C++'ta koşullu ifadeleri kullanmayı öğrenin.",
            "content": "C++ programında if-else ifadelerini kullanarak...",
            "note": "Görev için detaylı talimatlar Türkçe olarak burada...",
            "hint":"Labı çözmek için ipucu"
        }
    ],
    "quest": {
        "difficulty": 1, // Zorluk seviyesi (1: kolay, 3: zor)
        "funcName": "checkEvenOdd", // Yazılacak fonksiyonun adı
        "tests": [
            {
                "input": [4],
                "output": ["Even"]
            },
            {
                "input": [7],
                "output": ["Odd"]
            }
        ],
        "codeTemplates": [
            {
                "programmingID": 1,
                "templatePath": "object/paths/c++/4/template.txt" // Template dosyasının yolu
            }
        ]
    }
}
```

#### Açıklama

- **`id`**: Path ID'si. Path klasör ismiyle eşleşmelidir.
- `languages`: Path hakkında dil bazlı açıklamalar. Kullanıcı arayüzünde farklı dillerde görüntülenir.
    - 'lang': Dili belirtir.
    - 'title': Sorunun başlığı.
    - 'description': Sorunun açıklaması.
    - 'content (paths)': Bu soruyu soru yapan kısım. Burada soruyu açıklıyorsunuz. Bu sadece path oluştururken gereklidir.
    - 'note': Bu ise soru nasıl yapılır? diye açıkladığınız öğrettiğiniz kısım.
    - 'hint (labs)': Bu kısım lab sorularında az da olsa destek vermek için açılan kısım. Burada lab'larda ki hint kısmını yazıyorsunuz. Örnek: 'bu soruyu çözmek için recursive function ları öğrenin.'
- `quest`: Görev tanımları:
	- `funcName`: Kullanıcının yazacağı fonksiyonun adı. (Eğer fonksiyonun adı boş ise veya fonksiyonun adı **main** ise [**main.sh**](#mainsh-dosyası-yazımı) dosyasının oluşturulması gerekmektedir.)
	- `tests`: Kullanıcının kodunun test edileceği örnekler (girdi/çıktı eşleşmeleri).
	- `programmingID`: Inventory dosyasında eklenen dilin `ID` değeri.
	- `codeTemplates`: Bu path için kullanılacak template dosyasının yolu. Burada her soru için bir template yazılır, bu template sayesinde biz bu sorunun nasıl başladığını, nasıl kontrol edileceğini ve çıktı olarak ne döneceğini anlarız. Eğer bu template yanlış yapılırsa soru düzgün bir şekilde çözülemez, hatalar verir. Bu template'lerin yazılış şekilleri aşağıdaki gibidir.

Bu bilgilerin eklenmesi zorunludur. Örneklere de bakılarak ekleme yapılabilir.

## Code Template Yapısı - Path & Labs

Bu dosya, kullanıcının kodunu yazacağı alan, Docker ortamında çalıştırma mantığı ve test mekanizmasını içerir. Roadmap'ler tek dil içerdiği için onlar için tek bir tane template yeterlidir. Örnek olarak c++ için bir path yazıyorsanız c++ için bir template yazmanız bu soru için yeterlidir.
Eğer bir lab yazıyorsanız, o labı görmek istediğiniz her dil için ayrı template yazmak zorundasınız. Ardından 'quest.json' içerisinde 'codeTemplates' dizisi içinde doğru 'id' ve 'path' ile belirtirseniz otomatik olarak doğru programlama diline doğru soru eklenecektir. Şimdi nasıl template yazabileceğimize bakalım.

**Bölümler:**

- **FRONTEND**: Kullanıcının ön yüzde kod yazacağı şablon.
- **DOCKER**: Kodun test edilip derlenebilmesi için oluşturulan docker şablonu.
- **CHECK**: Test senaryoları.

**Örnek (C++ için)**

```txt
## FRONTEND

#include <iostream>
#include <string>

using namespace std;

string $funcname$(int n) {
    // Write your code here
}

int main() {
    int n;
    cout << "n değerini girin: ";
    cin >> n;

    string result = $funcname$(n);
    cout << "Fonksiyondan dönen sonuç: " << result << endl;
    return 0;
}

## DOCKER

$imps$

$usercode$

int main(){
    $checks$

    std::cout << "$success$|||" << result$res$<< "|||_|||_" << std::endl;
    return 0;
}

## CHECK

std::string result$rnd$ = $funcname$($input$);
if (result$rnd$ != $output$) {
    std::cout << "_|||" << result$rnd$ << "|||$out$|||_" << std::endl;
    return 0;
}

```

#### FRONTEND TAGI

Bu tagin aşağısında bulunan alan, kullanıcı arayüzünde kullanıcıdan kod yazılmasını istenen kod editörüne yansıtılır. Kullanıcı bu template üzerinde çalışmalara başlar.

- `$funcname$`: quest.json'da belirtilen fonksiyon adıyla değiştirilir.

#### DOCKER TAGI

Bu tagin aşağısında kullanıcının kodunu çalıştırmak için run edilecek kod yer alır.

`$usercode$`: Kullanıcının yazdığı kod. (Sadece istenen fonksiyonun içeriği alınır.)

`$checks$`: Test senaryoları. `#CHECKS` içerisinde yer alan kod `quest.json` dosyasındaki test verilerine dayanarak testler bu alana gelir.

`$imps$` : Bu değişken, kullanıcı tarafından eklenen kütüphanelerle backend tarafından ihtiyaç duyulan ek kütüphaneleri birleştirir.

- Örneğin, kullanıcı Go'da `fmt` ve `os` kütüphanelerini kullanıyorsa, `$imps$` alanında bu kütüphaneler yer alır.
- Backend ayrıca, Docker'ın çalışması için gerekli olan diğer kütüphaneleri de ekleyebilir.

`$success$`: Backend'de `Test Passed` alanı ile değiştirilir. Yani kullanıcı tüm testleri geçti anlamına gelir.

`result$res`: Testler de yer alan son testin sonucu anlamına gelir.

Backend'e gerekli formatta veri döndürmek zorunludur. Aksi takdirde backend bu dönüşü anlamlandıramaz ve front-end doğru çalışamaz. İstenen formatın detaylarını aşağıda bulabilirsiniz.

#### CHECK TAGI

`## CHECK` bölümü, kullanıcının kodunu test eden ve sonuçları değerlendiren kısımdır.

Bölüm yapısının bir örneği aşağıda yer almaktadır:

```txt
## CHECK

std::string result$rnd$ = $funcname$($input$);
if (result$rnd$ != $output$) {
    std::cout << "_|||" << result$rnd$ << "|||$out$|||_" << std::endl;
    return 0;
}
```

- `$out$` değeri `quest.json` dosyasında yer alan gerekli testin output değerini ifade eder.
- `result$rnd$` fonksiyonun her çıktısını aynı değişkene atmamak, sürekli aynı isimde değişken üretmemek adına yazılmıştır. Rastgele değişken isimleri oluşturulmasına yarar.

##### Çalışma Prensibi

Bu alan temel olarak `quest.json` dosyasında gerekli path'in test-case'lerinin test edildiği yerdir.
Fonksiyona gerekli input verilerek fonksiyondan dönen sonuç bir değişkene atılır ve kullanıcının yazdığı kodun dönüşü, beklenen dönüşe eşit mi diye bakılır. Eğer eşit değilse gerekli format ekrana yazdırılır ve return edilerek durdurulması sağlanır.

`#CHECK` alanı her test-case için `#DOCKER` içerisinde yer alan `$checks` alanına yazılır.

Kullanıcının yazdığı kod gerekli input ile beklenen outputu verdi ise yine gerekli formatta testlerinin başarılı olduğu yönünde dönüş yapılır.

#### Çıktı Formatı

Test sonuçları, `|||` ile ayrılmış formatta backend'e gönderilir. Backend bu verileri parse ederek frontend'e gönderir ve kullanıcıya gösterir.

**Örnek Çıktı**
`Test Passed|||5|||5|||`

- **Birinci Alan**: `Test Passed` (Test başarılı.) `template.txt` dosyasında `$success$` yazılması yeterlidir.

- **İkinci Alan**: Kullanıcının kodunun döndürdüğü sonuç (örneğin `5`).  `template.txt` dosyasında `result$rnd$` yazılması yeterlidir.

- **Üçüncü Alan**: Beklenen sonuç (örneğin `5`). `template.txt` dosyasında `$out$` yazılması yeterlidir.

- **Dördüncü Alan**: Hata mesajı (varsa).

### main.sh Dosyası Yazımı

Bu dosyanın pathler için eklenmesi gerekmektedir. Pathler içerisinde main fonksiyon veya fonksiyon adı olmadan bir kodun çalıştırılabilmesi amacıyla yazılmaktadır.

#### Dosya Dizini

main.sh dosyası:

```
object/ 
└── paths/ 
	├── c++/ 
	├── go/ 
	├── js/ 
	└── python/
```

dizinler altında yer almaktadır. Her dil için kendi klasörü altında ayrı bir sh dosyasının yazılması gerekmektedir.

#### main.sh Dosyası Nasıl Yazılmalıdır?

**Go dili için örnek bir yapı:**

```sh
#!/bin/bash
test=(-tests-) # test dizisi tanımlandı

export TERM=xterm  # TERM değişkeni ayarlandı

# Eğer test dizisi boşsa, bir kere çalıştır. Cevap gerekmeyen, öğrenmek için olan bir pathdir.
if [ ${#test[@]} -eq 0 ]; then
    result=$(go run ../main.go)
    echo "Test Passed|||$result|||_|||_"
    exit 0
fi

# Test döngüsü
for i in "${!test[@]}"; do
    expected_result="${test[$i]}"
    
    go install golang.org/x/tools/cmd/goimports@latest > /dev/null 2>&1
    goimports -w ../main.go > /dev/null 2>&1

    # GO dosyasını çalıştır 
    compile_output=$(go build -o main ../main.go 2>&1)

    if [ $? -ne 0 ]; then 
        echo "_|||_|||_|||$compile_output" 
        exit 1 
    fi

    result=$(go run ../main.go)  

    # Sonucu beklenen sonuç ile karşılaştır
    if [[ "$result" == "$expected_result" ]]; then
        echo "Test Passed|||$result|||_|||_"
    else
        echo "_|||$result|||$expected_result|||_"
        exit 2
    fi
done

```

Tek tek yazılışını inceleyelim.

- **#!/bin/bash** : Bunun bir bash scripti olduğunu belirtmek ve ona göre derlenmesini sağlamak için yazılması gereklidir.

- **test=(-tests-)** : Bu satır path içerisindeki `quest.json`'da tanımlanmış olan testlerin getirilmesi ve ona göre kontrol yapılması için gereklidir.

- *export TERM=xterm* : Terminalin uyumluluğunun ayarlanması için kullanılır.

- `if [${#test[@]} -eq 0 ]; then` : Bunun ile tanımlanmış bir test olup olmadığı kontrol edilir. Bu uzunluğun 0 olması durumunda `result=$(go run ../main.go)` ile dosya yalnızca derlenir ve çıktısı `echo "Test Passed|||$result||| _ ||| _ " exit 0` olarak geri döndürülür. Bu da çıktının her durumda doğru olmasını sağlar. `|||` ile ayırıp 4 bölümden oluşacak şekilde değer döndürülmesinin sebebi ise ilk kısımda geçtiği mesajını basmaktır. İkinci kısmın görevi kullanıcının ekrana bastığı metni döndürmektir. Diğer kısımlar ise burada kullanılmamaktadır; bu sebeple _ ile boş olduğu gösterilmektedir. En son ise exit 0 ile kod bitirilir ve bir sonraki satırların kontrol edilmesi önlenir. İstenen şart sağlanmıştır. fi yazılarak da if sonlandırılır.

- `for i in "${!test[@]}"; do` : Bu döngüye, eğer ilk if şartı sağlanmamışsa, yani test dizisi 0 değilse, girilir. Tek tek tüm testler denenir.

- `$expected_result="${test[$i]"` : Json içerisinde yazan, beklenen test sonucu buraya yazılacaktır.

- `go install golang.org/x/tools/cmd/goimports@latest > /dev/null 2>&1 goimports -w ../main.go > /dev/null 2>&1` : Bu yalnızca Go dili için gerekli bir kütüphanedir. Dilin dosyasının düzenlenmesi için gerekmektedir. Diğer dillerde gerekmemektedir.

- *`$compile_output`=`$(go build -o main ../main.go 2>&1)`* : Derleme sonucu buraya dönecektir. Eğr herhangi bir derleme hatası oluşursa bu kullanıcıya dönmektedir.

- *`$result`=`$(go run ../main.go)`* : Go dosyası çalıştırılır ve dönen sonuç buraya atılır.

- `if [[ "`$result`" == "`$expected_result`" ]]; then` : Beklenen test sonucu ile dönen sonuç karşılaştırılır. Doğru olması durumunda `echo "Test Passed|||$result||| _ ||| _" mesajı geri döndürülür. Hatalı olması durumunda ise else içerisine girilir ve *echo "_ |||$result|||$expected_result||| _" ` mesajı basılır. Burada hata olması sebebiyle beklenen sonuç da basılır. En son `|||` sonrasına yalnızca hata olması durumunda veri yazılır diğer durumlarda _ yazılmalıdır.

-  `if [ $? -ne 0 ]; then echo "_ ||| _ ||| _ |||$compile_output" exit 1 fi` : Bu blok ise her dilde yer almamaktadır. Derleme işleminin başarısız olup olmadığını kontrol eder. Başarısız olması durumunda hatanın dönmesini sağlar.

---

**Oluşturanlar:**

- Yusuf Küçükgökgözoğlu
- Çetin Boran Mesüm
- Melike Sena Çakır

---
