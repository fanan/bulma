package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Field(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("field"))}
	children = append(children, markup...)
	return elem.Div(children...)
}

func Control(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("control"))}
	children = append(children, markup...)
	return elem.Div(children...)
}

func Label(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("label"))}
	children = append(children, markup...)
	return elem.Label(children...)
}

func Input(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("input"))}
	children = append(children, markup...)
	return elem.Input(children...)
}

func Select(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("select")),
		elem.Select(
			markup...,
		),
	)
}

func TextArea(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("textarea"))}
	children = append(children, markup...)
	return elem.TextArea(children...)
}
