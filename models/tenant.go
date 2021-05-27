package models

import "github.com/gin-gonic/gin"

type Register struct {
	TenantName   string `gorm:"tenant_name" json:"tenant_name"`
	ContactName  string `gorm:"contact_name" json:"contact_name"`
	ContactEmail string `gorm:"contact_email" json:"contact_email"`
	ContactPhone string `gorm:"contact_phone" json:"contact_phone"`
}

func NewRegister(c *gin.Context) {

}
