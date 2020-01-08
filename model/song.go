package model

import "time"

type Song struct {
	Player     string
	SongName   string
	Album      string
	HotComment Comment
}

type Comment struct {
	UserName  string
	CreatedAt time.Time
	Content   string
	Like      int
}
