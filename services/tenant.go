package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

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

//GenRsaKey RSA公钥私钥产生
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
