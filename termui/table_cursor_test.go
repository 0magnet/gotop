package termui

import (
	"image"
	"testing"

	ui "github.com/gizak/termui/v3"
)

// The selected process row rendered as unreadable dark-on-dark: the style set
// Fg to the cursor colour and left Bg at ColorClear, then asked the terminal to
// REVERSE it. Reverse swaps the pair, so the text was drawn in ColorClear —
// the terminal's default background, black on any dark theme — over the cursor
// colour. The first row is selected at startup, so the top entry was the one
// nobody could read.
func TestSelectedRowIsReadable(t *testing.T) {
	tbl := NewTable()
	tbl.ShowCursor = true
	tbl.CursorColor = ui.ColorBlue
	tbl.Header = []string{"pid", "name"}
	tbl.Rows = [][]string{{"1", "alpha"}, {"2", "beta"}}
	tbl.ColWidths = []int{5, 10}
	tbl.ColGap = 1
	tbl.UniqueCol = 0
	tbl.SelectedRow = 0
	tbl.SetRect(0, 0, 30, 10)

	buf := ui.NewBuffer(tbl.GetRect())
	tbl.Draw(buf)

	// The selected row is drawn at Inner.Min.Y + 1 (header occupies +0).
	y := tbl.Inner.Min.Y + 1
	var got ui.Style
	found := false
	for x := tbl.Inner.Min.X; x < tbl.Inner.Max.X; x++ {
		c := buf.GetCell(image.Pt(x, y))
		if c.Rune == 'a' { // first rune of "alpha"
			got = c.Style
			found = true
			break
		}
	}
	if !found {
		t.Fatal("selected row text not found on its row — the draw moved")
	}

	if got.Modifier&ui.ModifierReverse != 0 {
		t.Errorf("selected row still uses ModifierReverse (style %+v): reverse over a "+
			"ColorClear background paints the text in the terminal default background", got)
	}
	if got.Bg != tbl.CursorColor {
		t.Errorf("selected row Bg = %v, want the cursor colour %v", got.Bg, tbl.CursorColor)
	}
	if got.Fg == ui.ColorClear || got.Fg == got.Bg {
		t.Errorf("selected row Fg = %v against Bg = %v: unreadable", got.Fg, got.Bg)
	}
}
