// Package user 业务
package user

import (
	"github.com/thoohv5/template/internal/service/user"
	"github.com/thoohv5/template/internal/service/user/entity"
	"github.com/thoohv5/template/pkg/http"

	"github.com/gin-gonic/gin"

	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/pkg/log"
)

type server struct {
	// 配置
	cf config.IConfig
	// 日志
	log log.ILog

	// 业务
	svr user.IService
}

// IServer 用户服务标准
type IServer interface {
	// Register 用户注册
	Register(gtx *gin.Context)
	// Detail 用户详情
	Detail(gtx *gin.Context)
	// List 用户列表
	List(gtx *gin.Context)
	// Edit 用户编辑
	Edit(gtx *gin.Context)
}

// New 创建
func New(cf config.IConfig, log log.ILog, svr user.IService) IServer {
	return &server{
		cf:  cf,
		log: log,

		svr: svr,
	}
}

// Register godoc
// @Summary 用户注册
// @Description 用户注册的接口
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param req body RegisterReq true "请求值"
// @Success 200 {object} http.ResponseEntity{data=RegisterResp}
// @Router /register [post]
func (s *server) Register(gtx *gin.Context) {
	resp := http.NewResponse(gtx)

	req := new(RegisterReq)
	if err := gtx.Bind(req); nil != err {
		resp.Error(err)
		return
	}

	// 创建
	ret, err := s.svr.Create(gtx, req.AccountType, req.Identity, req.Extra)
	if nil != err {
		resp.Error(err)
		return
	}
	resp.DefaultSuccess(&RegisterResp{
		UserIdentity: ret,
	})
}

// Detail godoc
// @Summary 用户详情
// @Description 用户的详情
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param req body DetailReq true "请求值"
// @Success 200 {object} http.ResponseEntity{data=DetailResp}
// @Router /detail [get]
func (s *server) Detail(gtx *gin.Context) {
	resp := http.NewResponse(gtx)

	req := new(DetailReq)
	if err := gtx.Bind(req); nil != err {
		resp.Error(err)
		return
	}

	// 详情
	ret, err := s.svr.Detail(gtx, req.AccountType, req.Identity)
	if nil != err {
		resp.Error(err)
		return
	}
	resp.DefaultSuccess(&DetailResp{
		Detail: ret,
	})
}

// List godoc
// @Summary 用户列表
// @Description 用户的列表
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param req body ListReq true "请求值"
// @Success 200 {object} http.ResponseEntity{data=ListResp}
// @Router /list [get]
func (s *server) List(gtx *gin.Context) {
	resp := http.NewResponse(gtx)

	req := new(ListReq)
	if err := gtx.Bind(req); nil != err {
		resp.Error(err)
		return
	}

	// 列表
	list, err := s.svr.List(gtx, &entity.ListParam{
		BasePage: http.BasePage{
			Start: req.Start,
			Limit: req.Limit,
		},
		Type: req.Type,
	})
	if nil != err {
		resp.Error(err)
		return
	}
	resp.DefaultSuccess(&ListResp{
		List: list,
	})
}

// Edit godoc
// @Summary 用户编辑
// @Description 编辑用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param req body EditReq true "请求值"
// @Success 200 {object} http.ResponseEntity
// @Router /edit [post]
func (s *server) Edit(gtx *gin.Context) {
	resp := http.NewResponse(gtx)

	req := new(EditReq)
	if err := gtx.Bind(req); nil != err {
		resp.Error(err)
		return
	}

	// 编辑
	err := s.svr.Edit(gtx, req.UserIdentity, &entity.EditParam{
		AccountType: req.AccountType,
		Identity:    req.Identity,
		Extra:       req.Extra,
	})
	if nil != err {
		resp.Error(err)
		return
	}
	resp.DefaultSuccess(nil)
}
