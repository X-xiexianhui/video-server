/**
    @author:xie
    @date:2022/1/29
    @note:
**/
package dao

import "testing"

var tempVid string

func clearTables() {
	dbConn.Exec("truncate video_server.users")
	dbConn.Exec("truncate video_server.comments")
	dbConn.Exec("truncate video_server.sessions")
	dbConn.Exec("truncate video_server.video_info")
	dbConn.Exec("truncate video_server.video_del_rec")
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
}
func testGetUser(t *testing.T) {
	pwd, err := GetUser("admin")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%s", pwd)
}

func testAddUser(t *testing.T) {
	err := AddUser("admin", "123456")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("admin", "123456")
	if err != nil {
		t.Logf("%v", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempVid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempVid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}
