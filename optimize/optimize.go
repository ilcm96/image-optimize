package optimize

import (
	"io"
	"net/http"
	"os"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	mimePNG  string = "image/png"
	mimeJPEG string = "image/jpeg"
	mimeGIF  string = "image/gif"
)

func Optimize(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	log.SetHeader(`{"time":"${time_rfc3339}","id":"` + requestId + `","level":"${level}","file":"${short_file}:${line}"}`)

	// Create unique path with request id
	path := "./images/" + requestId + "/"
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Error(err)
	}

	// Remove directory
	defer func() {
		err := os.RemoveAll(path)
		if err != nil {
			log.Error(err)
		}
	}()

	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "no such file")
	}

	src, err := file.Open()
	if err != nil {
		log.Error(err)
		return err
	}
	defer src.Close()

	filePath := path + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		log.Error(err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Error(err)
		return err
	}

	fileByteArr, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		return err
	}

	kind, err := filetype.Match(fileByteArr)
	if err != nil {
		log.Error(err)
		return err
	}

	var optimizedFilePath string
	if kind.MIME.Value == mimePNG {
		err, optimizedFilePath = optimizePng(filePath)
		if err != nil {
			log.Error(err)
			return err
		}
	} else if kind.MIME.Value == mimeJPEG {
		err, optimizedFilePath = optimizeJpeg(filePath)
		if err != nil {
			log.Error(err)
			return err
		}
	} else if kind.MIME.Value == mimeGIF {
		err, optimizedFilePath = optimizeGif(filePath)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	_ = c.Attachment(optimizedFilePath, appendStringOnFilePath(file.Filename, "optimized"))
	return nil
}
