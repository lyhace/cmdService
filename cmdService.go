package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/cmd", cmdHandler)
	http.ListenAndServe(":12580", nil)
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	uuidStr := uuid.NewV4()
	resultStr := "Only support Post"
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		checkErrs(err)
		defer r.Body.Close()

		shellStr := string(b)
		log.Println(uuidStr, "shellStr:", shellStr)
		if len(shellStr) > 0 && strings.HasPrefix(shellStr, "fs_cli") {
			cmd := exec.Command("sh", "-c", shellStr)
			//cmd := exec.Command(key)
			out, err := cmd.CombinedOutput()
			resultStr = string(out)
			if err != nil {
				log.Printf("%s Error: %s\n", uuidStr, err.Error())
			}

		} else {
			resultStr = "shellStr is nil"
		}
	}
	log.Println(uuidStr, "resultStr:", resultStr)
	fmt.Fprintf(w, resultStr)

}

//处理错误函数
func checkErrs(err error) {
	if err != nil {
		panic(err)
		log.Println("Read failed:", err)
	}
}
