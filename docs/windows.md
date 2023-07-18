**To run generative-shell, follow these steps:**

* Download the latest release from the [Releases](https://github.com/amitkrout/generative-shell/releases) section on GitHub. The release will include the cross-compiled binaries for multiple platforms.
* Download the windows binary `generative-shell-windows-amd64`.

**For Powershell user:**

* Open a powershell terminal and navigate to the location where you downloaded the binary package.
* Export the API key as an environment variable using the following command:
```shell
# Replace <your-api-key> with the actual API key you obtained from the Obtaining an OpenAI API Key section.
$env:OPENAI_API_KEY="<your-api-key>"
# IMPORTANT: Never forget to set the environment variable SHELL=PS,
# otherwise the prompt initialization will panic out.
$env:SHELL="PS"
```
* Run the binary:
```shell
.\generative-shell-windows-amd64
```

**For command prompt user:**

* Open a command prompt terminal and navigate to the location where you downloaded the binary package.
* Export the API key as an environment variable using the following command:
```shell
# Replace <your-api-key> with the actual API key you obtained from the Obtaining an OpenAI API Key section.
set OPENAI_API_KEY=<your-api-key>
# IMPORTANT: Never forget to set the environment variable SHELL=CMD,
# otherwise the prompt initialization will panic out.
set SHELL=CMD
```
* Run the binary:
```shell
.\generative-shell-windows-amd64
```
