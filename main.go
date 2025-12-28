package main

import (
	"context"
	"errors"
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/flamableassassin/tagitha/src/config"
	"github.com/urfave/cli/v3"
)

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
		},
		Action: entryPoint,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func entryPoint(ctx context.Context, cmd *cli.Command) error {
	configPath := cmd.String("configpath")
	_, err := config.Parse(configPath)
	return err
}
