package services

type Tenant struct {
	Status       int8   `gorm:"status"`
	TenantCode   string `gorm:"tenant_code"`
	TenantName   string `gorm:"tenant_name"`
	ContactName  string `gorm:"contact_name"`
	ContactEmail string `gorm:"contact_email"`
	ContactPhone string `gorm:"contact_phone"`
	AppKey       string `gorm:"app_key"`
	AppSecret    string `gorm:"app_secret"`
	Signature    string `gorm:"signature"`
}

type TenantAuthSetting struct {
	TenantCode         string `gorm:"tenant_code"`
	ApiCode            string `gorm:"api_code"`
	ApiChargeStatus    int    `gorm:"api_charge_status"`     //是否扣减
	ApiChargeCount     int    `gorm:"api_charge_count"`      //充值次数
	LeftApiChargeCount int    `gorm:"left_api_charge_count"` //剩余次数
}
