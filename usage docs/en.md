# CodinLab Project Documentation

This document is created to help users add content to the **CodinLab** project.

A default language file is included in the project. This file contains the existing languages (Go, JavaScript, Python, C++). If the language of the content you'll add does not exist, you must first add that language. If the language is already present, you can skip this step.

There are two separate categories for content in the project:

**Path**: The **steps of the roadmap** that provide a learning path for those who want to learn programming. It offers instructional and practical exercises in the desired language to help the learner improve in that language.

**Lab**: A **lab environment** designed to allow users to practice more on a specific topic based on the knowledge learned in the Path.

You can find more detailed information under the headings below.

## Add Language

### Directory Structure

The language files are located under the object folder in the project. The `inventory.json` file contains the default languages. To add a new language, include it in this file.

### How to Add a New Language?

First, let's examine the structure of the `inventory.json` file:

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

#### Explanation:

- **`id`**: The ID of the added language, which distinguishes it from others. The ID for the newly added language should be one more than the ID of the last added language.
- **`name`**: The name of the added language in uppercase. It is used as the title on the front end.
- **`dockerImage`**: The **docker image** name and version for the language you want to add. To find the image name, you can check [dockerhub](https://hub.docker.com/). Adding `latest` indicates the latest version of the image. You can also specify a specific version.
- **`labDir`**: This is fixed. You must always write **object/labs**. This is always the same as labs are separated only as templates by language.
- **`pathDir`**: This changes according to the language. It should be **object/paths/`<language_name>`**.
- **`iconPath`**: You need to add the language icon under **object/icons**. You can access the correct icon later by writing **images/`icon name`** on api.
- **`cmd`**: In this field, you need to write the command to compile and run the code in this language. This will allow the code written by the user to run in Docker. Commands start with "sh" followed by "-c".
   - **sh**: Runs the shell (command line).
   - **-c**: Indicates the command will be executed inside the shell.
   After this, the necessary compile & run command is written. For C++, it is **g++ -o main main.cpp && ./main**. It should be adjusted according to the added language.
- **`bashCmd`**: Unlike `cmd`, here you should write the necessary sh command to execute the **main.sh** file. What this file does is explained [here](#writing-the-mainsh-file).
- **`fileExtension`**: You need to write the file extension of the language you're adding. This is necessary for determining the file type that users will write in the language.
- **`monacoEditor`**: The file extension should be written here, as required by the Monaco Editor, to allow code customization according to the language.
- **`languages`**: The last property you need to write in the JSON file is the language-specific title and description in both Turkish and English.

   - **lang**: The language (tr or en)
   - **title**: The title of the added programming language.
   - **description**: The description of the added programming language.

After writing these properties, you can start adding labs or paths for the new language. When you add a new language, make sure at least one path is added to avoid errors.

## Add Path

### 1- Understanding the Roadmap Directory Structure

In the project's `objects/paths` directory, there are folders for supported programming languages. If the folder for the language you want to add is not present, do the following:

1. Add the new [language](#add-language) to the `inventory` file.
2. Then, create a folder with the name of the language in this directory and add new paths for that language.

An example of the file structure is shown below:

```
object/ 
└── paths/ 
	├── c++/ 
	├── go/ 
	├── js/ 
	└── python/
```

Each language folder contains numbered folders for the paths. These folders represent both the order of the paths and their `id`.

An example of the file structure is shown below:

```
c++/ 
├── 1/ 
├── 2/ 
├── 3/ 
└── 4/
```

### 2- Adding a New Path to the Roadmap

1. Go to the folder of the language you want to add (e.g., `c++`).
2. Check the last path folder (e.g., `4`). Create a new folder with the next number (e.g., `5`).
3. Add the following two files to the newly created folder:
	- `quest.json`: Contains task definitions and tests for the path.
	- `template.txt`: Contains the coding environment, instructions for running with Docker, and definitions for the tests.

## Add Lab

### 1- Understanding the Labs Directory Structure

In the project's `objects/labs` directory, each question has its own folder. All labs are contained within this folder. The programming language in which the question will be seen depends on the code templates.

1. Add a new language in the `inventory` file in the necessary format.

An example of the file structure is shown below:

```
object/ 
└── labs/ 
	├── 1
	    └── quest.json
	    └── c++.txt
	    └── js.txt
	├── 2
```

Here, two questions have been added to the labs folder.

### 2- Adding a New Lab

1. Create a new folder in the Labs directory with your desired name and create a 'quest.json' file.
2. Write the 'code templates' for this newly added lab. This will be explained in more detail later.

## Question Structure - Path & Lab

This file contains information, task definitions, and tests about the questions.

**Example Structure:**

```json
{
    "id": 4, // Path or Lab ID (Should match the sequential number for the new path or lab)
    "languages": [
        {
            "lang": "en",
            "title": "Conditional Statements",
            "description": "Understanding conditional statements in C++",
            "content": "Learn how to use if-else statements to make decisions...", // Only for path
            "note": "Detailed instructions about the task in English...",
            "hint":"Hint for solving the lab" // Only for labs
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
        "difficulty": 1, // Difficulty level (1: easy, 3: hard)
        "funcName": "checkEvenOdd", // The name of the function the user will write
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
                "templatePath": "object/paths/c++/4/template.txt" // Path to the template file
            }
        ]
    }
}
```

### Explanation

- **`id`**: The Path ID. This should match the folder name of the Path.
- **`languages`**: Language-specific descriptions for the Path. These are displayed in different languages in the user interface.
    - `lang`: Specifies the language.
    - `title`: The title of the problem.
    - `description`: The description of the problem.
    - `content (paths)`: This is the section where you explain what makes the problem a question. This is only necessary when creating the Path.
    - `note`: This is an instructional section where you explain how to approach solving the problem.
    - `hint (labs)`: This section provides hints for lab exercises. It's used to give small hints for solving the problem, for example: "Learn about recursive functions to solve this problem."
- **`quest`**: Task definitions:
    - `funcName`: The name of the function the user should write. (If the function name is empty or `main`, a **main.sh** file must be created.)
    - `tests`: Sample input/output pairs that will be used to test the user's code.
    - `programmingID`: The ID of the language in the inventory file.
    - `codeTemplates`: The path to the template file used for the task. A template file for each question describes how the problem begins, how it should be checked, and what the output will be. If the template is wrong, the question may not be solved correctly or will give errors. The format for these templates is described below.

These details must be included. You can refer to the examples for proper formatting.

## Code Template Structure - Path & Labs

This file contains the area where the user will write their code, the logic for running it in a Docker environment, and the test mechanism. Since roadmaps contain only one language, a single template is enough for them. For example, if you write a path for C++, one template for C++ will be enough for this question.  
If you are writing a lab, you must write a separate template for each language in which you want to see that lab. Then, if you specify the correct `id` and `path` in the `quest.json` file inside the `codeTemplates` array, the correct programming language will automatically be added to the question. Let's now look at how we can write templates.

### Sections

- **FRONTEND**: The template where the user will write code on the frontend.
- **DOCKER**: The Docker template used to run and compile the code.
- **CHECK**: Test scenarios.

#### Example (for C++)

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
    cout << "Enter value for n: ";
    cin >> n;

    string result = $funcname$(n);
    cout << "Result returned by function: " << result << endl;
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

#### FRONTEND TAG

The section below this tag is reflected in the user's code editor, where they will write the code. The user will start working on this template.

- `$funcname$`: This is replaced with the function name specified in `quest.json`.

#### DOCKER TAG

The section below this tag contains the code that will be executed to run the user's code.

- `$usercode$`: The user's code (only the content of the specified function is taken).
- `$checks$`: Test scenarios. The code in `#CHECKS` is added here based on the test data in `quest.json`.
- `$imps$`: This variable combines the libraries added by the user and any additional libraries required by the backend for Docker to function.
  - For example, if the user uses the `fmt` and `os` libraries in Go, these will appear in the `$imps$` section.
  - The backend may also add other necessary libraries for Docker to work.
- `$success$`: This is replaced with "Test Passed" in the backend, indicating that the user passed all tests.
- `result$res`: Represents the result of the final test in the `checks`.

It is mandatory to return data in the required format to the backend. Otherwise, the backend cannot interpret the response, and the frontend will not work correctly. Details of the required format are provided below.

#### CHECK TAG

The `## CHECK` section is responsible for testing the user's code and evaluating the results.

An example of the section structure is provided below:

## CHECK

```
std::string result$rnd$ = $funcname$($input$);
if (result$rnd$ != $output$) {
    std::cout << "_|||" << result$rnd$ << "|||$out$|||_" << std::endl;
    return 0;
}
```
- `$out$`: Represents the output value of the required test in the `quest.json` file.
- `result$rnd$`: This is used to avoid assigning each function's output to the same variable and creating variables with the same name. It helps in generating random variable names.

##### Working Principle

This section essentially tests the test cases for the necessary path in the `quest.json` file. The function is given the required input, and the result returned from the function is assigned to a variable. Then, the output of the user's code is compared with the expected output. If they are not equal, the necessary format is printed to the screen, and the execution is stopped by returning.

The `#CHECK` section is written into the `$checks` area in the `#DOCKER` section for each test case.

If the user's code returns the expected output with the given input, the backend will return a successful test result in the required format.

#### Output Format

Test results are sent to the backend in a format separated by `|||`. The backend parses these data and sends them to the frontend to display to the user.

**Example Output**  
`Test Passed|||5|||5|||`

- **First Field**: `Test Passed` (Test is successful.) It is enough to write `$success$` in the `template.txt` file.

- **Second Field**: The result returned by the user's code (e.g., `5`). It is enough to write `result$rnd$` in the `template.txt` file.

- **Third Field**: The expected result (e.g., `5`). It is enough to write `$out$` in the `template.txt` file.

- **Fourth Field**: The error message (if any).

### Writing the `main.sh` File

This file needs to be added for paths. It is written so that code can be executed within the paths without specifying the main function or function name.

#### Directory Structure

The `main.sh` file:

```
object/ 
└── paths/ 
	├── c++/ 
	├── go/ 
	├── js/ 
	└── python/
```

is located under these directories. A separate `sh` file must be written for each language under its respective folder.

#### How Should the main.sh File Be Written?

**Example structure for Go language:**

```sh
#!/bin/bash
test=(-tests-) # test dizisi tanımlandı

export TERM=xterm  # TERM değişkeni ayarlandı

# Eğer test dizisi boşsa, bir kere çalıştır. Cevap gerekmeyen öğrenmek için olan bir pathdir.
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

# Detailed Breakdown of `main.sh` File

Here’s an explanation of how each part of the `main.sh` file works:

- **`#!/bin/bash`**
This line is used to specify that the script should be interpreted using `bash`. It ensures that the script is executed correctly by the bash shell.

- **`test=(-tests-)`**
This line is used to retrieve the tests defined in the `quest.json` file within the specified path. It prepares the script to control and execute the tests based on this input.

- **`export TERM=xterm`**
This sets the terminal type to `xterm`, which ensures compatibility and proper functioning of terminal-based tools and outputs.

- **`if [${#test[@]} -eq 0 ]; then`**
This condition checks if there are any defined tests in the `test` array. If the length of the `test` array is 0, it means there are no tests. In this case, the code compiles and runs the file without running any tests.

- **`result=$(go run ../main.go)`**: The Go code is compiled and run.
- **`echo "Test Passed|||$result||| _ ||| _ "`**: This line returns the result in a format with four sections separated by `|||`. The first section is a success message, the second is the result from the user's code, and the last two are empty (represented by underscores). 
- **`exit 0`**: If no tests are defined, the script ends here.

- **`for i in "${!test[@]}"; do`**
If the first `if` condition is not met (i.e., there are tests to run), this loop iterates through all the defined tests.

- **`$expected_result="${test[$i]}"`**
This retrieves the expected result for each test case from the `quest.json` file.
- **`go install golang.org/x/tools/cmd/goimports@latest > /dev/null 2>&1 goimports -w ../main.go > /dev/null 2>&1`**
This command installs and runs the `goimports` tool, which is necessary for Go to format the code. It ensures that the Go file is properly formatted before running it. This step is specific to Go and not required for other languages.

- **`$compile_output=$(go build -o main ../main.go 2>&1)`**
This line compiles the Go code and stores any compilation errors in the `compile_output` variable. If there's an issue during compilation, the error message will be saved here.

- **`$result=$(go run ../main.go)`**
The Go code is executed, and the output of the program is stored in the `result` variable.

- **`if [[ "$result" == "$expected_result" ]]; then`**
This condition compares the actual result of the code with the expected result. If they match, it returns a success message:
- **`echo "Test Passed|||$result||| _ ||| _"`**: This prints the success message, including the result of the test.
If the results don’t match:
- **`else`**
- **`echo "_ |||$result|||$expected_result||| _"`**: This prints the result of the test along with the expected value. The underscores indicate the unused sections of the format.

- **`if [ $? -ne 0 ]; then echo "_ ||| _ ||| _ |||$compile_output" exit 1 fi`**
This block is used to check the exit status of the compilation. If the compilation fails (exit status not equal to 0), it will print the error message and return an error code (exit 1).

---

**Authors**

- Yusuf Küçükgökgözoğlu
- Çetin Boran Mesüm
- Melike Sena Çakır
  
---
