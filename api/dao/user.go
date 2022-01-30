//Package dao
/*
 @author:xie
   @date:2022/1/30
   @note:用户功能
*/
package dao

import (
	"log"
)

func AddUser(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	err = stmtIns.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetUser(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Panicf("%s", err)
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil {
		return "", err
	}
	err = stmtOut.Close()
	if err != nil {
		return "", err
	}
	return "", nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Panicf("%s", err)
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	err = stmtDel.Close()
	if err != nil {
		return err
	}
	return nil
}
