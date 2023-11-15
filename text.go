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

type Center struct {
	Content PrintAble
}

func (c *Center) String() string {
	return "<C>" + c.Content.String() + "</C>"
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

type Cut struct{}

func (Cut) String() string {
	return "<CUT>"
}

type Qrcode struct {
	Content string
}

func (q Qrcode) String() string {
	return "<QR>" + q.Content + "</QR>"
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

func (l Line) String() string {
	return l.content.String() + BR
}

type PrintContent struct {
	Lines []PrintAble
}

func (p *PrintContent) AddLine(contents ...PrintAble) {
	for _, content := range contents {
		p.Lines = append(p.Lines, Line{content})
	}
}

func (p *PrintContent) AddDiv(width int64) {
	p.AddLine(&Div{Width: width})
}

func (p *PrintContent) String() string {
	s := ""
	for _, line := range p.Lines {
		s += line.String()
	}
	return s
}
