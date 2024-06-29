package controllers

import (
	"Quiz-1/config"
	"Quiz-1/models"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.GetDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ImageURL, &article.ArticleLength, &article.CreatedAt, &article.UpdatedAt, &article.CategoryID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	json.NewEncoder(w).Encode(articles)
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
