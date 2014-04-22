package main

import(
    "net/http"
    "log"
    "encoding/json"
    "regexp"
    "flag"
    "dao"
)

type UserInfo struct {
    Name string
    Phone string
    Id int
}
var activityName = flag.String("name", "golang", "activity name")
var phoneReg = regexp.MustCompile("\\d{4,5}")


func main(){
    flag.Parse()
    log.Println("signIn start, activity name: ", activityName)
    http.Handle("/html/",http.FileServer(http.Dir("web")))
    http.Handle("/css/",http.FileServer(http.Dir("web")))
    http.Handle("/js/",http.FileServer(http.Dir("web")))
    http.Handle("/bootstrap/",http.FileServer(http.Dir("web")))
    http.HandleFunc("/signIn", signInHandler)	

    http.ListenAndServe(":8888",nil)	

}

type Result struct{
    Code int
    Msg string
    Data interface{}
}

func (result *Result) toJson() []byte{
    json, err := json.Marshal(result)
    if err == nil {
        return json
    }
    return nil
}

func signInHandler (rw http.ResponseWriter,req *http.Request){
    rw.Header().Set("Content-Type","application/json")
    userdb,_ := dao.Init()
    result := new(Result)
    err := req.ParseForm()
    if(nil != err){
        log.Println(err)
        result.Code = 1
        result.Msg = "parameter error"
        rw.Write(result.toJson())
        return
    }

    //phone invalid
    phone := req.FormValue("phone")
    if false == phoneReg.MatchString(phone){
        result.Code = 2
        result.Msg = "phone invalid: " + phone
        rw.Write(result.toJson())
        return
    }

    log.Println("phone is: ", phone)
    //sign up
    username := req.FormValue("username")
    if "" != username {
        //		user := &UserInfo{username,phone,0}
        userdb.SaveUser(username, phone)
        userdb.SaveSigninLog(phone)		
        result.Code = 0
        result.Msg = "sign in success: " + username
        log.Println(result.Msg)
        rw.Write(result.toJson())
        return
    }

    //not sign up
    user,err_get := userdb.GetUser(phone)
    if err_get != nil || user == nil {
        log.Println("user is nil")
        result.Code = 3
        result.Msg = phone + ": record not exist, need signup "
        rw.Write(result.toJson())
        return
    }

    //sign in
    userdb.SaveSigninLog(user.Phone)
    result.Code = 0
    result.Msg = ("sign in success: " + user.Name)
    log.Println(result.Msg)
    rw.Write(result.toJson())
    return 
}

