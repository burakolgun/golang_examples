package gui

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"os"
)

var channel = make(chan bool)

func ExampleQt() {
	ui.RunEx(os.Args, mainWindow)
}

func mainWindow() {
	//first button
	buttonOne := ui.NewPushButton()
	buttonOne.SetText("First Button")
	buttonTwo := ui.NewPushButton()
	buttonTwo.SetText("Clear")

	//Create a text box
	textBox := ui.NewPlainTextEdit()
	textBox.SetReadOnly(true)
	textBox.AppendPlainText("Can you try the buttons?")
	go listenButton(textBox)

	//if buttonOne clicked
	buttonOne.OnClicked(func() {
		channel <- true
	})

	buttonTwo.OnClicked(func() {
		textBox.Clear()
	})

	//create horizontal row
	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(buttonOne)
	hbox.AddWidget(buttonTwo)

	//create vertical lines
	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddWidget(textBox)

	//Show the elements on the window
	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.Show()
}

func listenButton(edit *ui.QPlainTextEdit) {
	i := 0
	for {
		<-channel
		i++
		edit.AppendPlainText(fmt.Sprintf("Clicked -- %d", i))
	}
}
