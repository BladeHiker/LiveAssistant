package main

import (
	_ "LiveAssistant/backend"
	"github.com/go-qamel/qamel"
)

func main() {
	engine := qamel.NewEngine()
	engine.Load("qrc:/res/main.qml")
}
