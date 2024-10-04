
# Kubernetes Config Checker

## Overview

[![BuildStatus](https://github.com/NineKama/Kubernetes-Config-Checker/actions/workflows/go.yml/badge.svg)](https://github.com/NineKama/Kubernetes-Config-Checker/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/NineKama/Kubernetes-Config-Checker/graph/badge.svg?token=HZ3N93I48S)](https://codecov.io/github/NineKama/Kubernetes-Config-Checker)

Kubernetes Config Checker is a tool designed to validate Kubernetes YAML files against best practices. This CLI tool checks for missing resource requests and limits, security context configurations, namespace usage, and much more. It helps ensure that your Kubernetes configurations are secure and optimized for production use.

## Features

- **Resource Validation**: Ensures that all containers define CPU and memory requests and limits.
- **Security Context Validation**: Verifies that pods and containers have proper security contexts.
- **Namespace Validation**: Encourages usage of namespaces other than the default.
- **ServiceAccount Validation**: Ensures that all pods are configured with a service account.
- **ConfigMap/Secret Validation**: Verifies that ConfigMaps and Secrets are immutable.

## Requirements

- [Go](https://golang.org/doc/install) (version 1.19 or later)

## Installation

### Clone the Repository

```bash
git clone <repository-url>
cd KubernetesConfigCheck
```

### Build the Application

```bash
go build -o k8s-checker
```

## Usage

Once the tool is built, you can use it to validate your Kubernetes YAML files.

### Command Syntax

```bash
./k8s-checker validate --file=<path-to-yaml-file>
```

### Example

To validate a file named `example.yaml` in the current directory:

```bash
./k8s-checker validate --file=example.yaml
```

### Combining Build and Validate

To build and run the tool in a single line:

```bash
go build -o k8s-checker && ./k8s-checker validate --file=example.yaml
```

## Available Commands

- **validate**: This is the primary command to validate Kubernetes YAML files. It checks for best practices and logs any issues it finds.

### Flags

- `--file`: Specifies the path to the Kubernetes YAML file that you want to validate.

## Example Output

```
Validating example.yaml...
Container nginx should define CPU/Memory limits and requests.
Pod nginx-pod should use a service account.
ConfigMap/Secret nginx-config should be immutable.
```

## Best Practices Validated

1. **Resource Limits and Requests**: Ensures that all containers have defined resource requests and limits.
2. **Security Context**: Verifies that security contexts are configured for both pods and containers.
3. **Privileged Mode**: Ensures that privileged containers are disabled.
4. **ReadOnlyRootFilesystem**: Ensures that containers have a read-only root filesystem.
5. **ServiceAccount Usage**: Checks that each pod has a specified service account.
6. **Namespace Usage**: Ensures pods are not running in the default namespace.
7. **ConfigMap/Secret Immutability**: Ensures ConfigMaps and Secrets are immutable.

## Troubleshooting

### Error: "The system cannot find the file specified"
Make sure that the file specified with the `--file` flag exists in the correct location. If the file is in a different directory, provide the full path to the file.

### Error: "validate: The system cannot find the file specified"
Ensure that you are correctly using the `validate` subcommand and specifying the file path with the `--file` flag.

## Contributing

We welcome contributions! Please feel free to submit a pull request or open an issue if you encounter any bugs or have suggestions for improvements.

