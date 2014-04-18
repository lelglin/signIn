package main
import(
	"net/http"
	"log"
)

func main(){
	log.Println("signIn start")
	http.Handle("/html/",http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8888",nil)	

}
