package optimize

import "strings"

// ex) /image-optimize/image.png to /image-optmize/image-optimized.png
func appendStringOnFilePath(filePath, str string) string {
	names := strings.Split(filePath, ".")
	basename := ""
	extension := "." + names[len(names)-1]
	for _, name := range names[:len(names)-1] {
		basename += name + "."
	}
	appendedName := basename[:len(basename)-1] + "-" + str

	return appendedName + extension
}
