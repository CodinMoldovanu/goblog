package db

import (
	"log"

	"github.com/codinmoldovanu/goblog/models"
)

//GrabPosts and relay
// func GrabPosts() []models.Post {
// 	db := createConn()
// 	defer db.Close()
// 	posts := []models.Post{}
// 	rows, err := db.Query(`SELECT * FROM posts ORDER BY posted DESC LIMIT 10 OFFSET $1`, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for rows.Next() {
// 		var row models.Post
// 		err := rows.Scan(&row.ID, &row.Author, &row.Posted, &row.Updated, &row.Title, &row.Content, &row.Status, &row.URL, &row.Excerpt, &row.Meta)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		posts = append(posts, row)
// 	}
// 	return posts
// }
//GrabPosts and transform to json and send
func GrabPosts() []models.Post {
	db := createConn()
	defer db.Close()
	posts := []models.Post{}
	rows, err := db.Query(`SELECT * FROM posts ORDER BY posted DESC LIMIT 10 OFFSET $1`, 0)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var row models.Post
		err := rows.Scan(&row.ID, &row.Author, &row.Posted, &row.Updated, &row.Title, &row.Content, &row.Status, &row.URL, &row.Excerpt, &row.Meta)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, row)
	}

	return posts
}

//GrabPost and return it to router
func GrabPost(postURL string) models.Post {
	db := createConn()
	defer db.Close()
	row := models.Post{}
	db.QueryRow(`SELECT * FROM posts WHERE url = $1`, postURL).Scan(&row.ID, &row.Author, &row.Posted, &row.Updated, &row.Title, &row.Content, &row.Status, &row.URL, &row.Excerpt, &row.Meta)
	return row
}

//UpdatePost based on input data and save in DB
func UpdatePost(post models.Post) models.Post {
	db := createConn()
	defer db.Close()
	row := models.Post{}
	err := db.QueryRow(`UPDATE posts SET content = $1, title = $2 WHERE ID = $3`, post.Content, post.Title, post.ID).Scan(&row.ID, &row.Author, &row.Posted, &row.Updated, &row.Title, &row.Content, &row.Status, &row.URL, &row.Excerpt, &row.Meta)
	if err != nil {
		log.Print("An error occured while updating a post: ")
		log.Print(err.Error())
	}
	return row
}

//NewPost to insert a new post into the DB
func NewPost(post models.Post) models.Post {
	db := createConn()
	defer db.Close()
	row := models.Post{}
	err := db.QueryRow(`INSERT INTO posts VALUES ($1, $2, $3, NOW(), $4, $5, $6, $7, $8, NULL)`, post.ID, post.Author, post.Posted, post.Title, post.Content, post.Status, post.URL, post.Excerpt).Scan(&row.ID, &row.Author, &row.Posted, &row.Updated, &row.Title, &row.Content, &row.Status, &row.URL, &row.Excerpt, &row.Meta)
	if err != nil {
		log.Print("An error occured while inserting a post: ")
		log.Print(err.Error())
	}
	return row
}
