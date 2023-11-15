package feieyun

type PrintAble interface {
	String() string
	Text() string
	Length() int
	Width() int
}

const BR = "<BR>"

type Text struct {
	Content string
}

func (t *Text) Text() string {
	return t.Content
}

func (t *Text) Width() int {
	var length int
	for _, c := range t.Content {
		if IsChineseChar(string(c)) {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}

func (t *Text) String() string {
	return t.Content
}

func (t *Text) Length() int {
	return len(t.Content)
}

type Center struct {
	Content PrintAble
}

func (c *Center) Text() string {
	return c.Content.Text()
}

func (t *Center) Width() int {
	return t.Content.Width()
}

func (c *Center) Length() int {
	return c.Content.Length()
}

func (c *Center) String() string {
	return "<C>" + c.Content.String() + "</C>"
}

type Bold struct {
	Content PrintAble
}

func (b *Bold) Text() string {
	return b.Content.Text()
}

func (t *Bold) Width() int {
	return t.Content.Width() * 2
}

func (c *Bold) Length() int {
	return c.Content.Length()
}

func (b *Bold) String() string {
	return "<B>" + b.Content.String() + "</B>"
}

type CenterBold struct {
	Content PrintAble
}

func (c *CenterBold) Text() string {
	return c.Content.Text()
}

func (c *CenterBold) Width() int {
	return c.Content.Width() * 2
}

func (c *CenterBold) Length() int {
	return c.Content.Length()
}

func (c CenterBold) String() string {
	return "<CB>" + c.Content.String() + "</CB>"
}

type Cut struct{}

func (Cut) Text() string {
	return ""
}

func (Cut) String() string {
	return "<CUT>"
}

func (Cut) Width() int {
	return 0
}

func (Cut) Length() int {
	return 0
}

type Qrcode struct {
	Content string
}

func (q Qrcode) String() string {
	return "<QR>" + q.Content + "</QR>"
}

func NewDiv(width int) *Div {
	return &Div{width}
}

type Div struct {
	width int
}

func (d *Div) Text() string {
	return ""
}

func (d *Div) Width() int {
	return d.Length()
}

func (d *Div) Length() int {
	return d.width
}

func (d *Div) String() string {
	div := ""
	for i := 0; i < int(d.Width()); i++ {
		div += "-"
	}
	div += BR
	return div
}

type Line struct {
	Content PrintAble
}

func (l Line) Text() string {
	return l.Content.Text()
}

func (l Line) Width() int {
	return l.Content.Width()
}

func (l Line) Length() int {
	return l.Content.Length()
}

func (l Line) String() string {
	return l.Content.String() + BR
}

type PrintContent struct {
	Lines []PrintAble
}

func (p *PrintContent) AddLines(contents ...PrintAble) {
	for _, content := range contents {
		p.Lines = append(p.Lines, Line{content})
	}
}

func (p *PrintContent) AddDiv(width int) {
	p.AddLines(NewDiv(width))
}

func (p *PrintContent) String() string {
	s := ""
	for _, line := range p.Lines {
		s += line.String()
	}
	return s
}

type Row struct {
	Columns []Column
}

func (r *Row) Width() int {
	width := 0
	for _, column := range r.Columns {
		width += column.Width()
	}
	return width
}

func (r *Row) Length() int {
	return r.Width()
}

func (r *Row) AddColumns(columns ...Column) {
	r.Columns = append(r.Columns, columns...)
}

func (r *Row) String() {

}

type Column struct {
	W int
	P PrintAble
}

func (c Column) Strings() []string {
	return []string{}
}

func (c Column) String() string {
	return ""
}

func (c Column) Width() int {
	return c.W
}

func (c Column) Length() int {
	return c.Width()
}
