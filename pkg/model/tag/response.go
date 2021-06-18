package tag

// SingleTag
// wrap the tag model
type SingleTag struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

// ListResponse
// tag list response
type ListResponse struct {
	List  []SingleTag `json:"list"`
	Total uint64      `json:"total"`
}
