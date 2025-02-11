package boomer

import (
	"fmt"

	"boomer/globals"
	"boomer/model"

	"github.com/myzhan/boomer"
)

func standalone(tc model.TestConfiguration, tasks []*boomer.Task) {
	fmt.Println("Boomer: Standalone Mode")
	globals.GlobalBoomer = boomer.NewStandaloneBoomer(int(tc.NumberOfUsers), tc.SpawnRate)
	go stopStandalone(int(tc.TestDuration))
	globals.GlobalBoomer.Run(tasks...)
}
