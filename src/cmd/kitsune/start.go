package kitsune

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kodmain/kitsune/src/internal/env"
	"github.com/kodmain/kitsune/src/internal/storages/fs"
	"github.com/spf13/cobra"
)

var forceRun bool

func init() {
	startCmd.Flags().BoolVarP(&forceRun, "forground", "f", false, "run service in forground")
}

func createLog(logName string) *os.File {
	file, err := fs.CreateFile(filepath.Join(env.PATH_LOGS, logName))
	if err != nil {
		fmt.Println("Impossible de créer le fichier kitsune.log :", err)
		os.Exit(1)
	}

	return file
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start all kitsune-services",
	Run: func(cmd *cobra.Command, args []string) {

		var serviceSupervisor string = filepath.Join(env.PATH_SERVICES, "supervisor")
		var exec *exec.Cmd = exec.Command(serviceSupervisor)
		var err error = nil

		if forceRun {
			fmt.Println("Starting the kitsune")
			exec.Stdout = os.Stdout
			exec.Stderr = os.Stderr
			err = exec.Run()
		} else {
			fmt.Println("Starting the kitsune as a services")
			exec.Stdout = createLog("kitsune.log")
			exec.Stderr = createLog("errors.log")
			err = exec.Start()
		}

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}
