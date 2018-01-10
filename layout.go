package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Section(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("section"))}
	children = append(children, markup...)
	return elem.Section(
		children...,
	)
}

func Container(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("container"))}
	children = append(children, markup...)
	return elem.Div(
		children...,
	)
}

func FluidContainer(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("container", "is-fluid"))}
	children = append(children, markup...)
	return elem.Div(
		children...,
	)
}
