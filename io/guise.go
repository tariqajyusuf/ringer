package io

type Guise struct {
	name        string `yaml:"name"`
	description string `yaml:"description"`
	platforms   []struct {
		platform    string `yaml:"platform"`
		packageName string `yaml:"packageName"`
	} `yaml:"platforms,flow"`
}
