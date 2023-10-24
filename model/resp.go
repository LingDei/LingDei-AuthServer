package model

type OperationResp struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
}

type UsersResp struct {
	Code  uint   `json:"code"`
	Users []User `json:"users"`
}

type TokenResp struct {
	Code  uint   `json:"code"`
	Token string `json:"token"`
	Role  int    `json:"role"`
}

type AdminProfileResp struct {
	Code     uint    `json:"code"`
	UserID   string  `json:"userID"`
	Username string  `json:"username"`
	Role     int     `json:"role"`
	Profile  Profile `json:"profile"`
}

type UIDResp struct {
	Code uint   `json:"code"`
	ID   string `json:"id"`
}

type ProfileResp struct {
	Code    uint    `json:"code"`
	Profile Profile `json:"profile"`
}
