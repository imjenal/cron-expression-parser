package internal

import (
	"fmt"
	"strings"
)

// PrintCronFields prints expanded cron fields
func PrintCronFields(fieldNames []string, fieldValues [][]int, command []string) {
	fmt.Println("Expanded Cron Fields:")
	for i, fieldName := range fieldNames {
		if fieldName != "command" {
			fmt.Printf("%-20s  %v\n", fieldName, fieldValues[i])
		}
	}
	commandStr := strings.Join(command, " ")
	fmt.Printf("command %34s\n", commandStr)
}
