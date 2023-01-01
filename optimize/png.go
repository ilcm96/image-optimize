package optimize

import (
	"os/exec"

	"github.com/labstack/gommon/log"
)

func optimizePng(filePath string) (error, string) {
	err := runPngquant(filePath)
	if err != nil {
		log.Error(err)
		return err, ""
	}

	pngquantFilePath := appendStringOnFilePath(filePath, "or8")
	err = runOxipng(pngquantFilePath)
	if err != nil {
		log.Error(err)
		return err, ""
	}

	return nil, pngquantFilePath
}

func runPngquant(filePath string) error {
	cmd := exec.Command("./bin/pngquant", "-f", "--speed", "11", filePath)
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func runOxipng(filePath string) error {
	cmd := exec.Command("./bin/oxipng", "-o", "0", "--strip", "all", "-i", "0", "--quiet", filePath)
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
