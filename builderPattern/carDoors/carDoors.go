package cardoors

import (
	"strings"
)

type carDoors struct {
	front bool
	back  bool
	right bool
	left  bool
}

type CarDoorsBuilder struct {
	code string
}

func NewCarDoorsBuilder() *CarDoorsBuilder {
	return new(CarDoorsBuilder)
}

func (builder *CarDoorsBuilder) AddLeft() *CarDoorsBuilder {
	builder.code += "L"
	return builder
}

func (builder *CarDoorsBuilder) AddRight() *CarDoorsBuilder {
	builder.code += "R"
	return builder
}

func (builder *CarDoorsBuilder) AddFront() *CarDoorsBuilder {
	builder.code += "F"
	return builder
}

func (builder *CarDoorsBuilder) AddBack() *CarDoorsBuilder {
	builder.code += "B"
	return builder
}

func (builder *CarDoorsBuilder) Build() *carDoors {
	code := builder.code
	return &carDoors{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}
