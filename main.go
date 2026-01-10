package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/mail"
	"os"
	"slices"
	"strings"

	"github.com/flamableassassin/tagitha/src/config"
	"github.com/flamableassassin/tagitha/src/terraform"
	"github.com/urfave/cli/v3"
)

var log = slog.With(slog.Group("main`"))

func main() {
	cmd := &cli.Command{
		Name:        "Tagitha",
		Description: "Auto tag Terraform resources",
		Suggest:     true,
		Authors: []any{
			mail.Address{Name: "FlammableAssassin", Address: "lighter@highlyflammable.tech"},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "configpath",
				Usage:     "Path to the config file. Should have a yaml/yml extension",
				TakesFile: true,
				Required:  true,
				OnlyOnce:  true,
				Validator: func(s string) error {
					// Sanity check to see if it is a yaml file. Doing lowercase in case someone put caps in extension
					lowerPath := strings.ToLower(s)
					if !strings.HasSuffix(lowerPath, ".yaml") && !strings.HasSuffix(lowerPath, ".yml") {
						return errors.New("Config file doesn't have a .yaml/.yml file extension")
					}

					if _, err := os.Stat(s); err != nil {
						return err
					}

					return nil
				},
			},
			&cli.StringFlag{
				Name:        "loglevel",
				Usage:       "Modify the level of logging Tagitha uses. (DEBUG,INFO,WARN,ERROR)",
				DefaultText: slog.LevelWarn.String(),
				Category:    "Logging",
				OnlyOnce:    true,
				Required:    false,
				Validator: func(level string) error {
					level = strings.ToUpper(level)
					if !slices.Contains([]string{"DEBUG", "INFO", "WARN", "ERROR"}, level) {
						return fmt.Errorf("Invalid log level %q. Available levels", level)
					}
					return nil
				},
			},
		},
		Action: entryPoint,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func entryPoint(ctx context.Context, cmd *cli.Command) error {
	// Setting logging
	loggingLevel := cmd.String("loglevel")
	createLogger(loggingLevel, "")

	configPath := cmd.String("configpath")

	actualConfig, err := config.Parse(configPath)
	if err != nil {
		return err
	}

	var terraformBlocks []terraform.TaggableBlock
	for _, dir := range actualConfig.Directories {
		tempBlocks, err := terraform.ParseDirectory(dir)
		if err != nil {
			return err
		}
		terraformBlocks = append(terraformBlocks, tempBlocks...)
	}

	return nil
}

// TODO Add format to cli flag
func createLogger(loggingLevel string, format string) {
	loggingOptions := &slog.HandlerOptions{}

	switch strings.ToUpper(loggingLevel) {
	case "DEBUG":
		loggingOptions.Level = slog.LevelDebug
		break
	case "INFO":
		loggingOptions.Level = slog.LevelInfo
		break
	case "ERROR":
		loggingOptions.Level = slog.LevelError
		break
	default:
		loggingOptions.Level = slog.LevelWarn
		break
	}

	var loggingHandler slog.Handler

	switch strings.ToUpper(format) {
	case "JSON":
		loggingHandler = slog.NewJSONHandler(os.Stdout, loggingOptions)
		break
	default:
		loggingHandler = slog.NewTextHandler(os.Stdout, loggingOptions)
	}

	logger := slog.New(loggingHandler)
	slog.SetDefault(logger)

}
