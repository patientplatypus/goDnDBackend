package dungeon

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/patientplatypus/gorest/config"
)

type RaceTypes struct {
	Name                 string           `json:"name"`
	AbilityScoreIncrease AbilityScoreType `json:"abilityscoreincrease"`
	Age                  []int            `json:"age"`
	Alignment            []string         `json:"alignment"`
	Size                 SizeType         `json:"size"`
	Speed                int              `json:"speed"`
	Languages            []string         `json:"languages"`
	Race                 RacesType        `json:"racetype"`
	SubRace              []SubracesType   `json:"subracetype"`
}

type RacesType struct {
	RaceName   string         `json:"racename"`
	RaceTraits []RacialTraits `json:"racialtraits"`
}

type SubracesType struct {
	SubraceName          string           `json:"subracename"`
	AbilityScoreIncrease AbilityScoreType `json:"abilityscoreincrease"`
	SubracialTraits      []RacialTraits   `json:"subracialtraits"`
}

type RacialTraits struct {
	TraitName        string   `json:"traitname"`
	TraitDescription string   `json:"traitdescription"`
	HasProficiency   []string `json:"hasproficiency"`
	HasResistance    []string `json:"hasresistance"`
	HasAdvantage     []string `json:"hasadvantage"`
}

type AbilityScoreType struct {
	Ability  []string `json:"ability"`
	Increase []int    `json:"increase"`
}

type SizeType struct {
	Height string `json:"height"`
	Weight string `json:"weight"`
	Size   string `json:"size"`
}

func DwarfSet(wg *sync.WaitGroup) {
	dwarf := RaceTypes{
		Name: "Dwarf",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Constitution"},
			Increase: []int{2},
		},
		Age:       []int{50, 350},
		Alignment: []string{"lawful", "good"},
		Size: SizeType{
			Height: "4-5 feet",
			Weight: "150 pounds",
			Size:   "Medium",
		},
		Speed:     25,
		Languages: []string{"Common", "Dwarvish"},
		Race: RacesType{
			RaceName: "Dwarf",
			RaceTraits: []RacialTraits{
				{TraitName: "Darkvision",
					TraitDescription: "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light"},
				{TraitName: "Tool Proficiency",
					TraitDescription: "Proficiency: smith's tools, brewer's supplies, mason's tools",
					HasProficiency:   []string{"smith's tools", "brewer's supplies", "mason's tools"}},
				{TraitName: "Dwarven Resilience",
					TraitDescription: "You have advantage on saving throws against poison, and you have resistance against poison damage",
					HasAdvantage:     []string{"poison"},
					HasResistance:    []string{"poison"}},
				{TraitName: "Dwarven Combat Training",
					TraitDescription: "Proficiency: battleaxe, handaxe, light hammer, warhammer",
					HasProficiency:   []string{"battleaxe", "handaxe", "light hammer", "warhammer"}},
				{TraitName: "Stonecunning",
					TraitDescription: "Whenever you make an Intelligence (History) check related to the origin of stonework, you are considered proficient in the History skill and add double your proficiency bonus to the check, instead of your normal proficiency bonus."},
			},
		},
		SubRace: []SubracesType{
			{SubraceName: "Hill Dwarf",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Wisdom"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Dwarven Toughness",
						TraitDescription: "hit point maximum 1 increases by 1 every level"},
				},
			},
			{SubraceName: "Mountain Dwarf",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Strength"},
					Increase: []int{2},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Dwarven Armor Training",
						TraitDescription: "Proficiency: light armor, medium armor",
						HasProficiency:   []string{"light armor", "medium armor"}},
				},
			},
		},
	}
	dataraces, err := json.Marshal(dwarf)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("dwarf", dataraces, wg)
}

