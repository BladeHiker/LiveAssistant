package main

import (
	_ "LiveAssistant/backend"
	"github.com/go-qamel/qamel"
	"os"
)

func main() {
	//Create Application
	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Live Assistant")

	engine := qamel.NewEngine()
	engine.Load("qrc:/res/main.qml")

	// Exec app
	app.Exec()
}
