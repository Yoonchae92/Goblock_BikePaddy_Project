package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"

	"bytes"
	"strings"

	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//! board--------------------------------------------------------------
type Board struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
	Content   string
}

type PassedData struct {
	PostData []Board
	Target   string
	Value    string
	PageList []string
	Page     string
}

var (
	gormDB *gorm.DB
	//go:embed web
	staticContent embed.FS
)

const (
	MaxPerPage = 100
)

//!  session.go------------------------------------------------------------------------
func getUser(w http.ResponseWriter, req *http.Request) User {
	fmt.Println("getUser()")
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u User

	un, err := ReadSession(db, c.Value)
	if err != nil {
		log.Fatal(err)
	}
	UpdateCurrentTime(db, un)
	u, _ = ReadUserById(db, un)
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	fmt.Println("alreadyLoggedIn()")
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un, err := ReadSession(db, c.Value)
	if err != nil {
		return false
	}

	UpdateCurrentTime(db, un)

	_, err = ReadUserById(db, un)
	if err != nil {
		return false
	}

	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return true
}

//!  session.go------------------------------------------------------------------------

//!   crud.go----------------------------------------
// Topic table columns
type User struct {
	Id           string
	Password     string
	Name         string
	Created      string
	Totaltime    string
	Trytime      string
	Recovery     string
	Frontcount   string
	Backcount    string
	AverageRPM   string
	AverageSpeed string
	Distance     string
	Musclenum    string
	Kcalorynum   string
}

type Input struct {
	Id       string
	Password string
}

// CustomError: error type struct
type CustomError struct {
	Code    string
	Message string
}

func (e *CustomError) Error() string {
	return e.Code + ", " + e.Message
}

func (e *CustomError) StatusCode() int {
	result, _ := strconv.Atoi(e.Code)
	return result
}

// Create1 insert data to db
func Create1(db *sql.DB) {
	// Create 1
	insert, err := db.Query("INSERT INTO topic (title, description, created, author, profile) VALUES ('GOPHER', 'Hello Golang', NOW(), 'techno', 'dev')")
	checkError(err)
	defer insert.Close()
}

func CreateSession(db *sql.DB, sessionId string, userId string) {
	stmt, err := db.Prepare("insert into session values (?, ?, ?)")
	checkError(err)
	defer stmt.Close()
	_, err = stmt.Exec(sessionId, userId, time.Now().Format("2006-01-02 15:04:05"))
	checkError(err)

}

// Create2 insert data to db
func CreateUser(db *sql.DB, req *http.Request) *CustomError { //! 이거는 어디껀가
	// req.ParseForm()
	id := req.PostFormValue("id")
	password := req.PostFormValue("password")
	name := req.PostFormValue("name")
	t := time.Now().Format("2006-01-02 15:04:05")
	// Create 2
	stmt, err := db.Prepare("insert into user (id, password, name, created,totaltime,trytime,recovery,backcount,averageRPM,averageSpeed,distance,musclenum,kcalorynum) values (?, ?, ?, ?,?,?,?,?,?,?,?,?,?)")
	checkError(err)
	defer stmt.Close()

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err = stmt.Exec(id, bs, name, t)
	if err != nil {
		fmt.Println("error:", err)
		return &CustomError{Code: "1062", Message: "already exists id."}
	}
	return nil
}

func ReadSession(db *sql.DB, sessionId string) (string, error) {
	fmt.Println("ReadSession()")
	row, err := db.Query("select user_id from session where session_id = ?", sessionId)
	checkError(err)
	defer row.Close()
	var userId string

	for row.Next() {
		err = row.Scan(&userId)
		if err != nil {
			log.Fatal(err)
		}
	}
	return userId, nil
}

func ReadUserById(db *sql.DB, userId string) (User, error) {

	fmt.Println("ReadUserById()")
	row, err := db.Query("select * from user where id = ?", userId)
	//row, err := db.Query("select * from user")

	checkError(err)
	defer row.Close()

	var user = User{} //! 배열로 받아서 모든 테이블 정보 가져오기 해야함

	for row.Next() {
		err := row.Scan(&user.Id, &user.Password, &user.Created, &user.Name, &user.Totaltime, &user.Trytime, &user.Recovery, &user.Frontcount, &user.Backcount, &user.AverageRPM, &user.AverageSpeed, &user.Distance, &user.Musclenum, &user.Kcalorynum)
		if err != nil {
			log.Fatal(err) //! 2021/11/4  이유
		}
	}

	return user, nil
}

