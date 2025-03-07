package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",

	ErrorExistUser:             "用户已存在",
	ErrorFailEncryption:        "加密失败",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthToken:             "token签发失败",
	ErrorAuthTokenWrong:        "token验证错误",
	ErrorAuthCheckTokenTimeout: "token已过期",
	ErrorUploadFail:            "图片上传失败",
	ErrorSendEmail:             "发送邮件失败",

	ErrorProductImgUpload: "图片上传错误",

	ErrorFavoriteExist: "收藏夹错误",

	ErrorCartExist: "购物车错误",
}

// GetMsg获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
