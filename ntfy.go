package main

import (
	"fmt"
	"net/http"
	"strings"
)

var NotifyChannel string

func SendAwlNotification(binColors []string) error {
	var bins = make([]string, 0, 5)
	for _, color := range binColors {
		switch color {
		case "pink":
			bins = append(bins, "🟣 Pink")
		case "gelb":
			bins = append(bins, "🟡 Gelb")
		case "blau":
			bins = append(bins, "🔵 Blau")
		case "grau":
			bins = append(bins, "⚫ Grau")
		case "braun":
			bins = append(bins, "🟤 Braun")
		}
	}

	var message string
	if len(bins) == 1 {
		message = "Morgen wird folgende Tonne abgeholt: "
	} else {
		message = "Morgen werden folgende Tonnen abgeholt: "
	}
	message += strings.Join(bins, ", ")
	return SendNotification(message)
}

func SendNotification(message string) error {
	topic := fmt.Sprintf("https://ntfy.sh/awl-neuss-%s", NotifyChannel)
	_, err := http.Post(topic, "text/plain", strings.NewReader(message))
	return err
}

func SendErr(err error) error {
	fmt.Println(err)
	topic := fmt.Sprintf("https://ntfy.sh/awl-neuss-%s-err", NotifyChannel)
	_, err = http.Post(topic, "text/plain", strings.NewReader(err.Error()))
	return err
}
