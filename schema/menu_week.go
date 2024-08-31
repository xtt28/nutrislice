package schema

import "time"

type MenuWeek struct {
	StartDate time.Time `json:"start_date"`
	MenuTypeID int `json:"menu_type_id"`
	ID int `json:"id"`
	LastUpdated time.Time `json:"last_updated"`
	BoldAllEntreesEnabled bool `json:"bold_all_entrees_enabled"`
	Published bool `json:"published"`
	DisplayName string `json:"display_name"`
}