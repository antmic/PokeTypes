package main

import (
	"fmt"
)

var singleTypes = [18]string{"normal", "fire", "water", "electric", "grass", "ice", "fighting", "poison", "ground", "flying", "psychic", "bug", "rock", "ghost", "dragon", "dark", "steel", "fairy"}

var unusedDoubleTypes = [9][2]string{{"normal", "ice"}, {"normal", "bug"}, {"normal", "rock"}, {"normal", "steel"}, {"fire", "fairy"}, {"ice", "poison"}, {"ground", "fairy"}, {"bug", "dragon"}, {"rock", "ghost"}}

var doubleTypes = [144][2]string{{"normal", "fire"}, {"normal", "water"}, {"normal", "electric"}, {"normal", "grass"}, {"normal", "fighting"}, {"normal", "poison"}, {"normal", "ground"}, {"normal", "flying"}, {"normal", "psychic"}, {"normal", "ghost"}, {"normal", "dragon"}, {"normal", "dark"}, {"normal", "fairy"}, {"fire", "water"}, {"fire", "electric"}, {"fire", "grass"}, {"fire", "ice"}, {"fire", "fighting"}, {"fire", "poison"}, {"fire", "ground"}, {"fire", "flying"}, {"fire", "psychic"}, {"fire", "bug"}, {"fire", "rock"}, {"fire", "ghost"}, {"fire", "dragon"}, {"fire", "dark"}, {"fire", "steel"}, {"water", "electric"}, {"water", "grass"}, {"water", "ice"}, {"water", "fighting"}, {"water", "poison"}, {"water", "ground"}, {"water", "flying"}, {"water", "psychic"}, {"water", "bug"}, {"water", "rock"}, {"water", "ghost"}, {"water", "dragon"}, {"water", "dark"}, {"water", "steel"}, {"water", "fairy"}, {"electric", "grass"}, {"electric", "ice"}, {"electric", "fighting"}, {"electric", "poison"}, {"electric", "ground"}, {"electric", "flying"}, {"electric", "psychic"}, {"electric", "bug"}, {"electric", "rock"}, {"electric", "ghost"}, {"electric", "dragon"}, {"electric", "dark"}, {"electric", "steel"}, {"electric", "fairy"}, {"grass", "ice"}, {"grass", "fighting"}, {"grass", "poison"}, {"grass", "ground"}, {"grass", "flying"}, {"grass", "psychic"}, {"grass", "bug"}, {"grass", "rock"}, {"grass", "ghost"}, {"grass", "dragon"}, {"grass", "dark"}, {"grass", "steel"}, {"grass", "fairy"}, {"ice", "fighting"}, {"ice", "ground"}, {"ice", "flying"}, {"ice", "psychic"}, {"ice", "bug"}, {"ice", "rock"}, {"ice", "ghost"}, {"ice", "dragon"}, {"ice", "dark"}, {"ice", "steel"}, {"ice", "fairy"}, {"fighting", "poison"}, {"fighting", "ground"}, {"fighting", "flying"}, {"fighting", "psychic"}, {"fighting", "bug"}, {"fighting", "rock"}, {"fighting", "ghost"}, {"fighting", "dragon"}, {"fighting", "dark"}, {"fighting", "steel"}, {"fighting", "fairy"}, {"poison", "ground"}, {"poison", "flying"}, {"poison", "psychic"}, {"poison", "bug"}, {"poison", "rock"}, {"poison", "ghost"}, {"poison", "dragon"}, {"poison", "dark"}, {"poison", "steel"}, {"poison", "fairy"}, {"ground", "flying"}, {"ground", "psychic"}, {"ground", "bug"}, {"ground", "rock"}, {"ground", "ghost"}, {"ground", "dragon"}, {"ground", "dark"}, {"ground", "steel"}, {"flying", "psychic"}, {"flying", "bug"}, {"flying", "rock"}, {"flying", "ghost"}, {"flying", "dragon"}, {"flying", "dark"}, {"flying", "steel"}, {"flying", "fairy"}, {"psychic", "bug"}, {"psychic", "rock"}, {"psychic", "ghost"}, {"psychic", "dragon"}, {"psychic", "dark"}, {"psychic", "steel"}, {"psychic", "fairy"}, {"bug", "rock"}, {"bug", "ghost"}, {"bug", "dark"}, {"bug", "steel"}, {"bug", "fairy"}, {"rock", "dragon"}, {"rock", "dark"}, {"rock", "steel"}, {"rock", "fairy"}, {"ghost", "dragon"}, {"ghost", "dark"}, {"ghost", "steel"}, {"ghost", "fairy"}, {"dragon", "dark"}, {"dragon", "steel"}, {"dragon", "fairy"}, {"dark", "steel"}, {"dark", "fairy"}, {"steel", "fairy"}}

func generatePairs(list []string) [][2]string {
	var pairs [][2]string

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			pairs = append(pairs, [2]string{list[i], list[j]})
		}
	}

	// remove unused double types from pairs
	for _, unusedPair := range unusedDoubleTypes {
		for i, pair := range pairs {
			if (pair[0] == unusedPair[0] && pair[1] == unusedPair[1]) || (pair[0] == unusedPair[1] && pair[1] == unusedPair[0]) {
				pairs = append(pairs[:i], pairs[i+1:]...)
			}
		}
	}
	return pairs
}

func showPairs() {
	pairs := generatePairs(singleTypes[:])
	fmt.Println("Pairs len:", len(pairs))
	for _, pair := range pairs {
		fmt.Printf("{\"%s\", \"%s\"},\n", pair[0], pair[1])
	}
}
