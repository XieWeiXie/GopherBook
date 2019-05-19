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
	"strings"
	"time"
	"unicode"
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
			WhatIsUp:       "Python",
		},
		{
			Avatar:         "http://images.org/456",
			OriginWeChatID: uuid(),
			ID:             2,
			CreatedAt:      time.Now(),
			Telephone:      "987654321",
			Gender:         0,
			WhatIsUp:       "Golang",
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
	req.Header.Set("Authorization", "Bearer profile")
	a := make(template.FuncMap)
	a["createdAt"] = createdAtHandle
	tmp := template.New("index.html")
	tmp.Funcs(a)
	tmp, _ = tmp.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/table.html"))
	tmp.Funcs(a).Execute(writer, NewPersonRecords())
}

type loginInfo struct {
	UserName string  `json:"user_name"`
	Password string  `json:"password"`
	Error    []error `json:"error"`
}

func login(writer http.ResponseWriter, req *http.Request) {
	currentPath, _ := os.Getwd()
	temp, _ := template.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/login.html"))
	var lgInfo loginInfo
	if req.Method == http.MethodGet {
		temp.Execute(writer, lgInfo)
		return
	}

	if err := req.ParseForm(); err != nil {
		return
	}
	UserName := req.PostFormValue("username")
	Password := req.PostFormValue("password")

	lgInfo.UserName = UserName
	lgInfo.Password = Password

	errFunc := func(values string) error {
		v := values
		if len(v) < 8 || len(v) == 0 {
			return fmt.Errorf("the length should be larger 8")
		}
		if unicode.IsNumber(rune(v[0])) {
			return fmt.Errorf("should not start number")
		}
		return nil
	}

	if errFunc(lgInfo.UserName) == nil && errFunc(lgInfo.Password) == nil {
		http.Redirect(writer, req, "/", http.StatusSeeOther)
		return
	} else {
		if errFunc(lgInfo.UserName) != nil {
			lgInfo.Error = append(lgInfo.Error, fmt.Errorf("username is not suitable"))
		}
		if errFunc(lgInfo.Password) != nil {
			lgInfo.Error = append(lgInfo.Error, fmt.Errorf("password is not suitable"))
		}
		log.Println(lgInfo)
		temp.Execute(writer, lgInfo)
	}
}

func logout(writer http.ResponseWriter, req *http.Request) {
	req.Header.Del("Authorization")
	http.Redirect(writer, req, "/login", http.StatusSeeOther)
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")
		if authorization == "" {
			http.Redirect(writer, request, "/login", http.StatusSeeOther)
			return
		}
		stringList := strings.Split(authorization, " ")
		if len(stringList) != 2 {
			writer.Write([]byte("Authorization Format: Authorization: Bearer xxx"))
			return
		}
		next.ServeHTTP(writer, request)
	}
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	now := time.Now()
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.RequestURI, now.Format("2006/01/02 -15:04:05"))
		next.ServeHTTP(writer, request)
	}

}

type router struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type routers []router

func apis(writer http.ResponseWriter, req *http.Request) {
	var rs routers
	rs = []router{
		{
			"/",
			"/",
			http.MethodGet,
		},
		{
			"persons",
			"/persons",
			http.MethodGet,
		}, {
			"person get",
			"/person/get",
			http.MethodGet,
		}, {
			"logout",
			"/logout",
			http.MethodGet,
		},
		{
			"person post",
			"/person/post",
			http.MethodPost,
		}, {
			"login",
			"/login",
			http.MethodPost,
		},
		{
			"person patch",
			"/person/patch",
			http.MethodPatch,
		},
	}
	currentPath, _ := os.Getwd()
	temp, _ := template.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/apis.html"))
	temp.Execute(writer, rs)
}

type progressStatus struct {
	Now  float64 `json:"now"`
	Year int     `json:"year"`
}

func progress(writer http.ResponseWriter, req *http.Request) {
	var proStatus progressStatus
	monthDays := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	now := time.Now()
	y, m, d := now.Date()
	ok := 0
	sum := monthDays[time.Month(m)-1] + d
	if (y%400 == 0) || ((y%4 == 0) && (y%100 != 0)) {
		ok = 1
	}
	if (ok == 1) && (time.Month(m) > 2) {
		sum += 1
	}
	proStatus.Now, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(sum)/float64(365)*100), 64)
	proStatus.Year = y
	currentPath, _ := os.Getwd()
	temp, _ := template.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/progress.html"))
	temp.Execute(writer, proStatus)
}

type content struct {
	Title string   `json:"title"`
	Text  []string `json:"text"`
}

type contents []content

func textHandler(values []string) string {
	return strings.Join(values, " ")
}
func home(writer http.ResponseWriter, req *http.Request) {
	var c contents
	c = []content{
		{
			Title: "net/http 内置库的使用",
			Text: []string{
				"http.HandleFunc",
				"http.Handle",
				"http.ServeMux",
				"http.Server",
			},
		},
		{
			Title: "template 的使用",
			Text: []string{
				"渲染文件",
				"模版语法",
				"if、else、range",
				"函数调用",
				"模版继承",
			},
		},
		{
			Title: "bootstrap",
			Text: []string{
				"栅格系统",
				"导航栏",
				"表格",
				"进度条",
			},
		},
	}
	currentPath, _ := os.Getwd()
	temp := template.New("index.html")
	temp.Funcs(template.FuncMap{"text": textHandler})
	t, _ := temp.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/home.html"))
	t.Funcs(template.FuncMap{"text": textHandler}).Execute(writer, c)
}

func main() {
	http.HandleFunc("/", logger(home))
	http.HandleFunc("/persons", logger(getHandler))
	http.HandleFunc("/person/post", logger(postHandler))
	http.HandleFunc("/person/patch", logger(patchHandler))
	http.HandleFunc("/person/get", logger(getProfile))
	http.HandleFunc("/login", logger(login))
	http.HandleFunc("/logout", logger(logout))
	http.HandleFunc("/apis", logger(apis))
	http.HandleFunc("/progress", logger(progress))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