func ElfSet(wg *sync.WaitGroup) {
	elf := RaceTypes{
		Name: "Elf",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Dexterity"},
			Increase: []int{2},
		},
		Age:       []int{100, 750},
		Alignment: []string{"good", "chaotic", "evil"},
		Size: SizeType{
			Height: "5-6 feet",
			Weight: "slender",
			Size:   "Medium",
		},
		Speed:     30,
		Languages: []string{"Common", "Elvish"},
		Race: RacesType{
			RaceName: "Elf",
			RaceTraits: []RacialTraits{
				{TraitName: "Darkvision",
					TraitDescription: "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light"},
				{TraitName: "Keen Senses",
					TraitDescription: "Proficiency: Perception",
					HasProficiency:   []string{"Perception"}},
				{TraitName: "Fey Ancestry",
					TraitDescription: "You have advantage on saving throws against being charmed, and magic can't put you to sleep",
					HasAdvantage:     []string{"charmed"}},
				{TraitName: "Trance",
					TraitDescription: "Resting only takes 4 hours."},
			},
		},
		SubRace: []SubracesType{
			{SubraceName: "High Elf",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Intelligence"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Elf Weapon Training",
						TraitDescription: "Proficiency: longsword, shortsword, shortbow, longbow",
						HasProficiency:   []string{"longsword", "shortsword", "shortbow", "longbow"}},
					{TraitName: "Cantrip",
						TraitDescription: "You know one cantrip of your choice from the wizard spell list. Intelligence is your spellcasting ability for it."},
					{TraitName: "Extra Language", TraitDescription: "You can speak, read, and write one extra language of your choice."},
				},
			},
			{SubraceName: "Wood Elf",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Wisdom"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Elf Weapon Training",
						TraitDescription: "Proficiency: longsword, shortsword, shortbow, longbow",
						HasProficiency:   []string{"longsword", "shortsword", "shortbow", "longbow"}},
					{TraitName: "Fleet of Foot",
						TraitDescription: "Speed: 35"},
					{TraitName: "Mask of the Wild",
						TraitDescription: "You can attempt to hide even when you are only lightly obscured by foliage, heavy rain, falling snow, mist, and other natural phenomena."},
				},
			},
			{SubraceName: "Dark Elf (Drow)",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Charisma"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Superior Darkvision",
						TraitDescription: "Your darkvision has a range of 120 feet"},
					{TraitName: "Sunlight Sensitivity",
						TraitDescription: "You have disadvantage on attack roles and on Wisdom (Perception) checks that rely on sight when you, the target of your attack, or whatever you are trying to percieve is in direk sunlight."},
					{TraitName: "Drow Magic",
						TraitDescription: "You know the dancing lights cantrip. When you reach 3rd level, you can cast the faerie fire spell once with this trait and regain the ability to do so when you finish a long rest. When you reach fifth level, you can cast the darkness spell once with this trait and regain the ability to do so when you finish a long rest. Charisma is your spell casting ability for these spells."},
					{TraitName: "Drow Weapon Training",
						TraitDescription: "Proficiency: rapiers, shortswords, hand crossbows",
						HasProficiency:   []string{"rapiers", "shortswords", "hand crossbows"}},
				},
			},
		},
	}
	dataraces, err := json.Marshal(elf)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("elf", dataraces, wg)
}

func HalflingSet(wg *sync.WaitGroup) {
	halfling := RaceTypes{
		Name: "Halfling",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Dexterity"},
			Increase: []int{2},
		},
		Age:       []int{20, 250},
		Alignment: []string{"lawful good"},
		Size: SizeType{
			Height: "3 feet",
			Weight: "40 pounds",
			Size:   "Small",
		},
		Speed:     25,
		Languages: []string{"Common", "Halfling"},
		Race: RacesType{
			RaceName: "Halfling",
			RaceTraits: []RacialTraits{
				{TraitName: "Lucky",
					TraitDescription: "When you roll a 1 on the d20 for an attack roll, ability check, or saving throw, you can reroll the die and must use the new roll"},
				{TraitName: "Brave",
					TraitDescription: "You have advantage on saving throws against being frightened",
					HasAdvantage:     []string{"frightened"}},
				{TraitName: "Halfling Nimbleness",
					TraitDescription: "You can move through the space of any creature that is of a size larger than yours"},
			},
		},
		SubRace: []SubracesType{
			{SubraceName: "Lightfoot",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Charisma"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Naturally Stealthy",
						TraitDescription: "You can attempt to hide even when you are obscured only by a creature that is at least one size larger than you."},
				},
			},
			{SubraceName: "Stout",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Constitution"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Stout Resilience",
						TraitDescription: "You have advantage on saving throws agains poison, and you have resistance against poison damage",
						HasAdvantage:     []string{"Poision"},
						HasResistance:    []string{"Poison"}},
				},
			},
		},
	}
	dataraces, err := json.Marshal(halfling)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("halfling", dataraces, wg)
}

