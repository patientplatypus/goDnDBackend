package main

import (
	"net/http"

	"github.com/patientplatypus/gorest/users"

	"github.com/patientplatypus/gorest/dungeon_db"

	"github.com/patientplatypus/gorest/character"

	"github.com/patientplatypus/gorest/createcharacter"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"ClassType",
		"POST",
		"/character/class",
		character.ClassType,
	},
	Route{
		"RaceType",
		"POST",
		"/character/race",
		character.RaceType,
	},
	Route{
		"BackgroundType",
		"POST",
		"/character/background",
		character.BackgroundType,
	},
	Route{
		"RegisterUser",
		"POST",
		"/users/registeruser",
		users.RegisterUser,
	},
	Route{
		"UserLogin",
		"POST",
		"/users/userlogin",
		users.UserLogin,
	},
	Route{
		"UserLogout",
		"POST",
		"/users/userlogout",
		users.UserLogout,
	},
	Route{
		"DeleteAll",
		"GET",
		"/users/deleteallusers",
		users.DeleteAllUsers,
	},
	Route{
		"QueryAll",
		"GET",
		"/users/queryallusers",
		users.QueryAllUsers,
	},
	Route{
		"DeleteAllDungeonTypes",
		"GET",
		"/populate/deletealldungeontypes",
		dungeon.DeleteAllDungeonTypes,
	},
	Route{
		"PopulateDB",
		"GET",
		"/populate/populatedb",
		dungeon.PopulateDB,
	},
	Route{
		"NewCharacter",
		"POST",
		"/createcharacter/newcharacter",
		createcharacter.NewCharacter,
	},
	Route{
		"SavedCharacters",
		"GET",
		"/createcharacter/savedcharacters",
		createcharacter.SavedCharacters,
	},
	Route{
		"DeleteAllUsersCharacters",
		"GET",
		"/createcharacter/deletealluserscharacters",
		createcharacter.DeleteAllUsersCharacters,
	},
}
