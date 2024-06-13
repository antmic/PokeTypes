package main

import (
	"slices"
	"sort"
)

// Constants
const NORMAL = 1.0
const SUPER_EFFECTIVE = 1.6
const INEFFECTIVE = 0.625
const VERY_INEFFECTIVE = 0.390625

var attackTypes = [18][1]string{{"normal"}, {"fire"}, {"water"}, {"electric"}, {"grass"}, {"ice"}, {"fighting"}, {"poison"}, {"ground"}, {"flying"}, {"psychic"}, {"bug"}, {"rock"}, {"ghost"}, {"dragon"}, {"dark"}, {"steel"}, {"fairy"}}

var attackerTypes = [162][]string{{"normal"}, {"fire"}, {"water"}, {"electric"}, {"grass"}, {"ice"}, {"fighting"}, {"poison"}, {"ground"}, {"flying"}, {"psychic"}, {"bug"}, {"rock"}, {"ghost"}, {"dragon"}, {"dark"}, {"steel"}, {"fairy"}, {"normal", "fire"}, {"normal", "water"}, {"normal", "electric"}, {"normal", "grass"}, {"normal", "fighting"}, {"normal", "poison"}, {"normal", "ground"}, {"normal", "flying"}, {"normal", "psychic"}, {"normal", "ghost"}, {"normal", "dragon"}, {"normal", "dark"}, {"normal", "fairy"}, {"fire", "water"}, {"fire", "electric"}, {"fire", "grass"}, {"fire", "ice"}, {"fire", "fighting"}, {"fire", "poison"}, {"fire", "ground"}, {"fire", "flying"}, {"fire", "psychic"}, {"fire", "bug"}, {"fire", "rock"}, {"fire", "ghost"}, {"fire", "dragon"}, {"fire", "dark"}, {"fire", "steel"}, {"water", "electric"}, {"water", "grass"}, {"water", "ice"}, {"water", "fighting"}, {"water", "poison"}, {"water", "ground"}, {"water", "flying"}, {"water", "psychic"}, {"water", "bug"}, {"water", "rock"}, {"water", "ghost"}, {"water", "dragon"}, {"water", "dark"}, {"water", "steel"}, {"water", "fairy"}, {"electric", "grass"}, {"electric", "ice"}, {"electric", "fighting"}, {"electric", "poison"}, {"electric", "ground"}, {"electric", "flying"}, {"electric", "psychic"}, {"electric", "bug"}, {"electric", "rock"}, {"electric", "ghost"}, {"electric", "dragon"}, {"electric", "dark"}, {"electric", "steel"}, {"electric", "fairy"}, {"grass", "ice"}, {"grass", "fighting"}, {"grass", "poison"}, {"grass", "ground"}, {"grass", "flying"}, {"grass", "psychic"}, {"grass", "bug"}, {"grass", "rock"}, {"grass", "ghost"}, {"grass", "dragon"}, {"grass", "dark"}, {"grass", "steel"}, {"grass", "fairy"}, {"ice", "fighting"}, {"ice", "ground"}, {"ice", "flying"}, {"ice", "psychic"}, {"ice", "bug"}, {"ice", "rock"}, {"ice", "ghost"}, {"ice", "dragon"}, {"ice", "dark"}, {"ice", "steel"}, {"ice", "fairy"}, {"fighting", "poison"}, {"fighting", "ground"}, {"fighting", "flying"}, {"fighting", "psychic"}, {"fighting", "bug"}, {"fighting", "rock"}, {"fighting", "ghost"}, {"fighting", "dragon"}, {"fighting", "dark"}, {"fighting", "steel"}, {"fighting", "fairy"}, {"poison", "ground"}, {"poison", "flying"}, {"poison", "psychic"}, {"poison", "bug"}, {"poison", "rock"}, {"poison", "ghost"}, {"poison", "dragon"}, {"poison", "dark"}, {"poison", "steel"}, {"poison", "fairy"}, {"ground", "flying"}, {"ground", "psychic"}, {"ground", "bug"}, {"ground", "rock"}, {"ground", "ghost"}, {"ground", "dragon"}, {"ground", "dark"}, {"ground", "steel"}, {"flying", "psychic"}, {"flying", "bug"}, {"flying", "rock"}, {"flying", "ghost"}, {"flying", "dragon"}, {"flying", "dark"}, {"flying", "steel"}, {"flying", "fairy"}, {"psychic", "bug"}, {"psychic", "rock"}, {"psychic", "ghost"}, {"psychic", "dragon"}, {"psychic", "dark"}, {"psychic", "steel"}, {"psychic", "fairy"}, {"bug", "rock"}, {"bug", "ghost"}, {"bug", "dark"}, {"bug", "steel"}, {"bug", "fairy"}, {"rock", "dragon"}, {"rock", "dark"}, {"rock", "steel"}, {"rock", "fairy"}, {"ghost", "dragon"}, {"ghost", "dark"}, {"ghost", "steel"}, {"ghost", "fairy"}, {"dragon", "dark"}, {"dragon", "steel"}, {"dragon", "fairy"}, {"dark", "steel"}, {"dark", "fairy"}, {"steel", "fairy"}}

