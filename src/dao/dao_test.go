// dao_test.go
package dao

import (
    "testing"
)

func TestSaveUser( t *testing.T) {
    userdb, err := Init()
    defer Close(userdb)
    if err != nil {
        t.Error("db init failed")
        return
    }
    userinfo := UserInfo{"å´è´µé”‹", "1603", 1}
    err = SaveUser(userinfo, userdb)
    if err != nil {
        t.Error("insert failed")
    }
}

func TestSaveSigninLog( t *testing.T) {
    userdb, err := Init()
    defer Close(userdb)
    if err != nil {
        t.Error("db init failed")
        return
    }
    userinfo := UserInfo{"å´è´µé”‹", "1603", 1}
    err = SaveSigninLog(userinfo.phone, userdb)
    if err != nil {
        t.Error("insert failed")
    }
}

func TestGetUser( t *testing.T) {
    userdb, err := Init()
    defer Close(userdb)
    if err != nil {
        t.Error("db init failed")
        return
    }
//    userinfo := &UserInfo{}
    phone := "1603"
    userinfo, err1 := GetUser(phone, userdb)
    if err1 != nil {
        t.Error("select failed")
    } else if userinfo == nil {
        t.Log("no phone record is 1603")
    } else {
        t.Log("id is:", userinfo.id, " name is:", userinfo.name)
    }
}

