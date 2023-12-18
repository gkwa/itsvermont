package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}"
)

var (
	verbose   bool
	logFormat string
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")

	flag.StringVar(&logFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	var logLevel slog.Level
	logLevel = slog.LevelInfo
	if verbose {
		logLevel = slog.LevelDebug
	}

	opts := littlecow.OpinionatedHandlerOptions(logLevel, littlecow.RemoveTimestampAndTruncateSource)

	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stderr, opts)
	if logFormat == "json" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	code := {{ cookiecutter.project_slug }}.Main()
	os.Exit(code)
}
