package optimize

import (
	"os"
	"os/exec"

	"github.com/labstack/gommon/log"
)

func optimizeJpeg(filePath string) (error, string) {
	err := runMozJPEG(filePath)
	if err != nil {
		log.Error(err)
		return err, ""
	}

	return nil, appendStringOnFilePath(filePath, "optimized")
}

func runMozJPEG(filePath string) error {
	cmd := exec.Command("./bin/cjpeg-static", "-quality", "80", filePath)

	optimizedFile, err := os.Create(appendStringOnFilePath(filePath, "optimized"))
	if err != nil {
		log.Error(err)
		return err
	}
	defer optimizedFile.Close()
	cmd.Stdout = optimizedFile

	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
