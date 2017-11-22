package dungeon

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/patientplatypus/gorest/config"
)

type ClassFeatures struct {
	BaseHitPoints int                `json:"basehitpoints"`
	Proficiencies ProficiencyTypes   `json:"array"`
	Equipment     []EquipmentChoices `json:"equipment"`
}

type ProficiencyTypes struct {
	Armor           []string `json:"armor"`
	Weapons         []string `json:"weapons"`
	Tools           []string `json:"tools"`
	SavingThrows    []string `json:"savingthrows"`
	Skills          []string `json:"skills"`
	BaseSkillNumber int      `json:"baseskillnumber"`
}

type EquipmentChoices struct {
	ChoiceA []string `json:"choicea"`
	ChoiceB []string `json:"choiceb"`
	ChoiceC []string `json:"choicec"`
	ChoiceD []string `json:"choiced"`
}

//1
func BarbarianSet(wg *sync.WaitGroup) {
	log.Print("Now inside BarbarianSet")
	barbarian := ClassFeatures{
		BaseHitPoints: 12,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor", "medium armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Strength", "Constitution"},
			Skills:          []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"greataxe"},
				ChoiceB: []string{"any martial melee weapon"}},
			{ChoiceA: []string{"two handaxes"},
				ChoiceB: []string{"any simple weapon"}},
			{ChoiceA: []string{"an explorer's pack", "four javelins"}},
		}}
	dataclasses, err := json.Marshal(barbarian)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("barbarian", dataclasses, wg)
}

//2
func BardSet(wg *sync.WaitGroup) {
	log.Print("Now inside BardSet")
	bard := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor"},
			Weapons:         []string{"simple weapons", "hand crossbows", "longswords", "rapiers", "shortswords"},
			Tools:           []string{"three musical instruments of your choice"},
			SavingThrows:    []string{"Dexterity", "Charisma"},
			Skills:          []string{"Athletics", "Acrobatics", "Sleight of Hand", "Stealth", "Arcana", "History", "Investigation", "Nature", "Religion", "Animal Handling", "Insight", "Medicine", "Perception", "Survival", "Deception", "Intimidation", "Performance", "Persuasion"},
			BaseSkillNumber: 3},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"rapier"},
				ChoiceB: []string{"longsword"},
				ChoiceC: []string{"any simple weapon"}},
			{ChoiceA: []string{"diplomat's pack"},
				ChoiceB: []string{"entertainer's pack"}},
			{ChoiceA: []string{"lute"},
				ChoiceB: []string{"any musical instrument"}},
			{ChoiceA: []string{"leather armor", "dagger"}},
		}}
	dataclasses, err := json.Marshal(bard)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("bard", dataclasses, wg)
}

//3
func ClericSet(wg *sync.WaitGroup) {
	log.Print("Now inside ClericSet")
	cleric := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor", "medium armor", "shields"},
			Weapons:         []string{"simple weapons"},
			Tools:           []string{"None"},
			SavingThrows:    []string{"Wisdom", "Charisma"},
			Skills:          []string{"History", "Insight", "Medicine", "Persuasion", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"mace"},
				ChoiceB: []string{"warhammer (if proficient)"}},
			{ChoiceA: []string{"scale male"},
				ChoiceB: []string{"leather armor"},
				ChoiceC: []string{"chain male (if proficient)"}},
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"any simple weapon"}},
			{ChoiceA: []string{"priest's pack", "explorer's pack"}},
			{ChoiceA: []string{"shield", "holy symbol"}},
		}}
	dataclasses, err := json.Marshal(cleric)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("cleric", dataclasses, wg)
}

//4
func DruidSet(wg *sync.WaitGroup) {
	log.Print("Now inside DruidSet")
	druid := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor (no metal)", "medium armor (no metal)", "shields (no metal)"},
			Weapons:         []string{"clubs", "daggers", "darts", "javelins", "maces", "quarterstaffs", "scimitars", "sickles", "slings", "spears"},
			Tools:           []string{"Herbalism kit"},
			SavingThrows:    []string{"Wisdom", "Intelligence"},
			Skills:          []string{"Arcana", "Animal Handling", "Insight", "Medicine", "Nature", "Perception", "Religion", "Survival"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"wooden shield"},
				ChoiceB: []string{"any simple weapon"}},
			{ChoiceA: []string{"scimitar"},
				ChoiceB: []string{"any simple melee weapon"}},
			{ChoiceA: []string{"leather armor", "explorer's pack", "druidic focus"}},
		}}
	dataclasses, err := json.Marshal(druid)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("druid", dataclasses, wg)
}

