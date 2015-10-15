package widgets

import (
	"strings"

	"github.com/kildevaeld/go-acsii"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Checkbox struct {
	Message             string
	Choices             []string
	Value               []string
	SelectedIndicator   string
	UnselectedIndicator string
	Config              WidgetConfig
}

func contains(haystack []string, needle string) int {
	for i, n := range haystack {
		if needle == n {
			return i
		}
	}
	return -1
}

func (c *Checkbox) Run() {
	choices := c.Choices
	var results []string
	config := c.Config
	if config.Writer == nil {
		config = DefaultConfig
		c.Config = config
	}

	if c.SelectedIndicator == "" {
		c.SelectedIndicator = acsii.CheckboxCircleOn
	}
	if c.UnselectedIndicator == "" {
		c.UnselectedIndicator = acsii.CheckboxCircleOff
	}

	writer := config.Writer

	cursor := acsii.Cursor{writer}

	cursor.Hide()

	write(writer, "%s\n", config.MessageColor.Color(c.Message))

	for i, s := range choices {
		if i == len(choices)-1 {
			c.highlight_line(results, s)
		} else {
			c.print_line(results, s)
		}
		write(writer, "\n")
	}
	l := len(choices)

	cursor.Up(1)
	curPos := l - 1
	for {
		a, k, e := tm.GetChar()
		if e != nil {
			return
		}

		tm.HandleSignals(a)

		if k == tm.UpKeyCode && curPos != 0 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(results, choices[curPos])

			curPos = curPos - 1
			cursor.Up(1).Backward(len(choices[curPos+1]) + 3)

			c.highlight_line(results, choices[curPos])

		} else if k == tm.DownKeyCode && curPos < l-1 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(results, choices[curPos])

			curPos = curPos + 1
			cursor.Down(1).Backward(len(choices[curPos-1]) + 3)

			c.highlight_line(results, choices[curPos])
		} else if a == tm.Enter {
			break
		} else if a == tm.Space {
			cursor.Backward(len(choices[curPos]) + 3)
			if i := contains(results, choices[curPos]); i > -1 {
				results = append(results[:i], results[i+1:]...)
				c.print_line(results, choices[curPos])
			} else {
				results = append(results, choices[curPos])
				c.select_line(choices[curPos])
			}
		}
	}
	c.Value = results

	cursor.Down(l - curPos)

	for l > -1 {
		cursor.Up(1)
		write(writer, tm.ClearLine)
		//c.Theme.Write([]byte(tm.ClearLine))
		l = l - 1
	}
	vals := strings.Join(results, ", ")
	write(writer, "%s %s\n", config.MessageColor.Color(c.Message), config.HighlightColor.Color(vals))

	cursor.Show()
	return
}

func (c *Checkbox) highlight_line(results []string, s string) {
	i := c.UnselectedIndicator
	if contains(results, s) > -1 {
		i = c.SelectedIndicator
	}
	write(c.Config.Writer, c.Config.HighlightColor.Color(" %s %s", i, s))
}

func (c *Checkbox) print_line(results []string, s string) {
	i := c.UnselectedIndicator
	if contains(results, s) > -1 {
		i = c.SelectedIndicator
	}
	write(c.Config.Writer, c.Config.StdinColor.Color(" %s %s", i, s))
	//c.Theme.Printf("   %s", s)
}

func (c *Checkbox) select_line(s string) {
	write(c.Config.Writer, " %s %s", c.Config.HighlightColor.Color(c.SelectedIndicator), c.Config.StdinColor.Color(s))
}
