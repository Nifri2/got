package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type FolderStructure struct {
	Meta      bool
	Blob      bool
	Commits   bool
	FTree     bool
	Branches  bool
	Statefile bool
	Stage     bool
}

func InitGotProject(cmd *cobra.Command, args []string) {
	var err Err
	var ex string
	var base string = ".got"

	ex, err.E = os.Executable()
	err.Handle()

	current_path := filepath.Dir(ex)

	log.Info("Initializing got project in", current_path)

	for {
		if !ensureGotProjectPath() {

			log.Info("No .got project path found, creating one!")
			err.E = os.Mkdir(base, 0755)
			err.HandleFatal()

		} else {

			log.Info("Found .got project path!")
			break

		}
	}

	var fs FolderStructure
	fs.ensureStructure()
	log.Debugf("%+v\n", fs)

	fs.setupStructure()

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

func (f *FolderStructure) ensureStructure() {
	// Check if the structure is already present
	// If not, set bool to false
	// If yes, set bool to true

	// Check if the meta folder is present
	_, err := os.Stat(".got/meta")
	if err != nil {
		f.Meta = false
		log.Debug("Meta folder not found")
	} else {
		f.Meta = true
		log.Debug("Meta folder found")
	}

	// Check if the blob folder is present
	_, err = os.Stat(".got/blob")
	if err != nil {
		f.Blob = false
		log.Debug("Blob folder not found")
	} else {
		f.Blob = true
		log.Debug("Blob folder found")
	}

	// Check if the commits folder is present
	_, err = os.Stat(".got/blob/commits")
	if err != nil {
		f.Commits = false
		log.Debug("Commits folder not found")
	} else {
		f.Commits = true
		log.Debug("Commits folder found")
	}

	// Check if the ftree folder is present
	_, err = os.Stat(".got/blob/ftree")
	if err != nil {
		f.FTree = false
		log.Debug("FTree folder not found")
	} else {
		f.FTree = true
		log.Debug("FTree folder found")
	}

	// Check if the branches folder is present
	_, err = os.Stat(".got/blob/branches")
	if err != nil {
		f.Branches = false
		log.Debug("Branches folder not found")
	} else {
		f.Branches = true
		log.Debug("Branches folder found")
	}

	// Check if the statefile is present
	_, err = os.Stat(".got/statefile")
	if err != nil {
		f.Statefile = false
		log.Debug("Statefile not found")
	} else {
		f.Statefile = true
		log.Debug("Statefile found")
	}

	// Check if the stage Folder is present
	_, err = os.Stat(".got/stage")
	if err != nil {
		f.Stage = false
		log.Debug("Stage folder not found")
	} else {
		f.Stage = true
		log.Debug("Stage folder found")
	}

}

// See README.md#Architecture#Init#Folder-Structure

func (f *FolderStructure) setupStructure() {
	var err Err

	var base string = ".got"
	log.Info("Setting up project structure...")

	// Create the metadata folder
	if !f.Meta {
		err.E = os.Mkdir(filepath.Join(base, "meta"), 0755)
		err.HandleWarn()
	}

	// Create the blobs folder
	if !f.Blob {
		err.E = os.Mkdir(filepath.Join(base, "blob"), 0755)
		err.HandleWarn()
	}

	// Create the commits folder
	if !f.Commits {
		err.E = os.Mkdir(filepath.Join(base, "blob", "commits"), 0755)
		err.HandleError()
	}

	// Create the ftree folder
	if !f.FTree {
		err.E = os.Mkdir(filepath.Join(base, "blob", "ftree"), 0755)
		err.HandleError()
	}

	// Create the branches folder
	if !f.Branches {
		err.E = os.Mkdir(filepath.Join(base, "blob", "branches"), 0755)
		err.HandleError()
	}

	// Create the Staging Area
	if !f.Stage {
		err.E = os.Mkdir(filepath.Join(base, "stage"), 0755)
		err.HandleError()
	}

	// Create the statefile
	if !f.Statefile {
		var statefile Statefile
		statefile.CurrectBranch = "main"
		statefile.Branches = []string{"main"}
		statefile.Commits = []Commits{}
		statefile.Write()
	}

	log.Info("Project structure setup complete!")
}
