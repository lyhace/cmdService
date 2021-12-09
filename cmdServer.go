package main

import (
	"fmt"
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
	resultStr := "Only support Post"
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		checkErrs(err)
		defer r.Body.Close()

		shellStr := string(b)
		if len(shellStr) > 0 && strings.HasPrefix(shellStr, "fs_cli") {
			cmd := exec.Command("sh", "-c", shellStr)
			//cmd := exec.Command(key)
			out, err := cmd.CombinedOutput()
			resultStr = string(out)
			if err != nil {
				log.Printf("%s shellStr: %s\n", shellStr, err.Error())
			} else {
				log.Print("shellStr:", shellStr)
			}

		} else {
			resultStr = "shellStr is nil"
		}
	}
	log.Println(resultStr)
	fmt.Fprintf(w, resultStr)

}

//处理错误函数
func checkErrs(err error) {
	if err != nil {
		panic(err)
		log.Println("Read failed:", err)
	}
}
