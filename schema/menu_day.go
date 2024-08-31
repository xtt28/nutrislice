package schema

import "time"

type MenuDay struct {
	Date                time.Time                `json:"date"`
	HasUnpublishedMenus bool                     `json:"has_unpublished_menus"`
	MenuInfo            map[string]MenuInfoValue `json:"menu_info"`
	MenuItems           []MenuItem               `json:"menu_items"`
}
