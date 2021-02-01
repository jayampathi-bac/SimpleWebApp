package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// User Struct (Model)
type User struct {
	ID       string   `json:"id"`
	UserName string   `json:"username"`
	Email    string   `json:"email"`
	Address  string   `json:"address"`
}

// enable cross Origin policy
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

// Init users var as a slice User Struct
var users []User



//Get All Users
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users [] User
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/simplecruddb")
	row, err := db.Query("select * from user")
	if err != nil {
		panic(err.Error())
	}else{
		for row.Next(){
			var id string
			var username string
			var email string
			var address string
			err2 := row.Scan(&id, &username, &email, &address)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				user := User{
					ID:id,
					UserName:username,
					Email: email,
					Address: address,
				}
				users = append(users, user)
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(users)
}


//get User
func getUser(w http.ResponseWriter, r *http.Request) {
	var userOb User

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/simplecruddb")
	row, err := db.Query("select * from user where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}else{
		for row.Next(){
			var id string
			var username string
			var email string
			var address string
			err2 := row.Scan(&id, &username, &email, &address)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				user := User{
					ID:      id,
					UserName:    username,
					Email:     email,
					Address: address,
				}
				userOb = user
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(userOb)


}

//Create User
func createUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	var user User
	_= json.NewDecoder(r.Body).Decode(&user)
	/*customer.ID = strconv.Itoa(rand.Intn(1000))
	customers = append(customers,customer)*/
	//json.NewEncoder(w).Encode(customer)


	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/simplecruddb")
	insert, err := db.Query("INSERT INTO user (id, username, email, address) VALUES (?, ?, ?, ?)", user.ID, user.UserName, user.Email, user.Address)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	json.NewEncoder(w).Encode(user)
}

//Update User
func updateUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	var user User
	_= json.NewDecoder(r.Body).Decode(&user)

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/simplecruddb")
	update, err := db.Query("update user set username=? , email=? , address=?  where id= ?", user.UserName, user.Email, user.Address, params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
	json.NewEncoder(w).Encode(user)
}

//Delete User
func deleteUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/simplecruddb")
	delete, err := db.Query("delete from user where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	json.NewEncoder(w).Encode(users)
}

func main() {
	//Router Initialization
	r := mux.NewRouter()
	fmt.Println("Server Running...")
	////Mock Data -@todo -implement DB
	//users = append(users, User{ID: "1", UserName: "CJ", Email: "alphacj@gmail.com", Address: "Horana"})
	//users = append(users, User{ID: "2", UserName: "Derek", Email: "derek@gmail.com", Address: "USA"})

	// Route Handlers
	r.HandleFunc("/api/users", getUsers).Methods("GET","OPTIONS")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET","OPTIONS")
	r.HandleFunc("/api/users", createUser).Methods("POST","OPTIONS")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT","OPTIONS")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE","OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", r))

}
