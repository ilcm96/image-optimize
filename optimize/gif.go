package optimize

import (
	"os/exec"

	"github.com/labstack/gommon/log"
)

func optimizeGif(filePath string) (error, string) {
	err := runGifsicle(filePath)
	if err != nil {
		log.Error(err)
		return err, ""
	}

	return nil, appendStringOnFilePath(filePath, "optimized")
}

func runGifsicle(filePath string) error {
	cmd := exec.Command("./bin/gifsicle", "-O3", "--lossy=80", "-o", appendStringOnFilePath(filePath, "optimized"), filePath)
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
