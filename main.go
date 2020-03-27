package main

import (
	_ "LiveAssistant/backend"
	"github.com/go-qamel/qamel"
	"os"
)

func main() {
	// Create application
	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Live Assistant")

	// Create a QML viewer
	view := qamel.NewViewerWithSource("qrc:/res/main.qml")
	view.Show()

	// Exec app
	app.Exec()
}
