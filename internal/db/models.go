package db

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	User_id    int    `json:"user_id"`
	Username   string `json:"username"`
	UserRank   string `json:"user_rank"`
	UserRating int    `json:"user_rating"`
}
