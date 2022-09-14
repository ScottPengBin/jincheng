package valida

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func LoadZh()  {
	if v,ok :=binding.Validator.Engine().(*validator.Validate);ok{
		zht := zh.New()

		enT := en.New()

		uni := ut.New(enT,zht)

		trans,ok = uni.GetTranslator("zh")

		err := zhTranslations.RegisterDefaultTranslations(v,trans)

		if err!= nil{
			fmt.Println(err)
			return
		}
		return
	}
	return
}

func Trans(err error)  string{
	errs,ok:=err.(validator.ValidationErrors)
	if !ok{
		return "不是合法的数据格式"
	}
	var  errMsg string
	for _, fieldError := range errs {
		errMsg += fieldError.Translate(trans)
	}
	return errMsg
}