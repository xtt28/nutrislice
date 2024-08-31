package schema

type MenuDay struct {
	Date                string                `json:"date"`
	HasUnpublishedMenus bool                     `json:"has_unpublished_menus"`
	MenuInfo            map[string]MenuInfoValue `json:"menu_info"`
	MenuItems           []MenuItem               `json:"menu_items"`
}
