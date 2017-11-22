package createcharacter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/patientplatypus/gorest/config"
)

func ReturnCharacterStart(charactername string, w http.ResponseWriter, r *http.Request) {
	returncharacter := ACharacter{
		CharactersName: charactername,
		CharacterData: CharacterStruct{
			UserInput: UserInput{
				Choices: []Choices{
					{ChoiceName: "character class",
						Options:          []string{"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"},
						PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterclass"}},
					{ChoiceName: "character race",
						Options:          []string{"Dwarf", "Elf", "Halfling", "Human", "Human", "Dragonborn", "Gnome", "Half Elf", "Half Orc", "Tiefling"},
						PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterrace"}},
					{ChoiceName: "character background",
						Options:          []string{"Acolyte", "Charlatan", "Criminal", "Entertainer", "Folk Hero", "Guild Artisan", "Hermit", "Noble", "Outlander", "Sage", "Sailor", "Soldier", "Urchin"},
						PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterrace"}},
				}}},
	}

	json.NewEncoder(w).Encode(returncharacter)
}

func RegisterCharacter(charactername string, w http.ResponseWriter, r *http.Request) {
	log.Print("inside register character")
	log.Print("Charactername:", charactername)

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished)
	<-finished
	log.Print("Continuing with the rest of the program")

	var cnt int
	config.DB.QueryRow(`select count(*) from usercharacters`).Scan(&cnt)
	log.Print("count: ", cnt)
	if cnt == 0 {
		log.Print("username is not in usercharacters adding now")

		usercharacter := UserCharacters{
			Character: []ACharacter{
				{CharactersName: charactername,
					CharacterData: CharacterStruct{
						UserInput: UserInput{
							Choices: []Choices{
								{ChoiceName: "character class",
									Options:          []string{"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"},
									PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterclass"}},
								{ChoiceName: "character race",
									Options:          []string{"Dwarf", "Elf", "Halfling", "Human", "Human", "Dragonborn", "Gnome", "Half Elf", "Half Orc", "Tiefling"},
									PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterrace"}},
								{ChoiceName: "character background",
									Options:          []string{"Acolyte", "Charlatan", "Criminal", "Entertainer", "Folk Hero", "Guild Artisan", "Hermit", "Noble", "Outlander", "Sage", "Sailor", "Soldier", "Urchin"},
									PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterrace"}},
							},
						}}},
			}}

		datacharacter, err := json.Marshal(usercharacter)
		if err != nil {
			fmt.Println("Oh no! There was an error:", err)
			return
		}

		config.DB.QueryRow("insert into usercharacters (usernames, characters) values ($1, $2);", Username, datacharacter)

		ReturnCharacterStart(charactername, w, r)

	} else {
		var databasecharacters UserCharacters
		var datarows []byte
		config.DB.QueryRow("SELECT characters FROM usercharacters where usernames = $1", Username).Scan(&datarows)
		json.Unmarshal(datarows, &databasecharacters)
		log.Print("Here are the usernames in the database:", databasecharacters)

		var namealreadyused bool

		for j := 0; j < len(databasecharacters.Character); j++ {
			if databasecharacters.Character[j].CharactersName == charactername {
				namealreadyused = true
			}
			if j == len(databasecharacters.Character) && namealreadyused != true {
				namealreadyused = false
			}
		}
		if namealreadyused != true {
			newcharacter := ACharacter{
				CharactersName: charactername,
				CharacterData: CharacterStruct{
					UserInput: UserInput{
						Choices: []Choices{
							{ChoiceName: "character class",
								Options:          []string{"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"},
								PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterclass"}},
							{ChoiceName: "character race",
								Options:          []string{"Dwarf", "Elf", "Halfling", "Human", "Human", "Dragonborn", "Gnome", "Half Elf", "Half Orc", "Tiefling"},
								PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterrace"}},
							{ChoiceName: "character background",
								Options:          []string{"Acolyte", "Charlatan", "Criminal", "Entertainer", "Folk Hero", "Guild Artisan", "Hermit", "Noble", "Outlander", "Sage", "Sailor", "Soldier", "Urchin"},
								PlacementAddress: []string{"character", "characterdata", "raceclassbackground", "characterrace"}},
						}}},
			}
			databasecharacters.Character = append(databasecharacters.Character, newcharacter)

			appendcharacter, err := json.Marshal(databasecharacters)
			if err != nil {
				fmt.Println("Oh no! There was an error:", err)
				return
			}

			config.DB.Query("UPDATE usercharacters SET characters=$1 where usernames=$2;", appendcharacter, Username)

			ReturnCharacterStart(charactername, w, r)

		} else {
			log.Print("Character name already in use!")
		}
	}
}
