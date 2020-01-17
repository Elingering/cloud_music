package model

type RawData struct {
	HasMore     bool `json:"hasMore"`
	HotComments []struct {
		User struct {
			Nickname string `json:"nickname"`
		} `json:"user"`
		Content    string `json:"content"`
		Time       int64  `json:"time"`
		LikedCount int    `json:"likedCount"`
		Liked      bool   `json:"liked"`
	} `json:"hotComments"`
	Code int `json:"code"`
}

type SongComment struct {
	Player     string `json:"player"`
	SongName   string `json:"songName"`
	Album      string `json:"album"`
	Nickname   string `json:"nickname"`
	Content    string `json:"content"`
	Time       string `json:"time"`
	LikedCount int    `json:"likedCount"`
}
