package main

import (
	"context"
	"log"
	"os"

	"github.com/kunejc/psdn4/vreme"
	"github.com/urfave/cli/v3"
)

func zazeni(minTemp, maxTemp float64, stMeritev int) {
	// stMeritev := 3
	// var minTemp float64 = -50
	// var maxTemp float64 = 50

	postaje := make(map[string]vreme.Postaja)

	vreme.DodajPostajo(postaje, "LJ01", "Glavna LJ", "Ljubljana")
	vreme.DodajMeritev(postaje, "LJ01", 12.4, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ01", 10, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ01", 14.5, minTemp, maxTemp)

	vreme.DodajPostajo(postaje, "LJ02", "Rezervna LJ", "Ljubljana")
	vreme.DodajMeritev(postaje, "LJ02", 12, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ02", 10, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ02", 13.5, minTemp, maxTemp)

	vreme.DodajPostajo(postaje, "MB01", "Glavna MB", "Maribor")
	vreme.DodajMeritev(postaje, "MB01", 13.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB01", 12.3, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB01", 14.1, minTemp, maxTemp)

	vreme.DodajPostajo(postaje, "CE01", "Glavna CE", "Celje")
	vreme.DodajMeritev(postaje, "CE01", 11, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "CE01", 10.7, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "CE01", 11.9, minTemp, maxTemp)

	vreme.IzpisMeritev(postaje)

	vreme.DodajMeritev(postaje, "LJ01", 12.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ02", 12.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB01", 13.1, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "CE01", -49, minTemp, maxTemp)

	vreme.IzpisMeritev(postaje)

	vreme.IzpisPorocila(postaje, stMeritev)

}

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Upravljanje vremenskih postaj in njihovih meritev",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnizja dovoljena temperatura meritve (°C)",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvisja dovoljena temperatura meritve (°C)",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "potrebno stevilo meritev za izracun statistike",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			zazeni(cmd.Float("minTemp"), cmd.Float("maxTemp"), cmd.Int("stMeritev"))
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
