package resources

// response format
type Response struct {
    ID      uint  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}