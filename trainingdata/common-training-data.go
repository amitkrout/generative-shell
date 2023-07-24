package trainingdata

var (
	PromptCommon = `
Input: Splunk query for all 500 errors
Output: index=<app_index> status=500

Input: Splunk query to find out all successful logs
Output: index=<app_index> <your_text_search_query> status=200
`
)
