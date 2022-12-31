package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type CallerInfo struct {
	PackageName string
	FuncName    string
	FileName    string
	Line        int
}

func retrieveCallInfo() CallerInfo {
	pc, file, line, _ := runtime.Caller(2)
	//_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return CallerInfo{
		PackageName: packageName,
		FuncName:    funcName,
		FileName:    file,
		Line:        line,
	}
}

const (
	PATH = "./logs"
)

func createIfNotExists(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	err := createIfNotExists(PATH)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/logs.log", PATH), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
}

func Trace(line string) {
	if os.Getenv("VERBOSE_LEVEL") == "TRACE" {
		callerInfo := retrieveCallInfo()
		log.Printf("TRACE %s:%d %s %s", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, line)
		fmt.Printf("TRACE %s:%d %s %s\n", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, line)
	}
}

func Debug(line string) {
	if os.Getenv("VERBOSE_LEVEL") == "TRACE" || os.Getenv("VERBOSE_LEVEL") == "DEBUG" {
		callerInfo := retrieveCallInfo()
		log.Printf("DEBUG %s:%d %s %s", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, line)
		fmt.Printf("DEBUG %s:%d %s %s\n", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, line)
	}
}

func Info(line string) {
	if os.Getenv("VERBOSE_LEVEL") == "TRACE" || os.Getenv("VERBOSE_LEVEL") == "DEBUG" || os.Getenv("VERBOSE_LEVEL") == "INFO" {
		callerInfo := retrieveCallInfo()
		log.Printf("INFO %s:%d %s %s", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, line)
		fmt.Printf("INFO %s:%d %s %s\n", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, line)
	}
}

func Error(short string, err error) {
	if os.Getenv("VERBOSE_LEVEL") == "TRACE" || os.Getenv("VERBOSE_LEVEL") == "DEBUG" || os.Getenv("VERBOSE_LEVEL") == "INFO" || os.Getenv("VERBOSE_LEVEL") == "ERROR" {
		callerInfo := retrieveCallInfo()
		log.Printf("ERROR %s:%d %s %s\n%s", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, short, err)
		fmt.Printf("ERROR %s:%d %s %s\n%s\n", callerInfo.FileName, callerInfo.Line, callerInfo.FuncName, short, err)
	}
}
