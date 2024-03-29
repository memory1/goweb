package model

import (
	"log"
	"database/sql"
	"fmt"
	"encoding/base64"
	"crypto/sha512"
	"time"
)
const passwordSalt="a99vvMMsdfsfa242341"

type User struct {
	ID int
	Email string
	Password string
	FirstName string
	LastName string
	LastLogin *time.Time
}

func Login(email, password string)(*User, error) {
	result := &User{}
	hasher :=sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	fmt.Println("hasher sum: ", hasher.Sum(nil))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println("base64.URLEncoding.EncodeToString:", pwd)
	row := db.QueryRow(`select id, email,firstname, lastname from public.user 
		where email = $1 and password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email,&result.FirstName, &result.LastName)
	switch  {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
		
	}
	t := time.Now()
	_, err = db.Exec(`
		update public.user set lastlogin = $1
		where id = $2`,t, result.ID)
	
	if err != nil {
		log.Printf("Failed to update the login time for user %v to %v:%v", result.Email, t, err)
	}
	
	return result, nil
}

