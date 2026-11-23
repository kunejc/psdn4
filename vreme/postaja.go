// Package vreme omogoca upravljanje z vremenskimi postajami in njihovimi meritvami.
package vreme

import "fmt"

// Postaja predstavlja vremesnko postajo z imenom, krajem in meritvami
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajMeritev doda postaji z dano oznako novo meritev, ce je ta veljavna, drugace izpise obvestilo.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("Postaja ne obstaja.")
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Meritev ni veljavna.")
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje. Če postaja ne obstaja ali ima manj kot 3 meritve, naj funkcija vrne ok == false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	vsota := 0.0

	for _, temperatura := range postaja.meritve {
		if temperatura < najnizja {
			najnizja = temperatura
		}
		if temperatura > najvisja {
			najvisja = temperatura
		}
		vsota += temperatura
	}

	povprecna = vsota / float64(len(postaja.meritve))
	return najnizja, najvisja, povprecna, true
}

// MERITVE:
// KP01 - Koper Luka (Koper): [24.5 27 26.1 28.4]

// IzpisMeritev izpiše oznake, imena in kraje vseh postaj ter njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("\nMERITVE:")
	for key, postaja := range postaje {
		fmt.Printf("%v - %v (%v): [ ", key, postaja.ime, postaja.kraj)
		for i := range postaja.meritve {
			fmt.Print(postaja.meritve[i])
			fmt.Print(" ")
		}
		fmt.Print("]\n")
	}
}

// če je povprečje >= 25 °C, naj izpiše Vroče;
// če je povprečje med 15 °C in 25 °C, naj izpiše Toplo;
// če je povprečje med 5 °C in 15 °C, naj izpiše Hladno;
// če je povprečje < 5 °C, naj izpiše Mrzlo.
// Če statistike ni mogoče izračunati, naj izpiše premalo meritev.

// IzpisPorocila za vsako postajo izpiše ime ter najnižjo, najvišjo in povprečno temperaturo.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	fmt.Println("\nPOROČILO:")
	for key, postaja := range postaje {
		fmt.Println("Postaja:", postaja.ime)

		min, max, povprecje, ok := statistika(postaje, key, stMeritev)

		if !ok {
			fmt.Println("	Premalo meritev.")
			continue
		}

		fmt.Println("	Najnižja temperatura:", min)
		fmt.Println("	Najvišja temperatura:", max)
		fmt.Println("	Povprečna temperatura:", povprecje)

		switch {
		case povprecje >= 25:
			fmt.Println("	Vroče")
		case povprecje >= 15:
			fmt.Println("	Toplo")
		case povprecje >= 5:
			fmt.Println("	Hladno")
		default:
			fmt.Println("	Mrzlo")
		}
		// fmt.Print("\n")
	}
}

// DodajPostajo v slovar doda novo postajo s podano oznako, imenom in krajem.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: nil,
	}
}

// func main() {
// 	var postaje map[string]Postaja
// 	postaje = make(map[string]Postaja)

// 	postaje["LJ01"] = Postaja{
// 		ime:     "Glavna LJ",
// 		kraj:    "Ljubljana",
// 		meritve: []float64{12.4, 10, 14.5},
// 	}

// 	postaje["LJ02"] = Postaja{
// 		ime:     "Rezervna LJ",
// 		kraj:    "Ljubljana",
// 		meritve: []float64{12, 10, 13.5},
// 	}

// 	postaje["MB01"] = Postaja{
// 		ime:     "Glavna MB",
// 		kraj:    "Maribor",
// 		meritve: []float64{13.5, 12.3, 14.1},
// 	}

// 	postaje["CE01"] = Postaja{
// 		ime:     "Glavna CE",
// 		kraj:    "Celje",
// 		meritve: []float64{11, 10.7, 11.9},
// 	}

// 	// izpisMeritev(postaje)

// 	// dodajMeritev(postaje, "LJ01", 12.5)
// 	// dodajMeritev(postaje, "LJ02", 12.5)
// 	// dodajMeritev(postaje, "MB01", 13.1)
// 	// dodajMeritev(postaje, "CE01", -49)

// 	// izpisMeritev(postaje)

// 	// izpisPorocila(postaje)

// }
