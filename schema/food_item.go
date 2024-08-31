package schema

type FoodItem struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Subtext string `json:"subtext"`
	ImageURL string `json:"image_url"`
	HoverpicURL string `json:"hoverpic_url"`
	// Price
	Ingredients string `json:"ingredients"`
	FoodCategory string `json:"food_category"`
	FoodHighlightMessage string `json:"food_highlight_message"`
	FileURL string `json:"file_url"`
	DownloadLabel string `json:"download_label"`
	// RoundedNutritionInfo FoodItemNutritionInfo `json:"rounded_nutrition_info"`
	ServingSizeInfo FoodItemServingSizeInfo `json:"serving_size_info"`
	HasNutritionInfo bool `json:"has_nutrition_info"`
	// Icons FoodItemIconData `json:"icons"`
	IconsApproved bool `json:"icons_approved"`
	NestedFoods []FoodItem `json:"nested_foods"`
	// AggregatedData
	OrderingEnabled bool `json:"ordering_enabled"`
	// FoodSizes
	DSCaloriesOverride string `json:"ds_calories_override"`
	SyncedID string `json:"synced_id"`
	// SyncPlaceholder
	HasOptionsOrSides bool `json:"has_options_or_sides"`
	Digest string `json:"digest"`
	POSItemID string `json:"pos_item_id"`
	// SmartRecipeID
	HasSubfoods bool `json:"has_subfoods"`
	// MealPlanPrice
	UseCustomSizes bool `json:"use_custom_sizes"` 
}

type FoodItemServingSizeInfo struct {
	ServingSizeAmount string `json:"serving_size_amount"`
	ServingSizeUnit string `json:"serving_size_unit"`
}