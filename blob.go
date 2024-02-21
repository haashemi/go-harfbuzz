package harfbuzz

import (
	"github.com/haashemi/go-harfbuzz/hb"
)

// Blob holds the low-level Blob.
type Blob struct{ b hb.Blob }

// Blob returns the raw blob used for low-level operations
func (b *Blob) Blob() hb.Blob { return b.b }

// Data returns the blob's data.
func (b *Blob) Data() []byte { return []byte(hb.BlobGetData(b.b)) }

// Length returns the length of the blob's data.
func (b *Blob) Length() uint { return hb.BlobGetLength(b.b) }

// Close destroys the blob and frees the memory.
func (b *Blob) Close() { hb.BlobDestroy(b.b) }

// NewBlob returns a new Blob from the data of file from path.
func NewBlob(filename string) *Blob {
	return &Blob{b: hb.BlobCreateFromFile(filename)}
}

// NewBlobFromBytes returns a new Blob from the data.
func NewBlobFromBytes(data []byte) *Blob {
	return &Blob{b: hb.BlobCreate(data, hb.MemoryModeDuplicate, nil, nil)}
}
