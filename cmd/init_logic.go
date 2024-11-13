package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func InitGotProject(cmd *cobra.Command, args []string) {

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	current_path := filepath.Dir(ex)

	log.Info("Initializing got project in", current_path)

	for {
		if !ensureGotProjectPath() {

			log.Info("No .got project path found, creating one!")
			os.Mkdir(".got", 0755)

		} else {

			log.Info("Found .got project path!")
			break

		}
	}
}

func ensureGotProjectPath() bool {
	var stat os.FileInfo
	var err error

	stat, err = os.Stat(".got")
	if err != nil {
		return false
	}

	log.Debugf("%+v\n", stat)
	return true
}

// See README.md#Architecture#Init#Folder-Structure

func setupStructure() {
	var base string = ".got"
	log.Info("Setting up project structure")

	// Create the metadata folder
	os.Mkdir(filepath.Join(base, "meta"), 0755)

	// Create the blobs folder
	os.Mkdir(filepath.Join(base, "blob", "commits"), 0755)
	os.Mkdir(filepath.Join(base, "blob", "ftree"), 0755)

}
