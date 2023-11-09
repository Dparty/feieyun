package feieyun

type Content interface {
	String() string
}

type Text struct {
	content string
}

func (t *Text) String() string {
	return t.content
}

func (t *Text) SetContent(content Content) {
	t.content = content.String()
}

type Center struct {
	content string
}

func (c *Center) String() string {
	return "<C>" + c.content + "</C>"
}

func (c *Center) SetContent(content Content) {
	c.content = content.String()
}

type Div struct {
	Width int64
}

func (d *Div) String() string {
	div := ""
	for i := 0; i < int(d.Width); i++ {
		div += "-"
	}
	return div
}

type Line struct {
	content Content
}

func (l *Line) String() string {
	return l.content.String() + "<BR>"
}
