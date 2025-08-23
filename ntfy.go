package main

import (
	"fmt"
	"net/http"
	"strings"
)

var NotifyChannel string

func SendAwlNotification(binColors []string, pink, yellow, blue, gray, brown bool) error {
	var bins = make([]string, 0, 5)
	for _, color := range binColors {
		switch color {
		case "pink":
			if pink {
				bins = append(bins, "🟣 Pink")
			}
		case "gelb":
			if yellow {
				bins = append(bins, "🟡 Gelb")
			}
		case "blau":
			if blue {
				bins = append(bins, "🔵 Blau")
			}
		case "grau":
			if gray {
				bins = append(bins, "⚫ Grau")
			}
		case "braun":
			if brown {
				bins = append(bins, "🟤 Braun")
			}
		}
	}

	if len(bins) == 0 {
		return nil
	}

	var message string
	if len(bins) == 1 {
		message = "Morgen wird folgende Tonne abgeholt: "
	} else {
		message = "Morgen werden folgende Tonnen abgeholt: "
	}
	message += "\n"
	message += strings.Join(bins, "\n")
	return SendNotification(message)
}

func SendNotification(message string) error {
	topic := fmt.Sprintf("https://ntfy.sh/awl-neuss-%s", NotifyChannel)
	_, err := http.Post(topic, "text/plain", strings.NewReader(message))
	return err
}

func SendErr(err error) error {
	fmt.Println(err)
	topic := "https://ntfy.sh/awl-neuss-err"
	_, err = http.Post(topic, "text/plain", strings.NewReader(NotifyChannel + "\n" + err.Error()))
	return err
}
