**To run generative-shell, follow these steps:**

* Download the latest release from the [Releases](https://github.com/amitkrout/generative-shell/releases) section on GitHub. The release will include the cross-compiled binaries for multiple platforms.
* Download the linux binary `generative-shell-linux-amd64`.
* Open a terminal and navigate to the location where you downloaded the binary package.
* Export the API key as an environment variable using the following command:
```shell
# Replace <your-api-key> with the actual API key you obtained from the Obtaining an OpenAI API Key section.
export OPENAI_API_KEY=<your-api-key>
```
* Run the binary:
```shell
chmod +x generative-shell-linux-amd64
./generative-shell-linux-amd64
```
