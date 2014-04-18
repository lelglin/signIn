package main

import(
	"net/http"
	"log"
	"encoding/json"
	"regexp"
)

var phoneReg = regexp.MustCompile("\\d{4}")

func main(){
	log.Println("signIn start")
	http.Handle("/html/",http.FileServer(http.Dir("web")))
	
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
        if req.Method == "GET" {
		log.Println("GET Method")
	}
	result := new(Result)
        err := req.ParseForm()
        if(nil != err){
                log.Println(err)
                result.Code = 1
                result.Msg = "parameter error"
                rw.Write(result.toJson())
		return
        }

	phone := req.FormValue("phone")
	if phoneReg.MatchString(phone){
		log.Println("phone match")
	}else{
		log.Println("phone is nil")
		result.Code = 2
		result.Msg = "phone is blank"	
		rw.Write(result.toJson())
		return
	}

}

