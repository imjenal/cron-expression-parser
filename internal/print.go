package internal

import (
	"fmt"
	"strings"
)

func PrintCronFields(fieldNames []string, fieldValues [][]int, command []string) {
	for i, fieldName := range fieldNames {
		times := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(fieldValues[i])), " "), "[]")
		printField(fieldName, times)
	}
	// Print the command with no extra spacing
	commandStr := strings.Join(command, " ")
	printField("command", commandStr)
}

// printField is a helper function that prints a field name and its associated value(s).
func printField(fieldName string, fieldValue string) {
	fmt.Printf("%-14s %s\n", fieldName, fieldValue)
}
