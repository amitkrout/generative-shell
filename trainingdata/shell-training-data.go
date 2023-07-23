package trainingdata

var (
	TemplateShell = `
Input: %s
Output:
`

	PromptShell = `
Input: List files
Output: ls -a

Input: Count files in a directory
Output: ls -a | wc -l

Input: Disk space used by home directory
Output: du -sh ~

Input: Replace foo with bar in all .go files
Output: find . -name "*.go" | xargs sed -i 's/foo/bar/g'

Input: Delete the test subdirectory and everything underneath it
Output: rm -rf test

Input: Splunk query for all 500 errors
Output: index=<app_index> status=500

Input: Splunk query to find out all successful logs
Output: index=<app_index> <your_text_search_query> status=200
`
)