//5
func FighterSet(wg *sync.WaitGroup) {
	log.Print("Now inside FighterSet")
	fighter := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"all armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Strength", "Constitution"},
			Skills:          []string{"Acrobatics", "Animal Handling", "Athletics", "History", "Insight", "Intimidation", "Perception", "Survival"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"chain mail"},
				ChoiceB: []string{"leather armor", "longbow", "20 arrows"}},
			{ChoiceA: []string{"any martial weapon", "shield"},
				ChoiceB: []string{"two martial weapons"}},
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"two handaxes"}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"}},
		}}
	dataclasses, err := json.Marshal(fighter)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("fighter", dataclasses, wg)
}

//6
func MonkSet(wg *sync.WaitGroup) {
	log.Print("Now inside MonkSet")
	monk := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"none"},
			Weapons:         []string{"simple weapons", "shortswords"},
			Tools:           []string{"one type of artisan's tools or one musical instrument"},
			SavingThrows:    []string{"Strength", "Dexterity"},
			Skills:          []string{"Acrobatics", "Athletics", "History", "Insight", "Religion", "Stealth"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"shortsword"},
				ChoiceB: []string{"any simple weapon"}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"}},
			{ChoiceA: []string{"10 darts"}},
		}}
	dataclasses, err := json.Marshal(monk)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("monk", dataclasses, wg)
}

//7
func PaladinSet(wg *sync.WaitGroup) {
	log.Print("Now inside PaladinSet")
	paladin := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"all armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Wisdom", "Charisma"},
			Skills:          []string{"Athletics", "Insight", "Intimidation", "Medicine", "Persuasion", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"martial weapon", "shield"},
				ChoiceB: []string{"two martial weapons"}},
			{ChoiceA: []string{"five javelins"},
				ChoiceB: []string{"any simple melee weapon"}},
			{ChoiceA: []string{"priest's pack"},
				ChoiceB: []string{"explorer's pack"}},
			{ChoiceA: []string{"chain mail", "holy symbol"}},
		}}
	dataclasses, err := json.Marshal(paladin)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("paladin", dataclasses, wg)
}

//8
func RangerSet(wg *sync.WaitGroup) {
	log.Print("Now inside RangerSet")
	ranger := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor", "medium armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Strength", "Dexterity"},
			Skills:          []string{"Animal Handling", "Athletics", "Insight", "Investigation", "Nature", "Perception", "Stealth", "Survival"},
			BaseSkillNumber: 3},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"scale mail"},
				ChoiceB: []string{"leather armor"}},
			{ChoiceA: []string{"two shortswords"},
				ChoiceB: []string{"two simple melee weapon"}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"}},
			{ChoiceA: []string{"longbow", "20 arrows"}},
		}}
	dataclasses, err := json.Marshal(ranger)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("ranger", dataclasses, wg)
}

//9
func RogueSet(wg *sync.WaitGroup) {
	log.Print("Now inside RogueSet")
	rogue := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor"},
			Weapons:         []string{"simple weapons", "hand crossbows", "longswords", "rapiers", "shortswords"},
			Tools:           []string{"thieves' tools"},
			SavingThrows:    []string{"Intelligence", "Dexterity"},
			Skills:          []string{"Acrobatics", "Athletics", "Deception", "Insight", "Intimidation", "Investigation", "Perception", "Performance", "Persuasion", "Sleight of Hand", "Stealth"},
			BaseSkillNumber: 4},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"rapier"},
				ChoiceB: []string{"shortsword"}},
			{ChoiceA: []string{"shortbow", "20 arrows"},
				ChoiceB: []string{"shortsword"}},
			{ChoiceA: []string{"burglar's pack"},
				ChoiceB: []string{"dungeoneer's pack"},
				ChoiceC: []string{"explorer's pack"}},
			{ChoiceA: []string{"leather armor", "two daggers", "thieves' tools"}},
		}}
	dataclasses, err := json.Marshal(rogue)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("rogue", dataclasses, wg)
}

//10
func SorcererSet(wg *sync.WaitGroup) {
	log.Print("Now inside SorcererSet")
	sorcerer := ClassFeatures{
		BaseHitPoints: 6,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"none"},
			Weapons:         []string{"daggers", "darts", "slings", "quarterstaffs", "light crossbows"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Charisma", "Constitution"},
			Skills:          []string{"Arcana", "Deception", "Insight", "Intimidation", "Persuasion", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"any simple weapon"}},
			{ChoiceA: []string{"component pouch"},
				ChoiceB: []string{"arcane focus"}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"}},
			{ChoiceA: []string{"two daggers"}},
		}}
	dataclasses, err := json.Marshal(sorcerer)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("sorcerer", dataclasses, wg)
}

