// takes in code as arg from go
//run go build on code given

package utils

import (
	"log"
	"os"

	cp "github.com/otiai10/copy"
)

//CopyDir creates exe from file passed in as arg
func CopyDir() {

	//copies contents of .hidden to workspace
	hiddenDir := os.Getenv("BUILDER_HIDDEN_DIR")
	workspaceDir := os.Getenv("BUILDER_WORKSPACE_DIR")
	//exec.Command("cp", "-a", hiddenDir+"/.", workspaceDir).Run()
	err := cp.Copy(hiddenDir, workspaceDir)
	if err != nil {
		//log.Println("Error in copy hidden directory.")
		log.Println(err)

	}
}
