package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Button(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{vecty.Markup(vecty.Class("button"))}
	children = append(children, markup...)
	return elem.Button(
		children...,
	)
}

func PrimaryButton(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{IsPrimary}
	children = append(children, markup...)
	return Button(
		children...,
	)
}

func InfoButton(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{IsInfo}
	children = append(children, markup...)
	return Button(
		children...,
	)
}

func WarningButton(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{IsWarning}
	children = append(children, markup...)
	return Button(
		children...,
	)
}
func SuccessButton(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{IsSuccess}
	children = append(children, markup...)
	return Button(
		children...,
	)
}
func LinkButton(markup ...vecty.MarkupOrChild) *vecty.HTML {
	var children = []vecty.MarkupOrChild{IsLink}
	children = append(children, markup...)
	return Button(
		children...,
	)
}
