// Code generated by hertz generator. DO NOT EDIT.

package core_api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	core_api "github.com/xh-polaris/openapi-core-api/biz/adaptor/controller/core_api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_base := root.Group("/base", _baseMw()...)
		_base.POST("/create", append(_createbaseinterfaceMw(), core_api.CreateBaseInterface)...)
		_base.GET("/delete", append(_deletebaseinterfaceMw(), core_api.DeleteBaseInterface)...)
		_base.POST("/get", append(_getbaseinterfacesMw(), core_api.GetBaseInterfaces)...)
		_base.POST("/update", append(_updatebaseinterfaceMw(), core_api.UpdateBaseInterface)...)
	}
	{
		_call := root.Group("/call", _callMw()...)
		_call.POST("/", append(_callinterfaceMw(), core_api.CallInterface)...)
	}
	{
		_full := root.Group("/full", _fullMw()...)
		_full.POST("/buy", append(_buyfullinterfaceMw(), core_api.BuyFullInterface)...)
		_full.POST("/create", append(_createfullinterfaceMw(), core_api.CreateFullInterface)...)
		_full.GET("/delete", append(_deletefullinterfaceMw(), core_api.DeleteFullInterface)...)
		_full.POST("/get", append(_getfullinterfacesMw(), core_api.GetFullInterfaces)...)
		_full.POST("/margin", append(_updatemarginMw(), core_api.UpdateMargin)...)
		_full.POST("/update", append(_updatefullinterfaceMw(), core_api.UpdateFullInterface)...)
	}
	{
		_gradient := root.Group("/gradient", _gradientMw()...)
		_gradient.POST("/amount", append(_getamountMw(), core_api.GetAmount)...)
		_gradient.POST("/create", append(_creategradientMw(), core_api.CreateGradient)...)
		_gradient.GET("/get", append(_getgradientMw(), core_api.GetGradient)...)
		_gradient.POST("/update", append(_updategradientMw(), core_api.UpdateGradient)...)
	}
	{
		_key := root.Group("/key", _keyMw()...)
		_key.GET("/delete", append(_deletekeyMw(), core_api.DeleteKey)...)
		_key.POST("/generate", append(_generatekeyMw(), core_api.GenerateKey)...)
		_key.POST("/get", append(_getkeysMw(), core_api.GetKeys)...)
		_key.POST("/hosts", append(_updatehostsMw(), core_api.UpdateHosts)...)
		_key.GET("/refresh", append(_refreshkeyMw(), core_api.RefreshKey)...)
		_key.POST("/update", append(_updatekeyMw(), core_api.UpdateKey)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/info", append(_getuserinfoMw(), core_api.GetUserInfo)...)
		_user.POST("/info", append(_setuserinfoMw(), core_api.SetUserInfo)...)
		_user.POST("/sign_in", append(_signinMw(), core_api.SignIn)...)
		_user.POST("/sign_up", append(_signupMw(), core_api.SignUp)...)
	}
}
