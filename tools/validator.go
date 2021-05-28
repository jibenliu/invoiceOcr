package tools

import (
	zhW "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"invoiceOcr/models"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var Validate *validator.Validate
var trans ut.Translator

func init() {
	zh := zhW.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	_ = zhTranslations.RegisterDefaultTranslations(Validate, trans)

	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	_ = Validate.RegisterValidation("InStatusRange", InStatusRange)
	_ = Validate.RegisterValidation("IsMobile", IsMobile)
	_ = Validate.RegisterValidation("CheckDate", CheckDate)
	_ = Validate.RegisterValidation("IsDateRange", IsDateRange)
	_ = Validate.RegisterValidation("CheckDateAfterNow", CheckDateAfterNow)
	_ = Validate.RegisterValidation("CheckDateBeforeNow", CheckDateBeforeNow)
	_ = Validate.RegisterTranslation(
		"InStatusRange",
		trans,
		registerTranslator("InStatusRange", "{0}不合法"),
		translate,
	)
	_ = Validate.RegisterTranslation(
		"InOperateTypeRange",
		trans,
		registerTranslator("InOperateTypeRange", "{0}不合法"),
		translate,
	)
	_ = Validate.RegisterTranslation(
		"IsMobile",
		trans,
		registerTranslator("IsMobile", "{0}不是合法的手机号"),
		translate,
	)
	_ = Validate.RegisterTranslation(
		"CheckDate",
		trans,
		registerTranslator("CheckDate", "{0}不满足2006-01-02这种格式"),
		translate,
	)
	_ = Validate.RegisterTranslation(
		"IsDateRange",
		trans,
		registerTranslator("IsDateRange", "{0}不满足2006-01-02~2006-02-02这种格式"),
		translate,
	)

	_ = Validate.RegisterTranslation(
		"CheckDateAfterNow",
		trans,
		registerTranslator("CheckDateAfterNow", "{0}只能输入今后的日期"),
		translate,
	)

	_ = Validate.RegisterTranslation(
		"CheckDateBeforeNow",
		trans,
		registerTranslator("CheckDateBeforeNow", "{0}只能输入今天之前的日期"),
		translate,
	)

}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// Translate 翻译错误信息
func Translate(err error, obj interface{}) map[string]string {
	var result = make(map[string]string)
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		st := reflect.TypeOf(obj)
		field, _ := st.FieldByName(err.StructField())
		tag := field.Tag.Get("tag")
		if tag != "" {
			result[err.Field()] = strings.Replace(err.Translate(trans), err.Field(), tag, -1)
		} else {
			result[err.Field()] = err.Translate(trans)
		}
	}
	return RemoveTopStruct(result)
}

func InStatusRange(fl validator.FieldLevel) bool {
	for _, status := range models.StatusRange {
		if fl.Field().Int() == int64(status) {
			return true
		}
	}
	return false
}

func IsMobile(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	result, _ := regexp.MatchString(`^(0?1[3|4|5|7|8][0-9]\d{4,8})$`, fl.Field().String())
	if result {
		return true
	}
	return false
}

func IsDateRange(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	result, _ := regexp.MatchString(`^\d{4}[/-]\d{1,2}[/-]\d{1,2}\~\d{4}[/-]\d{1,2}[/-]\d{1,2}$`, fl.Field().String())
	if result {
		return true
	}
	return false
}

func CheckDate(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}

func CheckDateAfterNow(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

func CheckDateBeforeNow(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.After(time.Now()) {
		return false
	}
	return true
}
