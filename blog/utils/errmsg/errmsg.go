package errmsg

type ErrCode int

const (
	SUCCEED = 200
	ERROR   = 500

	//CODE=1000 用户模块的错误
	ERROR_USERNAME_USED       = 1001
	ERROR_PASSWORD_INCORRECT  = 1002
	ERROR_USER_DOES_NOT_EXIST = 1003
	/*
		token登陆的错误码
	*/
	/*
		session登陆的错误码
	*/

	//CODE=2000 文章模块的错误
	ERROR_ARTICLE_DOES_NOT_EXIST       = 2001 //文章不存在
	ERROR_ARTICLE_TITLE_ALREADY_EXISTS = 2002 //文章标题重复

	//CODE=3000 分类模块的错误
	ERROR_CATEGORY_DOES_NOT_EXIST = 3001 //分类不存在
	ERROR_CATEGORY_ALREADY_EXIXTS = 3002 //分类已存在

	//CODE=4000 tag模块的错误
	ERROR_TAG_DOES_NOT_EXIST = 4001 //标签不存在
	ERROR_TAG_ALREADY_EXIXTS = 4002 //标签已存在
)

var codeMsg = map[ErrCode]string{
	SUCCEED: "OK",
	ERROR:   "Fail",

	ERROR_USERNAME_USED:       "用户名重复",
	ERROR_PASSWORD_INCORRECT:  "密码错误",
	ERROR_USER_DOES_NOT_EXIST: "用户不存在",

	ERROR_ARTICLE_DOES_NOT_EXIST:       "文章不存在",
	ERROR_ARTICLE_TITLE_ALREADY_EXISTS: "文章标题重复",

	ERROR_CATEGORY_DOES_NOT_EXIST: "分类不存在",
	ERROR_CATEGORY_ALREADY_EXIXTS: "分类已存在",

	ERROR_TAG_DOES_NOT_EXIST: "tag已存在",
	ERROR_TAG_ALREADY_EXIXTS: "tag已存在",
}

func GetErrMsg(code ErrCode) string {
	return codeMsg[code]
}
