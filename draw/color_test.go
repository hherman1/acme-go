package draw

import (
	"image"
	"image/color"
	"testing"
)

type AtTest struct {
	im         *Image
	p          image.Point
	r, g, b, a uint8
}

var (
	r1  = image.Rect(0, 0, 1, 1)
	r2  = image.Rect(1, 1, 2, 2)
	p0  = image.Pt(0, 0)
	p1  = image.Pt(101, 101) // Far outside source of the replicated box.
	top = Color(0x884422FF)
	bot = Color(0x773311FF)
	gry = Color(0x30303030)
	AA  = Color(0xAAAAAAAA)
	AB  = Color(0xABABABAB)
	DD  = Color(0xDDDDDDDD)
)

var atTests = []AtTest{
	// GREY1
	{alloc(r1, GREY1, true, DBlack), p0, 0x00, 0x00, 0x00, 0xFF},
	{alloc(r1, GREY1, true, DWhite), p0, 0xFF, 0xFF, 0xFF, 0xFF},
	// GREY2
	{alloc(r1, GREY2, true, DBlack), p0, 0x00, 0x00, 0x00, 0xFF},
	{alloc(r1, GREY2, true, DWhite), p0, 0xFF, 0xFF, 0xFF, 0xFF},
	{alloc(r1, GREY2, true, AA), p0, 0xAA, 0xAA, 0xAA, 0xFF},
	{alloc(r2, GREY2, true, AA), p1, 0xAA, 0xAA, 0xAA, 0xFF},
	// GREY4
	{alloc(r1, GREY4, true, DBlack), p0, 0x00, 0x00, 0x00, 0xFF},
	{alloc(r1, GREY4, true, DWhite), p0, 0xFF, 0xFF, 0xFF, 0xFF},
	{alloc(r1, GREY4, true, DD), p0, 0xDD, 0xDD, 0xDD, 0xFF},
	{alloc(r2, GREY4, true, AA), p1, 0xAA, 0xAA, 0xAA, 0xFF},
	// GREY8
	{alloc(r1, GREY8, true, DBlack), p0, 0x00, 0x00, 0x00, 0xFF},
	{alloc(r1, GREY8, true, DWhite), p0, 0xFF, 0xFF, 0xFF, 0xFF},
	{alloc(r1, GREY8, true, AB), p0, 0xAB, 0xAB, 0xAB, 0xFF},
	{alloc(r1, GREY8, true, AA), p0, 0xAA, 0xAA, 0xAA, 0xFF},
	{alloc(r2, GREY8, true, AB), p1, 0xAB, 0xAB, 0xAB, 0xFF},
	// CMAP8Cannot represent all 8-bit values accurately.
	{alloc(r1, CMAP8, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, CMAP8, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, CMAP8, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, CMAP8, true, top), p0, 0x88, 0x44, 0x00, 0xFF},
	{alloc(r1, CMAP8, true, bot), p0, 0x88, 0x44, 0x00, 0xFF},
	{alloc(r1, CMAP8, true, gry), p0, 0x33, 0x33, 0x33, 0xFF},
	{alloc(r2, CMAP8, true, gry), p1, 0x33, 0x33, 0x33, 0xFF},
	// RGB15 Cannot represent all 8-bit values accurately.
	{alloc(r1, RGB15, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, RGB15, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, RGB15, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, RGB15, true, top), p0, 0x8C, 0x42, 0x21, 0xFF},
	{alloc(r1, RGB15, true, bot), p0, 0x73, 0x31, 0x10, 0xFF},
	{alloc(r1, RGB15, true, gry), p0, 0x31, 0x31, 0x31, 0xFF},
	{alloc(r2, RGB15, true, top), p1, 0x8C, 0x42, 0x21, 0xFF},
	// RGB16 Cannot represent all 8-bit values accurately.
	{alloc(r1, RGB16, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, RGB16, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, RGB16, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, RGB16, true, top), p0, 0x8C, 0x42, 0x21, 0xFF},
	{alloc(r1, RGB16, true, bot), p0, 0x73, 0x31, 0x10, 0xFF},
	{alloc(r1, RGB16, true, gry), p0, 0x31, 0x31, 0x31, 0xFF},
	{alloc(r2, RGB16, true, top), p1, 0x8C, 0x42, 0x21, 0xFF},
	// RGB24
	{alloc(r1, RGB24, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, RGB24, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, RGB24, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, RGB24, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, RGB24, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, RGB24, true, gry), p0, 0x30, 0x30, 0x30, 0xFF},
	{alloc(r2, RGB24, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
	// BGR24
	{alloc(r1, BGR24, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, BGR24, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, BGR24, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, BGR24, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, BGR24, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, BGR24, true, gry), p0, 0x30, 0x30, 0x30, 0xFF},
	{alloc(r2, BGR24, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
	// RGBA32
	{alloc(r1, RGBA32, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, RGBA32, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, RGBA32, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, RGBA32, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, RGBA32, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, RGBA32, true, gry), p0, 0x30, 0x30, 0x30, 0x30},
	{alloc(r2, RGBA32, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
	// ARGB32
	{alloc(r1, ARGB32, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, ARGB32, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, ARGB32, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, ARGB32, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, ARGB32, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, ARGB32, true, gry), p0, 0x30, 0x30, 0x30, 0x30},
	{alloc(r2, ARGB32, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
	// ABGR32
	{alloc(r1, ABGR32, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, ABGR32, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, ABGR32, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, ABGR32, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, ABGR32, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, ABGR32, true, gry), p0, 0x30, 0x30, 0x30, 0x30},
	{alloc(r2, ABGR32, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
	// XRGB32
	{alloc(r1, XRGB32, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, XRGB32, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, XRGB32, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, XRGB32, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, XRGB32, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, XRGB32, true, gry), p0, 0x30, 0x30, 0x30, 0xFF},
	{alloc(r2, XRGB32, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
	// XBGR32
	{alloc(r1, XBGR32, true, DRed), p0, 0xFF, 0x00, 0x00, 0xFF},
	{alloc(r1, XBGR32, true, DGreen), p0, 0x00, 0xFF, 0x00, 0xFF},
	{alloc(r1, XBGR32, true, DBlue), p0, 0x00, 0x00, 0xFF, 0xFF},
	{alloc(r1, XBGR32, true, top), p0, 0x88, 0x44, 0x22, 0xFF},
	{alloc(r1, XBGR32, true, bot), p0, 0x77, 0x33, 0x11, 0xFF},
	{alloc(r1, XBGR32, true, gry), p0, 0x30, 0x30, 0x30, 0xFF},
	{alloc(r2, XBGR32, true, top), p1, 0x88, 0x44, 0x22, 0xFF},
}

func alloc(r image.Rectangle, pix Pix, repl bool, color Color) *Image {
	i, err := display().AllocImage(r, pix, repl, color)
	if err != nil {
		panic(err)
	}
	return i
}

func TestAt(t *testing.T) {
	for i, test := range atTests {
		r, g, b, a := test.im.At(test.p.X, test.p.Y).RGBA()
		got := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		want := color.RGBA{test.r, test.g, test.b, test.a}
		if got != want {
			t.Errorf("%d: got %x want %x", i, got, want)
		}
	}
}