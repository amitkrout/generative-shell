# generative-shell

## Introduction

This project is an interactive shell prompt developed using Golang (Go libarary for OpenAI) and the OpenAI endpoint. It takes prompts in a single simple English sentence and suggests shell commands based on the input.

## Development Environment

* Golang 1.20 or higher
* OpenAI API Key

#### Obtaining an OpenAI API Key

To use the OpenAI API, you will need an API key. Here's how you can obtain it:

* Visit the [OpenAI website](https://openai.com/) and sign in to your account (or create a new one if you don't have an account yet).
* Navigate to the API section or visit the [API settings page](https://platform.openai.com/account/api-keys).
* Follow the instructions provided by OpenAI to generate your API key.
* Once you have generated your API key, make sure to securely store it as it will be used to authenticate your requests to the OpenAI API.

## How to run generative-shell

* [Darwin](docs/darwin.md)
* [Linux](docs/linux.md)
* [Windows](docs/windows.md)

## How it works - Video reference

[bash terminal](https://github.com/amitkrout/generative-shell/blob/main/tutorials/bash.mp4)

## Training Data

The model used in this project has been trained on a specific dataset to provide shell command suggestions. If you would like to customize the training data or improve the model's performance, you can modify the `*-training-template.go` file located in the `trainingdata` directory in the root folder. This file contains example of input prompts and their corresponding expected output in template.

## Contributing

If you find any issue or have suggestions for improvements, we welcome contributions from community, Please follow the guidelines outlined in the [CONTRIBUTION](docs/CONTRIBUTING.md) file to submit pull requests or report issues.

## Acknowledgments

We would like to express our gratitude to OpenAI for providing the powerful GPT language model and the resources necessary to develop this project. Their commitment to advancing AI technology has made this project possible.
