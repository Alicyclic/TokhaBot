package games

import (
	"fmt"
	"log"
	"math/rand"
)

type artifactStat int
type artifactSlot int
type artifactSet string

func (a *Artifact) randomizeSubstats() {
	numRolls := 3 + 5
	if rand.Float32() <= 0.25 {
		numRolls++
	}

	a.SubStats = [MaxSubstats]*ArtifactSubstat{}
	possibleStats := weightedSubstats(a.MainStat)
	var subs [MaxSubstats]artifactStat
	for i := 0; i < numRolls; i++ {
		if i < MaxSubstats {
			stat := weightedRand(possibleStats)
			subs[i] = stat
			a.SubStats[i] = &ArtifactSubstat{Stat: stat, Rolls: 1}
			delete(possibleStats, stat)
		} else {
			// Rest of rolls
			index := rand.Intn(MaxSubstats)
			a.SubStats[index].Rolls += 1
		}
	}

	for _, substat := range a.SubStats {
		substat.randomizeValue()
	}
}

func (s *ArtifactSubstat) randomizeValue() {
	s.Value = 0
	for i := 0; i < s.Rolls; i++ {
		s.Value = s.Value + s.Stat.RandomRollValue()
	}
}

type ArtifactSubstat struct {
	Stat  artifactStat
	Rolls int
	Value float32
}

func (s *ArtifactSubstat) String() string {
	return fmt.Sprintf("%s: %.1f", s.Stat, s.Value)
}

type Artifact struct {
	Set      artifactSet
	Slot     artifactSlot
	MainStat artifactStat
	SubStats [MaxSubstats]*ArtifactSubstat
}

func (a Artifact) SubsQuality(subValue map[artifactStat]float32) float32 {
	var quality float32
	for _, sub := range a.SubStats {
		quality += float32(sub.Rolls) * subValue[sub.Stat]
	}
	return quality
}

func (a *Artifact) ranzomizeMainStat() {
	switch a.Slot {
	case SlotFlower:
		a.MainStat = HP
	case SlotPlume:
		a.MainStat = ATK
	case SlotSands:
		a.MainStat = weightedRand(sandsWeightedStats)
	case SlotGoblet:
		a.MainStat = weightedRand(gobletWeightedStats)
	case SlotCirclet:
		a.MainStat = weightedRand(circletWeightedStats)
	}
}

func RandomArtifact() *Artifact {
	var artifact Artifact
	artifact.randomizeSet(allArtifactSets...)
	artifact.randomizeSlot()
	artifact.ranzomizeMainStat()
	artifact.randomizeSubstats()
	return &artifact
}

func (a *Artifact) randomizeSet(options ...artifactSet) {
	a.Set = options[rand.Intn(len(options))]
}

func (a *Artifact) randomizeSlot() {
	a.Slot = artifactSlot(rand.Intn(5))
}

func FormatGenshinArtifact(artifact *Artifact) string {
	return fmt.Sprintf(`
**%s**
**%s (%s)**
 • %s: %.1f
 • %s: %.1f
 • %s: %.1f
 • %s: %.1f
		`, artifact.Set, artifact.Slot, artifact.MainStat,
		artifact.SubStats[0].Stat, artifact.SubStats[0].Value,
		artifact.SubStats[1].Stat, artifact.SubStats[1].Value,
		artifact.SubStats[2].Stat, artifact.SubStats[2].Value,
		artifact.SubStats[3].Stat, artifact.SubStats[3].Value,
	)
}

func weightedRand(weightedVals map[artifactStat]int) artifactStat {
	sum := 0
	for _, weight := range weightedVals {
		sum += weight
	}

	i := rand.Intn(sum)
	for value, weight := range weightedVals {
		i -= weight
		if i < 0 {
			return value
		}
	}

	log.Println("fatal error in WeightedRand: should never reach this log")
	return 0
}

const (
	MaxSubstats              = 4
	HP          artifactStat = iota
	ATK
	DEF
	HPP
	ATKP
	DEFP
	EnergyRecharge
	ElementalMastery
	CritRate
	CritDmg
	PyroDMG
	ElectroDMG
	CryoDMG
	HydroDMG
	AnemoDMG
	GeoDMG
	PhysDMG
	HealingBonus

	flatSubstatWeight   = 150
	commonSubstatWeight = 100
	critSubstatWeight   = 75
)

const (
	SlotFlower artifactSlot = iota
	SlotPlume
	SlotSands
	SlotGoblet
	SlotCirclet
)

func (t artifactSlot) String() string {
	switch t {
	case SlotFlower:
		return "Flower of Life"
	case SlotPlume:
		return "Plume of Death"
	case SlotSands:
		return "Sands of Eon"
	case SlotGoblet:
		return "Goblet of Eonothem"
	case SlotCirclet:
		return "Circlet of Logos"
	}
	return "Unknown"
}

