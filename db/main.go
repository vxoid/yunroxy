package main

import (
	"fmt"

	"github.com/vxoid/yunroxy/yunroxyDB"
)
var sql, Err = yunroxyDB.NewApiDb("yunroxyDB/db.db")

func main() {

	//yunroxyDB.Db.Create(&yunroxyDB.User{ApiKey: "Obamna"})
	if Err != nil {
		fmt.Println(Err, "Failed connection")
	}
	ben := yunroxyDB.ApiDb.IsApiKey(yunroxyDB.ApiDb{}, "Obamna")
	fmt.Println(ben)
	

	//db.Unscoped().Delete(&yunroxyDB.User{})

}
