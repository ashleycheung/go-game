package physics

// A generic bounding box
type BBox struct {
	TopLeft     Vector `json:"topLeft"`
	BottomRight Vector `json:"bottomRight"`
}

func (b BBox) Size() Vector {
	return b.BottomRight.Subtract(b.TopLeft)
}

func (b BBox) ToSizePosition() (size, position Vector) {
	size = b.Size()
	position = b.TopLeft.Add(size.Scale(0.5))
	return
}
