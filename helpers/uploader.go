package helpers

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

func Uploader(file *multipart.FileHeader, err error) (string, bool) {
	var fileType, fileName string
	isSuccess := true
	if err != nil {
		isSuccess = false
	} else {
		src, err := file.Open()
		if err != nil {
			isSuccess = false
		} else {
			fileByte, _ := ioutil.ReadAll(src)
			fileType = http.DetectContentType(fileByte)
			if fileType == "application/pdf" {
				fileName = "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
			} else if fileType == "text/plain; charset=utf-8" {
				fileName = "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"
			} else {
				isSuccess = false
			}

			err = ioutil.WriteFile(fileName, fileByte, 0777)
			if err != nil {
				isSuccess = false
			}

		}
		defer src.Close()
	}
	return fileName, isSuccess
}
