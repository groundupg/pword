package main

import (
	"fmt"
	"os"
	"slices"
	"flag"
)

func PasswordPath() (string, error) {
	p, flag := os.LookupEnv("XDG_DATA_HOME")
	p = p + "/.local/share/pword"
	if flag == false {
		p = os.Getenv("HOME") + "/.local/share/pword"
	}
	if _, err := os.Stat(p); os.IsNotExist(err) == true {
		err = os.Mkdir(p, 0700)
		if err != nil {
			return "", err
		}
	}
	return p, nil
}

func MakePasswordFile(name string) error {
	p, err := PasswordPath()
	if err != nil {
		return err
	}
	f, err := os.Create(p + name)
	if err != nil {
		return err
	}
	if err = f.Chmod(0600); err != nil {
		return err
	}
	return nil
}

func NewPassword(file, key, pass string) error {
	p, err := PasswordPath()
	if err != nil {
		return err
	}
	fp := p + "/" + file
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	f.WriteString(key + " " + pass + "\n")
	return nil
}

func GetFile(file string) ([]byte, error) {
	p, err := PasswordPath()
	if err != nil {
		return nil, err
	}
	fp := p + "/" + file
	b, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetPassword(file, key, pass string) error {
	b, err := GetFile(file)
	if err != nil {
		return err
	}
	fmt.Println(b)
	return nil
}


func HandleArgs(args []string) error {
	ns := "user"
	if slices.Contains(args, "-ns") {
		ns = args[slices.Index(args, "-ns") + 1]
	}
	
	switch args[0] {
		case "-n":
			err := NewPassword(, args[1])
	}
}

func main() {
	var ns, new string
	args := os.Args[1:]
	fmt.Println(len(args))
}
