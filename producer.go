package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

// This function will produce event
func newProducer(i int){
	//var err error
	fmt.Println(i)
	reader := bufio.NewReader(os.Stdin)
	kafka := newKafkaSyncProducer()

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, "->")
		cmd := args[0]
		fmt.Println(cmd)

		switch cmd{
		case "create":
			if len(args) == 2{
				accName := args[1]
				event := newCreateAccountEvent(accName)
				sendMsg(kafka, event)
			}else{
					fmt.Println("Only specify create->Account Name")
			}
		}

	}
}
