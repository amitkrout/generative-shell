package pkg

var (
	TemplatePS = `
Input: %s
Output:
`

	PromptPS = `
Input: List files and directories in the current directory
Output: Get-ChildItem

Input: List files only in the current directory
Output: Get-ChildItem -File

Input: List files and directories in path C:\Path\To\Directory
Output: Get-ChildItem -Path C:\Path\To\Directory

Input: Disk space used by home directory
Output: Get-ChildItem -Path $HOME -Recurse -File | Measure-Object -Property Length -Sum

Input: Replace foo with bar in all .go files in C:\Path\To\Directory
Output: Get-ChildItem -Path "C:\Path\To\Directory" -Filter "*.go" | ForEach-Object { (Get-Content $_.FullName) | ForEach-Object { $_ -replace "foo", "bar" } | Set-Content $_.FullName }

Input: Delete the test subdirectory and everything underneath it
Output: Remove-Item -Path "C:\Path\To\ParentDirectory\test" -Recurse -Force

Input: Splunk query for all 500 errors
Output: index=<app_index> status=500

Input: Splunk query to find out all successful logs
Output: index=<app_index> <your_text_search_query> status=200
`
)
