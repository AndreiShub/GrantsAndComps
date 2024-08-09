package swagger

type InlineResponse2001 struct {

	Grant *Grant `json:"grant"`

	FiltersMapping *map[string]interface{} `json:"filters_mapping"`

	FiltersOrder []string `json:"filters_order"`
}
