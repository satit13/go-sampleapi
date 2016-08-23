package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
	"html"
)

const (
	db_host = "tcp(nava.work:3306)"
	db_name = "sample"
	db_user = "root"
	db_pass = "mypass"
)

type User struct {
	Id int
	Name string
	Secret []byte
}

func main() {

	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", getUser ).Methods("GET"); log.Println("/")
	//log.Fatal(http.ListenAndServe(":8081", nil))


	http.HandleFunc("/users", getUser)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Hello(w http.ResponseWriter , r *http.Request){
	log.Println("Hello func work")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getUser(w http.ResponseWriter , r *http.Request)  {
	log.Println("get user function")
	var  dsn = db_user + ":" + db_pass + "@" + db_host + "/" + db_name + "?parseTime=true"
	db := sqlx.MustConnect("mysql",dsn)
	var sql string
	sql = "select id,name,secret from user"
	// close connecct when finished exec
	//user := User{}
	log.Println(sql)
	var users []User
	err := db.Select(&users, sql)
	if err != nil {
		log.Println("Error Cannot Query SQL Statement :  ",err, &users )
	}

	juser ,_ := json.Marshal(&users)
	log.Println(&juser)
	fmt.Fprintf(w,"%s",string(juser))
}

