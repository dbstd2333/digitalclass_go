// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloudclass_go/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/class/login",
				Handler: classloginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/get/subject",
				Handler: classgetsubjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/get/msg",
				Handler: classgetmsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/get/seat",
				Handler: classgetseatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/get/group",
				Handler: classgetgroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/get/daily",
				Handler: classgetdailyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/upload/readed",
				Handler: classupreadedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/class/get/phone",
				Handler: classgetphoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/student/login",
				Handler: studentloginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/student/change/group",
				Handler: studentchangegroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/student/change/seat",
				Handler: studentchangeseatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/student/upload/phone",
				Handler: studentupphoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/login",
				Handler: teacherloginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/change/subject",
				Handler: teacherchangesubjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/change/seat",
				Handler: teacherchangeseatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/change/group",
				Handler: teacherchangegroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/change/daily",
				Handler: teacherchangedailyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/upload/msg",
				Handler: teacheruploadmsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/teacher/get/readclass",
				Handler: teachergetreadclassHandler(serverCtx),
			},
		},
	)
}
