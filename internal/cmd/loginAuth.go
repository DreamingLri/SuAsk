package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
	v1 "suask/api/login/v1"
	"suask/internal/consts"
	"suask/internal/dao"
	"suask/internal/model/entity"
	"suask/utility"
	"suask/utility/response"
)

func LoginToken() (gfToken *gtoken.GfToken, err error) {
	gfToken = &gtoken.GfToken{
		CacheMode:       consts.CacheMode,
		ServerName:      consts.ServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFuncFrontend,
		LoginAfterFunc:  loginAfterFunc,
		LogoutPath:      "/user/logout",
		MultiLogin:      true,
		AuthAfterFunc:   authAfterFunc,
		AuthPaths:       g.SliceStr{},
	}
	err = gfToken.Start()
	return
}

func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	email := r.Get("email").String()
	password := r.Get("password").String()

	ctx := context.TODO()
	// 输入参数为空
	if (password == "" && name == "") || (password == "" && email == "") {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	userInfo := entity.Users{}
	var err error
	if name != "" {
		err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Name, name).Scan(&userInfo)
	}
	if email != "" {
		err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Email, email).Scan(&userInfo)
	}
	// 查不到用户
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(err.Error()))
		r.ExitAll()
	}
	// 密码校验失败
	if utility.EncryptPassword(password, userInfo.Salt) != userInfo.PasswordHash {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	return strconv.Itoa(userInfo.Id), userInfo
}

func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		userId := respData.GetString("userKey")
		userInfo := entity.Users{}
		err := dao.Users.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		var avatarURL string
		err = dao.Files.Ctx(context.TODO()).WherePri(userInfo.AvatarFileId).Scan(&avatarURL)
		if err != nil {
			avatarURL = consts.DefaultAvatarURL
		}
		data := &v1.LoginRes{
			Token: respData.GetString("token"),
			Type:  consts.TokenType,
		}
		data.Name = userInfo.Name
		data.Email = userInfo.Email
		data.Nickname = userInfo.Nickname
		data.Introduction = userInfo.Introduction
		data.Id = userInfo.Id
		data.Role = userInfo.Role
		data.AvatarURL = avatarURL
		data.ThemeId = userInfo.ThemeId
		response.JsonExit(r, 0, "login success", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo v1.LoginRes
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//fmt.Printf("resp", userInfo)
	r.SetCtxVar(consts.CtxId, userInfo.Id)
	r.SetCtxVar(consts.CtxName, userInfo.Name)
	r.SetCtxVar(consts.CtxRole, userInfo.Role)
	r.Middleware.Next()
}