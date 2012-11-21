package gokeeper

import (
	"bufio"
	"fmt"
	"os"
	"terminal"
)

var (
	STORAGE_PATH string = ".gokeeper.db"
	KEY          []byte
)

func GetPass(msg string) string {
	fmt.Printf(msg)
	key, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		return ""
	}
	return string(key)
}

func GetInput(msg string) string {
	fmt.Printf(msg)
	res, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(res)
}

func add(s *Storage) error {
	key := GetInput("Key : ")
	data := []byte(GetInput("Data : "))
	err := s.Put(key, data)
	//s.Save()
	return err
}

func list(s *Storage) {
	for key, _ := range s.Data {
		fmt.Println(key)
	}
}

func del(s *Storage) {
	key := GetInput("Key : ")
	delete(s.Data, key)
	//s.Save()
}

func show(s *Storage, key string) error {
	data, err := s.Get(key)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func Main() {
	KEY = Skein256([]byte(GetPass("Master key : ")))
	var command string = ""
	storage := NewStorage(STORAGE_PATH)
	for {
		command = GetInput("> ")
		switch command {
		case "add", "a":
			err := add(storage)
			if err != nil {
				fmt.Println(err)
			}
		case "list", "l":
			list(storage)
		case "del", "d":
			del(storage)
		case "save", "s":
			storage.Save()
		case "quit", "q":
			return
		default:
			err := show(storage, command)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
