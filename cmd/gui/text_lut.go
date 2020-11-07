package main

import (
)

var fg_color_lut [64]byte
var bg_color_lut [64]byte
var f_color_lut [16][4]byte
var b_color_lut [16][4]byte

// use MapRGBA here to fit to pixel format!
func pseudoInit() {
	fg_color_lut = [64]byte{
		0x00, 0x00, 0x00, 0xFF,
		0x00, 0x00, 0x80, 0xFF,
		0x00, 0x80, 0x00, 0xFF,
		0x80, 0x00, 0x00, 0xFF,
		0x00, 0x80, 0x80, 0xFF,
		0x80, 0x80, 0x00, 0xFF,
		0x80, 0x00, 0x80, 0xFF,
		0x80, 0x80, 0x80, 0xFF,
		0x00, 0x45, 0xFF, 0xFF,
		0x13, 0x45, 0x8B, 0xFF,
		0x00, 0x00, 0x20, 0xFF,
		0x00, 0x20, 0x00, 0xFF,
		0x20, 0x00, 0x00, 0xFF,
		0x20, 0x20, 0x20, 0xFF,
		0x60, 0x60, 0x60, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF,
	}

	bg_color_lut = [64]byte{
		0x00, 0x00, 0x00, 0xFF,
		0x00, 0x00, 0x80, 0xFF,
		0x00, 0x80, 0x00, 0xFF,
		0x80, 0x00, 0x00, 0xFF,
		0x00, 0x20, 0x20, 0xFF,
		0x20, 0x20, 0x00, 0xFF,
		0x20, 0x00, 0x20, 0xFF,
		0x20, 0x20, 0x20, 0xFF,
		0x1E, 0x69, 0xD2, 0xFF,
		0x13, 0x45, 0x8B, 0xFF,
		0x00, 0x00, 0x20, 0xFF,
		0x00, 0x20, 0x00, 0xFF,
		0x40, 0x00, 0x00, 0xFF,
		0x30, 0x30, 0x30, 0xFF,
		0x40, 0x40, 0x40, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF,
	}
	f_color_lut = [16][4]byte{
		{0x00, 0x00, 0x00, 0xFF},
		{0x00, 0x00, 0x80, 0xFF},
		{0x00, 0x80, 0x00, 0xFF},
		{0x80, 0x00, 0x00, 0xFF},
		{0x00, 0x80, 0x80, 0xFF},
		{0x80, 0x80, 0x00, 0xFF},
		{0x80, 0x00, 0x80, 0xFF},
		{0x80, 0x80, 0x80, 0xFF},
		{0x00, 0x45, 0xFF, 0xFF},
		{0x13, 0x45, 0x8B, 0xFF},
		{0x00, 0x00, 0x20, 0xFF},
		{0x00, 0x20, 0x00, 0xFF},
		{0x20, 0x00, 0x00, 0xFF},
		{0x20, 0x20, 0x20, 0xFF},
		{0x60, 0x60, 0x60, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}

	b_color_lut = [16][4]byte{
		{0x00, 0x00, 0x00, 0xFF},
		{0x00, 0x00, 0x80, 0xFF},
		{0x00, 0x80, 0x00, 0xFF},
		{0x80, 0x00, 0x00, 0xFF},
		{0x00, 0x20, 0x20, 0xFF},
		{0x20, 0x20, 0x00, 0xFF},
		{0x20, 0x00, 0x20, 0xFF},
		{0x20, 0x20, 0x20, 0xFF},
		{0x1E, 0x69, 0xD2, 0xFF},
		{0x13, 0x45, 0x8B, 0xFF},
		{0x00, 0x00, 0x20, 0xFF},
		{0x00, 0x20, 0x00, 0xFF},
		{0x40, 0x00, 0x00, 0xFF},
		{0x30, 0x30, 0x30, 0xFF},
		{0x40, 0x40, 0x40, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
}
