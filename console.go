package main

import (
	"os/exec"
	"os"
	"runtime"
)

var clearFuncMap = map[string]func(){}

func init() {
	clearFuncMap["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearFuncMap["darwin"] = clearFuncMap["linux"]
	clearFuncMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearConsole() {
	if clearFunc, ok := clearFuncMap[runtime.GOOS]; ok {
		clearFunc()
	}
}
