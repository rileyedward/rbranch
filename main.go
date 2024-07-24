package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Window Size */
const WIDTH = 5
const HEIGHT = 12

/* Lipgloss Styles */
var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(1)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#7B68EE"))
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 1, 1)
)

/* Command Flags */
var cFlag bool
var dFlag bool

/*
Initializes and parses the different command
flags that can be passed to the program.
*/
func initializeCommandFlags() {
	flag.BoolVar(&cFlag, "c", false, "Copy")
	flag.BoolVar(&dFlag, "d", false, "Delete")
	flag.Parse()
}

/*
Parses the output of the `git branch` command
of the current git repository and returns an
array of the available branches.
*/
func getAvailableBranches() []string {
	output, err := exec.Command("git", "branch").CombinedOutput()

	if err != nil {
		fmt.Println("You must be in a git repository to run this command")
	}

	var branches []string

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		branch := strings.TrimSpace(strings.TrimPrefix(line, "*"))

		if len(branch) > 0 {
			branches = append(branches, branch)
		}
	}

	return branches
}

/*
Configures the selection list and basic settings
for the selection list window that will be
displayed to the user to select an available branch.
*/
func configureSelectionList(items []list.Item) list.Model {
	selectionList := list.New(items, itemDelegate{}, WIDTH, HEIGHT)
	selectionList.Title = "Select a branch"
	selectionList.SetShowStatusBar(false)
	selectionList.SetFilteringEnabled(false)

	return selectionList
}

/*
Uses the list package from bubbletea to create
a list of items from the array of available git
branches that can be selected by the user.
*/
type item string

func (i item) FilterValue() string { return "" }

func buildBranchSelectionList(branches []string) []list.Item {
	var items []list.Item

	for _, branch := range branches {
		items = append(
			items,
			list.Item(item(branch)),
		)
	}

	return items
}

/*
Defines the delegate and the styles for the
selection list that will be displayed to the
user to select an available branch.
*/
type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

/* Bubbletea Selection Model */
type model struct {
	list     list.Model
	choice   string
	quitting bool
}

/*
Initializes the model that is used
by the bubbletea package.
*/
func (m model) Init() tea.Cmd {
	return nil
}

/*
Updates the model based on the input from
the user or the selection they have made.
*/
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		/* Quit the program */
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		/* Make a selection */
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

/*
Handles the view of the model and the action
that follows the user's selection from the
available branches and the given options.
*/
func (m model) View() string {
	if m.choice != "" {
		executeGitCommand(m.choice)
	}

	if m.quitting {
		return quitTextStyle.Render("See ya later!")
	}

	return "\n" + m.list.View()
}

/*
Run a simple git command based on the flags that
were passed in by the user.
*/
func executeGitCommand(selectedBranch string) {
	/* Copy */
	if cFlag {
		copyBranch(selectedBranch)
		return
	}

	/* Delete */
	if dFlag {
		deleteBranch(selectedBranch)
		return
	}

	/* Checkout (Default) */
	checkoutBranch(selectedBranch)
}

/* Checkout Git Branch */
func checkoutBranch(selectedBranch string) {
	_, err := exec.Command("git", "checkout", selectedBranch).CombinedOutput()

	if err != nil {
		quitTextStyle.Render("There was an unexpected error while checking out this branch.")
		return
	}
}

/* Copy Git Branch */
func copyBranch(selectedBranch string) {
	clipboard.WriteAll(selectedBranch)
}

/* Delete Git Branch */
func deleteBranch(selectedBranch string) {
	_, err := exec.Command("git", "branch", "-D", selectedBranch).CombinedOutput()

	if err != nil {
		quitTextStyle.Render("There was an unexpected error while deleting out this branch.")
		return
	}
}

func main() {
	initializeCommandFlags()
	branches := getAvailableBranches()
	items := buildBranchSelectionList(branches)
	selectionList := configureSelectionList(items)

	m := model{
		list: selectionList,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("fatal")
		os.Exit(1)
	}
}
