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
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		checkErrs(err)
		defer r.Body.Close()

		shellStr := string(b)

		if len(shellStr) > 0 && strings.HasPrefix(shellStr, "fs_cli") {
			cmd := exec.Command("sh", "-c", shellStr)
			//cmd := exec.Command(key)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s shellStr: %s\n", shellStr, err.Error())
			}
			log.Printf("shellStr "+shellStr+": %s\n", out)
		} else {
			log.Println("shellStr is nil")
		}

	} else {
		log.Println("Only support Post")
		fmt.Fprintf(w, "Only support post")
	}

}

//处理错误函数
func checkErrs(err error) {
	if err != nil {
		panic(err)
		log.Println("Read failed:", err)
	}
}
