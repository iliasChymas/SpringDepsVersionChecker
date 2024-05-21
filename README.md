## Dependency Checker

## Overview

Dependency Checker helps you ensure that your Maven project dependencies are up-to-date. It compares the versions specified in your `pom.xml` file against the latest recommended versions from an official repository and outputs any discrepancies in a JSON file.
**Note**: Ensure that the pom.xml file is in the same directory from where you run the program.

## Prerequisites

- Go 1.16 or later
- A valid `pom.xml` file in the same directory where the application is run

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/dependency-checker.git
    cd dependency-checker
    ```

2. Build the project:
    ```sh
    go build -o dependency-checker main.go
    ```

## Usage

Run the `dependency-checker` with the following options:

```sh
./dependency-checker -output=output.json -version=5.3.8
 SpringDepsVersionChecker
