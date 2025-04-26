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
	// Get petition id
	vars := mux.Vars(r)
	idStr := vars["id"]

	petitionID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}

	// Get query params
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort") // "best" or "newest"

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	orderBy := "created_at DESC" // default newest
	if sortBy == "best" {
		orderBy = "(upvotes - downvotes) DESC"
	}

	query := fmt.Sprintf(`
		SELECT id, user_id, comment_text, upvotes, downvotes, created_at
		FROM petition_comments
		WHERE petition_id = ? AND deleted_at IS NULL
		ORDER BY %s
		LIMIT ? OFFSET ?
	`, orderBy)

	rows, err := database.DB.Query(query, petitionID, limit, offset)
	if err != nil {
		http.Error(w, "❌ Failed to fetch comments", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []models.PetitionComment

	for rows.Next() {
		var c models.PetitionComment
		err := rows.Scan(&c.ID, &c.UserID, &c.CommentText, &c.Upvotes, &c.Downvotes, &c.CreatedAt)
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

type VoteCommentInput struct {
	VoteType string `json:"vote_type"` // "upvote" or "downvote"
}

func VoteOnComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid comment ID", http.StatusBadRequest)
		return
	}

	var input VoteCommentInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	var field string
	if input.VoteType == "upvote" {
		field = "upvotes"
	} else if input.VoteType == "downvote" {
		field = "downvotes"
	} else {
		http.Error(w, "❌ Invalid vote type", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(`
		UPDATE petition_comments
		SET `+field+` = `+field+` + 1, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`, commentID)

	if err != nil {
		http.Error(w, "❌ Failed to vote on comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Voted successfully on comment!",
	})
}
