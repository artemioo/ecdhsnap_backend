package handler

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/gorilla/sessions"
)

var (
	key       = ([]byte(os.Getenv("SESSION_KEY")))
	hashedKey = md5.Sum([]byte(key))
	hkey      = hashedKey[:16]
	store     = sessions.NewCookieStore(hkey)
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	store = sessions.NewCookieStore(hkey)
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   4000,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}

	session, err := store.Get(r, "session-id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user ecdhsnap.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	session.Values["UserId"] = user.Id
	session.Values["authenticated"] = true

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookieHeader := w.Header()["Set-Cookie"][0]
	sessions_id := strings.Split(cookieHeader, ";")[0]

	sessionIdMap := map[string]string{"session_id": sessions_id[11:]}
	sessionIdJSON, err := json.Marshal(sessionIdMap)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	sessionIdStr := string(sessionIdJSON)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sessionIdStr))

}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	session, err_s := store.Get(r, "session-id")
	if err_s != nil {
		http.Error(w, "Internal Server Error: "+err_s.Error(), http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		log.Printf("Error saving session: %v", err)
	}
}
