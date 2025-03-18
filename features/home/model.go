package home

import tea "github.com/charmbracelet/bubbletea"

type HomeState int
type model struct {
	screenState  HomeState
	menu_choices []string
	cursor       int
	selected     map[int]struct{}
}

func InitialModel() model {
	return model{
		menu_choices: []string{"Pomodoro Timer"},
		selected:     make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
