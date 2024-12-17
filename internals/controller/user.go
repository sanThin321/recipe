package controller

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"main.go/internals/model"
	httpresponse "main.go/util/httprespones"
	// "database/sql"
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	// "reflect"
	// "strconv"
	// "net/http"
	// "time"
	// "github.com/gorilla/mux"
	// "golang.org/x/crypto/bcrypt"
	// "main.go/internals/model"
	// httpresponse "main.go/util/httprespones"
)

func Register_user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	hashedPassword, err := httpresponse.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	c_err := user.Create()

	if c_err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// no error
	httpresponse.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "admin added"})
}

func User_login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	storedHashedPassword, err := model.GetUserHashedPassword(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			httpresponse.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		} else {
			httpresponse.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(user.Password))

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:    "recipe-cookie",
		Value:   "Ema_datshi",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	user_email := http.Cookie{
		Name:    "email",
		Value:   user.Email,
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}

	User_id := http.Cookie{
		Name:    "id",
		Value:   strconv.Itoa(user.User_id), // Convert int to string
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &user_email)
	http.SetCookie(w, &User_id)

	httpresponse.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "success login"})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})
	httpresponse.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the "my-cookie" cookie from the request
	cookie, err := r.Cookie("recipe-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			// No cookie found, redirect to login page or return an error
			httpresponse.RespondWithError(w, http.StatusSeeOther, "cookie not found")
			return false
		}
		// Some other error occurred
		httpresponse.RespondWithError(w, http.StatusInternalServerError, "internal server error")
		return false
	}
	// Verify the cookie value
	if cookie.Value != "Ema_datshi" {
		// Invalid cookie value, redirect to login page or return an error
		httpresponse.RespondWithError(w, http.StatusSeeOther, "cookie does not match")
		return false
	}
	return true
}

func Get_user(w http.ResponseWriter, r *http.Request) {
	var user_email model.User_email

	json.NewDecoder(r.Body).Decode(&user_email)

	defer r.Body.Close()

	var user model.User
	sql_err := user.User_get(user_email.Email)

	if sql_err != nil {
		httpresponse.RespondWithError(w, http.StatusUnauthorized, sql_err.Error())
		fmt.Println(sql_err)
		return
	}

	httpresponse.RespondWithJSON(w, http.StatusOK, user)

}

func getPnumber(pnumberParam string) (int, error) {
	phonenumber, err := strconv.Atoi(pnumberParam)
	if err != nil {
		return 0, err
	}
	return phonenumber, nil
}

func Update_user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")

	pnumber := mux.Vars(r)["user_id"]
	phonenumber, numErr := getPnumber(pnumber)

	fmt.Println(pnumber)

	if numErr != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}

	var profileData struct {
		ProfilePicture string `json:"profilepicture"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		Contact        string `json:"contact"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&profileData); err != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	// fmt.Println(profileData)

	// Decode base64 image data
	imageData, err := base64.StdEncoding.DecodeString(profileData.ProfilePicture[len("data:image/png;base64,"):])
	if err != nil {
		fmt.Print(err)
		httpresponse.RespondWithError(w, http.StatusBadRequest, "invalid base64 data")
		return
	}

	sellerProfile := &model.SellerProfile{
		Name:    profileData.Name,
		Email:   profileData.Email,
		Contact: profileData.Contact,
		Image:   imageData,
		User_id: phonenumber,
	}

	updateErr := sellerProfile.UpdatePic()
	if updateErr != nil {
		fmt.Println(updateErr)
	}
	httpresponse.RespondWithJSON(w, http.StatusOK, "user succcessfully updated")

}

func Delete_user(w http.ResponseWriter, r *http.Request) {
	pnumber := mux.Vars(r)["user_id"]
	phonenumber, numErr := getPnumber(pnumber)

	fmt.Println("pnumber")
	fmt.Println(pnumber)
	fmt.Printf("Type of intValue: %s\n", reflect.TypeOf(pnumber))
	fmt.Println("phonenumber")
	fmt.Println(phonenumber)
	fmt.Printf("Type of intValue: %s\n", reflect.TypeOf(phonenumber))

	if numErr != nil {
		httpresponse.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}

	del_id := &model.Del_type{Id: phonenumber}

	del_error := del_id.Delete_user_id()
	if del_error != nil {
		fmt.Println(del_error)
	}

	fmt.Println("user deleted")
	httpresponse.RespondWithJSON(w, http.StatusOK, "user succcessfully deleted")

}
