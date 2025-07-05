package menu

import (
	"fmt"

	"github.com/CKroes97/go-gadgematron/internal/modules/timecheck"
	// Later you can add more modules here, e.g. pi, weather, etc.
)

func ShowMainMenu() {
	fmt.Println("=== Go-Gadgematron ===")
	fmt.Println("Choose a module:")
	fmt.Println("1) Check time")
	fmt.Print("Enter choice: ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	switch choice {
	case 1:
		timecheck.Run()
	default:
		fmt.Println("Invalid choice")
	}
}
