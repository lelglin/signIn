diff --git a/src/dao/dao.go b/src/dao/dao.go
index 928f7b3..f14d8d3 100644
--- a/src/dao/dao.go
+++ b/src/dao/dao.go
@@ -18,7 +18,7 @@ type UserDB struct {
 
 func Init()(*UserDB, error) {
     user := new(UserDB)
-    db,err:=sql.Open("mysql","root:123456@tcp(localhost:3306)/golang?charset=utf8")
+    db,err:=sql.Open("mysql","root:@tcp(192.168.1.191:3306)/golang?charset=utf8")
     if err != nil {
         fmt.Println("database initialize error : ",err.Error())
         return nil, err
@@ -83,12 +83,6 @@ func SaveSigninLog(phone string, user *UserDB)(error) {
     }
 }
 
-func checkErr(err error){
-    if err!=nil{
-        panic(err)
-    }
-}
-
 func Close(user *UserDB) {
     if user.db != nil {
         user.db.Close()
