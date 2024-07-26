# Go Harfbuzz bindings

> [!NOTE]
> go-harfbuzz is currently work-in-progress. The current goal is to implement all Harfbuzz v8.3 methods and then start to support newer versions.

go-harfbuzz is a Golang binding to the Harfbuzz text shaping engine.

## Requirements:

- Having `CGO` enabled
- Having `pkg-config` installed.
- Having `harfbuzz-dev` installed.
  - version: `8.3` or above.
  - note: `pkg-config` should find it.

## Installation:

```
go get -u github.com/haashemi/go-harfbuzz@main
```

## Usage:

```go
package main

import (
	"fmt"

	hb "github.com/haashemi/go-harfbuzz"
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

```

## Contributions

All types of contributions are highly appreciated.