func HumanSet(wg *sync.WaitGroup) {
	human := RaceTypes{
		Name: "Human",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"},
			Increase: []int{1, 1, 1, 1, 1, 1},
		},
		Age:       []int{18, 76},
		Alignment: []string{"none"},
		Size: SizeType{
			Height: "5-6 feet",
			Weight: "150 pounds",
			Size:   "Medium",
		},
		Speed:     30,
		Languages: []string{"Common", "1Extra"},
		Race: RacesType{
			RaceName: "Human",
		},
	}
	dataraces, err := json.Marshal(human)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("human", dataraces, wg)
}

func DragonbornSet(wg *sync.WaitGroup) {
	dragonborn := RaceTypes{
		Name: "Dragonborn",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Strength", "Charisma"},
			Increase: []int{2, 1},
		},
		Age:       []int{15, 80},
		Alignment: []string{"none"},
		Size: SizeType{
			Height: ">6 feet",
			Weight: "250 pounds",
			Size:   "Medium",
		},
		Speed:     30,
		Languages: []string{"Common", "Draconic"},
		Race: RacesType{
			RaceName: "Dragonborn",
			RaceTraits: []RacialTraits{
				{TraitName: "Breath Weapon",
					TraitDescription: "You can use your action to exhale destructive energy. Your draconic ancestry determines the size, shape, and damage type of the exhalation. When you use your breath weapon each creature in the area of the exhalation must make a saving thow, the type of which is determined by your draconic ancestry. The DC for this saving throw equals 8 + your Constitution modifier + your proficiency bonus. A creature takes 2d6 damage on a failed save, and half as much damage on a successful one. The damage increases to 3d6 at 6th level, 4d6 at 11th level, and 5d6 at 16th level."},
			},
		},
		SubRace: []SubracesType{
			{SubraceName: "Black",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "5 by 30 ft line (Dex save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Acid",
						HasResistance:    []string{"Acid"}},
				},
			},
			{SubraceName: "Blue",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "5 by 30 ft line (Dex save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Lightning",
						HasResistance:    []string{"Lightning"}},
				},
			},
			{SubraceName: "Brass",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "5 by 30 ft line (Dex save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Fire",
						HasResistance:    []string{"Fire"}},
				},
			},
			{SubraceName: "Bronze",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "5 by 30 ft line (Dex save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Lightening",
						HasResistance:    []string{"Lightening"}},
				},
			},
			{SubraceName: "Copper",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "5 by 30 ft line (Dex save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Acid",
						HasResistance:    []string{"Acid"}},
				},
			},
			{SubraceName: "Gold",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "15 ft cone (Dex Save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Fire",
						HasResistance:    []string{"Fire"}},
				},
			},
			{SubraceName: "Green",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "15 ft cone (Con Save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Poison",
						HasResistance:    []string{"Poison"}},
				},
			},
			{SubraceName: "Red",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "15 ft cone (Dex Save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Fire",
						HasResistance:    []string{"Fire"}},
				},
			},
			{SubraceName: "Silver",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "15 ft cone (Con Save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Cold",
						HasResistance:    []string{"Cold"}},
				},
			},
			{SubraceName: "White",
				SubracialTraits: []RacialTraits{
					{TraitName: "Breath Weapon",
						TraitDescription: "15 ft cone (Con Save)"},
					{TraitName: "Damage Type",
						TraitDescription: "Cold",
						HasResistance:    []string{"Cold"}},
				},
			},
		},
	}
	dataraces, err := json.Marshal(dragonborn)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("dragonborn", dataraces, wg)
}

