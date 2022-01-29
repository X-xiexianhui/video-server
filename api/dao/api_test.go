/**
    @author:xie
    @date:2022/1/29
    @note:
**/
package dao

import "testing"

func clearTables() {
	dbConn.Exec("truncate video_server.comments")
	dbConn.Exec("truncate video_server.sessions")
	dbConn.Exec("truncate video_server.video_info")
	dbConn.Exec("truncate video_server.video_del_rec")
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("Del", testDeleteUserCredential)
	t.Run("ReGet", testReGetUserCredential)
}
func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("admin")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%s", pwd)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("admin", "123456")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func testDeleteUserCredential(t *testing.T) {
	err := DeleteUserCredential("admin", "123456")
	if err != nil {
		t.Logf("%v", err)
	}
}
func testReGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("admin")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%s", pwd)
}
