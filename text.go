package feieyun

type PrintAble interface {
	String() string
}

const BR = "<BR>"

type Text struct {
	Content string
}

func (t *Text) String() string {
	return t.Content
}

func (t *Text) SetContent(content string) {
	t.Content = content
}

type Center struct {
	Content PrintAble
}

func (c *Center) String() string {
	return "<C>" + c.Content.String() + "</C>"
}

func (c *Center) SetContent(content PrintAble) {
	c.Content = content
}

type Bold struct {
	Content PrintAble
}

func (b *Bold) String() string {
	return "<B>" + b.Content.String() + "</B>"
}

type CenterBold struct {
	Content PrintAble
}

func (c CenterBold) String() string {
	return "<CB>" + c.Content.String() + "</CB>"
}

type Div struct {
	Width int64
}

func (d *Div) String() string {
	div := ""
	for i := 0; i < int(d.Width); i++ {
		div += "-"
	}
	div += BR
	return div
}

type Line struct {
	content PrintAble
}

func (l *Line) String() string {
	return l.content.String() + BR
}
