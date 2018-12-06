package main

import (
	"time"
	"log"
)

func main()  {
	LongRunningFunction()
}

func LongRunningFunction()  {
	defer TimeToken(time.Now(),"LongRunningFunction")
	time.Sleep(2*time.Second)
}

func TimeToken(t  time.Time,name string)  {
	elapsed := time.Since(t)
	log.Printf("Time: %s tooks %s\n",name,elapsed)
}