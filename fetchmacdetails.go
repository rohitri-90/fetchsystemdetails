package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

func main() {

	getMacAddressDetails()

}

type ApiResponse struct {
	VendorDetails struct {
		Oui            string `json:"oui"`
		IsPrivate      bool   `json:"isPrivate"`
		CompanyName    string `json:"companyName"`
		CompanyAddress string `json:"companyAddress"`
		CountryCode    string `json:"countryCode"`
	} `json:"vendorDetails"`
}

func getMacAddressDetails() {

	mac_address := os.Getenv("MAC_ADDRESS")
	mac_api_token := os.Getenv("MAC_API_TOKEN")
	if mac_api_token == "" {
		logrus.Info("token is required to access the Api")
	}
	if mac_address == "" {
		logrus.Info("Please pass the mac address")
	}
	request, err := http.NewRequest("GET", "https://api.macaddress.io/v1?output=json&search="+mac_address+"", nil)
	if err != nil {
		logrus.Info("error reading the api")
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Authentication-Token", mac_api_token)
	resp, err := http.DefaultClient.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	apiResponse := ApiResponse{}
	json.Unmarshal(body, &apiResponse)
	logrus.Info("Company Name :", apiResponse.VendorDetails.CompanyName)
	logrus.Info("Company Address :", apiResponse.VendorDetails.CompanyAddress)

}
