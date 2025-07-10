package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const AWL_URL string = "https://buergerportal.awl-neuss.de/api/v1/calendar"

func GetStreetNumbers() (map[string]int, error) {
	url := AWL_URL + "/townarea-streets"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var streets []struct {
		Number int `json:"strasseNummer"`
		Name string `json:"strasseBezeichnung"`
	}
	err = json.Unmarshal(content, &streets)
	if err != nil {
		return nil, err
	}

	nameToNumber := make(map[string]int, len(streets))
	for _, street := range streets {
		name := strings.ToLower(street.Name)
		name = strings.Replace(name, " ", "-", -1)
		name = strings.Replace(name, "--", "-", -1)
		nameToNumber[name] = street.Number
	}

	return nameToNumber, nil
}

func AwlTomorrow(street, home int) ([]string, error)  {
	startDate := time.Now().AddDate(0, 0, 1)
	startDateFmt := startDate.Format("Jan 02 2006")

	requestUrl := fmt.Sprintf(AWL_URL + "?startMonth=%s&streetNum=%d&homeNumber=%d", url.QueryEscape(startDateFmt), street, home)
	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var jsonBody map[string]map[string][]string
	err = json.Unmarshal(content, &jsonBody)
	if err != nil {
		return nil, fmt.Errorf("could not parse json:\n%s\n%s", err.Error(), content)
	}

	if len(jsonBody) != 1 {
		return nil, fmt.Errorf("unexpected awl response:\n%s", string(content))
	}

	currentMonth := fmt.Sprintf("%d-%d", startDate.Month()-1, startDate.Year())
	currentDay := startDate.Format("2")

	month, ok := jsonBody[currentMonth]
	if !ok {
		return nil, fmt.Errorf("could not access month (%s):\n%s", currentMonth, string(content))
	}

	day, ok := month[currentDay]
	if !ok {
		return nil, nil
	}

	return day, nil
}

