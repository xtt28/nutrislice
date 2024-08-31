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
	RoundedNutritionInfo FoodItemNutritionInfo `json:"rounded_nutrition_info"`
	ServingSizeInfo FoodItemServingSizeInfo `json:"serving_size_info"`
	HasNutritionInfo bool `json:"has_nutrition_info"`
	Icons FoodItemIconsData `json:"icons"`
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

type FoodItemNutritionInfo struct {
	Calories       float64 `json:"calories"`
	Fat            float64 `json:"g_fat"`
	SaturatedFat   float64 `json:"g_saturated_fat"`
	TransFat       float64 `json:"g_trans_fat"`
	Cholesterol    float64 `json:"mg_cholesterol"`
	Carbs          float64 `json:"g_carbs"`
	AddedSugar     float64 `json:"g_added_sugar"`
	Sugar          float64 `json:"g_sugar"`
	Potassium      float64 `json:"mg_potassium"`
	Sodium         float64 `json:"mg_sodium"`
	Fiber          float64 `json:"g_fiber"`
	Protein        float64 `json:"g_protein"`
	Iron           float64 `json:"mg_iron"`
	Calcium        float64 `json:"mg_calcium"`
	VitaminC       float64 `json:"mg_vitamin_c"`
	VitaminA       float64 `json:"iu_vitamin_a"`
	RetinolEquivalents float64 `json:"re_vitamin_a"`
	MicrogramsVitaminA float64 `json:"mcg_vitamin_a"`
	VitaminD       float64 `json:"mg_vitamin_d"`
	MicrogramsVitaminD float64 `json:"mcg_vitamin_d"`
}

type FoodItemIconsData struct {
	FoodIcons []FoodIcon `json:"food_icons"`
}

type FoodIcon struct {
	Name string `json:"name"`
}