package boomer

import (
	"fmt"
	"strconv"

	"boomer/globals"
	"boomer/utils"

	"github.com/myzhan/boomer"
)

func distributed(tasks []*boomer.Task) {
	fmt.Println("Boomer: Distributed Mode")
	port, err := strconv.ParseInt(utils.GetEnv("LOCSUT_PORT", "5557"), 10, 32)
	if err != nil {
		port = 5557
	}
	globals.GlobalBoomer = boomer.NewBoomer(utils.GetEnv("LOCUST_ADDRESS", "localhost"), int(port))
	globals.GlobalBoomer.Run(tasks...)
	waitForQuit()
}
