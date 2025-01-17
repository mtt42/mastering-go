package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	fmt.Println(LOGFILE)

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	LstdFlags := log.Ldate | log.LstdFlags
	iLog := log.New(f, "LNumj", LstdFlags)
	iLog.Println("Mastering Go, 4th Edition")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry.")
}
