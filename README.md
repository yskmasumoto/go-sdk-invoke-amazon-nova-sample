# Go SDK Invoke Amazon Nova Sample

This repository contains a sample project demonstrating how to use the AWS SDK for Go V2 to invoke an Amazon Nova Pro model. It includes a simple example of building a request, invoking the model, and handling the response.

## Features

- Example usage of Amazon Nova Pro model invocation via Go.
- Minimal setup to quickly integrate and test the Amazon Nova service.
- Simple code structure for extending and integrating additional features.

## Requirements

- [Go](https://golang.org/dl/) 1.24.0 or later
- Valid AWS credentials (configured via environment variables or other methods)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yskmasumoto/go-sdk-invoke-amazon-nova-sample.git
    ```
2. Change to the project directory:
    ```sh
    cd go-sdk-invoke-amazon-nova-sample
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Configure your AWS credentials and any necessary Amazon Nova settings in your environment.
2. Build the project:
    ```sh
    go build -o invoke-nova
    ```
3. Run the executable:
    ```sh
    ./invoke-nova
    ```

During execution, the `InvokeBedrock` function in `invoke.go` is called, which sends a sample prompt ("What is the capital of France?") to the Amazon Nova Pro model. The resulting output is then logged.

## Project Structure

- `main.go` – The entry point of the application that sets up the context and calls `InvokeBedrock`.
- `invoke.go` – Contains the implementation of `InvokeBedrock`, which builds the request, invokes the model, and processes the response.
- `go.mod` – Go module definition and dependency management.
- [README.md](http://_vscodecontentref_/0) – Project documentation.
