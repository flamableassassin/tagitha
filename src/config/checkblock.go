package config

import "github.com/flamableassassin/tagitha/src/terraform"

func ShouldTag(block *terraform.TaggableBlock, configFilters *any) (bool, error) {

	return false, nil
}

func notExpression(block *terraform.TaggableBlock, configFilters *any) (bool, error) {
	passed, err := ShouldTag(block, configFilters)
	return !passed, err
}

// func andExpression(block *terraform.TaggableBlock, configFilters any) (bool, error) {

// }