//11
func WarlockSet(wg *sync.WaitGroup) {
	log.Print("Now inside WarlockSet")
	warlock := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor"},
			Weapons:         []string{"simple weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Charisma", "Wisdom"},
			Skills:          []string{"Arcana", "Deception", "History", "Intimidation", "Investigation", "Nature", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"any simple weapon"}},
			{ChoiceA: []string{"component pouch"},
				ChoiceB: []string{"arcane focus"}},
			{ChoiceA: []string{"scholar's pack"},
				ChoiceB: []string{"dungeoneer's pack"}},
			{ChoiceA: []string{"leather armor", "any simple weapon", "two daggers"}},
		}}
	dataclasses, err := json.Marshal(warlock)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("warlock", dataclasses, wg)
}

//12
func WizardSet(wg *sync.WaitGroup) {
	log.Print("Now inside WizardSet")
	wizard := ClassFeatures{
		BaseHitPoints: 6,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"none"},
			Weapons:         []string{"daggers", "darts", "slings", "quarterstaffs", "light crossbows"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Intelligence", "Wisdom"},
			Skills:          []string{"Arcana", "History", "Insight", "Investigation", "Investigation", "Medicine", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"quarterstaff"},
				ChoiceB: []string{"dagger"}},
			{ChoiceA: []string{"component pouch"},
				ChoiceB: []string{"arcane focus"}},
			{ChoiceA: []string{"scholar's pack"},
				ChoiceB: []string{"explorer's pack"}},
			{ChoiceA: []string{"spellbook"}},
		}}
	dataclasses, err := json.Marshal(wizard)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("wizard", dataclasses, wg)
}

func CheckClass(classesname string, dataclasses []byte, wg *sync.WaitGroup) {
	if err := config.DB.QueryRow("SELECT classesname FROM dungeon_classes WHERE classesname=$1", classesname).Scan(&classesname); err == nil {
		//found name
		log.Print("The chosen class already exists in dungeon_classes!")
		log.Print("Value of all entries in dataclassesbase is: ")
		rows, error := config.DB.Query("Select classesname, dataclasses from dungeon_classes")
		defer rows.Close()
		for rows.Next() {
			var classesname string
			var dataclasses []byte
			error = rows.Scan(&classesname, &dataclasses)
			if error != nil {
				panic(error)
			}
			var obj ClassFeatures
			if err := json.Unmarshal(dataclasses, &obj); err != nil {
				panic(err)
			}
			// log.Print("classname: ", classname, " obj ", obj)
			// fmt.Printf("classname: %+v\n", classname, " obj: %+v\n", obj)
			// fmt.Println("classname: ", classname, " obj ", obj)
			fmt.Printf("%+v\n", obj)
		}
		wg.Done()
	} else if err.Error() == "sql: no rows in result set" {
		log.Print("value of err: ", err)
		log.Print("no rows found of class, inserting into db")
		// log.Print("right before value of barbariandataclasses")
		// log.Print("value of barbariandataclasses: ", barbariandataclasses)
		config.DB.QueryRow("INSERT INTO dungeon_classes VALUES ($1, $2);", classesname, dataclasses)
		wg.Done()
	}
}

func ClassType() {
	log.Print("inside classtype of dungeon")

	var wg sync.WaitGroup
	wg.Add(12)

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished, "classes")
	<-finished
	log.Print("Continuing with the rest of the program")

	//1
	go BarbarianSet(&wg)
	log.Print("Back inside ClassType after 1")
	//2
	go BardSet(&wg)
	log.Print("Back inside ClassType after 2")
	//3
	go ClericSet(&wg)
	log.Print("Back inside ClassType after 3")
	//4
	go DruidSet(&wg)
	log.Print("Back inside ClassType after 4")
	//5
	go FighterSet(&wg)
	log.Print("Back inside ClassType after 5")
	//6
	go MonkSet(&wg)
	log.Print("Back inside ClassType after 6")
	//7
	go PaladinSet(&wg)
	log.Print("Back inside ClassType after 7")
	//8
	go RangerSet(&wg)
	log.Print("Back inside ClassType after 8")
	//9
	go RogueSet(&wg)
	log.Print("Back inside ClassType after 9")
	//10
	go SorcererSet(&wg)
	log.Print("Back inside ClassType after 10")
	//11
	go WarlockSet(&wg)
	log.Print("Back inside ClassType after 11")
	//12
	go WizardSet(&wg)
	log.Print("Back inside ClassType after 12")

	wg.Wait()
	log.Print("Class Done")

}
