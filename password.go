package widgets

import (
	"github.com/kildevaeld/go-acsii"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Password struct {
	Message   string
	Value     string
	Config    WidgetConfig
	Indicator string
}

func (c *Password) Run() {

	config := c.Config
	if config.Writer == nil {
		config = DefaultConfig
	}

	if c.Indicator == "" {
		c.Indicator = acsii.Bullet
	}

	writer := config.Writer

	cursor := acsii.Cursor{writer}

	write(writer, "%s ", config.MessageColor.Color(c.Message))

	//c.Theme.Printf("%s ", label)
	x := 0

	buffer := ""

	for {
		a, _, _ := tm.GetChar()
		tm.HandleSignals(a)
		if a == tm.Backspace {
			if x == 0 {
				continue
			}

			write(writer, "\b \b")

			x--
			buffer = buffer[0:x]
			continue

		} else if a == tm.Enter {
			c.Value = buffer
			break
		}

		buffer += string(a)

		write(writer, config.StdinColor.Color(c.Indicator))

		x++
	}

	cursor.Backward(x)
	str := ""
	for x > 0 {
		str += "*"
		x--
	}
	write(writer, "%s\n", config.HighlightColor.Color(str))

}