// Read select all data from db
func ReadUser(db *sql.DB, req *http.Request) (User, *CustomError) {
	// Read
	id, pw := req.PostFormValue("id"), req.PostFormValue("password")
	rows, err := db.Query("select * from user where id = ?", id)
	checkError(err)
	defer rows.Close()

	var user = User{}

	if !rows.Next() {
		return user, &CustomError{Code: "401", Message: "ID doesn't exist."}
	} else {
		_ = rows.Scan(&user.Id, &user.Password, &user.Created, &user.Name, &user.Totaltime, &user.Trytime, &user.Recovery, &user.Frontcount, &user.Backcount, &user.AverageRPM, &user.AverageSpeed, &user.Distance, &user.Musclenum, &user.Kcalorynum)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
	if err != nil {
		return user, &CustomError{Code: "401", Message: "uncorrect password."}
	}

	return user, nil
}

// Update change data from db
func Update(db *sql.DB) {
	// Update
	stmt, err := db.Prepare("update topic set profile=? where profile=?")
	checkError(err)

	res, err := stmt.Exec("developer", "dev")
	checkError(err)

	a, err := res.RowsAffected()
	checkError(err)

	fmt.Println(a, "rows in set")
}

func UpdateCurrentTime(db *sql.DB, sessionID string) {
	stmt, err := db.Prepare("UPDATE session SET `current_time`=? WHERE `user_id`=?")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), sessionID)
	checkError(err)
}

func CleanSessions(db *sql.DB) {

	var sessionID string
	var currentTime string
	rows, err := db.Query("select session_id, current_time from session")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&sessionID, &currentTime)
		if err != nil {
			log.Fatal(err)
		}
		t, _ := time.Parse("2006-01-02 15:04:05", currentTime)
		if time.Now().Sub(t) > (time.Second * 30) {
			DeleteSession(db, sessionID)
		}
	}

	dbSessionCleaned = time.Now()
}

func DeleteSession(db *sql.DB, sessionID string) {
	stmt, err := db.Prepare("delete from session where `session_id`=?")
	checkError(err)

	_, err = stmt.Exec(sessionID)
	checkError(err)

}

// Delete delete data from db
func Delete(db *sql.DB) {
	// Delete
	stmt, err := db.Prepare("delete from user where `id`=?")
	checkError(err)

	res, err := stmt.Exec(5)
	checkError(err)

	a, err := res.RowsAffected()
	checkError(err)
	fmt.Println(a, "rows in set")
}

func pingDB(db *sql.DB) {
	err := db.Ping()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func crud() {
	fmt.Println("Go MYSQL Tutorial")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)

	// Connect to mysql server
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	pingDB(db)

}

//!   crud.go----------------------------------------

func write(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		content := r.PostFormValue("content")

		newPost := Board{Title: title, Author: author, Content: content}
		gormDB.Create(&newPost)

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	tpl.ExecuteTemplate(w, "write.gohtml", nil)
}

