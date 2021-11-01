package model

type memoryHandler struct {
	memberMap   map[string]*Member
	Community   map[string]*Community
	File        map[string]*File
	My_dataMap  map[string]*My_data
	Workout_log map[string]*Workout_log
}

//4개 func을 만든다
func (m *memoryHandler) GetMembers() []*Member {
	list := []*Member{}
	for _, v := range m.memberMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) GetMemberAdmin() []*Member {
	list := []*Member{}
	for _, v := range m.memberMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddMember(id string, pswd string, name string, birth string, gender string, email string, area string, bike_info string, career string, club string) *Member {
	member := &Member{id, pswd, name, birth, gender, email, area, bike_info, career, club}
	m.memberMap[id] = member
	return member
}

func (m *memoryHandler) LoginChk(id string, pswd string) *Member {
	member := &Member{id, pswd, "", "", "", "", "", "", "", ""}
	m.memberMap[id] = member
	return member
}

func (m *memoryHandler) RemoveMember(id string) bool {
	if _, ok := m.memberMap[id]; ok { // memberMap id 값이 있으면
		delete(m.memberMap, id) //지우고
		return true
	}
	return false
}

func (m *memoryHandler) Close() {

}

// func MembersMemoryHandler() MembersDBHandler {
// 	m := &memoryHandler{}
// 	m.memberMap = make(map[string]*Member) // map을 초기화
// 	return m
// }

func (m *memoryHandler) GetCommunity() []*Community {
	list := []*Community{}
	for _, v := range m.Community {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddCommunity(board_id string, title string, content string, id string, date string, file_id string, good string) *Community {
	Community := &Community{board_id, title, content, id, date, file_id, good}
	m.Community[board_id] = Community
	return Community
}

func (m *memoryHandler) RemoveCommunity(board_id string) bool {
	if _, ok := m.Community[board_id]; ok { // memberMap id 값이 있으면
		delete(m.Community, board_id) //지우고
		return true
	}
	return false
}

func CommunityMemoryHandler() CommunityDBHandler {
	m := &memoryHandler{}
	m.Community = make(map[string]*Community)
	return m
}

func (m *memoryHandler) GetFile() []*File {
	list := []*File{}
	for _, v := range m.File {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddFile(file_id string, name string, location string, size string) *File {
	File := &File{file_id, name, location, size}
	m.File[file_id] = File
	return File
}

func (m *memoryHandler) RemoveFile(file_id string) bool {
	if _, ok := m.File[file_id]; ok { // memberMap id 값이 있으면
		delete(m.File, file_id) //지우고
		return true
	}
	return false
}

func FileMemoryHandler5() FileDBHandler {
	m := &memoryHandler{}
	m.File = make(map[string]*File)
	return m
}

func (m *memoryHandler) GetMyData() []*My_data {
	list := []*My_data{}
	for _, v := range m.My_dataMap {
		list = append(list, v)
	}
	return list
}

/*
func (m *memoryHandler3) AddMyData(data_id string, totaltime float64, trytime float64, recoverytime float64, frontcount float64, backcount float64, avaregyRPB float64, averagespeed float64, distance float64, musclenum float64, kcalorynum float64, date int, week int, month int, year int) *My_data {
	My_data := &My_data{data_id, totaltime, trytime, recoverytime, frontcount, backcount, avaregyRPB, averagespeed, distance, musclenum, kcalorynum, date, week, month, year}
	m.My_dataMap[data_id] = My_data
	return My_data
}
*/

func MydataMemoryHandler() MydataDBHandler {
	m := &memoryHandler{}
	m.My_dataMap = make(map[string]*My_data)
	return m
}

func (m *memoryHandler) GetWorkOutlog() []*Workout_log {
	list := []*Workout_log{}
	for _, v := range m.Workout_log {
		list = append(list, v)
	}
	return list
}

/*
func (m *memoryHandler2) AddWorkOutlog(workout_id string, work_id int, start int, end int, distance int) *Workout_log {
	workout := &Workout_log{workout_id, work_id, start, end, distance}
	m.Workout_log[workout_id] = workout
	return workout
}
*/
func WorkoutMemoryHandler2() WorkoutDBHandler {
	m := &memoryHandler{}
	m.Workout_log = make(map[string]*Workout_log)
	return m
}
