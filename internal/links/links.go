package links

import (
	"log"

	database "github.com/cristianortiz/gqlgen-hackernews/internal/pkg/db/migrations/mysql"
	"github.com/cristianortiz/gqlgen-hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// Save() insert in database a new Link with Title and Adress
func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row Inserted!")
	return id
}

// GetAll() returns all the links from DB
func GetAll() []Link {
	stmt, err := database.Db.Prepare("SELECT ID,Title,Address FROM Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