func board(w http.ResponseWriter, r *http.Request) {
	var b []Board

	page := r.FormValue("page")
	if page == "" {
		page = "1"
	}

	pageInt, _ := strconv.Atoi(page)

	if keyword := r.FormValue("v"); keyword != "" {
		target := r.FormValue("target")

		switch target {
		case "title":
			q := gormDB.Where("title LIKE ?", fmt.Sprintf("%%%s%%", keyword)).Find(&b)
			pg := paginator.New(adapter.NewGORMAdapter(q), MaxPerPage)
			pg.SetPage(pageInt)

			if err := pg.Results(&b); err != nil {
				panic(err)
			}
			pgNums, _ := pg.PageNums()
			pageSlice := getPageList(page, pgNums)

			temp := PassedData{
				PostData: b,
				Target:   target,
				Value:    keyword,
				PageList: pageSlice,
				Page:     page,
			}

			tpl.ExecuteTemplate(w, "board.gohtml", temp)
			return
		case "author":
			q := gormDB.Where("author LIKE ?", fmt.Sprintf("%%%s%%", keyword)).Find(&b)
			pg := paginator.New(adapter.NewGORMAdapter(q), MaxPerPage)
			pg.SetPage(pageInt)

			if err := pg.Results(&b); err != nil {
				panic(err)
			}
			pgNums, _ := pg.PageNums()
			pageSlice := getPageList(page, pgNums)

			temp := PassedData{
				PostData: b,
				Target:   target,
				Value:    keyword,
				PageList: pageSlice,
				Page:     page,
			}

			tpl.ExecuteTemplate(w, "board.gohtml", temp)
			return
		}
	}

	q := gormDB.Order("id desc").Find(&b)
	pg := paginator.New(adapter.NewGORMAdapter(q), MaxPerPage)

	pg.SetPage(pageInt)

	if err := pg.Results(&b); err != nil {
		panic(err)
	}

	pgNums, _ := pg.PageNums()
	pageSlice := getPageList(page, pgNums)

	temp := PassedData{
		PostData: b,
		PageList: pageSlice,
		Page:     page,
	}

	tpl.ExecuteTemplate(w, "board.gohtml", temp)
}
func edit(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/edit/")
	var b Board

	gormDB.First(&b, id)

	if r.Method == http.MethodPost {

		gormDB.Model(&b).Updates(Board{Title: r.PostFormValue("title"), Author: r.PostFormValue("author"), Content: r.PostFormValue("content")})
		var byteBuf bytes.Buffer
		byteBuf.WriteString("/post/")
		byteBuf.WriteString(id)
		http.Redirect(w, r, byteBuf.String(), http.StatusSeeOther)

	}

	tpl.ExecuteTemplate(w, "write.gohtml", b)
}
func post(w http.ResponseWriter, r *http.Request) {
	// id := r.FormValue("id")
	id := strings.TrimPrefix(r.URL.Path, "/post/")

	var b Board
	gormDB.First(&b, id)

	tpl.ExecuteTemplate(w, "post.gohtml", b)
}

func getPageList(p string, limit int) []string {
	page, _ := strconv.Atoi(p)
	var result []string

	for i := page - 2; i <= page+2; i++ {
		if i > 0 && i <= limit {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

//! board.go --------------------------------- 이까지가 board.go

//! account.go ---------------------------------
const (
	//추가
	user     = "root"
	password = "1234"
	//port     = "3307"
	database = "user"
	host     = "127.0.0.1"
)

//! account.go ---------------------------------

// var db *sql.DB
// var tpl *template.Template
var (
	db               *sql.DB
	tpl              *template.Template
	dbSessionCleaned time.Time
)

//go:embed web
var content embed.FS

const sessionLength int = 60

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*"))
	dbSessionCleaned = time.Now()
}

func main() {
	fmt.Println("Head")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, password, host, database)
	var err error
	fmt.Println("connection check..")
	// Connect to mysql server
	db, err = sql.Open("mysql", connectionString)
	fmt.Println("Connecting to DB..")
	checkError(err)
	defer db.Close()
	//바꾼코드
	err = db.Ping()
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	gormDB.AutoMigrate(&Board{})
	//원래코드
	//pingDB(db)
	fmt.Println("Successfully Connected to DB")

	http.HandleFunc("/", login)

	http.HandleFunc("/write", write)
	http.HandleFunc("/board/", board)
	http.HandleFunc("/post/", post)
	http.HandleFunc("/edit/", edit)

	http.HandleFunc("/index2", index2) //! 뭐여

	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/index", index)
	http.HandleFunc("/logout", logout)
	http.Handle("/web/", http.FileServer(http.FS(staticContent)))
	fmt.Println("Listening...ss")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// var b []Board

	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "dashboard.html", u) //! html로 바꾸는법~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

}

func index2(w http.ResponseWriter, req *http.Request) {

	// var b []Board

	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "user.html", u) //! html로 바꾸는법~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

}

func login(w http.ResponseWriter, req *http.Request) { //! ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~``
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		user, err := ReadUser(db, req)
		if err != nil {
			errMsg := map[string]interface{}{"error": err}
			tpl.ExecuteTemplate(w, "login3.html", errMsg)
			return
		}
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		CreateSession(db, c.Value, user.Id)
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login3.html", nil)
}

func signUp(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}

	if req.Method == http.MethodPost {
		err := CreateUser(db, req)
		if err != nil {
			errMsg := map[string]interface{}{"error": err}
			tpl.ExecuteTemplate(w, "signup.gohtml", errMsg)
		} else {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		return
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete session
	DeleteSession(db, c.Value)

	//
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionCleaned) > (time.Second * 30) {
		go CleanSessions(db)
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

//! main에 남아야 할 내용들...........................................
