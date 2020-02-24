package factory_function

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
	CreateReader() Reader
}


type GIFReaderFactory struct {}

func (f *GIFReaderFactory) CreateReader() Reader {
	return new(GIFReader)
}


type JPEGReaderFactory struct {}

func (f *JPEGReaderFactory) CreateReader() Reader {
	return new(JPEGReader)
}


func CreateGIFFactory() *GIFReaderFactory {
	return &GIFReaderFactory{}
}

func CreateJPEGFactory() *JPEGReaderFactory {
	return &JPEGReaderFactory{}
}