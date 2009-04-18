package main

import "flag"

func gridCmd(args []string) {
	fs := flag.NewFlagSet("grid", flag.ExitOnError)
	fs.Int("year", 0, "year")
	fs.Parse(args)
}
// grid reads the csv
