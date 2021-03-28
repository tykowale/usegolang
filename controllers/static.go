package controllers

import "lenslocked.com/views"

type Static struct {
	HomeView     *views.View
	ContactView  *views.View
	NotFoundView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:     views.NewView("bootstrap", "static/home"),
		ContactView:  views.NewView("bootstrap", "static/contact"),
		NotFoundView: views.NewView("bootstrap", "static/404"),
	}
}
