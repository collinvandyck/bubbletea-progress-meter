package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model, err := tea.NewProgram(newMeter(),
		tea.WithInputTTY(),
	).Run()
	if err == nil {
		err = model.(meter).err
	}
	if err != nil {
		log.Fatal(err)
	}
}

func newMeter() meter {
	return meter{
		in:     bufio.NewScanner(os.Stdin),
		widget: progress.New(progress.WithDefaultGradient()),
	}
}

type meter struct {
	in     *bufio.Scanner
	err    error
	pct    float64
	widget progress.Model
}

type update struct {
	pct float64
	err error
}

// Init implements tea.Model.
func (m meter) Init() tea.Cmd {
	return m.updateCmd
}

func (m meter) updateCmd() tea.Msg {
	if m.in.Scan() {
		s := strings.TrimSpace(m.in.Text())
		pct, err := strconv.ParseFloat(s, 64)
		return update{pct, err}
	}
	time.Sleep(time.Millisecond * 100)
	return nil
}

// Update implements tea.Model.
func (m meter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "ctrl+d":
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, tea.Quit
	case update:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
		m.pct = msg.pct
		if m.pct >= 1.0 {
			return m, tea.Quit
		}
		return m, m.updateCmd
	}
	return m, nil
}

// View implements tea.Model.
func (m meter) View() string {
	return m.widget.ViewAs(m.pct)
}
