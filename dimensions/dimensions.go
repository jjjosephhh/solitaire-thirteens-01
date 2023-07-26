package dimensions

type Dimensions struct {
	Width  int32
	Height int32
}

func NewDimensions(width, height int32) *Dimensions {
	return &Dimensions{Width: width, Height: height}
}
