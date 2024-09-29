package services

import (
	"strings"
)

// MenuItem represents a menu item to be displayed in the sidebar
type MenuItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	HRef        string `json:"href"`
	Icon        string `json:"icon"`
	IsActive    bool   `json:"is_active"`
	ActiveClass string `json:"active_class"`
	IsAvailable bool   `json:"is_available"`
}

// MenuSection represents a section in the sidebar
type MenuSection struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	MenuItems   []MenuItem `json:"menu_items"`
	IsAvailable bool       `json:"is_available"`
}

// GetMenu returns the menu items
func GetMenu() []MenuSection {
	sections := []MenuSection{
		{
			ID:   1,
			Name: "Menu",
			MenuItems: []MenuItem{
				{
					ID:          1,
					Name:        "Dashboard",
					HRef:        "/dashboard",
					Icon:        "bi-grid-fill",
					IsAvailable: true,
				},
			},
			IsAvailable: true,
		},
		{
			ID:   2,
			Name: "Settings",
			MenuItems: []MenuItem{
				{
					ID:          1,
					Name:        "Users",
					HRef:        "/users",
					Icon:        "bi-person-fill",
					IsAvailable: true,
				},
			},
		},
	}
	return sections
}

// MarkMenuItemActive marks the menu item as active based on the activeHref
func MarkMenuItemActive(activeHref string) []MenuSection {
	menu := GetMenu()
	//check if the activeHref is a subpath of any menu item
	for i, section := range menu {
		for j, item := range section.MenuItems {
			if strings.Contains(activeHref, item.HRef) {
				menu[i].MenuItems[j].IsActive = true
				menu[i].MenuItems[j].ActiveClass = "active"
				break
			}
			if item.HRef == activeHref {
				menu[i].MenuItems[j].IsActive = true
				menu[i].MenuItems[j].ActiveClass = "active"
				break
			}
		}
	}
	return menu
}
