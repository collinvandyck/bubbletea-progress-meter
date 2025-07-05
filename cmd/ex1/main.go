package main

import (
	"log"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Progress interface {
	// return a value between 0 and 1
	Done() (float64, error)
}

type ProgressFunc func() (float64, error)

func (p ProgressFunc) Done() (float64, error) {
	return p()
}

type tickMsg time.Time

func main() {
	// here you'd pass in your own func or something that implements Progress
	var mut sync.Mutex
	var value float64
	go func() {
		for i := range 100 {
			time.Sleep(10 * time.Millisecond)
			mut.Lock()
			value = float64(i+1) / 100
			mut.Unlock()
		}
	}()
	progMeter := newMeter(ProgressFunc(func() (float64, error) {
		mut.Lock()
		defer mut.Unlock()
		return value, nil
	}))
	prog := tea.NewProgram(progMeter)
	m, err := prog.Run()
	if err == nil {
		err = m.(meter).err
	}
	if err != nil {
		log.Fatal(err)
	}
}

func newMeter(prog Progress) meter {
	return meter{
		callback: prog,
		widget:   progress.New(progress.WithDefaultGradient()),
	}
}

type meter struct {
	callback Progress
	widget   progress.Model // the widget
	pct      float64
	err      error
}

func (m meter) Init() tea.Cmd {
	return tea.Batch(tickCmd(), m.widget.Init())
}

// Update implements tea.Model.
func (m meter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "ctrl+d":
			return m, tea.Quit
		}
	case tickMsg:
		pct, err := m.callback.Done()
		if err != nil {
			m.err = err
			return m, tea.Quit
		}
		if pct >= 1.0 {
			return m, tea.Quit
		}
		m.pct = pct
		return m, tickCmd()
	}
	return m, nil
}

func (m meter) View() string {
	return m.widget.ViewAs(m.pct)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
