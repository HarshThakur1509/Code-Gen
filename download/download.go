package download

import (
	"Code_Gen/global"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	repoURL := "https://github.com/HarshThakur1509/gin-gorm-api/archive/refs/heads/master.zip"
	resp, err := http.Get(repoURL)
	if err != nil {
		http.Error(w, "Failed to fetch repository", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Set headers to prompt download
	w.Header().Set("Content-Disposition", "attachment; filename=Boilerplate.zip")
	w.Header().Set("Content-Type", "application/zip")
	w.WriteHeader(http.StatusOK)

	// Copy the response body to the writer
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Failed to send repository", http.StatusInternalServerError)
		return
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Path string
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to Read Body", http.StatusBadRequest)
		return
	}
	path := body.Path
	global.PATH = path
	fmt.Println(global.PATH)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{"path": path})
}
