# generative-shell

## Introduction

This project is an interactive shell prompt developed using Golang (Go libarary for OpenAI) and the OpenAI endpoint. It takes prompts in a single simple English sentence and suggests shell commands based on the input.

## Development Environment

* Golang 1.20 or higher
* OpenAI API Key

## Training Data

The model used in this project has been trained on a specific dataset to provide shell command suggestions. If you would like to customize the training data or improve the model's performance, you can modify the `training-template.go` file located in the `pkg` directory. This file contains example of input prompts and their corresponding expected output in template.

## Contributing

If you find any issue or have suggestions for improvements, we welcome contributions from community, Please follow the guidelines outlined in the [CONTRIBUTION](docs/CONTRIBUTING.md) file to submit pull requests or report issues.

## Acknowledgments

We would like to express our gratitude to OpenAI for providing the powerful GPT language model and the resources necessary to develop this project. Their commitment to advancing AI technology has made this project possible.
