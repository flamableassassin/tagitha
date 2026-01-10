package config

import (
	"errors"
	"os"

	"github.com/go-playground/validator"
	"github.com/goccy/go-yaml"
)

type ResourceConfig struct {
	// Resources to support if not supported already
	Add []string `yaml:"add" json:"add,omitempty"`
	// Resources ignore if they are no longer have a tag block
	Remove []string `yaml:"remove" json:"remove,omitempty"`
}

type TagGroup struct {
	// List of Tags to
	Values  []map[string]string `yaml:"values" json:"values" jsonschema:"required" validate:"required"`
	Filters any                 `yaml:"filters" json"filters"`
}

type Config struct {
	// Modify what resources Tagitha supports
	Resources ResourceConfig `yaml:"resources" json:"resources"`
	// The working directories Tagitha should look in
	Directories []string   `validate:"required" yaml:"directories" json:"directories" jsonschema:"required"`
	Tags        []TagGroup `yaml:"tags" json:"tags" jsonschema:"required" validate:"required"`
}

// TODO: Add validation so that a directory isn't already covered by another item in the slice
func Parse(configPath string) (Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, errors.New("Failed to read config file")
	}

	var configData Config

	// Validation
	validate := validator.New()

	if err := yaml.UnmarshalWithOptions(data, &configData, yaml.Validator(validate)); err != nil {
		return Config{}, err
	}

	return configData, nil
}
