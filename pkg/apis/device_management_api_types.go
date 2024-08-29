package apis

type DeviceManagementApiLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeviceManagementApiLoginResponse struct {
	Auth  bool   `json:"auth"`
	Token string `json:"token"`
}
