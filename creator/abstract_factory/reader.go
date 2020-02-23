package abstract_factory

import "fmt"

type Reader interface {
	Read(image Image) string
}


type GIFReader struct {}

func (r *GIFReader) Read(i Image) string {
	return fmt.Sprintf("GIFReader read:%s", i.GetContent())
}


type JPEGReader struct {}

func (r *JPEGReader) Read(i Image) string {
	return fmt.Sprintf("GIFReader read:%s", i.GetContent())
}




type ReaderFactory interface {
	CreateFactory() Reader
}


type GIFReaderFactory struct {}

func (f *GIFReaderFactory) CreateFactory() Reader {
	return new(GIFReader)
}


type JPEGReaderFactory struct {}

func (f *JPEGReaderFactory) CreateFactory() Reader {
	return new(JPEGReader)
}