package main

import "github.com/rivo/tview"

func main() {
	box := tview.NewBox().SetTitle("Log In").SetBorder(true)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
