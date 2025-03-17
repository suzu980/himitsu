package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	home "github.com/suzu980/himitsu/features/home"
)

func main() {
	p := tea.NewProgram(home.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Print("Whoops, an error happened")
		os.Exit(1)
	}

}
