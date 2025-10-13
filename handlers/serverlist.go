package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	ServerID    string   `json:"server_id"`
	Name        string   `json:"name"`
	IP          string   `json:"ip_address"`
	Gamemodes  []string `json:"gamemodes"`
	Description string   `json:"description"`
}

func ServerListHandler(w http.ResponseWriter, r *http.Request) {
     
	idparam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	file, err := os.ReadFile("data/data.json") // read the JSON file
	if err != nil {
		http.Error(w, "Could not open data file", http.StatusInternalServerError)
		return
	}

	// Strip UTF-8 BOM if present ("\xef\xbb\xbf") which breaks encoding/json
	file = bytes.TrimPrefix(file, []byte{0xEF, 0xBB, 0xBF})

	var servers []Server
	if err = json.Unmarshal(file, &servers); err != nil {
		// Try a common alternative layout: an object wrapping the array, e.g. {"servers": [...]}.
		var wrapper struct{
			Servers []Server `json:"servers"`
		}
		if werr := json.Unmarshal(file, &wrapper); werr == nil && len(wrapper.Servers) > 0 {
			servers = wrapper.Servers
		} else {
			// Return the actual JSON error to aid debugging (developer-friendly). Include a short sample of the file.
			sample := string(file)
			if len(sample) > 400 {
				sample = sample[:400] + "..."
			}
			http.Error(w, fmt.Sprintf("Could not parse data file: %v\nfile sample: %s", err, sample), http.StatusInternalServerError)
			return
		}
	}
    
	var found *Server
	for _, server := range servers {
		if server.ServerID == strconv.Itoa(id) {
			found = &server
			break
		}
	}
	if found == nil {
		http.Error(w, "Server not found", http.StatusNotFound)
		return
	}

	// Set a session cookie (example values used here, adjust as needed)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "your-session-value",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // set true if using HTTPS
		MaxAge:   3600,  // seconds
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(found)

	
}