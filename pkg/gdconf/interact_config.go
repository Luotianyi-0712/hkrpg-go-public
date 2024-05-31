package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type InteractConfig struct {
	InteractID  uint32 `json:"InteractID"`
	SrcState    string `json:"SrcState"`
	TargetState string `json:"TargetState"`
	IsEvent     bool   `json:"IsEvent"`
	// ItemCostList []uint32 `json:"ItemCostList"`
}

func (g *GameDataConfig) loadInteractConfig() {
	g.InteractConfigMap = make(map[uint32]*InteractConfig)
	interactConfigMap := make(map[string]*InteractConfig)
	playerElementsFilePath := g.excelPrefix + "InteractConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &interactConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for id, interactConfig := range interactConfigMap {
		g.InteractConfigMap[alg.S2U32(id)] = interactConfig
	}
	logger.Info("load %v InteractConfig", len(g.InteractConfigMap))
}

func GetInteractConfigMap() map[uint32]*InteractConfig {
	return CONF.InteractConfigMap
}

func GetInteractConfigById(id uint32) *InteractConfig {
	return CONF.InteractConfigMap[id]
}
