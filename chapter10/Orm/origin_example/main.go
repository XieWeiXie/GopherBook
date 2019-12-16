package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/Orm/origin_example/model"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	var config mysql.Config
	config = mysql.Config{
		User:      "root",
		Passwd:    "admin123",
		DBName:    "person",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		ParseTime: true,
	}
	dsn := config.FormatDSN()
	originURL := "root:admin123@/person?charset=utf8&parseTime=true&loc=Local"
	fmt.Println(dsn, originURL)

	db, err = sql.Open("mysql", originURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	//deleteAllTable()
	//insert()
}

func deleteTable() {
	stmt, _ := db.Prepare("DELETE  from  wechat_persons")
	stmt.Exec()

}

func records() []model.Person {
	rand.Seed(time.Now().UnixNano())
	var result []model.Person
	for i := 0; i < 10; i++ {
		one := model.Person{
			Base: model.Base{
				ID:        uint(rand.Int31()),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Avatar:        "https://images.pexels.com/photos/2326961/pexels-photo-2326961.jpeg?auto=format%2Ccompress&cs=tinysrgb&dpr=1&w=500",
			NickName:      fmt.Sprintf("%d", uint(rand.Int31())),
			AccountString: strconv.Itoa(int(rand.Int31n(100000))),
			Gender:        1,
			Location:      fmt.Sprintf("北京: %d", uint(rand.Int31n(100000))),
			Signal:        fmt.Sprintf("走自己的路，让别人去说吧: %d", uint(rand.Int31n(100))),
		}
		result = append(result, one)
	}
	return result
}

func insert() {
	for _, i := range records() {
		_, err := db.Exec("INSERT INTO wechat_persons (id, created_at, updated_at, avatar, nick_name, account_string, gender,location,signal_person ) values ( ?,?,?,?,?,?,?,?,? )",
			i.ID, i.CreatedAt, i.UpdatedAt, i.Avatar, i.NickName, i.AccountString, i.Gender, i.Location, i.Signal)
		if err != nil {
			log.Println(err)
			return
		}

	}
}

func query() []model.PersonSerializer {
	var (
		result []model.PersonSerializer
	)
	rows, err := db.Query("SELECT id, created_at, updated_at, avatar, nick_name, account_string,gender,location,signal_person FROM  wechat_persons")
	if err != nil {
		log.Println(err)
		return result
	}
	for rows.Next() {
		var one model.Person
		err := rows.Scan(&one.ID, &one.CreatedAt, &one.UpdatedAt, &one.Avatar, &one.NickName, &one.AccountString, &one.Gender, &one.Location, &one.Signal)
		if err != nil {
			break
		}
		result = append(result, one.JSONSerializer())

	}
	return result
}

func apiGet(writer http.ResponseWriter, request *http.Request) {
	var result []model.PersonSerializer
	result = query()
	var values = make(map[string]interface{})
	values["data"] = result
	//writer.Write(Json(values))
	json.NewEncoder(writer).Encode(values)
}

func apiPatch(writer http.ResponseWriter, request *http.Request) {
	var result = make(map[string]interface{})
	request.Header.Add("content-type", "application/json")
	if request.Method != http.MethodPatch {
		result["code"] = http.StatusBadRequest
		result["detail"] = fmt.Sprintf("method is not allow")
		json.NewEncoder(writer).Encode(Json(result))
		return
	}
	request.ParseForm()

	id := request.FormValue("id")
	nickName := request.PostFormValue("nick_name")
	location := request.PostFormValue("location")
	log.Println(id, nickName, location)

	row := db.QueryRow("SELECT id, nick_name, location FROM wechat_persons  where id = ?", id)
	var person model.Person
	err := row.Scan(&person.ID, &person.NickName, &person.Location)
	if err != nil {
		fmt.Println(err)
		return
	}

	if id == strconv.Itoa(int(person.ID)) && nickName != "" && location != "" {

		stmt, err := db.Prepare("UPDATE wechat_persons set nick_name = ? , location = ?, updated_at = ? where id = ?")
		if err != nil {
			log.Println(err)
			json.NewEncoder(writer).Encode(Json(result))
			return
		}
		_, err = stmt.Exec(nickName, location, time.Now(), id)
		row := db.QueryRow("SELECT id, created_at, updated_at, avatar, nick_name, account_string,gender, location, signal_person FROM wechat_persons  where id = ?", id)
		var one model.Person
		err = row.Scan(&one.ID, &one.CreatedAt, &one.UpdatedAt, &one.Avatar, &one.NickName, &one.AccountString, &one.Gender, &one.Location, &one.Signal)
		if err != nil {
			log.Println(err)
			json.NewEncoder(writer).Encode(Json(result))
			return
		}

		result["data"] = one.JSONSerializer()
		result["code"] = http.StatusOK
		json.NewEncoder(writer).Encode(result)
		return
	}
}

func Json(value map[string]interface{}) []byte {
	content, _ := json.Marshal(value)
	return content
}

func main() {
	http.HandleFunc("/", apiGet)
	http.HandleFunc("/person", apiPatch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
