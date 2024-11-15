package stage

import (
	helpers "github.com/nifri2/got/cmd/helpers"
	"github.com/spf13/cobra"
)

func Stage(cmd *cobra.Command, args []string) {

	// var err helpers.Err
	var sf helpers.Statefile

	sf.Read()

}
