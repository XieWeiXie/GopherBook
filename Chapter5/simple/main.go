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

func middlewareLogger(next http.Handler) http.Handler {
	now := time.Now()
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.RequestURI, now.Format("2006/01/02 -15:04:05"))
		next.ServeHTTP(writer, request)
	})
}

type singleSong struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Time   string `json:"time"`
	Album  string `json:"album"`
}

type songs []singleSong

type api struct {
	Title   string `json`
	Content string `json:"content"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Comment string `json:"comment"`
}

type apis []api

func song(writer http.ResponseWriter, req *http.Request) {
	var ss songs
	ss = []singleSong{
		{
			ID:     1,
			Name:   "全部都是你",
			Author: "DP龙猪",
			Time:   "03:23",
			Album:  "全部都是你",
		},
		{
			ID:     2,
			Name:   "对你的感觉",
			Author: "DP龙猪",
			Time:   "04:23",
			Album:  "对你的感觉",
		},
		{
			ID:     3,
			Name:   "我可不可以",
			Author: "DP龙猪",
			Time:   "05:23",
			Album:  "我可不可以",
		},
		{
			ID:     4,
			Name:   "围绕",
			Author: "DP龙猪",
			Time:   "03:23",
			Album:  "围绕",
		},
	}
	var aps apis
	aps = []api{
		{
			Title:   "获取人员信息",
			Content: "通过ID获取人员信息",
			Method:  http.MethodGet,
			Path:    fmt.Sprintf("/person/get?id=ID"),
			Comment: fmt.Sprintf("ID 选择 1 或者 2"),
		},
		{
			Title:   "获取所有人员信息",
			Content: "获取内置所有人员的信息",
			Method:  http.MethodGet,
			Path:    fmt.Sprintf("/persons"),
			Comment: fmt.Sprintf("无须传入请求参数"),
		},
		{
			Title:   "创建人员信息",
			Content: "传入参数 id 和 telephone 创建新人",
			Method:  http.MethodPost,
			Path:    fmt.Sprintf("/person/post"),
			Comment: fmt.Sprintf("传入参数 id 或者 telephone"),
		},
		{
			Title:   "更新人员信息",
			Content: "传入参数 id 更新人员 telephone 信息",
			Method:  http.MethodPatch,
			Path:    fmt.Sprintf("/person/patch?id=ID"),
			Comment: fmt.Sprintf("传入路径参数 id 和请求参数 telephone"),
		},
	}
	var all = struct {
		Songs songs
		APis  apis
	}{
		Songs: ss,
		APis:  aps,
	}
	currentPath, _ := os.Getwd()
	temp, err := template.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/song.html"))
	if err != nil {
		log.Println(err)
		return
	}
	err2 := temp.Execute(writer, all)
	if err2 != nil {
		log.Println(err2)
		return
	}
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
	Tag           string    `json:"tag"`
	Title         string    `json:"title"`
	Time          time.Time `json:"time"`
	Content       string    `json:"content"`
	CommentInt    int       `json:"comment_int"`
	CollectionInt int       `json:"collection_int"`
	ClickInt      int       `json:"click_int"`
}

type contents []content

func timeHandle(date time.Time) string {
	return date.Format("2006/01/02")
}

func home(writer http.ResponseWriter, req *http.Request) {
	var c contents
	c = []content{
		{
			Tag:           "Go",
			Title:         "How to learn Golang",
			Time:          time.Now(),
			Content:       "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
			CommentInt:    1,
			CollectionInt: 12,
			ClickInt:      100,
		},
		{
			Tag:           "Python",
			Title:         "How to learn Python",
			Time:          time.Now().Add(-24 * time.Hour),
			Content:       "Python is a programming language that lets you work quickly and integrate systems more effectively.",
			CommentInt:    2,
			CollectionInt: 34,
			ClickInt:      1000,
		},
		{
			Tag:           "Java",
			Title:         "How to learn Java",
			Time:          time.Now().Add(-24 * 2 * time.Hour),
			Content:       "Java is a general-purpose programming language that is class-based, object-oriented, and designed to have as few implementation dependencies as possible.",
			CommentInt:    3,
			CollectionInt: 124,
			ClickInt:      900,
		},
		{
			Tag:           "JavaScript",
			Title:         "How to learn JavaScript",
			Time:          time.Now().Add(-24 * 2 * 2 * time.Hour),
			Content:       "JavaScript often abbreviated as JS, is a high-level, interpreted programming language that conforms to the ECMAScript specification. JavaScript has curly-bracket syntax, dynamic typing, prototype-based object-orientation, and first-class functions.",
			CommentInt:    212,
			CollectionInt: 1224,
			ClickInt:      9030,
		},
	}
	currentPath, _ := os.Getwd()
	temp := template.New("index.html")
	t := temp.Funcs(template.FuncMap{"timeHandle": timeHandle})
	t, err := t.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/home.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(writer, c)
	if err != nil {
		panic(err)
	}
}

type passageContent struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
	Detail    string    `json:"detail"`
}

type side struct {
	Tag   string   `json:"tag"`
	Items []string `json:"items"`
}

func timeFormat(date time.Time) string {
	return fmt.Sprintf(date.Format(time.Stamp) + " By ")
}

func passage(writer http.ResponseWriter, request *http.Request) {

	var content = struct {
		passageContent
		side
	}{
		passageContent: passageContent{
			Title:     "How to learn golang",
			CreatedAt: time.Now(),
			Author:    "Go Team",
			Detail: `The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.`,
		},
		side: side{
			Tag: "状态",
			Items: []string{
				"用户数: 62",
				"分享数: 27",
				"评论数: 19",
				"收藏数: 12",
			},
		},
	}
	currentPath, _ := os.Getwd()
	temp := template.New("index.html")
	t := temp.Funcs(template.FuncMap{"time": timeFormat})
	t, err := t.ParseFiles(path.Join(currentPath, "GopherBook/Chapter5/simple/template/index.html"), path.Join(currentPath, "GopherBook/Chapter5/simple/template/passage.html"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(writer, content)
	if err != nil {
		panic(err)
	}
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}

func textHandle(value string) string {
	return fmt.Sprintf(value + " Yes")
}

func HelloTemplate(writer http.ResponseWriter, request *http.Request) {

	content := `
{{define "content"}}
			{{if .Content}}
				{{range .Content}}
				<p>{{ . | handle}}</p>
				{{end}}
			{{ else}}
			<button>No Result</button>
			{{end}}
{{end}}
`
	html := `
	<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			{{template "content" .}}
		</body>
	</html>
`
	temp := template.New("Hello")
	t2 := template.Must(template.Must(temp.Funcs(template.FuncMap{"handle": textHandle}).Parse(content)).Clone())
	t3, err := t2.Parse(html)
	if err != nil {
		panic(fmt.Sprintf("template fail : %s", err.Error()))
	}
	text := struct {
		Title   string
		Content []string
	}{
		Title:   "Hello Golang",
		Content: []string{"Go", "Python", "Java", "JavaScript"},
	}
	err = t3.Execute(writer, text)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", logger(home))
	http.Handle("/2", middlewareLogger(http.HandlerFunc(home)))
	http.HandleFunc("/persons", logger(getHandler))
	http.HandleFunc("/person/post", logger(postHandler))
	http.HandleFunc("/person/patch", logger(patchHandler))
	http.HandleFunc("/person/get", logger(getProfile))
	http.HandleFunc("/login", logger(login))
	http.HandleFunc("/logout", logger(logout))
	http.HandleFunc("/songs", logger(song))
	http.HandleFunc("/progress", logger(progress))
	http.HandleFunc("/passage", logger(passage))
	http.HandleFunc("/template", logger(HelloTemplate))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
