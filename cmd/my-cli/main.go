package main

import (
	"fmt"
	"log"
	"net"
	"os"

	lc "ARIS/cmd/ipinfo"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Artificial Recollection Interface System [A.R.I.S.]"
	app.Usage = "Let's you query IPs, name servers, and location"
	app.Author = "Daniel Tsang"
	app.Version = "1.0.0"
}

func commands() {
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP Addresses for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)

				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "lc",
			Usage: "Looks up the user location",
			Action: func(c *cli.Context) error {
				lc, err := lc.GetLocation()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("\n      Location of user found: %s, %s, %s, %s,%s\n", lc.Country, lc.City, lc.Region, lc.Zip, lc.Loc)
				return nil
			},
		},
		{
			Name:  "cn",
			Usage: "Looks up the cname of a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