var chart = map[string]map[string]float64{
	"normal": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     INEFFECTIVE,
		"ghost":    VERY_INEFFECTIVE,
		"dragon":   NORMAL,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
	"fire": {
		"normal":   NORMAL,
		"fire":     INEFFECTIVE,
		"water":    INEFFECTIVE,
		"electric": NORMAL,
		"grass":    SUPER_EFFECTIVE,
		"ice":      SUPER_EFFECTIVE,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      SUPER_EFFECTIVE,
		"rock":     INEFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   INEFFECTIVE,
		"dark":     NORMAL,
		"steel":    SUPER_EFFECTIVE,
		"fairy":    NORMAL,
	},
	"water": {
		"normal":   NORMAL,
		"fire":     SUPER_EFFECTIVE,
		"water":    INEFFECTIVE,
		"electric": NORMAL,
		"grass":    INEFFECTIVE,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   SUPER_EFFECTIVE,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     SUPER_EFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   INEFFECTIVE,
		"dark":     NORMAL,
		"steel":    NORMAL,
		"fairy":    NORMAL,
	},
	"electric": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    SUPER_EFFECTIVE,
		"electric": INEFFECTIVE,
		"grass":    INEFFECTIVE,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   VERY_INEFFECTIVE,
		"flying":   SUPER_EFFECTIVE,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    NORMAL,
		"dragon":   INEFFECTIVE,
		"dark":     NORMAL,
		"steel":    NORMAL,
		"fairy":    NORMAL,
	},
	"grass": {
		"normal":   NORMAL,
		"fire":     INEFFECTIVE,
		"water":    SUPER_EFFECTIVE,
		"electric": NORMAL,
		"grass":    INEFFECTIVE,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   INEFFECTIVE,
		"ground":   SUPER_EFFECTIVE,
		"flying":   INEFFECTIVE,
		"psychic":  NORMAL,
		"bug":      INEFFECTIVE,
		"rock":     SUPER_EFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   INEFFECTIVE,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
	"ice": {
		"normal":   NORMAL,
		"fire":     INEFFECTIVE,
		"water":    INEFFECTIVE,
		"electric": NORMAL,
		"grass":    SUPER_EFFECTIVE,
		"ice":      INEFFECTIVE,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   SUPER_EFFECTIVE,
		"flying":   SUPER_EFFECTIVE,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    NORMAL,
		"dragon":   SUPER_EFFECTIVE,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
	"fighting": {
		"normal":   SUPER_EFFECTIVE,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      SUPER_EFFECTIVE,
		"fighting": NORMAL,
		"poison":   INEFFECTIVE,
		"ground":   NORMAL,
		"flying":   INEFFECTIVE,
		"psychic":  INEFFECTIVE,
		"bug":      INEFFECTIVE,
		"rock":     SUPER_EFFECTIVE,
		"ghost":    VERY_INEFFECTIVE,
		"dragon":   NORMAL,
		"dark":     SUPER_EFFECTIVE,
		"steel":    SUPER_EFFECTIVE,
		"fairy":    INEFFECTIVE,
	},
	"poison": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    SUPER_EFFECTIVE,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   INEFFECTIVE,
		"ground":   INEFFECTIVE,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     INEFFECTIVE,
		"ghost":    INEFFECTIVE,
		"dragon":   NORMAL,
		"dark":     NORMAL,
		"steel":    VERY_INEFFECTIVE,
		"fairy":    SUPER_EFFECTIVE,
	},
	"ground": {
		"normal":   NORMAL,
		"fire":     SUPER_EFFECTIVE,
		"water":    NORMAL,
		"electric": SUPER_EFFECTIVE,
		"grass":    INEFFECTIVE,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   SUPER_EFFECTIVE,
		"ground":   NORMAL,
		"flying":   VERY_INEFFECTIVE,
		"psychic":  NORMAL,
		"bug":      INEFFECTIVE,
		"rock":     SUPER_EFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   NORMAL,
		"dark":     NORMAL,
		"steel":    SUPER_EFFECTIVE,
		"fairy":    NORMAL,
	},
	"flying": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": INEFFECTIVE,
		"grass":    SUPER_EFFECTIVE,
		"ice":      NORMAL,
		"fighting": SUPER_EFFECTIVE,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      SUPER_EFFECTIVE,
		"rock":     INEFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   NORMAL,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
	"psychic": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      NORMAL,
		"fighting": SUPER_EFFECTIVE,
		"poison":   SUPER_EFFECTIVE,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  INEFFECTIVE,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    NORMAL,
		"dragon":   NORMAL,
		"dark":     VERY_INEFFECTIVE,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
	"bug": {
		"normal":   NORMAL,
		"fire":     INEFFECTIVE,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    SUPER_EFFECTIVE,
		"ice":      NORMAL,
		"fighting": INEFFECTIVE,
		"poison":   INEFFECTIVE,
		"ground":   NORMAL,
		"flying":   INEFFECTIVE,
		"psychic":  SUPER_EFFECTIVE,
		"bug":      NORMAL,
		"rock":     INEFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   NORMAL,
		"dark":     SUPER_EFFECTIVE,
		"steel":    INEFFECTIVE,
		"fairy":    INEFFECTIVE,
	},
	"rock": {
		"normal":   NORMAL,
		"fire":     SUPER_EFFECTIVE,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      SUPER_EFFECTIVE,
		"fighting": INEFFECTIVE,
		"poison":   NORMAL,
		"ground":   INEFFECTIVE,
		"flying":   SUPER_EFFECTIVE,
		"psychic":  NORMAL,
		"bug":      SUPER_EFFECTIVE,
		"rock":     NORMAL,
		"ghost":    NORMAL,
		"dragon":   NORMAL,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
	"ghost": {
		"normal":   VERY_INEFFECTIVE,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  SUPER_EFFECTIVE,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    SUPER_EFFECTIVE,
		"dragon":   NORMAL,
		"dark":     INEFFECTIVE,
		"steel":    NORMAL,
		"fairy":    NORMAL,
	},
	"dragon": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      NORMAL,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    NORMAL,
		"dragon":   SUPER_EFFECTIVE,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    VERY_INEFFECTIVE,
	},
	"dark": {
		"normal":   NORMAL,
		"fire":     NORMAL,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      NORMAL,
		"fighting": INEFFECTIVE,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  SUPER_EFFECTIVE,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    SUPER_EFFECTIVE,
		"dragon":   NORMAL,
		"dark":     INEFFECTIVE,
		"steel":    NORMAL,
		"fairy":    INEFFECTIVE,
	},
	"steel": {
		"normal":   NORMAL,
		"fire":     INEFFECTIVE,
		"water":    INEFFECTIVE,
		"electric": INEFFECTIVE,
		"grass":    NORMAL,
		"ice":      SUPER_EFFECTIVE,
		"fighting": NORMAL,
		"poison":   NORMAL,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     SUPER_EFFECTIVE,
		"ghost":    NORMAL,
		"dragon":   NORMAL,
		"dark":     NORMAL,
		"steel":    INEFFECTIVE,
		"fairy":    SUPER_EFFECTIVE,
	},
	"fairy": {
		"normal":   NORMAL,
		"fire":     INEFFECTIVE,
		"water":    NORMAL,
		"electric": NORMAL,
		"grass":    NORMAL,
		"ice":      NORMAL,
		"fighting": SUPER_EFFECTIVE,
		"poison":   INEFFECTIVE,
		"ground":   NORMAL,
		"flying":   NORMAL,
		"psychic":  NORMAL,
		"bug":      NORMAL,
		"rock":     NORMAL,
		"ghost":    NORMAL,
		"dragon":   SUPER_EFFECTIVE,
		"dark":     SUPER_EFFECTIVE,
		"steel":    INEFFECTIVE,
		"fairy":    NORMAL,
	},
}

