package boomer

import (
	"fmt"

	"boomer/model"

	"github.com/myzhan/boomer"
)

func setBoomerMode(tc model.TestConfiguration, tasks []*boomer.Task) {
	switch tc.Mode {
	case "distributed":
		{
			distributed(tasks)
		}

	case "standalone":
		{
			standalone(tc, tasks)
		}

	default:
		{
			fmt.Printf("Invalid Boomer mode: [%s]", tc.Mode)
		}
	}
}
