package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type SqliteHandler struct {
	db *sql.DB // 멤버변수로 가진다
}

func (s *SqliteHandler) GetMembers() []*Member {
	members := []*Member{}                                                                                             //list를 만든다
	rows, err := s.db.Query("SELECT id, pswd, name, birth, gender, email, area, bike_info, career, club FROM members") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var member Member                                                                                                                                             //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&member.ID, &member.PSWD, &member.Name, &member.Birth, &member.Gender, &member.Email, &member.Area, &member.BikeINFO, &member.Career, &member.Club) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		members = append(members, &member)
	}
	log.Print(members[0])
	return members
}

func (s *SqliteHandler) AddMember(id string, pswd string, name string, birth string, gender string, email string, area string, bike_info string, career string, club string) *Member { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO members (id, pswd, name, birth, gender, email, area, bike_info, career, club) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id, pswd, name, birth, gender, email, area, bike_info, career, club)
	if err != nil {
		panic(err)
	}
	var member Member
	member.ID = id
	member.PSWD = pswd
	member.Name = name
	member.Birth = birth
	member.Gender = gender
	member.Email = email
	member.Area = area
	member.BikeINFO = bike_info
	member.Career = career
	member.Club = club
	return &member
}

func (s *SqliteHandler) GetLoginChk(id string, pw string) *Member {
	fmt.Println("GetLoginChk ::: id=" + id + " pw=" + pw)
	stmt, err := s.db.Prepare("SELECT id, name FROM members WHERE id= ? ")

	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	var member Member

	err = stmt.QueryRow(id).Scan(&member.ID, &member.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
		}
		return &member
	}
	return &member

	// stmt.Scan(&member.ID, &member.Name)
	// member.ID = id
	// return &member
}

func (s *SqliteHandler) RemoveMemberAdmin(id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := s.db.Prepare("DELETE FROM members WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

func (s *SqliteHandler) GetMemberAdmin() []*Member {
	members := []*Member{}                                                                                             //list를 만든다
	rows, err := s.db.Query("SELECT id, pswd, name, birth, gender, email, area, bike_info, career, club FROM members") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var member Member                                                                                                                                             //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&member.ID, &member.PSWD, &member.Name, &member.Birth, &member.Gender, &member.Email, &member.Area, &member.BikeINFO, &member.Career, &member.Club) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		members = append(members, &member)
	}
	log.Print(members[0])
	return members
}

func (d *SqliteHandler) RemoveMember(id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := d.db.Prepare("DELETE FROM members WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (d *SqliteHandler) Close() {
	d.db.Close()
}

func MemberSqliteHandler(filepath string) MembersDBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS members (
			id			TEXT  NOT NULL PRIMARY KEY,
			pswd		TEXT NOT NULL,
			name		TEXT NOT NULL,
			birth		DATE,
			gender		TEXT NOT NULL,
			email		TEXT NOT NULL,
			area		TEXT NOT NULL,
			bike_info TEXT,
			career TEXT,
			club TEXT
			)`)

	statement.Exec()
	return &SqliteHandler{db: database} // &sqliteHandler{}를 반환
}

func (d *SqliteHandler) GetCommunity() []*Community {
	community := []*Community{}                                                                        //list를 만든다
	rows, err := d.db.Query("SELECT board_id, title, content, id, date, file_id, good FROM community") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var Community Community                                                                                                                   //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&Community.Board_Id, &Community.Title, &Community.Content, &Community.ID, &Community.Date, &Community.File_Id, &Community.Good) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		community = append(community, &Community)
	}
	log.Print(community[0])
	return community
}

func (d *SqliteHandler) AddCommunity(board_id string, title string, content string, id string, date string, file_id string, good string) *Community { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := d.db.Prepare("INSERT INTO community (board_id, title, content, id, date, file_id, good) VALUES (?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(board_id, title, content, id, date, file_id, good)
	if err != nil {
		panic(err)
	}
	var Community Community
	Community.Board_Id = board_id
	Community.Title = title
	Community.Content = content
	Community.ID = id
	Community.Date = date
	Community.File_Id = file_id
	Community.Good = good

	return &Community
}

func (d *SqliteHandler) RemoveCommunity(board_id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := d.db.Prepare("DELETE FROM community WHERE board_id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(board_id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.

func CommunitySqliteHandler(filepath string) CommunityDBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS community (
			board_id			TEXT PRIMARY KEY,
			title		TEXT ,
			content		TEXT ,
			id		TEXT ,
			date		TEXT ,
			file_id		TEXT ,
			good		TEXT
			);`)

	statement.Exec()
	return &SqliteHandler{db: database} // &sqliteHandler{}를 반환
}

