package cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

func UploadFile(myFile multipart.File, header *multipart.FileHeader) (string, string, error) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", header.Filename)
	if err != nil {
		fmt.Println("Dosya okuma hatası - ", err.Error())
	}

	_, err = io.Copy(fileWriter, myFile)
	if err != nil {
		fmt.Println("Dosya kopyalama hatası - ", err.Error())
	}

	tokenWriter, err := bodyWriter.CreateFormField("token")
	if err != nil {
		fmt.Println("Token hatası - ", err.Error())
	}
	_, err = io.Copy(tokenWriter, strings.NewReader("abc"))
	if err != nil {
		fmt.Println("Token kopyalama hatası - ", err.Error())
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := "https://" + BestServer() + ".gofile.io/uploadFile"
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		fmt.Println("Istek hatası - ", err.Error())
	}

	body2, _ := ioutil.ReadAll(resp.Body)

	var uploadResponse UploadResponse
	json.Unmarshal(body2, &uploadResponse)

	return uploadResponse.Data.DownloadPage, uploadResponse.Data.FileId, nil
}
