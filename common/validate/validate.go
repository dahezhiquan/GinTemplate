package validate

import (
	"regexp"
	"unicode"
)

// VerifyMobile 校验手机号合法性
func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyPassword(password string) (b bool) {
	regex, _ := regexp.Compile(`^[A-Za-z0-9@#$%^&+-=!_?]{8,16}$`)
	if !regex.MatchString(password) {
		return false
	}

	var (
		hasLetter  bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasLetter && hasNumber && hasSpecial
}

func VerifyUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9 _-]{4,16}$", username); !ok {
		return false
	}
	return true
}

// QQ号正则判断

func VerifyQQFormat(qq string) bool {
	pattern := `^[1-9][0-9]{4,10}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(qq)
}

// 微信号正则判断

func VerifyWechatFormat(wechat string) bool {
	pattern := `^[a-zA-Z][-_a-zA-Z0-9]{5,19}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(wechat)
}
