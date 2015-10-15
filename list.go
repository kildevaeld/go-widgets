package widgets

import (
	"github.com/kildevaeld/go-acsii"
	tm "github.com/kildevaeld/prompt/terminal"
)

type List struct {
	Message   string
	Choices   []string
	Value     string
	Indicator string
	Config    WidgetConfig
}

func (c *List) Run() {
	choices := c.Choices
	config := c.Config
	if config.Writer == nil {
		config = DefaultConfig
		c.Config = config
	}

	if c.Indicator == "" {
		c.Indicator = acsii.Pointer
	}

	writer := config.Writer

	cursor := acsii.Cursor{writer}

	cursor.Hide()

	write(writer, "%s\n", config.MessageColor.Color(c.Message))

	for i, s := range choices {
		if i == len(choices)-1 {
			c.highlight_line(s)
		} else {
			c.print_line(s)
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
			c.print_line(choices[curPos])

			curPos = curPos - 1
			cursor.Up(1).Backward(len(choices[curPos+1]) + 3)

			c.highlight_line(choices[curPos])

		} else if k == tm.DownKeyCode && curPos < l-1 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(choices[curPos])

			curPos = curPos + 1
			cursor.Down(1).Backward(len(choices[curPos-1]) + 3)

			c.highlight_line(choices[curPos])
		} else if a == tm.Enter {
			break
		}
	}

	c.Value = choices[curPos]
	cursor.Down(l - curPos)

	for l > -1 {
		cursor.Up(1)
		write(writer, tm.ClearLine)
		//c.Theme.Write([]byte(tm.ClearLine))
		l = l - 1
	}
	write(writer, "%s %s\n", config.MessageColor.Color(c.Message), config.HighlightColor.Color(c.Value))

	cursor.Show()
	return
}

func (c *List) highlight_line(s string) {
	write(c.Config.Writer, c.Config.HighlightColor.Color(" %s %s", c.Indicator, s))
}

func (c *List) print_line(s string) {
	write(c.Config.Writer, c.Config.StdinColor.Color("   %s", s))
	//c.Theme.Printf("   %s", s)
}