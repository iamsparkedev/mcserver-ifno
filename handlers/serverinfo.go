package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	pinger "github.com/Raqbit/mc-pinger"
)

type ServerInfo struct {
	IP       string `json:"ip"`
	Version  string `json:"version"`
	Online   bool   `json:"online"`
	Players  struct {
		Online int `json:"online"`
		Max    int `json:"max"`
	} `json:"players"`
	MOTD string `json:"motd"`
}

func ServerInfoHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ip := query.Get("ip")
	if ip == "" {
		http.Error(w, "Missing 'ip' query parameter", http.StatusBadRequest)
		return
	}

	// Support optional port (host:port). Default Minecraft port is 25565.
	host := ip
	port := uint16(25565)
	if strings.Contains(ip, ":") {
		parts := strings.Split(ip, ":")
		host = parts[0]
		p, err := strconv.ParseUint(parts[1], 10, 16)
		if err == nil {
			port = uint16(p)
		}
	}

	p := pinger.New(host, port)
	info, err := p.Ping()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to ping server: %v", err), http.StatusInternalServerError)
		return
	}

	if info == nil {
		http.Error(w, "No info returned from ping", http.StatusInternalServerError)
		return
	}

	var serverInfo ServerInfo
	serverInfo.IP = ip
	serverInfo.Version = info.Version.Name
	serverInfo.Online = true
	serverInfo.Players.Online = int(info.Players.Online)
	serverInfo.Players.Max = int(info.Players.Max)
	serverInfo.MOTD = info.Description.Text

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serverInfo)

}