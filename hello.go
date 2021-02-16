package main

import (
	"fmt"
	"encoding/json"
)

type Runner struct {
	Name string
	Age int
}

type Runnig struct {
	Distance int
}

func changePos(x, y string)(string, string){
	return y, x
}

func main() {
	runner_1 := Runner{"james", 31}
	fmt.Println("Runner", runner_1)
	fmt.Println("Runner", runner_1.Age)
	panic(runner_1)
	res, err := json.Marshal(runner_1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Runner json", string(res))
}
