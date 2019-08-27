package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DefDomainRegisterType struct {
	domain      string
	protocol    string
	accountName string
	secret      string
	result      map[string]string
}

const (
	loadAPIProtocol   = "https://"
	loadAPIDomain     = "api.passivetotal.org"
	loadAPIDomainPath = "/v2/enrichment/subdomains"
)

//GetAPIRequestResponseData : call community.riskiq.com Search api load domain
func GetAPIRequestResponseData(domain string) {
	urls := fmt.Sprintf("%s%s%s", loadAPIProtocol, loadAPIDomain, loadAPIDomainPath)
	var request = []byte(`{"query":"` + domain + `"}`)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urls, bytes.NewBuffer(request))
	req.SetBasicAuth(LoadPerformUploadApiUserName(), LoadPerformUploadApiKey())
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	//fmt.Println(string(body))
	//jq := gojsonq.New().File(string(body)).From("subdomains")
	var dat map[string]interface{}
	json.Unmarshal([]byte(body), &dat)
	if extractDomain, ok := dat["subdomains"]; ok {
		ws := extractDomain.([]interface{})
		for i := 0; i < len(ws); i++ {
			joinDomain := fmt.Sprintf("%s.%s", ws[i], domain)
			fmt.Println(joinDomain)
		}
	}
}
