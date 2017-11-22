package config

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

var (
	// Store is the cookie store
	Store *sessions.CookieStore
	// Name is the session name
	Name string
)

// Session stores session level information
type Session struct {
	Options   sessions.Options `json:"Options"`   // Pulled from: http://www.gorillatoolkit.org/pkg/sessions#Options
	Name      string           `json:"Name"`      // Name for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.Get
	SecretKey string           `json:"SecretKey"` // Key for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.New
}

// Configure the session cookie store
func Configure_Sessions() {
	Store = sessions.NewCookieStore([]byte("adsads;fja;i4gna;nbeq09bjse"))
	Store.Options = &sessions.Options{
		Domain:   "",
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	// Store.Options = &s.Options
	// Name = s.Name
}

func Instance(r *http.Request) (*sessions.Session, error) {
	session, err := Store.Get(r, "cookies4dndyo")
	time.Sleep(time.Second * 2)
	return session, err
}

// Empty deletes all the current session values
func Empty(sess *sessions.Session) {
	// Clear out all stored values in the cookie
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}
