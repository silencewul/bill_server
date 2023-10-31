package constant

import "errors"

type ResponseCode int

const (
	CodeSuccess ResponseCode = 0
	CodeError   ResponseCode = 1
	//CodeInvalidParams       ResponseCode = 400
	//CodeUnauthorized        ResponseCode = 401
	//CodeResourceNotFound    ResponseCode = 404
	//CodeServerInternalError ResponseCode = 500
	//CodeCreateFail          ResponseCode = 2
	//CodeUpdateFail          ResponseCode = 3
	//CodeDeleteFail          ResponseCode = 4
	//
	//// 用户相关
	//CodeUserNotFound     ResponseCode = 1000
	//CodeUserNameInvalid  ResponseCode = 1001
	//CodeUserEmailTaken   ResponseCode = 1002
	//CodeUserNameTaken    ResponseCode = 1003
	//CodeUserPwd2Mismatch ResponseCode = 1004
	//CodeUserPwdMismatch  ResponseCode = 1005
	//
	//// 课程相关 从2000开始
	//CodeCourseNotFound          ResponseCode = 2000
	//CodeCourseCreateFail        ResponseCode = 2001
	//CodeCourseEditFail          ResponseCode = 2002
	//CodeCourseMemberRefreshFail ResponseCode = 2003
	//CodeCourseNoAuth            ResponseCode = 2004
	//CodeCourseMemberJoinedFail  ResponseCode = 2005
	//CodeCourseLessonLearnFail   ResponseCode = 2006
	//
	////课程分类
	//CodeCategoryNameExist ResponseCode = 2100
	//CodeCategoryCodeExist ResponseCode = 2101
	////标签
	//CodeTagNameExist ResponseCode = 2200
	//
	//// 杂项
	//CodeTokenNotFound     ResponseCode = 3000
	//CodeTokenExpired      ResponseCode = 3001
	//CodeTokenTypeMismatch ResponseCode = 3002
	//CodeTokenDataMismatch ResponseCode = 3003
	//CodeSmsCodeMismatch   ResponseCode = 3004
	//
	//CodeOauthNotSupported ResponseCode = 3100
	//
	////文件上传
	//CodeFileNotFound  ResponseCode = 4000
	//CodeFileTypeErr   ResponseCode = 4001
	//CodeFileSizeLimit ResponseCode = 4002
	//CodeFileNotUpload ResponseCode = 4003
	//CodeFileEmptyErr  ResponseCode = 4004
)

var (
	// 通用错误组
	ErrUnknown             = errors.New("未知错误")
	ErrSiteConfigErr       = errors.New("网站数据获取错误")
	ErrInvalidParams       = errors.New("参数无效")
	ErrUnauthorized        = errors.New("用户没有被认证")
	ErrPermissionInvalid   = errors.New("权限无效")
	ErrPermissionDenied    = errors.New("没有权限")
	ErrResourceNotFound    = errors.New("访问的资源不存在")
	ErrServerInternalError = errors.New("服务器内部错误")
	ErrTooManyRequest      = errors.New("操作频繁")
	ErrCreateFail          = errors.New("创建失败")
	ErrUpdateFail          = errors.New("更新失败")
	ErrDeleteFail          = errors.New("删除失败")
	ErrUnderDevelopment    = errors.New("还在开发中")

	// 用户相关
	ErrUserNotFound             = errors.New("用户不存在")
	ErrUserNameInvalid          = errors.New("用户名不合法")
	ErrUserEmailInvalid         = errors.New("邮箱不合法")
	ErrUserEmailTaken           = errors.New("邮箱已经被占用")
	ErrUserEmailOrMobileInvalid = errors.New("手机号或者邮箱不合法")
	ErrUserMobileInvalid        = errors.New("手机号不合法")
	ErrUserNameTaken            = errors.New("用户名已经被占用")
	ErrUserMobileTaken          = errors.New("手机号已经被占用")

	ErrUserPwdInvalid        = errors.New("用户密码不符合要求")
	ErrUserPwd2Mismatch      = errors.New("两次输入的密码不匹配")
	ErrUserPwdMismatch       = errors.New("密码错误")
	ErrUserOriginPwdMismatch = errors.New("原密码不正确")

	ErrRoleCodeTaken = errors.New("角色编码已经存在")
	ErrRoleNotFound  = errors.New("角色不存在")
	ErrRoleLocked    = errors.New("系统默认角色禁止修改/删除")

	ErrCodeInvalid              = errors.New("验证码格式不对")
	ErrRegisterSmsCodeInvalid   = errors.New("手机验证码错误")
	ErrRegisterEmailCodeInvalid = errors.New("邮箱验证码错误")

	// 课程相关

	// 杂项
	ErrTokenNotFound = errors.New("Token不存在")
	ErrTokenInvalid  = errors.New("Token无效")

	ErrTokenExpired      = errors.New("Token过期")
	ErrTokenTypeMismatch = errors.New("Token类型不匹配")
	ErrTokenDataMismatch = errors.New("Token数据不匹配")
	ErrSmsCodeMismatch   = errors.New("短信验证码错误,你还有%d次机会!")

	ErrOauthNotSupported = errors.New("不支持此类型的第三方登录")
)
