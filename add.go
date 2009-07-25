package main

import (
	"flag"
	"os"
	"time"
)

func addCmd(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	engine := fs.String("engine", "", "engine id")
	secs := fs.Int("seconds", 0, "duration")
	thrust := fs.Int("thrust", 0, "thrust, kN")
	result := fs.String("result", "ok", "ok, aborted, RUD")
	fs.Parse(args)
	f, _ := os.OpenFile("log/firings.csv", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	_ = engine; _ = secs; _ = thrust; _ = result
	_ = time.Now()
}
// reject unknown result
// append only, open with O_APPEND
// thrust in kN, not lbf
// usage on missing engine
