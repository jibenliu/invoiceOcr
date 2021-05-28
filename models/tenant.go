package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"invoiceOcr/services"
	"net/http"
)

type Register struct {
	TenantName   string `gorm:"tenant_name" json:"tenant_name" validate:"required,min=1,max=150"`
	ContactName  string `gorm:"contact_name" json:"contact_name" validate:"required,min=3,max=200"`
	ContactEmail string `gorm:"contact_email" json:"contact_email" validate:"required,email"`
	ContactPhone string `gorm:"contact_phone" json:"contact_phone" validate:"required,IsMobile"`
}

func NewRegister(c *gin.Context) {
	form := new(Register)
	err := c.ShouldBind(form)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tenant := services.Tenant{}
	err = DB.Where("tenant_name = ?", form.TenantName).First(&tenant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.AbortWithError(http.StatusBadRequest, errors.New("服务异常，请稍后重试"))
		return
	}
	if err == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("租户名称已被占用！"))
		return
	}
	tenant.TenantName = form.TenantName
	tenant.ContactName = form.ContactName
	tenant.ContactEmail = form.ContactEmail
	tenant.ContactPhone = form.ContactPhone
	DB.Create(&tenant)
}
