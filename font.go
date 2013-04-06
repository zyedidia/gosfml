// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. 	This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Graphics/Font.h> 
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Font struct {
	cptr *C.sfFont
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Font constructor
// Creates a new font from file
//
// The supported font formats are: TrueType, Type 1, CFF,
// OpenType, SFNT, X11 PCF, Windows FNT, BDF, PFR and Type 42.
// Note that this function know nothing about the standard
// fonts installed on the user's system, thus you can't
// load them directly.
func NewFontFromFile(filename string) (font *Font, err error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	font = &Font{C.sfFont_createFromFile(cFilename)}
	runtime.SetFinalizer(font, (*Font).Destroy)

	if font.cptr == nil {
		err = errors.New("NewFontFromFile: Cannot load font " + filename)
	}

	return
}

// Font constructor
// Creates a new font from memory
//
// The supported font formats are: TrueType, Type 1, CFF,
// OpenType, SFNT, X11 PCF, Windows FNT, BDF, PFR and Type 42.
func NewFontFromMemory(data []byte) (*Font, error) {
	if len(data) > 0 {
		font := &Font{C.sfFont_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))}
		runtime.SetFinalizer(font, (*Font).Destroy)
		return font, nil
	}
	return nil, errors.New("NewFontFromMemory: no data")
}

func (this *Font) Copy() *Font {
	font := &Font{C.sfFont_copy(this.cptr)}
	runtime.SetFinalizer(font, (*Font).Destroy)
	return font
}

func (this *Font) Destroy() {
	C.sfFont_destroy(this.cptr)
	this.cptr = nil
}

// Retrieve a glyph of the font
//
// 	codePoint:     Unicode code point of the character to get
// 	characterSize: Reference character size
// 	bold:          Retrieve the bold version or the regular one?
//
// return The glyph corresponding to codePoint and characterSize
func (this *Font) GetGlyph(codePoint uint, characterSize uint32, bold bool) (glyph Glyph) {
	glyph.fromC(C.sfFont_getGlyph(this.cptr, C.sfUint32(codePoint), C.uint(characterSize), goBool2C(bold)))
	return
}

//Get the kerning offset of two glyphs
//
// The kerning is an extra offset (negative) to apply between two
// glyphs when rendering them, to make the pair look more "natural".
// For example, the pair "AV" have a special kerning to make them
// closer than other characters. Most of the glyphs pairs have a
// kerning offset of zero, though.
//
// 	first:         Unicode code point of the first character
// 	second:        Unicode code point of the second character
// 	characterSize: Reference character size
//
// return Kerning value for first and second, in pixels
func (this *Font) GetKerning(first uint32, second uint32, characterSize uint) int {
	return int(C.sfFont_getKerning(this.cptr, C.sfUint32(first), C.sfUint32(second), C.uint(characterSize)))
}

// Get the line spacing
//
// Line spacing is the vertical offset to apply between two
// consecutive lines of text.
//
// 	characterSize: Reference character size
//
// return Line spacing, in pixels
func (this *Font) GetLineSpacing(characterSize uint) int {
	return int(C.sfFont_getLineSpacing(this.cptr, C.uint(characterSize)))
}

// Retrieve the texture containing the loaded glyphs of a certain size
//
// The contents of the returned texture changes as more glyphs
// are requested, thus it is not very relevant.
//
// 	characterSize: Reference character size
//
// Texture containing the glyphs of the requested size
func (this *Font) GetTexture(characterSize uint) Texture {
	return Texture{C.sfFont_getTexture(this.cptr, C.uint(characterSize))}
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Font) toCPtr() *C.sfFont {
	if this != nil {
		return this.cptr
	}
	return nil
}
