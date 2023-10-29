package model

type Profile struct {
	ID        string `json:"id" validate:"required" gorm:"primaryKey"`
	NickName  string `json:"nickname" validate:"required,min=2,max=25" gorm:"not null, default:用户名称"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

type AvatorResp struct {
	Code int    `json:"code"`
	Url  string `json:"url"`
}
