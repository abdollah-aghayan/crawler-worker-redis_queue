package config

import "fmt"

// only use for demo project
const (
	//HTTPPort http port
	HTTPPort = "6050"
)

//Init manual initialize
func Init() {
	fmt.Println("config loader...")
}
