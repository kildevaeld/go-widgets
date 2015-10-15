package widgets

import (
	"github.com/kildevaeld/go-acsii"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Input struct {
	Message string
	Value   string
	Config  WidgetConfig
}

func (c *Input) Run() {
	config := c.Config
	if config.Writer == nil {
		config = DefaultConfig
	}

	writer := config.Writer

	cursor := acsii.Cursor{writer}

	write(writer, "%s ", config.MessageColor.Color(c.Message))

	x := 0
	var buffer []byte

	for {
		a, k, _ := tm.GetChar()
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
			c.Value = string(buffer)
			break
		} else if k == tm.RightKeyCode {
			if x < len(buffer)-1 {
				x++
				cursor.Forward(1)
			}
			continue
		} else if k == tm.LeftKeyCode {
			if x > 0 {
				x--
				cursor.Backward(1)
			}
			continue
		} else if k == tm.UpKeyCode || k == tm.DownKeyCode {
			continue
		}

		if len(buffer) == x {
			buffer = append(buffer, byte(a))
		} else {
			buffer[x] = byte(a)
		}
		write(writer, config.StdinColor.Color(string(a)))

		x++
	}

	cursor.Backward(x)
	write(writer, "%s\n", config.HighlightColor.Color(string(buffer)))

}
