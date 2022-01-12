package domain

type APIResp struct {
	Application string `json:"application"`
	Param1      int    `json:"param1"`
	Param2      string `json:"param2"`
	Version     int    `json:"version"`
	id          int
}
