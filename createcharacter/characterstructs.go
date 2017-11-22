package createcharacter

type UserCharacters struct {
	Character []ACharacter `json:"character"`
}

type ACharacter struct {
	CharactersName string          `json:"charactername"`
	CharacterData  CharacterStruct `json:"characterdata"`
}

type CharacterStruct struct {
	RaceClassBackground      RaceClassBackground           `json:"raceclassbackground"`
	CharacterNumbers         CharacterNumbers              `json:"characternumbers"`
	Equipment                []string                      `json:"equipment"`
	SkillsProficienciesSaves SkillsProficienciesAdvantages `json:"skillsproficienciesadvantages"`
	Spells                   []string                      `json:"spells"`
	Fluff                    Fluff                         `json:"fluff"`
	CreationStage            string                        `json:"creationstage"`
	UserInput                UserInput                     `json:"userinput"`
}

type UserInput struct {
	Choices []Choices `json:"choices"`
	Answers []Answers `json:"answers"`
}

type Choices struct {
	ChoiceName       string   `json:"choicename"`
	Options          []string `json:"options"`
	PlacementAddress []string `json:"placementaddress"`
	PickNumber       int      `json:"picknumber"`
}

type Answers struct {
	Answers []string `json:"answers"`
}

type Fluff struct{}

type RaceClassBackground struct {
	Race       string `json:"characterrace"`
	SubRace    string `json:"charactersubrace"`
	Class      string `json:"characerclass"`
	Background string `json:"characterbackground"`
}

type SkillsProficienciesAdvantages struct {
	Proficiencies Proficiencies `json:"proficiencies"`
	Advantages    []string      `json:"advantages"`
	Resistances   []string      `json:"resistances"`
}

type Proficiencies struct {
	Armor        []string `json:"armor"`
	Weapons      []string `json:"weapons"`
	Tools        []string `json:"tools"`
	SavingThrows []string `json:"savingthrows"`
	Skills       []string `json:"skills"`
}

type CharacterNumbers struct {
	Level         int           `json:"level"`
	XP            int           `json:"xp"`
	GP            int           `json:"gp"`
	HitPoints     int           `json:"hitpoints"`
	Speed         int           `json:"speed"`
	CharacterAge  int           `json:"characterage"`
	AbilityScores AbilityScores `json:"abilityscores"`
}

type AbilityScores struct {
	StrengthBase     int `json:"strengthbase"`
	DexterityBase    int `json:"dexteritybase"`
	ConstitutionBase int `json:"constitutionbase"`
	IntelligenceBase int `json:"intelligencebase"`
	WisdomBase       int `json:"wisdombase"`
	CharismaBase     int `json:"charismabase"`

	StrengthAdded     int `json:"strengthadded"`
	DexterityAdded    int `json:"dexterityadded"`
	ConstitutionAdded int `json:"constitutionadded"`
	IntelligenceAdded int `json:"intelligenceadded"`
	WisdomAdded       int `json:"wisdomadded"`
	CharismaAdded     int `json:"charismaadded"`

	StrengthTotal     int `json:"strengthtotal"`
	DexterityTotal    int `json:"dexteritytotal"`
	ConstitutionTotal int `json:"constitutiontotal"`
	IntelligenceTotal int `json:"intelligencetotal"`
	WisdomTotal       int `json:"wisdomtotal"`
	CharismaTotal     int `json:"charismatotal"`

	StrengthModifier     int `json:"strengthmodifier"`
	DexterityModifier    int `json:"dexteritymodifier"`
	ConstitutionModifier int `json:"constitutionmodifier"`
	IntelligenceModifier int `json:"intelligencemodifier"`
	WisdomModifier       int `json:"wisdommodifier"`
	CharismaModifier     int `json:"charismamodifier"`
}
