package dungeon

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/patientplatypus/gorest/config"
)

type BackgroundTypes struct {
	HasProficiency []string           `json:"hasproficincy"`
	Languages      []string           `json:"languages"`
	Equipment      []EquipmentChoices `json:"equipment"`
}

func AcolyteSet(wg *sync.WaitGroup) {
	acolyte := BackgroundTypes{
		HasProficiency: []string{"Insight", "Religion"},
		Languages:      []string{"2Extra"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"holy symbol"}},
			{ChoiceA: []string{"prayer book"},
				ChoiceB: []string{"prayer wheel"}},
			{ChoiceA: []string{"5 sticks of incense"}},
			{ChoiceA: []string{"vestments"}},
			{ChoiceA: []string{"set of common clothes"}},
			{ChoiceA: []string{"pouch", "15GP"}},
		},
	}
	databackgrounds, err := json.Marshal(acolyte)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("acolyte", databackgrounds, wg)
}

func CharlatanSet(wg *sync.WaitGroup) {
	charlatan := BackgroundTypes{
		HasProficiency: []string{"Deception", "Sleight of Hand", "Disguise Kit", "Forgery Kit"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"set of fine clothes"}},
			{ChoiceA: []string{"disguise kit"}},
			{ChoiceA: []string{"ten stoppered bottles filled with colored liquid"},
				ChoiceB: []string{"Set of weighted dice"},
				ChoiceC: []string{"deck of marked cards"},
				ChoiceD: []string{"signet rign of imaginary Duke"}},
			{ChoiceA: []string{"pouch", "15GP"}},
		},
	}
	databackgrounds, err := json.Marshal(charlatan)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("charlatan", databackgrounds, wg)
}

func CriminalSet(wg *sync.WaitGroup) {
	criminal := BackgroundTypes{
		HasProficiency: []string{"Deception", "Stealth", "One type of gaming set", "theives' tools"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"crowbar"}},
			{ChoiceA: []string{"dark common clothes", "hood"}},
			{ChoiceA: []string{"pouch", "15GP"}},
		},
	}
	databackgrounds, err := json.Marshal(criminal)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("criminal", databackgrounds, wg)
}

func EntertainerSet(wg *sync.WaitGroup) {
	entertainer := BackgroundTypes{
		HasProficiency: []string{"Acrobatics", "Performance", "One type of musical instrument", "Disguise kit"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"musical instrument"}},
			{ChoiceA: []string{"love letter from admirer"},
				ChoiceB: []string{"lock of hair from admirer"},
				ChoiceC: []string{"trinket from admirer"}},
			{ChoiceA: []string{"costume"}},
			{ChoiceA: []string{"pouch", "15GP"}},
		},
	}
	databackgrounds, err := json.Marshal(entertainer)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("entertainer", databackgrounds, wg)
}

func FolkheroSet(wg *sync.WaitGroup) {
	folkhero := BackgroundTypes{
		HasProficiency: []string{"Animal Handling", "Survival", "One type of artisan's tools", "vehicles (land)"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"set of artisan's tools (your choice)"}},
			{ChoiceA: []string{"shovel"}},
			{ChoiceA: []string{"iron pot"}},
			{ChoiceA: []string{"common clothes"}},
			{ChoiceA: []string{"pouch", "15GP"}},
		},
	}
	databackgrounds, err := json.Marshal(folkhero)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("folkhero", databackgrounds, wg)
}

func GuildartisanSet(wg *sync.WaitGroup) {
	guildartisan := BackgroundTypes{
		HasProficiency: []string{"Insight", "Persuasion", "one type of artisan's tools"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"set of artisan's tools (your choice)"}},
			{ChoiceA: []string{"letter of introduction from your guild"}},
			{ChoiceA: []string{"set of traveller's clothes"}},
			{ChoiceA: []string{"pouch", "15GP"}},
		},
	}
	databackgrounds, err := json.Marshal(guildartisan)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("guildartisan", databackgrounds, wg)
}

func HermitSet(wg *sync.WaitGroup) {
	hermit := BackgroundTypes{
		HasProficiency: []string{"Medicine", "Religion", "Herbalism Kit"},
		Languages:      []string{"1Extra"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"scroll case stuffed full of notes from your studies of prayers"}},
			{ChoiceA: []string{"winter blanket"}},
			{ChoiceA: []string{"set of common clothes"}},
			{ChoiceA: []string{"herbalism kit"}},
			{ChoiceA: []string{"5GP"}},
		},
	}
	databackgrounds, err := json.Marshal(hermit)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("hermit", databackgrounds, wg)
}

func NobleSet(wg *sync.WaitGroup) {
	hermit := BackgroundTypes{
		HasProficiency: []string{"History", "Persuasion", "one type of gaming set"},
		Languages:      []string{"1Extra"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"set of fine clothes"}},
			{ChoiceA: []string{"signet ring"}},
			{ChoiceA: []string{"scroll of pedigree"}},
			{ChoiceA: []string{"purse", "25GP"}},
		},
	}
	databackgrounds, err := json.Marshal(hermit)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("hermit", databackgrounds, wg)
}

