package models

type Profile struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	PublicKeyB64  string `json:"publickey_base64"`
	PrivateKeyB64 string `json:"privatekey_base64"`
	ClientConfig  string `json:"client_config"`
	ClientQRCode  string `json:"client_qrcode"`
	IsActive      bool   `json:"is_active"`
}

type ProfileCreateForm struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