func (d *SqliteHandler) GetFile() []*File {
	file := []*File{}                                                        //list를 만든다
	rows, err := d.db.Query("SELECT file_id, name, content, location, size") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var File File                                                    //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&File.File_ID, &File.Name, &File.Location, &File.Size) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		file = append(file, &File)
	}
	log.Print(file[0])
	return file
}

func (d *SqliteHandler) AddFile(file_id string, name string, location string, size string) *File { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := d.db.Prepare("INSERT INTO file (file_id, name, location, size) VALUES (?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(file_id, name, location, size)
	if err != nil {
		panic(err)
	}
	var File File
	File.File_ID = file_id
	File.Name = name
	File.Location = location
	File.Size = size

	return &File
}

func (d *SqliteHandler) RemoveFile(file_id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := d.db.Prepare("DELETE FROM file WHERE file_id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(file_id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.

func FileSqliteHandler(filepath string) FileDBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS file (
			file_id			TEXT PRIMARY KEY,
			name		TEXT NOT NULL,
			location		TEXT NOT NULL,
			size		TEXT NOT NULL
			);`)

	statement.Exec()
	return &SqliteHandler{db: database} // &sqliteHandler{}를 반환
}

func (d *SqliteHandler) GetMyData() []*My_data {
	MyData := []*My_data{}                                                         //list를 만든다
	rows, err := d.db.Query("SELECT data_id, date, week, month, year FROM mydata") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var My_data My_data                                                                      //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&My_data.Data_id, &My_data.Date, &My_data.Week, &My_data.Month, &My_data.Year) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		MyData = append(MyData, &My_data)
	}
	log.Print(MyData[0])
	return MyData
}

/*
func (s *sqliteHandler3) AddMyData(data_id string, distance int, date int, week int, month int, year int, time int) *My_data { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO mydata (data_id, totaltime, trytime, recoverytime, frontcount, backcount, avaregyRPB, averagespeed, distance, musclenum, kcalorynum, date, week, month, year) VALUES (?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(data_id, distance, date, week, month, year, time)
	if err != nil {
		panic(err)
	}
	var My_data My_data
	My_data.Data_id = data_id
	My_data.Distance = distance
	My_data.Date = date
	My_data.Week = week
	My_data.Month = month
	My_data.Year = year
	My_data.Time = time

	return &My_data
}
*/

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func MydataSqliteHandler(filepath string) MydataDBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS mydata (
			data_id			TEXT PRIMARY KEY,
			date	INTEGER DEFAULT '0',
			week	INTEGER DEFAULT '0',
			month	INTEGER DEFAULT '0',
			year	INTEGER DEFAULT '0'
			);`)

	statement.Exec()
	return &SqliteHandler{db: database} // &sqliteHandler{}를 반환
}

func WorkoutMemoryHandler() WorkoutDBHandler {
	m := &memoryHandler{}
	m.Workout_log = make(map[string]*Workout_log)
	return m
}

func (d *SqliteHandler) GetWorkOutlog() []*Workout_log {
	workoutlog := []*Workout_log{}                                                                                                                                                   //list를 만든다
	rows, err := d.db.Query("SELECT workout_id, totaltime, trytime, recoverytime, frontcount, backcount, avaregyRPB, averagespeed, distance, musclenum, kcalorynum FROM workoutlog") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var workout Workout_log                                                                                                                                                                                                                      //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&workout.WORKOUT_ID, &workout.TotalTime, &workout.Trytime, &workout.RecoveryTime, &workout.FrontCount, &workout.Backcount, &workout.AvaregyRPB, &workout.AverageSpeed, &workout.Distance, &workout.MuscleNum, &workout.KcaloryNum) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		workoutlog = append(workoutlog, &workout)
	}
	return workoutlog
}

/*
func (s *sqliteHandler2) AddWorkOutlog(workout_id string, work_id int, start int, end int, distance int) *Workout_log { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO workout (workout_id, work_id, start, end, distance) VALUES (?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(workout_id, work_id, start, end, distance)
	if err != nil {
		panic(err)
	}
	var workout Workout_log
	workout.WORKOUT_ID = workout_id
	workout.ID = work_id
	workout.Start = start
	workout.End = end
	workout.Distance = distance

	return &workout
}
*/
// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.

func WorkoutSqliteHandler(filepath string) WorkoutDBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS workout (
			workout_id			TEXT PRIMARY KEY,
			totaltime		DOUBLE,
			trytime		DOUBLE,
			recoverytime		DOUBLE,
			frontcount		DOUBLE,
			backcount		DOUBLE,
			avaregyRPB		DOUBLE,
			averagespeed	DOUBLE,
			distance	DOUBLE,
			musclenum	DOUBLE,
			kcalorynum	DOUBLE
			);`)

	statement.Exec()
	return &SqliteHandler{db: database} // &sqliteHandler{}를 반환
}
