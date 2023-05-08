package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {

app := &cli.App{
	Commands: []*cli.Command{
		{
			Name:    "add",
			Usage:   "performs addition",
			Aliases: []string{"a"},
			Action: func(ctx *cli.Context) error {
				if ctx.Args().Len() != 1 { //if the input arguments less than 1
					log.Fatal("insufficient arguments")
				}
				inputs := ctx.Args().Get(0) //get all arguments
				separate := strings.Split(inputs, "+")//split the "+" from arguments now we have a slice for example like this [2] [2]
				if len(separate) != 2 { //if the arguments lenght is less than 2 
					return errors.New("invalid numbers")
				}
				x, err := strconv.Atoi(separate[0])//convert the first slice argument to integer
				if err != nil {
					return err
				}
				y, err := strconv.Atoi(separate[1])//convert the second an integer
				if err != nil {
					return err
				}
				fmt.Println("Result:", x+y) //print Result

				return nil
			},
		},
		{
			Name:    "sub",
			Usage:   "performs subtraction",
			Aliases: []string{"s"},
			Action: func(ctx *cli.Context) error {
				if ctx.Args().Len() != 1 {
					log.Fatal("insufficient arguments")
				}
				inputs := ctx.Args().Get(0) 
				separate := strings.Split(inputs, "-") 
				if len(separate) != 2 { 
					return errors.New("invalid arguments")
				}
				x, err := strconv.Atoi(separate[0]) 
				if err != nil {
					return err
				}
				y, err := strconv.Atoi(separate[1])
				if err != nil {
					return err
				}
				fmt.Println("Result:", x-y)
				return nil
			},
		},
		{
			Name:    "div",
			Usage:   "performs division",
			Aliases: []string{"d"},
			Action: func(ctx *cli.Context) error {
				if ctx.Args().Len() != 1 {
					log.Fatal("insufficient arguments")
				}
				inputs := ctx.Args().Get(0)
				separate := strings.Split(inputs, "/")
				if len(separate) != 2 {
					return errors.New("invalid argument count")
				}
				x, err := strconv.ParseFloat(separate[0], 64)
				if err != nil {
					return err
				}
				y, err := strconv.ParseFloat(separate[1], 64)
				if err != nil {
					return err
				}
				fmt.Println("Result: ", x/y)
				return nil
			},
		},
		{
			Name:    "mul",
			Usage:   "performs multiplication",
			Aliases: []string{"m"},
			Action: func(ctx *cli.Context) error {
				if ctx.Args().Len() != 1 {
					log.Fatal("invalid argument")
				}
				inputs := ctx.Args().Get(0)
				separate := strings.Split(inputs, "*")
				if len(separate) != 2 {
					return errors.New("invalid argument count")
				}
				x, err := strconv.Atoi(separate[0])
				if err != nil {
					return err
				}
				y, err := strconv.Atoi(separate[1])
				if err != nil {
					return err
				}
				fmt.Println("Result:", x*y)
				return nil
			},
		},
	},
}
   if err := app.Run(os.Args); err != nil {
	log.Fatal(err)
}
}
