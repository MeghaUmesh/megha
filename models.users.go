package main

import (
    "errors"
    "database/sql"
    "strings"
    "fmt"
    _"github.com/go-sql-driver/mysql"
    _"github.com/jinzhu/gorm"
)

type user struct {
    Username string `json:"username"`
    Password string `json:"-"`
}


/* var userList = []user{
    user{Username: "user1", Password: "pass1"},
    user{Username: "user2", Password: "pass2"},
    user{Username: "user3", Password: "pass3"},
} */

var db *sql.DB


func init(){
	var err error
	db, err = sql.Open("mysql", "root:Pengus@123@/megha")
	if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
        
    }
	//Migrate the schema
	//db.AutoMigrate(&todoModel{})
}

// Register a new user with the given username and password
func registerNewUser(username, password string) (*user, error) {
    if strings.TrimSpace(username) == "" {
        return nil, errors.New("The username can't be empty")
    }else if strings.TrimSpace(password) == "" {
        return nil, errors.New("The password can't be empty")
    }else if isUsernameAvailable(username) {
        return nil, errors.New("The username already taken")
    }

   // u := user{Username: username, Password: password}

    stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES (?, ?)")
	if err != nil {
		return nil,err
	}
	_ , err = stmt.Exec(username, password)

    return nil, err
}

// Check if the supplied username is available
/* func isUsernameAvailable(username string) bool {
    for _, u := range userList {
        if u.Username == username {
            return true
        }
    }
    return false
} */

func isUsernameAvailable(usernam string) bool {
   
   row := db.QueryRow("SELECT username FROM users where username = ?",usernam)
   var name string
   err := row.Scan(&name)
   fmt.Println(err)
    
    if  err == sql.ErrNoRows{
        return false
    }
    return true
}