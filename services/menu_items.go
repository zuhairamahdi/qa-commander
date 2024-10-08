package services

import (
	"strings"
)

// MenuItem represents a menu item to be displayed in the sidebar
type MenuItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	HRef        string `json:"href"`
	Icon        string `json:"icon" default:"bi-grid-fill"`
	IsActive    bool   `json:"is_active" default:"false"`
	IsAvailable bool   `json:"is_available" default:"true"`
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
				{
					ID:          2,
					Name:        "Projects",
					HRef:        "/projects",
					Icon:        "bi-folder-fill",
					IsAvailable: true,
				},
				{
					ID:          3,
					Name:        "Tasks",
					HRef:        "/tasks",
					Icon:        "bi-list-task",
					IsAvailable: true,
				},
				{
					ID:          4,
					Name:        "Reports",
					HRef:        "/reports",
					Icon:        "bi-file-earmark-bar-graph-fill",
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
				{
					ID:          2,
					Name:        "Roles",
					HRef:        "/roles",
					Icon:        "bi-person-lines-fill",
					IsAvailable: true,
				},
				{
					ID:          3,
					Name:        "Permissions",
					HRef:        "/permissions",
					Icon:        "bi-shield-fill",
					IsAvailable: true,
				},
				{
					ID:          4,
					Name:        "Site Settings",
					HRef:        "/site-settings",
					Icon:        "bi-gear-fill",
					IsAvailable: true,
				},
			},
			IsAvailable: true,
		},
		{
			ID:   3,
			Name: "User",
			MenuItems: []MenuItem{
				{
					ID:          1,
					Name:        "Profile",
					HRef:        "/profile",
					Icon:        "bi-person-circle",
					IsAvailable: true,
				},
				{
					ID:          2,
					Name:        "Logout",
					HRef:        "/logout",
					Icon:        "bi-box-arrow-right",
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
			if strings.Contains(activeHref, item.HRef) || item.HRef == activeHref {
				menu[i].MenuItems[j].IsActive = true
				break
			}
		}
	}
	return menu
}
