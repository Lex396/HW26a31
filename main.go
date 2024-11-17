package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"HW2021/filtering"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func initLoggers(logFile *os.File) {
	nmw := io.MultiWriter(os.Stdout, logFile)

	infoLog = log.New(nmw, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(nmw, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	logFile, err := os.OpenFile("logFile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		errorLog.Panicf("error opening file: %s", err)
	}

	initLoggers(logFile)
	defer logFile.Close()

	numbers := make(chan int)
	go dataSource(numbers)

	filterNegative := make(chan int)
	go filtering.FilterNegative(numbers, filterNegative)

	filterNumberNotMultipleThree := make(chan int)
	go filtering.FilterNumberNotMultipleThree(filterNegative, filterNumberNotMultipleThree)

	buff := make(chan int)
	go filtering.Buffering(filterNumberNotMultipleThree, buff)

	go consumer(buff)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	select {
	case sig := <-c:
		infoLog.Printf("Got %s signal. Aborting ... \n", sig)
		os.Exit(0)
	}

}

func dataSource(numb chan int) {
	menu := "Menu:\n \"Menu\" - open the menu\n \"buffer\" - enter the buffer value\n \"timer\" - enter the timer value\n \"exit\" - exiting the program"
	ent := "enter a number"
	fmt.Println(menu)
	fmt.Println(ent)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		switch scanner.Text() {
		case "menu":
			fmt.Println(menu)
			fmt.Println(ent)
		case "exit":
			infoLog.Println("exit from the program")
			os.Exit(0)
		case "buffer":
			fmt.Println("enter the buffer value")
			scanner.Scan()
			buff, err := strconv.Atoi(scanner.Text())
			if err != nil {
				errorLog.Println("only need to enter an integer")
				continue
			}
			filtering.BufferSize = buff
		case "timer":
			fmt.Println("enter a number after how many seconds to clear the buffer")
			scanner.Scan()
			sec, err := strconv.Atoi(scanner.Text())
			if err != nil {
				errorLog.Println("only need to enter an integer T")
				continue
			}
			filtering.TimeBufferClear = time.Duration(sec) * time.Second
		default:
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				errorLog.Println("only need to enter integers")
				continue
			}
			numb <- num
		}

	}

}

func consumer(numbIn chan int) {
	for numb := range numbIn {
		infoLog.Printf("Получены данные: %d\n", numb)
	}
}
