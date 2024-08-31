package schema

import "time"

type MenuItem struct {
	ID   int `json:"id"`
	Date time.Time `json:"date"`
	Position int `json:"position"`
	IsSectionTitle bool `json:"is_section_title"`
	Bold bool `json:"bold"`
	Featured bool `json:"featured"`
	Text string `json:"text"`
	NoLineBreak bool `json:"no_line_break"`
	BlankLine bool `json:"blank_line"`
	Food FoodItem `json:"food"`
	IsHoliday bool `json:"is_holiday"`
	FoodList []FoodItem `json:"food_list"`
	StationID int `json:"station_id"`
	IsStationHeader bool `json:"is_station_header"`
	StationIsCollapsible bool `json:"station_is_collapsible"`
	Image string `json:"image"`
	ImageDescription string `json:"image_description"`
	ImageAlt string `json:"image_alt"`
	ImageThumbnail string `json:"image_thumbnail"`
	Category string `json:"category"`
	Price string `json:"price"`
	// ServingSize
	ServingSizeAmount float64 `json:"serving_size_amount"`
	ServingSizeUnit string `json:"serving_size_unit"`
	// SmartRecipeID
	MenuID int `json:"menu_id"`
	FoodVariationID int `json:"food_variation_id"`
}