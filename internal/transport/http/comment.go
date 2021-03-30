package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ksupdev/updev-go-rest-api-course/internal/comment"
)

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// fmt.Fprintf(w, "Unable to parse UINT from ID")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		// fmt.Fprintf(w, "Error Retrieving Comment By ID")
		sendErrorResponse(w, "Error Retrieving Comment By ID", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
	// fmt.Fprintf(w, "%+v", comment)

}

// GetAllComments - retrieve all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()
	if err != nil {
		// fmt.Fprintf(w, "Failed to retrieve all comments")
		sendErrorResponse(w, "Failed to retrieve all comments", err)
		return
	}
	// fmt.Fprintf(w, "%+v", comments)
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}

}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		// fmt.Fprintf(w, "Failed to decode JSON Body")
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		// fmt.Fprintf(w, "Failed to post new comment")
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

	// fmt.Fprintf(w, "%+v", comment)
}

// UpdateComment - update a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		// fmt.Fprintf(w, "Failed to decode JSON Body")
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)

	comment, err = h.Service.UpdateComment(uint(commentID), comment)

	if err != nil {
		// fmt.Fprintf(w, "Faild to update comment")
		sendErrorResponse(w, "Faild to update comment", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

	// fmt.Fprintf(w, "%+v", comment)

}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// fmt.Fprintf(w, "Unable to parse UINT from ID")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	// fmt.Println("--- track --- ", i)

	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		// fmt.Fprintf(w, "Failed to delete comment by comment id")
		sendErrorResponse(w, "Failed to delete comment by comment id", err)
		return
	}
	// fmt.Fprintf(w, "Successfully deleted comment")

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfully deleted"}); err != nil {
		panic(err)
	}
}
