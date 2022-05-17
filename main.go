package main

import (
	"context"
	"eval/evaluate"

	"fmt"
	"os"
	"time"
)

func checkIfError(err error, where string) {
	if err == nil {
		if where != "" {
			fmt.Printf("Success: %s\n\n", where)
		}
		return
	}

	fmt.Printf("Error at %s:\n\x1b[31;1m%s\x1b[0m\n", where, fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func timerule(expr string) {
	startTime := time.Now()
	rule, err := evaluate.Compile(context.Background(), expr)
	compDur := time.Since(startTime).Seconds()
	checkIfError(err, "compile")

	startTime = time.Now()
	_ = rule()
	runDur := time.Since(startTime).Nanoseconds()
	checkIfError(err, "run")

	fmt.Printf("Took %f seconds to compile, %d nanoseconds to run.\n", compDur, runDur)
}

func main() {
	timerule(`1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2 OR
	1 EQ 2
	`)
}
