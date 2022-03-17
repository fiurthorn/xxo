package lib

// https://fonts.google.com/icons?selected=Material+Icons

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var NewGame *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionAutorenew)
	return icon
}()

var GameIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.AVGames)
	return icon
}()

var MenuIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()

var HomeIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionHome)
	return icon
}()

var SettingsIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionSettings)
	return icon
}()
