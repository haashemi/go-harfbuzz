package main

import (
	"fmt"

	"github.com/haashemi/go-harfbuzz/hb"
)

func main() {
	buf := hb.BufferCreate()
	defer hb.BufferDestroy(buf)

	hb.BufferAddUTF8(buf, "Hello World!")
	hb.BufferGuessSegmentProperties(buf)

	blob := hb.BlobCreateFromFile("path/to/font.ttf")
	defer hb.BlobDestroy(blob)

	face := hb.FaceCreate(blob, 0)
	defer hb.FaceDestroy(face)

	font := hb.FontCreate(face)
	defer hb.FontDestroy(font)

	hb.Shape(font, buf, nil)

	glyphsInfo := hb.BufferGetGlyphInfos(buf)
	glyphsPositions := hb.BufferGetGlyphPositions(buf)

	var cursorX, cursorY int32
	for i := 0; i < len(glyphsInfo); i++ {
		info := glyphsInfo[i]
		pos := glyphsPositions[i]

		fmt.Printf("index: %d\tinfo: %v\tpos: %v\n", i, info, pos)

		// glyphID := info.Codepoint
		// XOffset := pos.XOffset
		// YOffset := pos.YOffset
		// Draw the glyph here.

		cursorX += pos.XAdvance
		cursorY += pos.YAdvance
	}

	fmt.Println("X Advance:", cursorX)
	fmt.Println("Y Advance:", cursorY)
}
