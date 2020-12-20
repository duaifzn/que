package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Data struct {
	Ats AtsContent `json:"Ats"`
}
type AtsContent struct {
	Results ResultsContent `json:"Results"`
}
type ResultsContent struct {
	Result ResultContent `json:"Result"`
}
type ResultContent struct {
	Alexa AlexaContent `json:"Alexa"`
}
type AlexaContent struct {
	TopSites TopSitesContent `json:"TopSites"`
}
type TopSitesContent struct {
	Country CountryContent `json:"Country"`
}
type CountryContent struct {
	Sites SitesContent `json:"Sites"`
}
type SitesContent struct {
	Site []SiteContent `json:"Site"`
}
type SiteContent struct {
	DataUrl string `json:"DataUrl"`
}

func callAts(country string, count string) {
	const apikey = "nFib543s9B6O6Jjd365eh2AaSfKMwVSr4xTUmLaK"
	uri := "https://ats.api.alexa.com/api?Action=TopSites&Count=" + count + "&CountryCode=" + country + "&ResponseGroup=Country&Output=json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("x-api-key", apikey)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	bodyByte := []byte(string(body))
	data := Data{}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(data.Ats.Results.Result.Alexa.TopSites.Country.Sites.Site)
}

func main() {

	if len(os.Args[0:]) == 1 {
		return
	}

	var action string = os.Args[1]
	switch {
	case action == "top" && len(os.Args[0:]) == 3:
		var number string = os.Args[2]
		callAts("", number)

	case action == "country" && len(os.Args[0:]) == 3:
		var country string = os.Args[2]
		callAts(country, "20")
	default:
		fmt.Print("error")
	}

}
