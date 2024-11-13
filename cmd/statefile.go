package cmd

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
