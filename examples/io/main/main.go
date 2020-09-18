package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := []string{
		"텍스트0",
		"텍스트1",
	}
	//기존에 쓰여진것을 저장
	o := []string{}
	// 파일을 읽기 전용으로 연다
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	//한줄씩 읽어주는 scanner함수
	scanner := bufio.NewScanner(file)
	//읽을 줄이 남아있으면 true 아니면 false
	for scanner.Scan() {
		// scanner.Text() 한 줄 string타입 리턴
		o = append(o, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	o = append(o, s...)

	//파일을 쓰기 전용으로 연다
	file, err = os.Create("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range o {
		// file에다가 한줄씩 입력
		_, err := fmt.Fprintf(file, "%s\n", line)
		if err != nil {
			log.Fatal(err)
		}
	}

}
