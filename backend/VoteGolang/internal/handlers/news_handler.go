package handlers

import (
	"VoteGolang/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"VoteGolang/internal/database"
)

func ListNews(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, title, paragraph, photo, created_at
		FROM general_news
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch news", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var newsList []models.News

	for rows.Next() {
		var n models.News
		if err := rows.Scan(&n.ID, &n.Title, &n.Paragraph, &n.Photo, &n.CreatedAt); err != nil {
			http.Error(w, "❌ Error reading news", http.StatusInternalServerError)
			return
		}
		newsList = append(newsList, n)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newsList)
}

type CreateNewsInput struct {
	Title     string `json:"title"`
	Paragraph string `json:"paragraph"`
	Photo     string `json:"photo"`
}

func CreateNews(w http.ResponseWriter, r *http.Request) {
	var input CreateNewsInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO general_news (title, paragraph, photo, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, input.Title, input.Paragraph, input.Photo, time.Now(), time.Now())

	if err != nil {
		http.Error(w, "❌ Failed to create news", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ News article created!",
	})
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid news ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(`
		UPDATE general_news
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`, id)

	if err != nil {
		http.Error(w, "❌ Failed to delete news", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ News deleted successfully",
	})
}
