package swagger

type InlineResponse200 struct {

	Grants []Grant `json:"grants"`

	FiltersMapping *map[string]interface{} `json:"filters_mapping"`

	FiltersOrder []string `json:"filters_order"`

	Meta *PaginationMeta `json:"meta"`
}
