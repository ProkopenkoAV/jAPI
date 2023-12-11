# Jenkins CLI

Jenkins CLI is a command-line utility for interacting with Jenkins, providing functionality for job creation, deletion, and execution.

## Table of Contents

- [Introduction](#jenkins-cli)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Commands](#commands)
  - [`create`](#create-command)
  - [`delete`](#delete-command)
  - [`run`](#run-command)
- [Configuration](#configuration)
- [Contributing](#contributing)

## Features

- Create Jenkins jobs from XML configuration files.
- Delete existing Jenkins jobs.
- Run Jenkins jobs.

## Getting Started

### Installation

To install Jenkins CLI, follow these steps:

1. Clone the repository:

    ```sh
    git https://github.com/ProkopenkoAV/jenkins-cli.git
    ```

2. Build the executable:

    ```sh
    cd jenkins-cli
    go build -o jcli main.go
    ```

3. Move the executable to a directory in your system's PATH.

## Commands

### `create` command

The `create` command is used to create a Jenkins job.

```sh
jcli create -s <Jenkins_URL> -p <Jenkins_Port> -u <Jenkins_User> -t <Jenkins_Token> -j <Job_Name> -f <XML_File_Path>
```

### `delete` command

The `delete` command is used to delete a Jenkins job.

```sh
jcli delete -s <Jenkins_URL> -p <Jenkins_Port> -u <Jenkins_User> -t <Jenkins_Token> -j <Job_Name>
```

### `run` command

The `run` command is used to running a Jenkins job.

```sh
jcli run -s <Jenkins_URL> -p <Jenkins_Port> -u <Jenkins_User> -t <Jenkins_Token> -j <Job_Name>
```

## Configuration

- `-s`: Jenkins url
- `-p`: Jenkins port
- `-u`: Jenkins user
- `-t`: Jenkins token
- `-j`: Job name or pathToFile (Flag supports both command line input and file path.)

These options can be set using command-line flags.

## Contributing

I welcome contributions from the community. If you have any suggestions, feedback, or want to contribute to the project, please feel free to reach out to me via email at your prokopenkoartsiom1@gmail.com. I appreciate your input and look forward to collaborating with you!