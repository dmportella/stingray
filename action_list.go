package stingray

// ActionList The common list representation of resources within stingray api.
type ActionList struct {
	Children []struct {
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"children"`
}
