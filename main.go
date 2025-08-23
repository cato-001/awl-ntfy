package main

import (
	"fmt"
	"os"
	"strings"

	arg "github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Street string `arg:"positional,required"`
		Home int `arg:"positional,required"`
		Pink bool `arg:"--pink"`
		Yellow bool `arg:"--yellow"`
		Blue bool `arg:"--blue"`
		Gray bool `arg:"--gray"`
		Brown bool `arg:"--brown"`
	}

	parser, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		fmt.Println(err)
		return
	}
	parser.MustParse(os.Args[1:])

	if !(args.Pink || args.Yellow || args.Blue || args.Gray || args.Brown) {
		args.Pink, args.Yellow, args.Blue, args.Gray, args.Brown = true, true, true, true, true
	}

	replacer := strings.NewReplacer("ö", "oe", "ä", "ae", "ü", "ue", "ß", "ss")
	NotifyChannel = fmt.Sprintf("%s-%d", args.Street, args.Home)
	NotifyChannel = replacer.Replace(NotifyChannel)

	streetNumbers, err := GetStreetNumbers()
	if err != nil {
		_ = SendErr(err)
		return
	}

	streetNumber, ok := streetNumbers[args.Street]
	if !ok {
		fmt.Println("street could not be found:", args.Street)
	}

	tomorrow, err := AwlTomorrow(streetNumber, args.Home)
	if err != nil {
		_ = SendErr(err)
		return
	}

	if len(tomorrow) == 0 {
		return
	}

	err = SendAwlNotification(tomorrow, args.Pink, args.Yellow, args.Blue, args.Gray, args.Brown)
	if err != nil {
		_ = SendErr(err)
		return
	}
}
