package dispatch

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRouter() {
	s.Router.Use(clientIPMiddleware())
	s.Router.Any("/", s.HandleDefault)
	s.Router.Any("/index.html", s.HandleDefault)
	s.Router.Any("/sdk/dataUpload", s.SdkDataUploadHandler)
	s.Router.POST("/apm/dataUpload", s.apmdataUpload)
	s.Router.POST("/common/h5log/log/batch", s.commonh5log)

	// 调度
	s.Router.GET("/query_dispatch", s.QueryDispatchHandler)
	s.Router.GET("/query_dispatch/gucooing/az", s.QueryDispatchHandler)
	s.Router.GET("/query_gateway", s.QueryGatewayHandler)
	s.Router.GET("/query_gateway_capture", s.QueryGatewayHandlerCapture)
	s.Router.GET("/query_gateway_capture_cn", s.QueryGatewayHandlerCaptureCn)

	// 登录
	s.Router.POST("/account/risky/api/check", s.RiskyApiCheckHandler)
	Global := s.Router.Group("/hkrpg_:type")
	{
		Global.GET("/combo/granter/api/getConfig", s.ComboGranterApiGetConfigHandler) // 获取服务器配置
		Global.POST("/combo/granter/api/compareProtocolVersion", s.compareProtocolVersion)
		Global.POST("/mdk/shield/api/login", s.LoginRequestHandler)   // 账号登录
		Global.POST("/mdk/shield/api/verify", s.VerifyRequestHandler) // token登录
		Global.GET("/mdk/shield/api/loadConfig", s.loadConfig)
		Global.POST("/combo/granter/login/v2/login", s.V2LoginRequestHandler) // 获取combo token
		Global.GET("/mdk/agreement/api/getAgreementInfos", s.GetAgreementInfos)
		Global.POST("/combo/granter/login/beforeVerify", s.beforeVerify)
		Global.POST("/combo/red_dot/list", s.redDotList)
		Global.POST("/mdk/shopwindow/shopwindow/listPriceTier", s.listPriceTier)
		Global.POST("/mdk/shopwindow/shopwindow/listPriceTierV2", s.listPriceTier)
		Global.GET("/mdk/shopwindow/shopwindow/listPriceTierV2", s.listPriceTier)
	}
	// 杂
	s.Router.POST("/data_abtest_api/config/experiment/list", s.GetExperimentListHandler)
	s.Router.GET("/common/apicdkey/api/exchangeCdkey", s.ExchangeCdkey)
}

func (s *Server) HandleDefault(c *gin.Context) {
	c.String(200, "hkrpg-go")
}
