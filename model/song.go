package model

type Song struct {
	Player     string
	SongName   string
	Album      string
	HotComment Comment
}

type Comment struct {
	TopComments []interface{} `json:"topComments"`
	HasMore     bool          `json:"hasMore"`
	HotComments []HM          `json:"hotComments"`
	Total       int           `json:"total"`
	Code        int           `json:"code"`
}

type HM struct {
	User                User          `json:"user"`
	BeReplied           []interface{} `json:"beReplied"`
	PendantData         interface{}   `json:"pendantData"`
	ShowFloorComment    interface{}   `json:"showFloorComment"`
	Status              int           `json:"status"`
	CommentID           int           `json:"commentId"`
	Content             string        `json:"content"`
	Time                int64         `json:"time"`
	LikedCount          int           `json:"likedCount"`
	ExpressionURL       interface{}   `json:"expressionUrl"`
	CommentLocationType int           `json:"commentLocationType"`
	ParentCommentID     int           `json:"parentCommentId"`
	Decoration          struct{}      `json:"decoration"`
	RepliedMark         interface{}   `json:"repliedMark"`
	Liked               bool          `json:"liked"`
}

type User struct {
	LocationInfo interface{} `json:"locationInfo"`
	LiveInfo     interface{} `json:"liveInfo"`
	VipRights    struct {
		Associator   interface{} `json:"associator"`
		MusicPackage struct {
			VipCode int  `json:"vipCode"`
			Rights  bool `json:"rights"`
		} `json:"musicPackage"`
		RedVipAnnualCount int `json:"redVipAnnualCount"`
	} `json:"vipRights"`
	UserID     int         `json:"userId"`
	VipType    int         `json:"vipType"`
	ExpertTags interface{} `json:"expertTags"`
	RemarkName interface{} `json:"remarkName"`
	Nickname   string      `json:"nickname"`
	UserType   int         `json:"userType"`
	AuthStatus int         `json:"authStatus"`
	Experts    interface{} `json:"experts"`
	AvatarURL  string      `json:"avatarUrl"`
}
