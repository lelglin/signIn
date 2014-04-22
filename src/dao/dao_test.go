// dao_test.go
package dao

import (
    "testing"
)

func TestSaveUser( t *testing.T) {
    userdb, err := Init()
    defer userdb.Close()
    if err != nil {
        t.Error("db init failed")
        return
    }
    userinfo := UserInfo{"吴贵锋", "1603", 1}
    err = userdb.SaveUser(userinfo.Name, user.Phone)
    if err != nil {
        t.Error("insert failed")
    }
}

func TestSaveSigninLog( t *testing.T) {
    userdb, err := Init()
    defer userdb.Close()
    if err != nil {
        t.Error("db init failed")
        return
    }
    userinfo := UserInfo{"吴贵锋", "1603", 1}
    err = userdb.SaveSigninLog(userinfo.Phone)
    if err != nil {
        t.Error("insert failed")
    }
}

func TestGetUser( t *testing.T) {
    userdb, err := Init()
    defer userdb.Close()
    if err != nil {
        t.Error("db init failed")
        return
    }
//    userinfo := &UserInfo{}
    Phone := "1603"
    userinfo, err1 := userdb.GetUser(Phone)
    if err1 != nil {
        t.Error("select failed")
    } else if userinfo == nil {
        t.Log("no Phone record is 1603")
    } else {
        t.Log("id is:", userinfo.Id, " Name is:", userinfo.Name)
    }
}

