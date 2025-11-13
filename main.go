package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	NameSpace  string
	NewFlag    bool
	DeleteFlag bool
	ListFlag   bool
)

func InitFlags() {
	// Initialises flags
	flag.StringVar(&NameSpace, "ns", "user", "Namespace")
	flag.StringVar(&NameSpace, "namespace", "user", "Namespace")
	flag.BoolVar(&NewFlag, "n", false, "create a new password")
	flag.BoolVar(&NewFlag, "new", false, "create a new password")
	flag.BoolVar(&DeleteFlag, "d", false, "delete a password")
	flag.BoolVar(&DeleteFlag, "delete", false, "delete a password")
	flag.BoolVar(&ListFlag, "l", false, "list all passwords")
	flag.BoolVar(&ListFlag, "list", false, "list all passwords")
	flag.Parse()
}

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
	f, err := os.Create(p + "/" + name)
	if err != nil {
		return err
	}
	if err = f.Chmod(0600); err != nil {
		return err
	}
	fmt.Printf("CREATED NAMESPACE %s\n\r", name)
	return nil
}

func GetFile(file string) (string, error) {
	p, err := PasswordPath()
	if err != nil {
		return "", err
	}
	fp := p + "/" + file
	f, err := os.ReadFile(fp)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func HandleNewPassword(key, value string) error {
	if err := NewPassword(key, value); os.IsNotExist(err) == true {
		if err = MakePasswordFile(NameSpace); err != nil {
			return err
		}
		return NewPassword(key, value)
	}
	return nil
}

func HandleArgs() error {
	args := flag.Args()
	if NewFlag == true {
		return HandleNewPassword(args[0], args[1])
	}
	if DeleteFlag == true {
		return HandleDelete
	}
	return PrintPassword(args[0])
}

func main() {
	InitFlags()
	err := HandleArgs()
	if err != nil {
		fmt.Println(err)
	}
}
