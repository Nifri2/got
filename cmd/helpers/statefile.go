package helpers

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Statefile struct {
	Name          string    `yaml:"name"`
	CurrectBranch string    `yaml:"currectBranch"`
	Branches      []string  `yaml:"branches"`
	Commits       []Commits `yaml:"commits"`
	Error         *Err
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

func (s *Statefile) Read() {
	var err Err
	var yamlFile []byte

	// Read the statefile
	file := ".got/statefile"
	yamlFile, err.E = os.ReadFile(file)
	err.Handle()

	err.E = yaml.Unmarshal(yamlFile, s)
	err.Handle()
}
