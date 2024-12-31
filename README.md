[![Docker](https://github.com/Yavuzlar/CodinLab/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/Yavuzlar/CodinLab/actions/workflows/docker-publish.yml)

<p align="center">
<img src="https://github.com/Yavuzlar/CodinLab/blob/main/design/assets/logo/main-horizontal.png" alt="CodinLab Logo" height="200"> 
</p>
<br>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [What is CodinLab?](#what-is-codinlab)
- [Features](#features)
- [Installation](#installation)
  - [Install with DockerHub](#install-with-dockerhub)
  - [Manuel Installation](#manuel-installation)
- [Usage \& Credentials](#usage--credentials)
- [Roadmaps](#roadmaps)
- [Labs](#labs)
- [Contributing](#contributing)
  - [Content Update Instructions](#content-update-instructions)
    - [How to contribute/add or update content?](#how-to-contributeadd-or-update-content)
  - [Content Contributors](#content-contributors)
- [Supporters](#supporters)

## What is CodinLab?

CodinLab is an **open-source software lab** designed to help users learn programming from scratch. Developed with **Go programming language** as the backend and **Next.js** for the frontend, CodinLab provides comprehensive **roadmaps** and **hands-on labs** that allow users to learn the syntax and structure of various programming languages, along with a focus on specific topics within those languages. 

The project is easily extensible and deployable via **Docker Compose**, making it simple for developers to get started or contribute to the project.

![CodinLab Dashboard Screenshot](https://github.com/Yavuzlar/CodinLab/blob/main/design/assets/screenshots/codinlab.png)

## Features

**Language-Specific Roadmaps**: Roadmaps provide a step-by-step guide to learning languages.

**Algorithmic Labs**: Improve problem-solving skills through language-agnostic challenges that focus on data structures and algorithms.
  
**Progress Tracking**: Keep track of your development journey as you complete labs and projects within each roadmap.

**User Management**: Organize and manage different users with role-based permissions and customizable access control.

**Docker Compose Support**: Quick and straightforward setup with Docker Compose, making deployment a breeze for any user.

## Installation

**Note:** Please make sure your docker engine version is updated to the latest version.

### Install with DockerHub

To run CodinLab locally, follow these steps:

1. Use Docker Run to launch the application:

    ```bash
    docker run -p"80:80" -v "/var/run/docker.sock:/var/run/docker.sock" ghcr.io/yavuzlar/codinlab:latest
    ```

2. Once the application is up, open your browser and go to <a href="http://localhost/" target="_blank">`http://localhost`</a> or <a href="http://localhost/" target="_blank">`http://127.0.0.1`</a> to start using CodinLab.

### Manuel Installation

1. Clone the repo
   ```sh
    git clone https://github.com/Yavuzlar/CodinLab
   ```
2. Build docker image
   ```sh
    docker build -t yavuzlar/codinlab . -f ./docker/prod.Dockerfile
   ```
3. Run container
   ```sh
    docker run -d -p 80:80 yavuzlar/codinlab
   ```
4. Go to <a href="http://localhost/" target="_blank">`http://localhost`</a> or <a href="http://localhost/" target="_blank">`http://127.0.0.1`</a>

## Usage & Credentials

1. Apply installation instructions.
2. Open your browser and go to <a href="http://localhost/" target="_blank">`http://localhost`</a> or <a href="http://localhost/" target="_blank">`http://127.0.0.1`</a>
3. Use credentials to login.

   ```
   Username: admin
   Password: admin1234
   ```

## Roadmaps

CodinLab offers various **roadmaps** that guide users from basic to advanced programming concepts. These roadmaps are tailored to help you understand the following:

- C++
- Python
- Javascript
- Go programming

## Labs

CodinLab provides two types of labs:
- **Language-Specific Labs**: Hands-on labs designed to help users dive deep into specific programming languages with exercises and sample projects.
  
- **Algorithmic Labs**: Labs that focus on language-independent algorithmic challenges and data structures, helping users improve their problem-solving skills.

## Contributing

We welcome contributions from the community! Whether you want to add new content to roadmaps, develop labs or fix bugs, your efforts are invaluable. 

Join our journey to help everyone learn software development from scratch, through open-source collaboration with CodinLab!

### Content Update Instructions
#### How to contribute/add or update content?
[EN] Check content settings documents: [Instructions](https://github.com/Yavuzlar/CodinLab/blob/main/usage%20docs/en.md)

[TR] İçerikleri nasıl düzenleyebileceğinize bakın: [Kullanım Talimatları](https://github.com/Yavuzlar/CodinLab/blob/main/usage%20docs/tr.md)

### Content Contributors
- Hüseyin Tazegül
- Mustafa Batuhan Alun

## Supporters

<a href="https://sibervatan.org/" target="_blank"><img src="https://github.com/Yavuzlar/CodinLab/blob/main/design/assets/logo/sibervatan.png" alt="CodinLab Logo" height="64"></a>
&emsp;<a href="https://yavuzlar.org/" target="_blank"><img src="https://github.com/Yavuzlar/CodinLab/blob/main/design/assets/logo/yavuzlar.png" alt="CodinLab Logo" height="64"></a>