package pkg

var (
	Template = `
Input: %s
Output:
`

	Prompt = `
Input: List files
Output: ls -a

Input: Count files in a directory
Output: ls -a | wc -l
`
)
