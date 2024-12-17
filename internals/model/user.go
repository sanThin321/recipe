package model

import (
	"fmt"

	postgres "main.go/data"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	P_number string `json:"number"`
	Password string `json:"password"`
	Image    []byte `json:"image"`
	User_id  int    `json:"id"`
}

type User_info struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	Image   []byte `json:"image"`
	Cookie  string `json:"cookie"`
}

type SellerProfile struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	Image   []byte `json:"image"`
	User_id int    `json:"id"`
}

type User_email struct {
	Email string `json:"email"`
}

type Del_type struct {
	Id int
}

func (user *User) Create() error {
	sql := "INSERT INTO users (username, email, phone_number, password ) VALUES ($1, $2, $3, $4)"
	_, err := postgres.Db.Exec(sql, user.Username, user.Email, user.P_number, user.Password)
	if err != nil {
		panic(err)
	}
	fmt.Println("User inserted successfully")
	return err
}

func (user *User) Get() error {
	sql := "SELECT  email, password , id FROM users WHERE email=$1 and password=$2;"
	return postgres.Db.QueryRow(sql, user.Email, user.Password).Scan(&user.Email, &user.Password, &user.User_id)
}

func (user *User) User_get(email string) error {
	sql := "SELECT username, email , phone_number , image_data FROM users WHERE email=$1 "
	return postgres.Db.QueryRow(sql, email).Scan(&user.Username, &user.Email, &user.P_number, &user.Image)
}

func (user *User_info) Info_user_update() error {

	sqlStatement := `
        UPDATE users
        SET username = $1, phone_number = $2, image_data = $3, email = $4
        WHERE email = $5`

	result, err := postgres.Db.Exec(sqlStatement, user.Name, user.Contact, user.Image, user.Email, user.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Get the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return nil
}

// const queryUpdateProfile = "UPDATE seller_profile SET profile_picture=$1 WHERE contact_number=$2"
const queryUpdateProfile = "UPDATE users SET username = $1, phone_number = $2, image_data = $3 , email = $4 WHERE id = $5"

func (p *SellerProfile) UpdatePic() error {
	_, err := postgres.Db.Exec(queryUpdateProfile, p.Name, p.Contact, p.Image, p.Email, p.User_id)
	return err
}

const deletequery = "DELETE FROM public.users WHERE id = $1;"

func (id *Del_type) Delete_user_id() error {
	_, err := postgres.Db.Exec(deletequery, id.Id)
	return err
}

func GetUserHashedPassword(email string) (string, error) {
	var hashedPassword string
	err := postgres.Db.QueryRow("SELECT password FROM public.users WHERE email = $1", email).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
