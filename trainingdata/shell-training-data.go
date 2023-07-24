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
`
)
