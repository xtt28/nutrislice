package schema

type MenuWeek struct {
	StartDate             string    `json:"start_date"`
	MenuTypeID            int       `json:"menu_type_id"`
	ID                    int       `json:"id"`
	LastUpdated           string    `json:"last_updated"`
	BoldAllEntreesEnabled bool      `json:"bold_all_entrees_enabled"`
	Published             bool      `json:"published"`
	DisplayName           string    `json:"display_name"`
	Days                  []MenuDay `json:"days"`
}
