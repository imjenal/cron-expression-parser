package main

import (
	"cron-expression-parser/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: cron-expression-parser <cron_expression>")
		os.Exit(1)
	}

	cronExpression := os.Args[1]

	fieldNames, fieldValues, command, err := internal.ParseCronExpression(cronExpression)
	if err != nil {
		fmt.Println("error processing cron expression:", err)
		os.Exit(1)
	}

	internal.PrintCronFields(fieldNames, fieldValues, command)
}
