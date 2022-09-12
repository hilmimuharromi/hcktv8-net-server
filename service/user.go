package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Nama string `json:"name"`
}

var registeredUserList []User

type UserService struct {
}

type UserIface interface {
	Register(r *http.Request, w *http.ResponseWriter)
	GetUser(w http.ResponseWriter)
}

func NewUserService() UserIface {
	return &UserService{}
}

func (u *UserService) Register(r *http.Request, w *http.ResponseWriter) {
	decodeData := json.NewDecoder(r.Body)
	var newUser User
	err := decodeData.Decode(&newUser)
	if err != nil {
		fmt.Println("ERROR")
	}
	registeredUserList = append(registeredUserList, newUser)
	fmt.Fprint(*w, newUser.Nama+" berhasil didaftarkan.")
}

func (u *UserService) GetUser(w http.ResponseWriter) {
	var wg sync.WaitGroup
	wg.Add(len(registeredUserList))
	count := 0
	for _, value := range registeredUserList {
		go func(name string) {
			count = count + 1
			PrintName(&wg, &w, count, name)
		}(value.Nama)
	}
	wg.Wait()

	json, err := json.Marshal(registeredUserList)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func PrintName(wg *sync.WaitGroup, w *http.ResponseWriter, number int, name string) {
	fmt.Println(number, name)
	wg.Done()
}
