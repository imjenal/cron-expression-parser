package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestIntegration(t *testing.T) {
	cronExprParserPath := "/Users/jyotsna/GolandProjects/cron_expr_parser/cron_expr_parser"

	testCases := []struct {
		cmdInput       string
		expectedOutput string
	}{
		{
			cmdInput: "*/15 0 1,15 * 1-5 /usr/bin/find",
			expectedOutput: `
Expanded Cron Fields:
minute: [0 15 30 45]
hour: [0]
day of month: [1 15]
month: [1 2 3 4 5 6 7 8 9 10 11 12]
day of week: [1 2 3 4 5]
command: /usr/bin/find
`,
		},
		{
			cmdInput: "*/5 0 15 * 5 /usr/bin/find",
			expectedOutput: `
Expanded Cron Fields:
minute: [0 5 10 15 20 25 30 35 40 45 50 55]
hour: [0]
day of month: [15]
month: [1 2 3 4 5 6 7 8 9 10 11 12]
day of week: [5]
command: /usr/bin/find
`,
		},
		{
			cmdInput: "2-5 1 */10 * 5 /usr/bin/find",
			expectedOutput: `
Expanded Cron Fields:
minute: [2 3 4 5]
hour: [1]
day of month: [1 11 21 31]
month: [1 2 3 4 5 6 7 8 9 10 11 12]
day of week: [5]
command: /usr/bin/find
`,
		},
	}

	for _, testCase := range testCases {
		cmd := exec.Command(cronExprParserPath, testCase.cmdInput)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Error running cron_expr_parser: %v", err)
		}
		// Print the combined output for debugging
		fmt.Printf("Command Output:\n%s\n", output)

		actualOutput := strings.TrimSpace(string(output))
		expectedOutput := strings.TrimSpace(testCase.expectedOutput)
		if actualOutput != expectedOutput {
			t.Errorf("Unexpected output for input:\n%s\nGot:\n%s\nExpected:\n%s", testCase.cmdInput, actualOutput, expectedOutput)
		}
	}
}