func GnomeSet(wg *sync.WaitGroup) {
	gnome := RaceTypes{
		Name: "Gnome",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Intelligence"},
			Increase: []int{2},
		},
		Age:       []int{40, 500},
		Alignment: []string{"good", "lawful", "chaotic"},
		Size: SizeType{
			Height: "3-4 feet",
			Weight: "40 pounds",
			Size:   "Small",
		},
		Speed:     25,
		Languages: []string{"Common", "Gnomish"},
		Race: RacesType{
			RaceName: "Gnome",
			RaceTraits: []RacialTraits{
				{TraitName: "Darkvision",
					TraitDescription: "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim ligh"},
				{TraitName: "Gnome Cunning",
					TraitDescription: "You have advantage on all Intelligence, Wisdom, and Charisma saving throws against magic.",
					HasAdvantage:     []string{"Magic"}},
			},
		},
		SubRace: []SubracesType{
			{SubraceName: "Forest Gnome",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Dexterity"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Natural Illusionist",
						TraitDescription: "You know the natural illusionist cantrip. Intelligence is your spellcasting ability for it."},
					{TraitName: "Speak with Small Beasts",
						TraitDescription: "Through sounds and gestures you can communicate simple ideas with small or smaller beasts."},
				},
			},
			{SubraceName: "Rock Gnome",
				AbilityScoreIncrease: AbilityScoreType{
					Ability:  []string{"Constitution"},
					Increase: []int{1},
				},
				SubracialTraits: []RacialTraits{
					{TraitName: "Artificer's Lore",
						TraitDescription: "Whenever you make an Intelligence (History) check related to magic items, alchemical objects, or technological devices, you can add twice your proficiency bonus, instead of any proficiency bonus you normally apply"},
					{TraitName: "Tinker",
						TraitDescription: "You have proficiency with artisan’s tools (tinker’s tools). Using those tools, you can spend 1 hour and 10 gp worth of materials to construct a Tiny clockwork device (AC 5, 1 hp). The device ceases to function after 24 hours (unless you spend 1 hour repairing it to keep the device functioning), or when you use your action to dismantle it; at that time, you can reclaim the materials used to create it. You can have up to three such devices active at a time. When you create a device, choose one of the following options: Clockwork Toy: This toy is a clockwork animal, monster, or person, such as a frog, mouse, bird, dragon, or soldier. When placed on the ground, the toy moves 5 feet across the ground on each of your turns in a random direction. It makes noises as appropriate to the creature it represents. Fire Starter: The device produces a miniature flame, which you can use to light a Candle, torch, or campfire. Using the device requires your action. Music Box: When opened, this music box plays a single song at a moderate volume. The box stops playing when it reaches the song’s end or when it is closed."},
				},
			},
		},
	}
	dataraces, err := json.Marshal(gnome)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("gnome", dataraces, wg)
}

func HalfelfSet(wg *sync.WaitGroup) {
	halfelf := RaceTypes{
		Name: "Half Elf",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Charisma", "Other", "Other"},
			Increase: []int{2, 1, 1},
		},
		Age:       []int{20, 180},
		Alignment: []string{"chaotic"},
		Size: SizeType{
			Height: "5-6 feet",
			Weight: "150 pounds",
			Size:   "Medium",
		},
		Speed:     30,
		Languages: []string{"Common", "Elvish", "1Extra"},
		Race: RacesType{
			RaceName: "Half Elf",
			RaceTraits: []RacialTraits{
				{TraitName: "Fey Ancestry",
					TraitDescription: "You have advantage on saving throws against being charmed, and magic can't put you to sleep.",
					HasAdvantage:     []string{"Charmed"}},
				{TraitName: "Skill Versatility",
					TraitDescription: "Proficiency: 2 skills of your choice.",
					HasProficiency:   []string{"2Extra"}},
				{TraitName: "Darkvision",
					TraitDescription: "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light"},
			},
		},
	}
	dataraces, err := json.Marshal(halfelf)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("halfelf", dataraces, wg)
}

