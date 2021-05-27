package tag

/// SingleTag
/// wrap the tag model
type SingleTag struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

/// TagListResponse
/// tag list response
type TagListResponse struct {
	List  []SingleTag `json:"list"`
	Total uint64      `json:"total"`
}
