package web

// @hello
// yo yo yo yo
type Pet struct {
	// ID this is petid
	ID       int `json:"id"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	Status string `json:"status"`
	S      S
	Ss     Ss
}

type Pet2 struct {
	ID int `json:"id"`
}

type S struct {
	ID   int
	Name string
}

type Ss []S

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}
