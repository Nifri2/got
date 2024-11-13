package cmd

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Statefile struct {
	CurrectBranch string    `yaml:"currectBranch"`
	Branches      []string  `yaml:"branches"`
	Commits       []Commits `yaml:"commits"`
}
type Commits struct {
	Sha     int    `yaml:"sha"`
	Branch  string `yaml:"branch"`
	Message string `yaml:"message"`
}

func (s *Statefile) Write() {
	var err Err
	var yamlFile []byte

	// Write the statefile
	file := ".got/statefile"
	yamlFile, err.E = yaml.Marshal(s)
	err.Handle()

	err.E = os.WriteFile(file, yamlFile, 0644)
	err.Handle()
}
