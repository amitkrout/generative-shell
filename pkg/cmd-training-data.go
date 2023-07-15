package pkg

var (
	TemplateCMD = `
Input: %s
Output:
`

	PromptCMD = `
Input: List files and directories in the current directory
Output: dir

Input: List files only in the current directory
Output: dir /A:-D

Input: List files and directories in path C:\Path\To\Directory
Output: dir C:\Path\To\Directory

Input: Disk space used by home directory
Output: dir /s /a /-c "C:\Path\To\Home\Directory"

Input: Replace foo with bar in all .go files in C:\Path\To\Directory
Output: findstr /M /S /I /C:".go" "C:\Path\To\Directory\*.go" | powershell -Command "& {$files = Get-Content -Path stdin; foreach ($file in $files) { (Get-Content $file) -replace 'foo', 'bar' | Set-Content $file }}"

Input: Delete the test subdirectory and everything underneath it
Output: rmdir /s /q "C:\Path\To\Directory\test"

Input: Splunk query for all 500 errors
Output: index=<app_index> status=500

Input: Splunk query to find out all successful logs
Output: index=<app_index> <your_text_search_query> status=200
`
)