type attacker struct {
	pokemonType    []string
	damageReceived []float64
}

type attack struct {
	attackType          string
	attackEffectiveness float64
}

type optimalAttacker struct {
	pokemonType         []string
	attackType          string
	attackEffectiveness float64
	damageReceived      []float64
	difference          float64
}

func getEffectiveness(defenderType []string) ([]attack, []attacker, [][]optimalAttacker) {

	attacks := []attack{}

	// calculate the effectiveness of each attack type against the defender
	for _, element := range attackTypes {
		attackEffectiveness := 1.0
		for _, defender := range defenderType {
			attackEffectiveness *= chart[element[0]][defender]
		}
		attacks = append(attacks, attack{
			attackType:          element[0],
			attackEffectiveness: round2(attackEffectiveness),
		})
	}
	// sort the results by effectiveness
	sort.Slice(attacks, func(i, j int) bool {
		return attacks[i].attackEffectiveness > attacks[j].attackEffectiveness
	})

	// calculate the effectiveness of defender's counter attacks for every possible attacker type
	attackers := []attacker{}
	for _, attackerType := range attackerTypes {
		var damageReceived []float64

		for _, defType := range defenderType {
			if len(attackerType) == 1 {
				damageReceived = append(damageReceived, round2(chart[defType][attackerType[0]]))
			} else {
				damageReceived = append(damageReceived, round2(chart[defType][attackerType[0]]*chart[defType][attackerType[1]]))
			}
		}
		attackers = append(attackers, attacker{
			pokemonType:    attackerType,
			damageReceived: damageReceived,
		})
	}
	// sort the results ascending by the sum of the damage received
	sort.Slice(attackers, func(i, j int) bool {
		return sumSlice(attackers[i].damageReceived) < sumSlice(attackers[j].damageReceived)
	})

	// get attacks with effectiveness higher than 1.0
	var effectiveAttacks []attack
	for _, attack := range attacks {
		if attack.attackEffectiveness > 1.0 {
			effectiveAttacks = append(effectiveAttacks, attack)
		} else {
			break
		}
	}

	// get unique attackers with type that includes the most effective attack type and sum is less than 2.0
	var matchingAttackers []attacker
	for _, attacker := range attackers {
		for _, attack := range effectiveAttacks {
			if sumSlice(attacker.damageReceived) < 2.0 {
				if slices.Contains(attacker.pokemonType, attack.attackType) {
					if len(matchingAttackers) == 0 || slices.Compare(attacker.pokemonType, matchingAttackers[len(matchingAttackers)-1].pokemonType) != 0 {
						matchingAttackers = append(matchingAttackers, attacker)
					}
				}
			} else {
				break
			}
		}
	}

	// match the most effective attack type with the most effective attacker type
	optimalAttackers := []optimalAttacker{}
	for _, attacker := range matchingAttackers {
		for _, attackerAttackType := range attacker.pokemonType {
			for _, attack := range effectiveAttacks {
				if attack.attackType == attackerAttackType {
					optimalAttackers = append(optimalAttackers, optimalAttacker{
						pokemonType:         attacker.pokemonType,
						attackType:          attack.attackType,
						attackEffectiveness: attack.attackEffectiveness,
						damageReceived:      attacker.damageReceived,
						difference:          round2(attack.attackEffectiveness - sumSlice(attacker.damageReceived)),
					})
				}
			}
		}
	}

	// sort the results by the attack effectiveness - sum of the damage received
	sort.Slice(optimalAttackers, func(i, j int) bool {
		return optimalAttackers[i].attackEffectiveness-sumSlice(optimalAttackers[i].damageReceived) > optimalAttackers[j].attackEffectiveness-sumSlice(optimalAttackers[j].damageReceived)
	})

	// get pokemon with best attack effectiveness - sum of the damage received
	groupedOptimalAttackers := [][]optimalAttacker{}
	counter := 0
	for _, attacker := range optimalAttackers {
		if len(groupedOptimalAttackers) == 0 || groupedOptimalAttackers[len(groupedOptimalAttackers)-1][0].difference == attacker.difference {
			if len(groupedOptimalAttackers) <= counter {
				groupedOptimalAttackers = append(groupedOptimalAttackers, []optimalAttacker{})
			}
			groupedOptimalAttackers[counter] = append(groupedOptimalAttackers[counter], attacker)
		} else {
			counter++
			if len(groupedOptimalAttackers) <= counter {
				groupedOptimalAttackers = append(groupedOptimalAttackers, []optimalAttacker{})
			}
			groupedOptimalAttackers[counter] = append(groupedOptimalAttackers[counter], attacker)
		}
	}

	return attacks, attackers, groupedOptimalAttackers
}
