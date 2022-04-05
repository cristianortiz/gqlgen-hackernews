package users

import (
	"database/sql"
	"log"

	database "github.com/cristianortiz/gqlgen-hackernews/internal/pkg/db/migrations/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	print(stmt)
	if err != nil {
		log.Fatal(err)
	}
	//hashing the user password before stored it in DB
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

//HashPassword hashes the given password for a new user
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserIdByUsername check if a a user exists in DB by a given username
func GetUserIdByUsername(username string) (int, error) {
	//sql statementto get the ID of a given username
	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username=?")
	if err != nil {
		log.Fatal(err)
	}
	//Queryrow returns a pointer to sql.Row
	row := stmt.QueryRow(username)
	var Id int
	//copy the returned Id in roe into Id var declared above
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		//the username does not exists in DB
		return 0, err
	}

	return Id, nil

}
