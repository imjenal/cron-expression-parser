package main

import (
	"cron_expr_parser/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron_expr_parser <cron_expression>")
		os.Exit(1)
	}

	cronExpression := os.Args[1]

	fieldNames, fieldValues, command, err := internal.ParseCronExpression(cronExpression)
	if err != nil {
		fmt.Println("Error processing cron expression:", err)
		os.Exit(1)
	}

	internal.PrintCronFields(fieldNames, fieldValues, command)
}
