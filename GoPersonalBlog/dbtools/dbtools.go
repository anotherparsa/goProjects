package dbtools

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db = ConnectToDataBase()

type Posts struct {
	Id       int
	Title    string
	Text     string
	Author   string
	Category string
}

func ConnectToDataBase() *sql.DB {
	db, err := sql.Open("mysql", "root:simplepassword123@tcp(127.0.0.1:3306)/personalblog")
	if err != nil {
		panic(err)
	} else {
		return db
	}
}

//getting posts
func GetPosts() map[int]Posts {
	posts := make(map[int]Posts)
	query := "SELECT * FROM posts"
	get, _ := db.Query(query)
	var pPosts Posts
	for get.Next() {
		_ = get.Scan(&pPosts.Id, &pPosts.Title, &pPosts.Text, &pPosts.Author, &pPosts.Category)
		posts[pPosts.Id] = pPosts
	}
	return posts

}