func HalforcSet(wg *sync.WaitGroup) {
	halforc := RaceTypes{
		Name: "Half Orc",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Constitution", "Strength"},
			Increase: []int{1, 2},
		},
		Age:       []int{14, 75},
		Alignment: []string{"chaotic", "evil"},
		Size: SizeType{
			Height: "5-6 feet",
			Weight: "200 pounds",
			Size:   "Medium",
		},
		Speed:     30,
		Languages: []string{"Common", "Orc"},
		Race: RacesType{
			RaceName: "Half Orc",
			RaceTraits: []RacialTraits{
				{TraitName: "Menacing",
					TraitDescription: "Proficiency: Intimidation"},
				{TraitName: "Relentless Endurance",
					TraitDescription: "When you are reduced to 0 hit points but not killed outright, you can drop to 1 hit point instead. You can't use this feature again until after you have finished a long rest."},
				{TraitName: "Savage Attacks",
					TraitDescription: "When you score a critical hit with a melee weapon attack, you can roll one of the weapon's damage dice one additional time and add it to the extra damage of the critical hit."},
				{TraitName: "Darkvision",
					TraitDescription: "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light"},
			},
		},
	}
	dataraces, err := json.Marshal(halforc)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("halforc", dataraces, wg)
}

func TieflingSet(wg *sync.WaitGroup) {
	tiefling := RaceTypes{
		Name: "Tiefling",
		AbilityScoreIncrease: AbilityScoreType{
			Ability:  []string{"Intelligence", "Charisma"},
			Increase: []int{1, 2},
		},
		Age:       []int{20, 110},
		Alignment: []string{"chaotic", "evil"},
		Size: SizeType{
			Height: "5-6 feet",
			Weight: "150 pounds",
			Size:   "Medium",
		},
		Speed:     30,
		Languages: []string{"Common", "Infernal"},
		Race: RacesType{
			RaceName: "Tiefling",
			RaceTraits: []RacialTraits{
				{TraitName: "Hellish Resistance",
					TraitDescription: "You have resistance to fire damage.",
					HasResistance:    []string{"Fire"}},
				{TraitName: "Infernal Legacy",
					TraitDescription: "You know the thaumaturgy cantrip. When you reach third level, you can cast the hellish rebuke spell as a 2nd level spell once with this trait and regain the ability to do so when you finish a long rest. When you reach 5th level you can cast the darkness spell once with this trait and regain the ability to do so when you finish a long rest. Charisma is your spellcasting ability for this spell."},
				{TraitName: "Darkvision",
					TraitDescription: "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light"},
			},
		},
	}
	dataraces, err := json.Marshal(tiefling)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckRace("tiefling", dataraces, wg)
}

func CheckRace(racesname string, dataraces []byte, wg *sync.WaitGroup) {
	log.Print("inside CheckRace() and the value of race is ", racesname)
	if err := config.DB.QueryRow("SELECT racesname FROM dungeon_races WHERE racesname=$1", racesname).Scan(&racesname); err == nil {
		log.Print("The chosen race already exists in dungeon_races!")
		log.Print("Value of all entries in dataracesbase is: ")
		rows, error := config.DB.Query("Select racesname, dataraces from dungeon_races")
		defer rows.Close()
		for rows.Next() {
			var racesname string
			var dataraces []byte
			error = rows.Scan(&racesname, &dataraces)
			if error != nil {
				panic(error)
			}
			var obj RaceTypes
			if err := json.Unmarshal(dataraces, &obj); err != nil {
				panic(err)
			}
			log.Print("racename: ", racesname, " obj ", obj)
		}
		wg.Done()
	} else if err.Error() == "sql: no rows in result set" {
		log.Print("value of err: ", err)
		log.Print("no rows found of class, inserting into db")
		config.DB.QueryRow("INSERT INTO dungeon_races VALUES ($1, $2);", racesname, dataraces)
		wg.Done()
	} else {
		log.Print("There must have been some other error: ", err)
		wg.Done()
	}
}

func RaceType() {
	log.Print("inside racetype of dungeon")
	var wg sync.WaitGroup

	wg.Add(9)

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished, "races")
	<-finished
	log.Print("Continuing with the rest of the program")

	//1
	go DwarfSet(&wg)
	//2
	go ElfSet(&wg)
	//3
	go HalflingSet(&wg)
	//4
	go HumanSet(&wg)
	//5
	go DragonbornSet(&wg)
	//6
	go GnomeSet(&wg)
	//7
	go HalfelfSet(&wg)
	//8
	go HalforcSet(&wg)
	//9
	go TieflingSet(&wg)

	wg.Wait()
	log.Print("Race Done")
}
