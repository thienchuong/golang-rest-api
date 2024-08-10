package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thienchuong/golang-rest-api/db"
	"github.com/thienchuong/golang-rest-api/log"
	"github.com/thienchuong/golang-rest-api/models"
)

type handler struct {
	db  db.IDatabase
	log log.ILogger
}

// NewHandler creates a new handler
func NewHandler(db db.IDatabase, log log.ILogger) IHandler {
	return &handler{
		db:  db,
		log: log,
	}
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.db.GetAllBooks()
	if err != nil {
		h.log.Error(err, "Error getting all books")
		respondWithError(w, http.StatusInternalServerError, "Error getting all books")
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func (h *handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		h.log.Error(err, "Invalid book ID")
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	book, err := h.db.GetBookByID(id)
	if err != nil {
		h.log.Error(err, "Error getting book by ID")
		respondWithError(w, http.StatusInternalServerError, "Error getting book by ID")
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book, err := decodeBook(r)
	if err != nil {
		h.log.Error(err, "Error decoding book")
		respondWithError(w, http.StatusInternalServerError, "Error decoding book")
		return
	}

	b, err := h.db.CreateBook(*book)
	if err != nil {
		h.log.Error(err, "Error creating book")
		return
	}

	respondWithJSON(w, http.StatusCreated, b)
}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		h.log.Error(err, "Invalid book ID")
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := decodeBook(r)
	if err != nil {
		h.log.Error(err, "Error decoding book")
		respondWithError(w, http.StatusInternalServerError, "Error decoding book")
		return
	}

	b, err := h.db.UpdateBook(id, *book)
	if err != nil {
		h.log.Error(err, "Error updating book")
		respondWithError(w, http.StatusInternalServerError, "Error updating book")
		return
	}

	respondWithJSON(w, http.StatusOK, b)
}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		h.log.Error(err, "Invalid book ID")
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	err = h.db.DeleteBook(id)
	if err != nil {
		h.log.Error(err, "Error deleting book")
		respondWithError(w, http.StatusInternalServerError, "Error deleting book")
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func getID(r *http.Request) (int, error) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func decodeBook(r *http.Request) (*models.Book, error) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
