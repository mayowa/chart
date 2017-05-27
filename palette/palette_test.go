package palette

import (
	"image/color"
	"testing"
)

func TestPalette(t *testing.T) {
	pal := NewPalette("doesnotexist") // defaults to "white"
	if len(pal.Palette) != len(colorOrder) {
		t.Fatal("len(pal) != len(colorOrder)")
	}
	exp := color.RGBA{255, 255, 255, 255}
	if pal.Palette[0] != exp {
		t.Fatal("Expected color[0] to be white")
	}

	h := pal.GetHexColor("grid")
	if h != "#999" {
		t.Error("Unexpected hex grid color")
	}

	exp = color.RGBA{153, 153, 153, 255}
	c := pal.GetColor("grid")
	if c != exp {
		t.Error("Unexpected grid color")
	}

	c = pal.GetColor("doesnotexist")
	if c != UnknownColor {
		t.Error("Should have matched UnknownColor")
	}

	ch := pal.GetHexAxisColor(1)
	if ch != "#0000ff" {
		t.Error("Axis 1 color should be #00f")
	}

	n := pal.GetAxisColorName(0)
	if n != "area" {
		t.Error("First axis color should be area")
	}
}

func TestRandom(t *testing.T) {
	p := NewPalette("random")
	if p.Name != "random" {
		t.Errorf("Palette name should be random, is %#v", p)
	}

	p = NewPalette("hsl:180,0.5,0.5")
	if p.Name != "hsl:180,0.5,0.5" {
		t.Errorf("Palette name should be hsl:180,0.5,0.5, is %#v", p)
	}

	p = NewPalette("hsl:804020", "light")
	if p.Name != "hsl:804020" {
		t.Errorf("Palette name should be hsl:804020, is %#v", p)
	}
}
