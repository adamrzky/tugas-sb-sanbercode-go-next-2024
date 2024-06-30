package controllers

import (
	"Quiz-1/config"
	"Quiz-1/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func fetchFilteredArticles(w http.ResponseWriter, r *http.Request, categoryId string) {
	db, err := config.GetDB()
	if err != nil {
		http.Error(w, "Kesalahan koneksi database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var whereClauses []string
	var queryParams []interface{}

	if categoryId != "" {
		whereClauses = append(whereClauses, "category_id = ?")
		queryParams = append(queryParams, categoryId)
	}

	if title := r.URL.Query().Get("title"); title != "" {
		whereClauses = append(whereClauses, "LOWER(title) LIKE LOWER(?)")
		queryParams = append(queryParams, "%"+title+"%")
	}

	if minYear := r.URL.Query().Get("minYear"); minYear != "" {
		whereClauses = append(whereClauses, "YEAR(created_at) >= ?")
		queryParams = append(queryParams, minYear)
	}

	if maxYear := r.URL.Query().Get("maxYear"); maxYear != "" {
		whereClauses = append(whereClauses, "YEAR(created_at) <= ?")
		queryParams = append(queryParams, maxYear)
	}

	if minWord := r.URL.Query().Get("minWord"); minWord != "" {
		whereClauses = append(whereClauses, "CHAR_LENGTH(content) - CHAR_LENGTH(REPLACE(content, ' ', '')) + 1 >= ?")
		queryParams = append(queryParams, minWord)
	}

	if maxWord := r.URL.Query().Get("maxWord"); maxWord != "" {
		whereClauses = append(whereClauses, "CHAR_LENGTH(content) - CHAR_LENGTH(REPLACE(content, ' ', '')) + 1 <= ?")
		queryParams = append(queryParams, maxWord)
	}

	sortOrder := "ASC"
	if sortByTitle := r.URL.Query().Get("sortByTitle"); sortByTitle == "desc" {
		sortOrder = "DESC"
	}

	query := "SELECT id, title, content, image_url, article_length, created_at, updated_at FROM articles"
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	query += fmt.Sprintf(" ORDER BY title %s", sortOrder)

	rows, err := db.Query(query, queryParams...)
	if err != nil {
		http.Error(w, "Kesalahan eksekusi query: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the query and parameters for debugging
	log.Printf("Executing SQL Query: %s", query)
	log.Printf("With parameters: %+v", queryParams)
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ImageURL, &article.ArticleLength, &article.CreatedAt, &article.UpdatedAt); err != nil {
			http.Error(w, "Kesalahan membaca artikel: "+err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	json.NewEncoder(w).Encode(articles)
}
func GetArticlesByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryId := ps.ByName("id")
	fetchFilteredArticles(w, r, categoryId)
}

func GetAllArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fetchFilteredArticles(w, r, "")
}

func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.GetDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validation for image_url
	if !isValidURL(article.ImageURL) {
		http.Error(w, "Invalid URL format for image_url", http.StatusBadRequest)
		return
	}

	// Validation for title
	if len(strings.Fields(article.Title)) > 3 {
		http.Error(w, "Title cannot have more than three words", http.StatusBadRequest)
		return
	}

	// Set article_length based on content
	article.ArticleLength = determineArticleLength(article.Content)

	stmt, err := db.Prepare("INSERT INTO articles (title, content, image_url, article_length, category_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(article.Title, article.Content, article.ImageURL, article.ArticleLength, article.CategoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.GetDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := ps.ByName("id")

	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Additional validations as in CreateArticle

	stmt, err := db.Prepare("UPDATE articles SET title = ?, content = ?, image_url = ?, article_length = ?, category_id = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(article.Title, article.Content, article.ImageURL, article.ArticleLength, article.CategoryID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.GetDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := ps.ByName("id")

	stmt, err := db.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted article with ID %s", id)
}

// Helper functions
func isValidURL(urlStr string) bool {
	regex := `^https?://[a-z0-9\-]+(\.[a-z0-9\-]+)*(/.*)?$`
	return regexp.MustCompile(regex).MatchString(urlStr)
}

func determineArticleLength(content string) string {
	words := len(strings.Fields(content))
	characters := len(content)
	if words <= 100 || characters <= 400 {
		return "pendek"
	} else if words <= 200 || characters <= 800 {
		return "sedang"
	}
	return "panjang"
}
