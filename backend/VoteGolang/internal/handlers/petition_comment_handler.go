package handlers

import (
	"VoteGolang/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"VoteGolang/internal/database"
)

type CreateCommentInput struct {
	PetitionID  int    `json:"petition_id"`
	CommentText string `json:"comment_text"`
}

func CreatePetitionComment(w http.ResponseWriter, r *http.Request) {
	var input CreateCommentInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	username := r.Context().Value("username").(string)

	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	}

	_, err = database.DB.Exec(`
		INSERT INTO petition_comments (petition_id, user_id, comment_text, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, input.PetitionID, userID, input.CommentText, time.Now(), time.Now())

	if err != nil {
		http.Error(w, "❌ Failed to add comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Comment added!"})
}
func ListPetitionComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	var petitionID int
	_, err := fmt.Sscanf(idStr, "%d", &petitionID)
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT id, user_id, comment_text, created_at
		FROM petition_comments
		WHERE petition_id = ? AND deleted_at IS NULL
		ORDER BY created_at DESC
	`, petitionID)

	if err != nil {
		http.Error(w, "❌ Failed to fetch comments", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []models.PetitionComment

	for rows.Next() {
		var c models.PetitionComment
		err := rows.Scan(&c.ID, &c.UserID, &c.CommentText, &c.CreatedAt)
		if err != nil {
			http.Error(w, "❌ Error reading comment", http.StatusInternalServerError)
			return
		}
		comments = append(comments, c)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func DeletePetitionComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid comment ID", http.StatusBadRequest)
		return
	}

	// Soft delete the comment
	_, err = database.DB.Exec(`
		UPDATE petition_comments
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`, id)

	if err != nil {
		http.Error(w, "❌ Failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Comment deleted successfully",
	})
}
