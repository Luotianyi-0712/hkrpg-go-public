package dispatch

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *Server) ComboGranterApiGetConfigHandler(c *gin.Context) {
	getConfigrsq := new(constant.GranterApiGetConfig)

	data := &constant.GranterApiGetConfigData{
		Protocol:                true,
		QrEnabled:               true,
		LogLevel:                "INFO",
		AnnounceURL:             s.OuterAddr,
		PushAliasType:           0,
		DisableYsdkGuard:        true,
		EnableAnnouncePicPopup:  true,
		AppName:                 "崩坏RPG",
		FunctionalSwitchConfigs: make([]string, 0),
	}
	getConfigrsq.Retcode = 0
	getConfigrsq.Message = "OK"
	getConfigrsq.Data = data

	c.JSON(200, getConfigrsq)
}

func (s *Server) GetExperimentListHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"success\":true,\"message\":\"\",\"data\":[{\"code\":1000,\"type\":2,\"config_id\":\"14\",\"period_id\":\"6125_197\",\"version\":\"1\",\"configs\":{\"cardType\":\"direct\"}}]}")
}

func (s *Server) SdkDataUploadHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) apmdataUpload(c *gin.Context) {
	req := c.Request
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	logger.Debug("/apm/dataUpload", string(body))
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) commonh5log(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"success\",\"data\":null}")
}

func (s *Server) GetAgreementInfos(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"message\":\"OK\",\"data\":{\"marketing_agreements\":[]}}")
}

func (s *Server) beforeVerify(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"is_heartbeat_required\":false,\"is_realname_required\":false,\"is_guardian_required\":false}}")
}

func (s *Server) redDotList(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"infos\":[]}}")
}

func (s *Server) listPriceTier(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"suggest_currency\":\"CNY\",\"tiers\":[{\"tier_id\":\"Alternate_Tier_1\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_18\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"11800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_28\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"18800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_32\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"21800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_5\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"3000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_56\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"51800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_59\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"61800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_8\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"5000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_4\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"2800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_A\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"100\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_37\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"24300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_50\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"32800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_6\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"4000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_1\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"600\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_21\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"13800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_39\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"25300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_48\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"30800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_62\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"79800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_77\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"164800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_78\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"199800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_11\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"7300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_23\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"15300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_27\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"17800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_30\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"19800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_33\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"22300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_40\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"25800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_58\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"58800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_73\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"139800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_10\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"6800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_43\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"27300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_47\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"29800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_52\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"38800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_15\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"9800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_34\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"22800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_63\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"81800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_64\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"84800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_74\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"144800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_75\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"149800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_20\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"12800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_22\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"14800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_4\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"2500\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_44\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"27800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_53\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"41800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_31\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"20800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_66\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"99800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_B\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_49\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"31800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_60\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"64800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_69\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"114800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_71\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"124800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_9\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"6000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_13\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"8800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_26\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"16800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_38\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"24800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_70\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"119800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_72\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"129800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_17\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"11300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_29\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"19300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_42\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"26800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_45\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"28300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_55\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"48800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_7\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"4500\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_14\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"9300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_3\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_61\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"69800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_76\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"159800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_2\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1200\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_24\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"15800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_46\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"28800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_51\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"34800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_54\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"44800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_79\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"229800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_36\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"23800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_57\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"54800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_65\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"89800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_80\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"259800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_5\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"3000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_16\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"10800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_35\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"23300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_41\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"26300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_67\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"104800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_3\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_12\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"7800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_19\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"12300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_2\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1200\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_25\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"16300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_68\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"109800\",\"symbol\":\"￥\"}]}],\"price_tier_version\":\"0\"}}")
}

func (s *Server) ExchangeCdkey(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"msg\":\"兑换成功\"}}")
}
