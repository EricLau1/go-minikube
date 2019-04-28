package utils

import (
	"fmt"
	"time"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"html"
	"github.com/gorilla/mux"
	"github.com/badoux/checkmail"
)

func ShowTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d T%d:%d:%d", now.Day(), int(now.Month()), now.Year(),
		now.Hour(), now.Minute(), now.Second())
}

func HttpInfo(r *http.Request) {
	fmt.Printf("\r\n%s/\t %s%s\t %s\t %s\n", r.Method, r.Host, r.URL, r.Proto, ShowTime())
}

func Console(data interface{}) {
	byteArray, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray))
}

func ExtractId(r *http.Request) uint32 {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	return uint32(id)
}

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}
	return false
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}

func Trim(param string) string {
	return html.EscapeString(strings.TrimSpace(param))
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

