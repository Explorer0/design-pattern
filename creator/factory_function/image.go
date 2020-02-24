package factory_function

type Image interface {
	GetContent() string
}


type GIFImage struct {}

func (i *GIFImage) GetContent() string {
	return "This is GIF image."
}


type JPEGImage struct {}

func (i *JPEGImage) GetContent() string {
	return "This is GIF image."
}

