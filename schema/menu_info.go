package schema

type MenuInfoValue struct {
	Position int `json:"position"`
	SectionOptions MenuInfoValueSectionOptions `json:"section_options"`
}

type MenuInfoValueSectionOptions struct {
	DisplayName string `json:"display_name"`
	UseSectionTitle bool `json:"use_section_title"`
	SectionTitleCanExpandCollapse bool `json:"section_title_can_expand_collapse"`
}