var allArtifactSets = []artifactSet{
	"Gladiator's Finale",
	"Wanderer's Troupe",
	"Thundersoother",
	"Thundering Fury",
	"Maiden Beloved",
	"Viridescent Venerer",
	"Crimson Witch of Flames",
	"Lavawalker",
	"Noblesse Oblige",
	"Bloodstained Chivalry",
	"Archaic Petra",
	"Retracing Bolide",
	"Blizzard Strayer",
	"Heart of Depth",
	"Tenacity of the Millelith",
	"Pale Flame",
	"Emblem of Severed Fate",
	"Shimenawa's Reminiscence",
	"Husk of Opulent Dreams",
	"Ocean-Hued Clam",
}

func (s artifactStat) String() string {
	switch s {
	case HP:
		return "HP"
	case ATK:
		return "ATK"
	case DEF:
		return "DEF"
	case HPP:
		return "HP%"
	case ATKP:
		return "ATK%"
	case DEFP:
		return "DEF%"
	case EnergyRecharge:
		return "Energy Recharge%"
	case ElementalMastery:
		return "Elemental Mastery"
	case CritRate:
		return "CRIT Rate%"
	case CritDmg:
		return "CRIT DMG%"
	case PyroDMG:
		return "Pyro DMG%"
	case ElectroDMG:
		return "Electro DMG%"
	case CryoDMG:
		return "Cryo DMG%"
	case HydroDMG:
		return "Hydro DMG%"
	case AnemoDMG:
		return "Anemo DMG%"
	case GeoDMG:
		return "Geo DMG%"
	case PhysDMG:
		return "Physical DMG%"
	case HealingBonus:
		return "Healing Bonus%"
	}
	return "Unknown"
}

func (s artifactStat) RandomRollValue() float32 {
	var highRoll float32
	switch s {
	case HP:
		highRoll = 298.75
	case ATK:
		highRoll = 19.45
	case DEF:
		highRoll = 23.15
	case HPP:
		highRoll = 5.83
	case ATKP:
		highRoll = 5.83
	case DEFP:
		highRoll = 7.29
	case ElementalMastery:
		highRoll = 23.31
	case EnergyRecharge:
		highRoll = 6.48
	case CritRate:
		highRoll = 3.89
	case CritDmg:
		highRoll = 7.77
	}

	switch rand.Intn(4) {
	case 0:
		return highRoll * 0.7
	case 1:
		return highRoll * 0.8
	case 2:
		return highRoll * 0.9
	default:
		return highRoll
	}
}

var sandsWeightedStats = map[artifactStat]int{
	HPP:              2668,
	ATKP:             2666,
	DEFP:             2666,
	EnergyRecharge:   1000,
	ElementalMastery: 1000,
}

var gobletWeightedStats = map[artifactStat]int{
	HPP:              2125,
	ATKP:             2125,
	DEFP:             2000,
	PyroDMG:          500,
	ElectroDMG:       500,
	CryoDMG:          500,
	HydroDMG:         500,
	AnemoDMG:         500,
	GeoDMG:           500,
	PhysDMG:          500,
	ElementalMastery: 250,
}

var circletWeightedStats = map[artifactStat]int{
	HPP:              2200,
	ATKP:             2200,
	DEFP:             2200,
	CritRate:         1000,
	CritDmg:          1000,
	HealingBonus:     1000,
	ElementalMastery: 400,
}

func RandomArtifactOfSlot(slot artifactSlot) *Artifact {
	var artifact Artifact
	artifact.randomizeSet(allArtifactSets...)
	artifact.Slot = slot
	artifact.ranzomizeMainStat()
	artifact.randomizeSubstats()
	return &artifact
}

func weightedSubstats(mainStat artifactStat) map[artifactStat]int {
	weightedSubs := map[artifactStat]int{
		HP:               flatSubstatWeight,
		ATK:              flatSubstatWeight,
		DEF:              flatSubstatWeight,
		HPP:              commonSubstatWeight,
		ATKP:             commonSubstatWeight,
		DEFP:             commonSubstatWeight,
		EnergyRecharge:   commonSubstatWeight,
		ElementalMastery: commonSubstatWeight,
		CritRate:         critSubstatWeight,
		CritDmg:          critSubstatWeight,
	}
	delete(weightedSubs, mainStat)
	return weightedSubs
}

func GenerateRandomArtifacts() map[string]interface{} {
	return map[string]interface{}{
		"Flower":  FormatGenshinArtifact(RandomArtifactOfSlot(SlotFlower)),
		"Plume":   FormatGenshinArtifact(RandomArtifactOfSlot(SlotPlume)),
		"Sands":   FormatGenshinArtifact(RandomArtifactOfSlot(SlotSands)),
		"Goblet":  FormatGenshinArtifact(RandomArtifactOfSlot(SlotGoblet)),
		"Circlet": FormatGenshinArtifact(RandomArtifactOfSlot(SlotCirclet)),
	}
}
