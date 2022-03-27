package cloud

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BestServer() string {
	resp, _ := http.Get("https://api.gofile.io/getServer")

	body, _ := ioutil.ReadAll(resp.Body)

	//bodyString := string(body)
	//fmt.Println(bodyString)

	var serverResponse ServerResponse
	json.Unmarshal(body, &serverResponse)
	return serverResponse.Datan.Server
}
