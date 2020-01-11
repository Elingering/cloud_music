package model

type Song struct {
	Player     string
	SongName   string
	Album      string
	HotComment Comment
}

type Comment struct {
	HasMore bool          `json:"hasMore"`
	Content []CommentItem `json:"hotComments"`
	Code    int           `json:"code"`
}

type CommentItem struct {
	User       CommentUser `json:"user"`
	Content    string      `json:"content"`
	Time       int64       `json:"time"`
	LikedCount int         `json:"likedCount"`
}

type CommentUser struct {
	Nickname string `json:"nickname"`
}
