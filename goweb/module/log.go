package module

import (
	"fmt"
	"io"
	"log"
	"os"
)

// LOG is
type LOG struct {
	Title string
	Log   string
	Etc   string
}

func (l LOG) SetLog() {
	szLog := fmt.Sprintf("[%s]%s (%s)", l.Title, l.Log, l.Etc)
	println(szLog)

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

	log.Println(szLog)
}
