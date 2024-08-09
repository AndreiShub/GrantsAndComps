package swagger

type Grant struct {

	Id int32 `json:"id"`

	Title string `json:"title"`

	SourceUrl string `json:"source_url"`

	FilterValues *map[string]Object `json:"filter_values"`
}
