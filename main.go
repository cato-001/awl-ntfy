package main

import (
	"fmt"
	"os"

	arg "github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Street string `arg:"positional,required"`
		Home int `arg:"positional,required"`
	}

	parser, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		fmt.Println(err)
		return
	}
	parser.MustParse(os.Args[1:])

	NotifyChannel = fmt.Sprintf("%s-%d", args.Street, args.Home)

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

	err = SendAwlNotification(tomorrow)
	if err != nil {
		_ = SendErr(err)
		return
	}
}
