package handlers

import (
	"encoding/json"
	"mood-api/store"
	"net/http"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

func MoodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		handleGetMood(w)
	case "POST":
		handlePostMood(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetMood(w http.ResponseWriter) {
	currentDate := time.Now().Format("2006-01-02")

	moodDetails := store.CalculateMoodDetailsForDate(currentDate)
	json.NewEncoder(w).Encode(moodDetails)

}

func handlePostMood(w http.ResponseWriter, r *http.Request) {
	var moodEntry store.MoodEntry

	if err := json.NewDecoder(r.Body).Decode(&moodEntry); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	store.SaveMoodData(moodEntry.Mood)
	w.WriteHeader(http.StatusCreated)
}
