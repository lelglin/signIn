package dao

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
    Name string
    Phone string
    Id int
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
func (user *UserDB)GetUser(Phone string) (*UserInfo, error) {
    rows, err := user.db.Query("select id, Name from golangUserinfo where Phone =?", Phone)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        userinfo := &UserInfo{"", "", 0}
        rows.Scan(&userinfo.Id, &userinfo.Name)
        userinfo.Phone = Phone
        return userinfo, nil
    }
    return nil, nil
}


//func(user *UserDB) SaveUser(userinfo *UserInfo ) error {
func(user *UserDB) SaveUser( Name string, Phone string ) error {
    stmt,err := user.db.Prepare("insert ignore into golangUserinfo set Name = ?, Phone = ?, created_at = now() ")
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    defer stmt.Close()
    if res,err := stmt.Exec(Name, Phone); err == nil {
        if id,err := res.LastInsertId(); err == nil {
            fmt.Println("insert id : ",id);
        }
    } else {
        fmt.Println("insert db failed: ", err.Error())
        return err
    }
    return nil
}

func (user *UserDB)SaveSigninLog(Phone string)(error) {
    stmt,err := user.db.Prepare("insert into golangUserinfoEvent (user_id, signin_time) (select id, now() from golangUserinfo where Phone = ?)")
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    defer stmt.Close()
    if res,err := stmt.Exec(Phone); err == nil {
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

func (user *UserDB )Close() {
    user.db.Close()
}
