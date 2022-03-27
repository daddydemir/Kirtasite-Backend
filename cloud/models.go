package cloud

type ServerData struct {
	Server string `json:"server"`
}

type ServerResponse struct {
	Status string     `json:"status"`
	Datan  ServerData `json:"data"`
}

type UploadResponse struct {
	Status string     `json:"status"`
	Data   UploadData `json:"data"`
}

type UploadData struct {
	DownloadPage string `json:"downloadPage"`
	Code         string `json:"code"`
	ParentFolder string `json:"parentFolder"`
	FileId       string `json:"fileId"`
	FileName     string `json:"fileName"`
	Md5          string `json:"md5"`
	DirectLink   string `json:"directLink"`
	Info         string `json:"info"`
}
