package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DeletePassword(key string) error {
	line, err := GetLineNumber(key)
	if err != nil {
		return err
	}
	if line == 0 {
		fmt.Print("COULD NOT FIND PASSWORD\n")
	}
	ppath, err := PasswordPath()
	if err != nil {
		return err
	}
	cmd := exec.Command("sed", "-i", string(line)+"d", ppath)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("password deleted. ns: %v, key: %s", NameSpace, key)
	return nil
}

func PrintPassword(key string) error {
	f, err := GetFile(NameSpace)
	if err != nil {
		return err
	}
	split := strings.Split(f, "\n")
	for i := 0; i < len(split); i++ {
		if line := strings.Split(split[i], " "); line[0] == key {
			fmt.Println(line[1])
			return nil
		}
	}
	fmt.Printf("COULD NOT FIND PASSWORD FOR %s\n", key)
	return nil
}

func NewPassword(key, pass string) error {
	p, err := PasswordPath()
	if err != nil {
		return err
	}
	fp := p + "/" + NameSpace
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	f.WriteString("\n" + key + " " + pass)
	fmt.Printf("pword created. ns: %v, key: %s\n", NameSpace, key)
	return nil
}

func GetLineNumber(key string) (int, error) {
	f, err := GetFile(NameSpace)
	if err != nil {
		return 0, err
	}
	split := strings.Split(f, "\n")
	for i := 0; i < len(split); i++ {
		if line := strings.Split(split[i], " "); line[0] == key {
			return i + 1, nil
		}
	}
	fmt.Printf("COULD NOT FIND PASSWORD FOR %s\n", key)
	return 0, nil
}
