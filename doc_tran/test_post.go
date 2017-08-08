package main

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"os"
	"io/ioutil"
	"bytes"
	//"encoding/xml"
)

func jsonPrint() (jsonData []byte){
	mapD := map[string]string{"trash":"1", "A": "xXXXxx", "trash1":"1", "B": "yYYYEEyy", "trash2":"1"}
	mapB, _ := json.Marshal(mapD)
	return mapB
}



func main() {
	apiUrl := "http://192.168.10.107:8000"
	resource := "/J2X"
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	//tmp := jsonPrint()

	rawJSONData := []byte(`{"A": "xxx", "B": "yyy", "C": [{"C1": "some_data1"}, {"C2": "some_data2"}], "person": "Maria"}`)



	resp, err := http.Post(urlStr, http.DetectContentType(rawJSONData) , bytes.NewBuffer(rawJSONData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", urlStr, err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Printf(string(b))
}