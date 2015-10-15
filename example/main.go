package main

import "github.com/kildevaeld/go-widgets"

func main() {

	// Input
	input := widgets.Input{
		Message: "Enter name",
	}
	input.Run()

	confirm := widgets.Confirm{
		Message: "Confirm",
	}
	confirm.Run()

	password := widgets.Password{
		Message: "Password",
	}
	password.Run()

	list := widgets.List{
		Message: "List",
		Choices: []string{"Test", "Test 2"},
	}
	list.Run()

	check := widgets.Checkbox{
		Message: "Checkbox",
		Choices: []string{"Test", "Test 2"},
	}
	check.Run()
}
