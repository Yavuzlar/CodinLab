
<img src="https://github.com/Yavuzlar/CodinLab/blob/main/design/assets/logo/main-horizontal.png" alt="CodinLab Logo" >

## Table of Contents

- [What is CODINLAB?](#what-is-codinlab)
- [Features](#features)
- [Installation](#installation)
- [Roadmaps](#roadmaps)
- [Labs](#labs)
- [Contributing](#contributing)
- [How to Make your Labs Or Path?](#how-to-make-your-labs-or-path)


## What is CODINLAB?

CODINLAB is an **open-source software lab** designed to help users learn programming from scratch. Developed with **Golang** as the backend and **React.js** for the frontend, CODINLAB provides comprehensive **roadmaps** and **hands-on labs** that allow users to grasp the syntax and structure of various programming languages, along with a focus on specific topics within those languages. 

The project is easily extensible and deployable via **Docker Compose**, making it simple for developers to get started or contribute to the project.

![CODINLAB Dashboard Screenshot](https://github.com/Yavuzlar/CodinLab/blob/64-preparing-of-readmemd-file-of-the-project/design/assets/screenshots/codinlab.png)

## Features

- **Language-Specific Roadmaps**: Roadmaps provide a step-by-step guide to learning languages like:
  - Python
  - C++
  - Golang
  - Javascript
  
  Each roadmap offers focused learning on different areas such as:
    - Basic Syntax and Concepts
    - Object-Oriented Programming (OOP)
    - Functional Programming
    - Backend and Frontend Development

- **Algorithmic Labs**: Improve problem-solving skills through language-agnostic challenges that focus on data structures and algorithms.
  
- **Sample Projects**: Apply your skills by building real-world projects included in the roadmaps.
  
- **Progress Tracking**: Keep track of your development journey as you complete labs and projects within each roadmap.

- **User Management**: Organize and manage different users with role-based permissions and customizable access control.

- **Extensible Architecture**: Easily extend CODINLAB with new roadmaps, languages, or additional labs using the plugin architecture.

- **Docker Compose Support**: Quick and straightforward setup with Docker Compose, making deployment a breeze for any user.

## Installation

To run CODINLAB locally, follow these steps:

1. Use Docker Run to launch the application:

    ```bash
    docker run -p"80:80" -v "/var/run/docker.sock:/var/run/docker.sock" ghcr.io/yavuzlar/codinlab:latest
    ```

2. Once the application is up, open your browser and go to `http://localhost:80` to start using CODINLAB.

## Roadmaps

CODINLAB offers various **roadmaps** that guide users from basic to advanced programming concepts. These roadmaps are tailored to help you understand the following:

- **Programming Language Syntax**: Grasp the fundamental syntax and structures of different languages.
- **Specific Language Areas**: Focus on key areas such as backend development, functional programming, and object-oriented design.
- **Data Structures and Algorithms**: Master essential problem-solving techniques used in the real world for any software development role.

## Labs

CODINLAB provides two types of labs:
- **Language-Specific Labs**: Hands-on labs designed to help users dive deep into specific programming languages with exercises and sample projects.
  
- **Algorithmic Labs**: Labs that focus on language-independent algorithmic challenges and data structures, helping users improve their problem-solving skills.

## Contributing

We welcome contributions from the community! Whether you want to add new roadmaps, develop labs, or fix bugs, your efforts are invaluable. 
Join our journey to help everyone learn software development from scratch and beyond, through open-source collaboration with CODINLAB!
<br></br>
## How to Make your Labs Or Path?

### Inventory JSON Explanation
=======
[![Docker](https://github.com/Yavuzlar/CodinLab/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/Yavuzlar/CodinLab/actions/workflows/docker-publish.yml)
# CodinLab
CodinLab software development laboratory

## Installation
```bash
    docker run -p"80:80" -v "/var/run/docker.sock:/var/run/docker.sock" ghcr.io/yavuzlar/codinlab:latest
```

## Inventory JSON Explanation


This inventory JSON contains all the programming languages supported by CodinLab. You can add as many languages as you want to the array by filling in the required fields in the desired format.

- *id*  
    - This is a unique identifier for the programming language, represented as an integer.
- *name*  
    - The name of the programming language.
- *dockerImage*  
    - The Docker image used for this programming language. You can look at these docker images from docker hub.
- *labDir*  
    - The directory path where lab files for this language are stored.
- *pathDir*  
    - The directory path where the path files for this language are stored.
- *iconPath*  
    - The file path for the icon representing this programming language.
- *cmd*  
    - An array of command-line arguments to compile and run the main program. In this case, it includes the command to compile `main.cpp` and execute the resulting binary.
- *bashCmd*  
    - An array of command-line arguments to execute a Bash script. This command that will run the main.sh script that will check the questions in the roads.
- *fileExtension*  
    - The file extension associated with this programming language.
- *languages*  
    - This is an array that provides descriptions of the programming language in different languages. Each object within this array contains the following properties:
        - *lang*: Holds the value `tr` or `en`, indicating the language.
        - *title*: The name of the programming language.
        - *description*: A brief explanation of the programming language.

A JSON example of the C++ programming language is provided below.

```json
 {
        "id": 1,
        "name": "C++",
        "dockerImage": "gcc:latest",
        "labDir": "object/labs",
        "pathDir": "object/paths/c++",
        "iconPath": "images/c++.png",
        "cmd": ["sh", "-c", "g++ -o main main.cpp && ./main"],
        "bashCmd":  ["bash", "-c", "./main.sh"],
        "fileExtension": "cpp",
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
```


### Quest JSON Explanation


This JSON represents a path question or a lab question in a roadmap on codinlab; now we will explain it.

- *ID*  
    - Each question must have a unique ID value, which is of type `int`.  
    - It is essential that this ID does not match any ID from other JSON objects.
- *languages*  
    - Each object within this array contains the following properties:
        - *lang*: Holds the value `tr` or `en`, indicating the language.
        - *title*: The name of the question or path.
        - *description*: A brief explanation of the question or path.
        - *content*: The main content of the question.
        - *note*: A short piece of information related to the question.
- *quest*  
    - *difficulty*: An integer indicating the difficulty level of the quest (1 being the easiest).
    - *funcName*: A string representing the name of the function to be implemented.
    - *tests*: An array of test cases associated with the quest. Each test case includes:
        - *input*: An array representing the input for the test case.
        - *output*: An array containing the expected output for the test case.
    - *codeTemplates*: An array of code template objects that provide initial code structure. Each object includes:
        - *programmingID*: An integer representing the programming language ID.
        - *templatePath*: A string indicating the path to the code template file.

codeTemplates array contains the IDs of the programming languages you created in the inventory JSON, along with the corresponding file path of the code template for this question. The code template will be explained in more detail in the following section.

Below is an example of how to create a question JSON. This question supports two programming languages, as can be seen in the code templates section.

```json
{
	"id": 1,
	"languages": [
		{
			"lang": "en",
			"title": "Basic Syntax",
			"description": "Learn the basics of C++.",
			"content": "Understand the structure of a C++ program.",
			"note": "1. Ensure the helloworld function prints 'Hello, World!'.\n2. Include header files:\n\t**#include <iostream>**: Required for 'std::cout'.\n3. Use: `std::cout << \"Hello, World!\" << std::endl;`"
		},
		{
			"lang": "tr",
			"title": "Basit yazılış",
			"description": "C++'ın temellerini öğrenin.",
			"content": "C++ program yapısını anlayın.",
			"note": "1. helloworld fonksiyonunun 'Hello, World!' yazdırdığından emin olun.\n2. Header dosyalarını ekleyin:\n\t**#include <iostream>**: 'std::cout' için gereklidir.\n3. Kullanım: `std::cout << \"Hello, World!\" << std::endl;`"
		}
	],
	"quest": {
		"difficulty": 1,
		"funcName": "helloworld",
		"tests": [
			{
				"input": [],
				"output": ["Hello, World!"]
			}
		],
		"codeTemplates": [
			{
				"programmingID": 2,
				"templatePath": "object/labs/1/go.txt"
			},
            {
				"programmingID": 1,
				"templatePath": "object/labs/1/c++.txt"
			}
		]
	}
}
```


### Code Template Explanation

This file describes the template used in Codinlab for generating questions and testing solutions.

#### Constants


1. **$funcname$**: Holds the name of the function for the given problem.
2. **$imps$**: Contains the import statements from the user's submitted code. These imports can be customized depending on the Frontend or Docker template.
3. **$usercode$**: Stores the code coming from the frontend. If a `main` function exists, it is removed in this section.
4. **$rnd$**: Represents the test number. It is used to avoid variable name collisions when creating multiple tests.
5. **$input$**: Specifies the input data for the test cases. If the input is an array (`[]`), $input$ separates the elements by commas (`,`). In such cases, the array needs to be stored in a separate variable.
6. **$output$ / $out$**: Represents the expected output for the test cases.
7. **$checks$**: Contains the logic for checking a single test. This is later used in the Docker template, where it applies the test logic to all test inputs and outputs.
8. **$success$**: This constant is used to indicate a successful test. It should be printed at the end if all tests pass.


#### Sections


When writing templates, each section should begin with a header formatted as follows:

- `## DOCKER`: This section contains the Docker template, where necessary import statements and user code are specified for testing within a Docker environment.

- `## FRONTEND`: This section is for the Frontend template, which includes the relevant import statements and the function that will be tested based on user input.

- `## CHECK`: This section defines the logic for verifying the correctness of the tests. It includes the initialization of test inputs and the comparison of the function's output against the expected result. This ensures that all tests are properly validated.


##### FRONTEND



Defines the template to be sent to the frontend. In this section:

- The relevant `import` statements should be written.
- The function specified by $funcname$ should be implemented to return the correct output as described in the problem.


An example of a Frontend template for user code submission:

```cpp
#include <stdio.h>
#include <string>
#include <iostream>

using namespace std;

bool $funcname$(std::string s) {
    // Write your code here

    return false;
}

int main() {
    std::string input;

    std::cout << "Please enter a string: ";
    std::getline(std::cin, input); 

    std::cout << std::boolalpha << $funcname$(input) << std::endl;

    return 0;
}
```


##### DOCKER


Defines the template for Docker-based testing. In this section:

- $imps$ contains the import statements from the user's submitted code.
- Place the $usercode$ provided by the user.
- The test checks are performed using the logic in $checks$.
- The success message ($success$) is printed if all tests pass.

An example of a Docker Engine template used for testing:

```cpp
$imps$

using namespace std;

$usercode$

int main() {
    $checks$

    cout << "$success$" << endl;
}
```

##### CHECK



- The function's result is then should be compared against the expected output, which is stored in $output$. If the result does not match $output$, the test will fail.
- The test inputs are initialized using $input$. If the input is an array, each element will be joined with a comma (`,`), meaning all array elements will appear side by side.
- The function is tested by comparing the result against $output$.
- If the result does not match $output$, the test will fail and print an error message.

The following example shows how a test is structured when the input is an array, and the function takes a vector as a parameter:

```cpp
std::vector<string> inputArr$rnd$ = {$input$};

std::string result$rnd$ = $funcname$(inputArr$rnd$);

if (result$rnd$ != $output$) {
    std::cout << "Test Failed: Expected $out$, but got " << result$rnd$ << std::endl;
    exit($rnd$);
}
```

