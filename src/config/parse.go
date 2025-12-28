package config

import (
	"errors"
	"log"
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
	Values  any `yaml:"values" json:"values" jsonschema:"required"`
	Filters any `yaml:"filters" json"filters"`
}

type Config struct {
	// Modify what resources Tagitha supports
	Resources ResourceConfig `yaml:"resources" json:"resources,omitempty"`
	// The working directories. Can be an array of strings or a string
	Directories any        `validate:"directories-filter,required" yaml"directories" json:"directories" jsonschema:"required"`
	Tags        []TagGroup `yaml:"tags" json:"tags" jsonschema:"required"`
}

func Parse(configPath string) (Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
		return Config{}, errors.New("Failed to read config file")
	}

	var configData Config

	// Validation
	validate := validator.New()
	validate.RegisterValidation("directories-filter", validateDirectories)

	if err := yaml.UnmarshalWithOptions(data, &configData, yaml.Validator(validate)); err != nil {
		return Config{}, err
	}

	return configData, nil
}

func validateDirectories(fl validator.FieldLevel) bool {
	// Some basic validation on the directories to make sure they
	field := fl.Field()

	switch value := field.Interface().(type) {
	case string:
		return true

	case []any:
		// Looping over slice to check each item is a string
		for _, item := range value {
			if _, ok := item.(string); !ok {
				return false
			}
		}
		return true

	default:
		return false
	}
}