func OutlanderSet(wg *sync.WaitGroup) {
	outlander := BackgroundTypes{
		HasProficiency: []string{"Athletics", "Survival", "one type of musical instrument"},
		Languages:      []string{"1Extra"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"staff"}},
			{ChoiceA: []string{"hunting trap"}},
			{ChoiceA: []string{"trophy from an animal you killed"}},
			{ChoiceA: []string{"set of traveller's clothes"}},
			{ChoiceA: []string{"pouch", "10GP"}},
		},
	}
	databackgrounds, err := json.Marshal(outlander)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("outlander", databackgrounds, wg)
}

func SageSet(wg *sync.WaitGroup) {
	sage := BackgroundTypes{
		HasProficiency: []string{"Arcana", "History"},
		Languages:      []string{"2Extra"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"bottle of black ink"}},
			{ChoiceA: []string{"quill"}},
			{ChoiceA: []string{"small knife"}},
			{ChoiceA: []string{"a letter from a dead colleague posing a question you have not yet been able to answer"}},
			{ChoiceA: []string{"a set of common clothes"}},
			{ChoiceA: []string{"pouch", "10GP"}},
		},
	}
	databackgrounds, err := json.Marshal(sage)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("sage", databackgrounds, wg)
}

func SailorSet(wg *sync.WaitGroup) {
	sailor := BackgroundTypes{
		HasProficiency: []string{"Athletics", "Perception", "Navigator's tools", "vehicle's (water)"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"belaying pin (club)"}},
			{ChoiceA: []string{"50 feet of silk rope"}},
			{ChoiceA: []string{"lucky charm"}},
			{ChoiceA: []string{"set of common clothes"}},
			{ChoiceA: []string{"pouch", "10GP"}},
		},
	}
	databackgrounds, err := json.Marshal(sailor)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("sailor", databackgrounds, wg)
}

func SoldierSet(wg *sync.WaitGroup) {
	soldier := BackgroundTypes{
		HasProficiency: []string{"Athletics", "Intimidation", "one type of gaming set", "vehicle's (land)"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"insignia of rank"}},
			{ChoiceA: []string{"dagger from fallen enemy"},
				ChoiceB: []string{"broken blade from fallen enemy"},
				ChoiceC: []string{"piece of banner from fallen enemy"}},
			{ChoiceA: []string{"set of bone dice"},
				ChoiceB: []string{"deck of cards"}},
			{ChoiceA: []string{"set of common clothes"}},
			{ChoiceA: []string{"pouch", "10GP"}},
		},
	}
	databackgrounds, err := json.Marshal(soldier)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("soldier", databackgrounds, wg)
}

func UrchinSet(wg *sync.WaitGroup) {
	urchin := BackgroundTypes{
		HasProficiency: []string{"Sleight of Hand", "Stealth", "Disguise kit", "thieves' tools"},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"small knife"}},
			{ChoiceA: []string{"map of the city you grew up in"}},
			{ChoiceA: []string{"pet mouse"}},
			{ChoiceA: []string{"a token to remember your parents by"}},
			{ChoiceA: []string{"set of common clothes"}},
			{ChoiceA: []string{"pouch", "10GP"}},
		},
	}
	databackgrounds, err := json.Marshal(urchin)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckBackground("urchin", databackgrounds, wg)
}

func CheckBackground(backgroundsname string, databackgrounds []byte, wg *sync.WaitGroup) {
	log.Print("inside CheckBackground() and the value of background is ", backgroundsname)
	if err := config.DB.QueryRow("SELECT backgroundsname FROM dungeon_backgrounds WHERE backgroundsname=$1", backgroundsname).Scan(&backgroundsname); err == nil {
		log.Print("The chosen background already exists in dungeon_races!")
		log.Print("Value of all entries in databackgroundsbase is: ")
		rows, error := config.DB.Query("Select backgroundsname, databackgrounds from dungeon_backgrounds")
		defer rows.Close()
		for rows.Next() {
			var backgroundsname string
			var databackgrounds []byte
			error = rows.Scan(&backgroundsname, &databackgrounds)
			if error != nil {
				panic(error)
			}
			var obj BackgroundTypes
			if err := json.Unmarshal(databackgrounds, &obj); err != nil {
				panic(err)
			}
			log.Print("backgroundsname: ", backgroundsname, " obj ", obj)
		}
		wg.Done()
	} else if err.Error() == "sql: no rows in result set" {
		log.Print("value of err: ", err)
		log.Print("no rows found of class, inserting into db")
		config.DB.QueryRow("INSERT INTO dungeon_backgrounds VALUES ($1, $2);", backgroundsname, databackgrounds)
		wg.Done()
	} else {
		log.Print("There must have been some other error: ", err)
		wg.Done()
	}
}

func BackgroundType() {
	log.Print("inside background type of dungeon")
	var wg sync.WaitGroup
	wg.Add(13)

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished, "backgrounds")
	<-finished
	log.Print("Continuing with the rest of the program")

	//1
	go AcolyteSet(&wg)
	//2
	go CharlatanSet(&wg)
	//3
	go CriminalSet(&wg)
	//4
	go EntertainerSet(&wg)
	//5
	go FolkheroSet(&wg)
	//6
	go GuildartisanSet(&wg)
	//7
	go HermitSet(&wg)
	//8
	go NobleSet(&wg)
	//9
	go OutlanderSet(&wg)
	//10
	go SageSet(&wg)
	//11
	go SailorSet(&wg)
	//12
	go SoldierSet(&wg)
	//13
	go UrchinSet(&wg)

	wg.Wait()
	log.Print("Background Done")
}
