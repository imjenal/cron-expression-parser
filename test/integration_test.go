package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestIntegration(t *testing.T) {
	cronExprParserPath := "/Users/rajat/workspace/cron-expression-parser/cron-expression-parser"

	testCases := []struct {
		name           string
		cmdInput       string
		expectedOutput string
		expectError    bool
	}{
		{
			name:     "Valid cron expression with multiple fields and command",
			cmdInput: "*/15 0 1,15 * 1-5 /usr/bin/find",
			expectedOutput: `
minute         0 15 30 45
hour           0
day of month   1 15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    1 2 3 4 5
command        /usr/bin/find
`,
			expectError: false,
		},
		{
			name:     "Valid cron expression with single command",
			cmdInput: "*/5 0 15 * 5 /usr/bin/find",
			expectedOutput: `
minute         0 5 10 15 20 25 30 35 40 45 50 55
hour           0
day of month   15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    5
command        /usr/bin/find
`,
			expectError: false,
		},
		{
			name:     "Valid cron expression with ranges and steps",
			cmdInput: "2-5 1 */10 * 5 /usr/bin/find",
			expectedOutput: `
minute         2 3 4 5
hour           1
day of month   1 11 21 31
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    5
command        /usr/bin/find
`,
			expectError: false,
		},
		{
			name:           "Invalid range in minute field",
			cmdInput:       "70-90 0 15 * 5 /usr/bin/find",
			expectedOutput: `error processing cron expression: validation failed for field 'minute': range start '70' is out of bounds. value should be (0-59)`,
			expectError:    true,
		},
		{
			name:           "Invalid step in minute field",
			cmdInput:       "*/100 0 15 * 5 /usr/bin/find",
			expectedOutput: `error processing cron expression: validation failed for field 'minute': step value '100' is out of bounds. value should be (0-59)`,
			expectError:    true,
		},
		{
			name:           "Invalid range in day of week field",
			cmdInput:       "0 0 15 * 10-12 /usr/bin/find",
			expectedOutput: `error processing cron expression: validation failed for field 'day of week': range start '10' is out of bounds. value should be (0-6)`,
			expectError:    true,
		},
		{
			name:           "Valid cron expression with month and day of week names",
			cmdInput:       "0 0 1 Jan 1 /usr/bin/find",
			expectedOutput: `error processing cron expression: validation failed for field 'month': invalid value: Jan`,
			expectError:    true,
		},
		{
			name:     "Valid cron expression with boundary values",
			cmdInput: "59 23 31 12 0 /usr/bin/find",
			expectedOutput: `
minute         59
hour           23
day of month   31
month          12
day of week    0
command        /usr/bin/find
`,
			expectError: false,
		},
		{
			name:           "Invalid cron expression with fewer fields",
			cmdInput:       "*/15 0 1 * /usr/bin/find",
			expectedOutput: `error processing cron expression: invalid cron expression. It should have 6 fields`,
			expectError:    true,
		},
		{
			name:           "Invalid cron expression with extra field",
			cmdInput:       "*/15 0 1,15 * 1-5 /usr/bin/find extra_field",
			expectedOutput: `error processing cron expression: invalid cron expression. It should have 6 fields`,
			expectError:    true,
		},
	}

	for _, testCase := range testCases {
		cmd := exec.Command(cronExprParserPath, testCase.cmdInput)
		output, err := cmd.CombinedOutput()
		actualOutput := strings.TrimSpace(string(output))

		if testCase.expectError {
			if err == nil {
				t.Fatalf("Expected an error but got none for input: %s", testCase.cmdInput)
			}

			// Check the actual output against the expected error output
			expectedOutput := strings.TrimSpace(testCase.expectedOutput)
			if actualOutput != expectedOutput {
				t.Errorf("Unexpected output for input:\n%s\n\nGot:\n%s\n\nExpected:\n%s", testCase.cmdInput, actualOutput, expectedOutput)
			}

			// Skip further checks since an error was expected and correctly handled
			continue
		}

		// If no error is expected but the command failed
		if err != nil {
			t.Fatalf("Unexpected error running cron-expression-parser: %v", err)
		}

		expectedOutput := strings.TrimSpace(testCase.expectedOutput)
		if actualOutput != expectedOutput {
			t.Errorf("Unexpected output for input:\n%s\n\nGot:\n%s\n\nExpected:\n%s", testCase.cmdInput, actualOutput, expectedOutput)
		}
	}
}
