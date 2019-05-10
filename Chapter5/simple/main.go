package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type Person struct {
	Avatar         string    `json:"avatar"`
	OriginWeChatID string    `json:"origin_id"`
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Telephone      string    `json:"telephone"`
	Gender         int       `json:"gender"` // 男1 女0
	WhatIsUp       string    `json:"what_is_up"`
}

func uuid() string {
	b := []byte(fmt.Sprintf("%v", time.Now().UnixNano()))
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func NewPersonRecords() []Person {
	var persons []Person
	persons = []Person{
		{
			Avatar:         "http://images.org/123",
			OriginWeChatID: uuid(),
			ID:             1,
			CreatedAt:      time.Now(),
			Telephone:      "1234567890",
			Gender:         1,
			WhatIsUp:       "Hello Golang",
		},
		{
			Avatar:         "http://images.org/456",
			OriginWeChatID: uuid(),
			ID:             2,
			CreatedAt:      time.Now(),
			Telephone:      "987654321",
			Gender:         0,
			WhatIsUp:       "Hello Python",
		},
	}
	return persons
}

func postHandler(writer http.ResponseWriter, req *http.Request) {
	header := writer.Header()
	header.Add("Content-Type", "application/json")
	var person Person
	if req.Method == http.MethodPost {
		if err := req.ParseForm(); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		values := req.PostForm

		for k, v := range values {
			if k == "id" {
				id, _ := strconv.Atoi(v[0])
				person.ID = uint(id)
			}
			if k == "telephone" {
				person.Telephone = v[0]
			}
		}

	}
	defer json.NewEncoder(writer).Encode(person)

	writer.WriteHeader(http.StatusOK)

}

func getHandler(writer http.ResponseWriter, req *http.Request) {
	header := writer.Header()
	header.Add("Content-Type", "application/json")

	expire := time.Now().AddDate(0, 0, 1)
	cookie := &http.Cookie{Name: "expires", Value: "Get", Expires: expire}
	http.SetCookie(writer, cookie)
	writer.WriteHeader(http.StatusOK)
	log.Println(writer.Header().Get("Set-Cookie"))
	defer json.NewEncoder(writer).Encode(NewPersonRecords())

}

func patchHandler(writer http.ResponseWriter, req *http.Request) {
	var person Person
	if req.Method == http.MethodPatch {
		req.ParseForm()
		id := req.FormValue("id")
		intID, _ := strconv.Atoi(id)

		telephone := req.PostFormValue("telephone")
		log.Println(id, telephone)
		for _, i := range NewPersonRecords() {
			if i.ID == uint(intID) {
				person = i
				person.Telephone = telephone
				break
			}
		}
	}
	log.Println(person)
	writer.WriteHeader(http.StatusOK)
	defer json.NewEncoder(writer).Encode(person)

}

func createdAtHandle(createdAt time.Time) string {
	return createdAt.Format(time.ANSIC)
}
func getProfile(writer http.ResponseWriter, req *http.Request) {
	currentPath, _ := os.Getwd()
	a := make(template.FuncMap)
	a["createdAt"] = createdAtHandle
	tmp := template.New("index.html")
	tmp.Funcs(a)
	tmp, _ = tmp.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/table.html"))
	tmp.Funcs(a).Execute(writer, NewPersonRecords())
}

func login(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		return
	}
	UserName := req.PostFormValue("username")
	Password := req.PostFormValue("password")
	writer.WriteHeader(http.StatusOK)
	mapLogin := make(map[string]string)
	mapLogin["username"] = UserName
	mapLogin["password"] = Password
	currentPath, _ := os.Getwd()
	tmp, _ := template.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/login.html"))
	tmp.Execute(writer, mapLogin)
}

func main() {
	http.HandleFunc("/", getProfile)
	http.HandleFunc("/persons", getHandler)
	http.HandleFunc("/person/post", postHandler)
	http.HandleFunc("/person/patch", patchHandler)
	http.HandleFunc("/person/get", getProfile)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
