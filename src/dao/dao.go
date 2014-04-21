package dao

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
    name string
    phone string
    id int
}

type UserDB struct {
    db *sql.DB
}

func Init()(*UserDB, error) {
    user := new(UserDB)
    db,err:=sql.Open("mysql","root:@tcp(192.168.1.191:3306)/golang?charset=utf8")
    if err != nil {
        fmt.Println("database initialize error : ",err.Error())
        return nil, err
    }
    user.db = db
    return user, nil
}

// func PreExc()
func GetUser(phone string, user *UserDB ) (*UserInfo, error) {
    rows, err := user.db.Query("select id, name from golangUserinfo where phone =?", phone)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        userinfo := &UserInfo{"", "", 0}
        rows.Scan(&userinfo.id, &userinfo.name)
        userinfo.phone = phone
        return userinfo, nil
    }
    return nil, nil
}

func SaveUser(userinfo UserInfo, user *UserDB) error {
    stmt,err := user.db.Prepare("insert ignore into golangUserinfo set name = ?, phone = ?, created_at = now() ")
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    defer stmt.Close()
    if res,err := stmt.Exec(userinfo.name, userinfo.phone); err == nil {
        if id,err := res.LastInsertId(); err == nil {
            fmt.Println("insert id : ",id);
        }
    } else {
        fmt.Println("insert db failed: ", err.Error())
        return err
    }
    return nil
}

func SaveSigninLog(phone string, user *UserDB)(error) {
    stmt,err := user.db.Prepare("insert into golangUserinfoEvent (user_id, signin_time) (select id, now() from golangUserinfo where phone = ?)")
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    defer stmt.Close()
    if res,err := stmt.Exec(phone); err == nil {
        if id,err := res.LastInsertId(); err == nil {
            fmt.Println("insert id : ",id);
            return nil
        }
        return err

    } else {
        fmt.Println("insert db failed: ", err.Error())
        return err
    }
}

func Close(user *UserDB) {
    if user.db != nil {
        user.db.Close()
    }
}
