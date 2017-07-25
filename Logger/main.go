package main

import (
	"io"
	"log"
	"logger/module"
	"os"
)

func main() {
	module.ExampleWriter()

	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 표준로거를 파일로그로 변경
	//log.SetOutput(fpLog)
	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)

	run()
	//log.Fatalln("fatal log") // fatal이 되면 바로 종료되네..
	log.Println("End of Program")
}

func run() {
	log.Print("Test")
}
