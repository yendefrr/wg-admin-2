package models

type Profile struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	PublicKeyB64  string `json:"publickey_b64"`
	PrivateKeyB64 string `json:"privatekey_b64"`
	ClientConfig  string `json:"client_config"`
	ClientQRCode  string `json:"client_qr_code"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	IsActive      bool   `json:"is_active"`
}

type ProfileCreateForm struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
