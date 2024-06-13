package main

import "fmt"

func main() {
	defender := []string{"psychic", "ghost"}
	attacks, attackers, groupedOptimalAttackers := getEffectiveness(defender)
	fmt.Println(len(attacks))
	fmt.Println(len(attackers))
	for i, group := range groupedOptimalAttackers {
		fmt.Printf("Group %d: \n", i+1)
		for _, pokemon := range group {
			fmt.Println(pokemon)
		}
	}
}
