package games

import (
	"fmt"

	"github.com/icza/gox/mathx"
)

func times(att1, att2, att3 int64, offset float64) string {
	return fmt.Sprintf("%.1f", mathx.Round(float64(2.0*att1+2.0*att2+att3)*offset, 0.1)) + "%"
}

func div(att1, att2, att3 int64, offset float64) string {
	return fmt.Sprintf("%.1f", mathx.Round(float64(2.0*att1+2.0*att2+att3)/offset, 0.1)) + "%"
}

func Calculate(attributes map[string]int64) map[string]map[string]interface{} {
	strength := attributes["strength"]
	will := attributes["willpower"]
	intelligence := attributes["intelligence"]
	power := attributes["power"]
	agility := attributes["agility"]
	return map[string]map[string]interface{}{
		"damage": {
			"bringer": div(strength, will, power, 400.0),
			"giver":   div(strength, will, power, 200.0),
			"dealer":  times(strength, will, power, 0.0075),
		},
		"resist": {
			"ward":  times(strength, agility, power, 0.012),
			"proof": div(strength, agility, power, 125.0),
		},
		"critical": {
			"defender": times(intelligence, will, power, 0.024),
			"blocker":  times(intelligence, will, power, 0.02),
		},
		"pierce": {
			"breaker": div(strength, agility, power, 400.0),
			"piercer": times(strength, agility, power, 0.0015),
		},
		"stun": {
			"recal":  div(strength, intelligence, power, 125.0),
			"resist": div(strength, intelligence, power, 250.0),
		},
		"healing": {
			"lively": times(strength, agility, power, 0.0065),
			"healer": times(strength, agility, power, 0.003),
			"medic":  times(strength, agility, power, 0.0065),
		},
		"health": {
			"healthy": times(intelligence, agility, power, 0.003),
			"gift":    times(agility, will, power, 0.1),
			"add":     times(agility, will, power, 0.06),
		},
		"attributes": {
			"strength":  strength,
			"intellect": intelligence,
			"agility":   agility,
			"willpower": will,
			"power":     power,
			"happiness": strength + intelligence + agility + will + power,
		},
	}
}
