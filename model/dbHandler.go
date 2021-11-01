package model

type Member struct {
	ID       string `json:"id"`
	PSWD     string `json:"pswd"`
	Name     string `json:"name"`
	Birth    string `json:"birth"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Area     string `json:"area"`
	BikeINFO string `json:"bike_info"`
	Career   string `json:"career"`
	Club     string `json:"club"`
}

type Community struct {
	Board_Id string `json:"board_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ID       string `json:"id"`
	Date     string `json:"date"`
	File_Id  string `json:"file_id"`
	Good     string `json:"good"`
}

type File struct {
	File_ID  string `json:"file_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Size     string `json:"size"`
}

type My_data struct {
	Data_id string `json:"data_id"`
	Date    int    `json:"date"`
	Week    int    `json:"week"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
}

type Workout_log struct {
	WORKOUT_ID   string  `json:"workout_id"`
	TotalTime    float64 `json:"totaltime"`
	Trytime      float64 `json:"trytime"`
	RecoveryTime float64 `json:"recoverytime"`
	FrontCount   float64 `json:"frontcount"`
	Backcount    float64 `json:"backcount"`
	AvaregyRPB   float64 `json:"avaregyRPB"`
	AverageSpeed float64 `json:"averagespeed"`
	Distance     float64 `json:"distance"`
	MuscleNum    float64 `json:"musclenum"`
	KcaloryNum   float64 `json:"kcalorynum"`
}

type MembersDBHandler interface {
	GetMembers() []*Member
	GetMemberAdmin() []*Member
	AddMember(id string, pswd string, name string, birth string, gender string, email string, area string, bike_info string, career string, club string) *Member
	GetLoginChk(id string, pswd string) *Member
	RemoveMember(id string) bool
	RemoveMemberAdmin(id string) bool
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func MemberHandler(filepath string) MembersDBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return MemberSqliteHandler(filepath)
}

type CommunityDBHandler interface {
	GetCommunity() []*Community
	AddCommunity(board_id string, title string, content string, id string, date string, file_id string, good string) *Community
	RemoveCommunity(board_id string) bool
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func CommunityHandler(filepath string) CommunityDBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return CommunitySqliteHandler(filepath)
}

type FileDBHandler interface {
	GetFile() []*File
	AddFile(file_id string, name string, location string, size string) *File
	RemoveFile(file_id string) bool
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func FileHandler(filepath string) FileDBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return FileSqliteHandler(filepath)
}

type MydataDBHandler interface {
	GetMyData() []*My_data
	//AddMyData(data_id string, distance int, date int, week int, month int, year int, time int) *My_data
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func MydataHandler(filepath string) MydataDBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return MydataSqliteHandler(filepath)
}

type WorkoutDBHandler interface {
	GetWorkOutlog() []*Workout_log
	//AddWorkOutlog(workout_id string, work_id int, start int, end int, distance int) *Workout_log
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func WorkoutHandler(filepath string) WorkoutDBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return WorkoutSqliteHandler(filepath)
}
