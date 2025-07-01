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

	streetNumbers, err := GetStreetNumbers()
	if err != nil {
		fmt.Println(err)
		_ = SendErr(args.Street, err)
		return
	}

	streetNumber, ok := streetNumbers[args.Street]
	if !ok {
		fmt.Println("street could not be found:", args.Street)
	}

	tomorrow, err := AwlTomorrow(streetNumber, args.Home)
	if err != nil {
		fmt.Println(err)
		_ = SendErr(args.Street, err)
		return
	}

	if len(tomorrow) == 0 {
		return
	}

	err = SendAwlNotification(args.Street, tomorrow)
	if err != nil {
		fmt.Println(err)
		_ = SendErr(args.Street, err)
		return
	}
}
