package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"strconv"
)

func Column(size int, markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("column", "is-"+strconv.Itoa(size)))}
	children = append(children, markup...)
	return elem.Div(
		children...,
	)
}

func Column1(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(1, markup...) }
func Column2(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(2, markup...) }
func Column3(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(3, markup...) }
func Column4(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(4, markup...) }
func Column5(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(5, markup...) }
func Column6(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(6, markup...) }
func Column7(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(7, markup...) }
func Column8(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(8, markup...) }
func Column9(markup ...vecty.MarkupOrChild) *vecty.HTML  { return Column(9, markup...) }
func Column10(markup ...vecty.MarkupOrChild) *vecty.HTML { return Column(10, markup...) }
func Column11(markup ...vecty.MarkupOrChild) *vecty.HTML { return Column(11, markup...) }

func Columns(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("columns"))}
	children = append(children, markup...)
	return elem.Div(
		children...,
	)
}
