package player

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) SetTurnFoodSwitchCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetTurnFoodSwitchCsReq)
	rsp := &proto.SetTurnFoodSwitchScRsp{
		Retcode:     0,
		DGLLJFNEMOK: req.DGLLJFNEMOK,
		LNEKBEGKACP: req.LNEKBEGKACP,
	}
	g.Send(cmd.SetTurnFoodSwitchScRsp, rsp)
}

func (g *GamePlayer) SceneCastSkillCostMpCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneCastSkillCostMpCsReq)
	rsp := &proto.SceneCastSkillCostMpScRsp{
		CastEntityId: req.CastEntityId,
		Retcode:      0,
	}
	g.Send(cmd.SceneCastSkillCostMpScRsp, rsp)
}

func (g *GamePlayer) SceneEnterStageCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneEnterStageCsReq)
	rsp := &proto.SceneEnterStageScRsp{
		Retcode: 0,
	}
	battleBackup := &model.BattleBackup{
		IsBattle:   true,
		WorldLevel: g.GetPd().GetWorldLevel(),
		Sce:        new(model.SceneCastEntity),
		EventId:    req.EventId,
	}
	battleBackup.Sce.EvenIdList = []uint32{req.EventId}
	// 获取战斗角色
	avatarMap := make(map[uint32]*model.BattleAvatar, 0)
	planeEvent := gdconf.GetPlaneEventById(req.EventId, battleBackup.WorldLevel)
	if planeEvent == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.SceneEnterStageScRsp, rsp)
		return
	}
	stageConfig := gdconf.GetStageConfigById(planeEvent.StageID)
	if stageConfig == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.SceneEnterStageScRsp, rsp)
		return
	}
	if stageConfig.TrialAvatarList != nil {
		for _, trial := range stageConfig.TrialAvatarList {
			conf := gdconf.GetSpecialAvatarById(trial)
			if conf == nil {
				continue
			}
			baseAvatarId := conf.AvatarID
			avatarMap[baseAvatarId] = &model.BattleAvatar{
				AvatarId:     trial,
				BaseAvatarId: baseAvatarId,
				AvatarType:   spb.LineAvatarType_LineAvatarType_TRIAL,
				AssistUid:    0,
			}
		}
	}
	if len(avatarMap) == 0 {
		battleBackup.BattleAvatarList = g.GetPd().GetBattleAvatarMap(g.GetPd().GetBattleLineUp())
	} else {
		battleBackup.BattleAvatarList = avatarMap
	}
	// battleBackup.BattleAvatarList = avatarMap
	battleInfoPb := g.GetPd().GetSceneBattleInfo(battleBackup)
	rsp.BattleInfo = battleInfoPb
	// 记录战斗
	g.GetPd().AddBattleBackup(battleBackup)
	g.Send(cmd.SceneEnterStageScRsp, rsp)
}

/***********************************攻击事件处理***********************************/

func (g *GamePlayer) SceneCastSkillCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneCastSkillCsReq)
	rsp := &proto.SceneCastSkillScRsp{
		CastEntityId: req.CastEntityId, // 攻击唯一id
	}
	battleBackup := &model.BattleBackup{
		IsBattle:   false,
		Sce:        new(model.SceneCastEntity), // 添加参与此次攻击的实体
		WorldLevel: g.GetPd().GetWorldLevel(),
	}

	// HitTargetEntityIdList 被攻击目标
	// AssistMonsterEntityIdList 击中列表
	// AssistMonsterEntityInfo 战斗群
	// AttackedByEntityId 攻击发起者

	// 添加攻击发起者
	g.GetPd().GetMem([]uint32{req.AttackedByEntityId}, battleBackup.Sce)
	// 添加被攻击者
	var isAttac = false
	if req.AssistMonsterEntityInfo != nil {
		for _, info := range req.AssistMonsterEntityInfo {
			entityIdList := make([]uint32, 0)
			for _, entityId := range info.EntityIdList {
				if entityId == req.AttackedByEntityId {
					isAttac = true
					continue
				}
				entityIdList = append(entityIdList, entityId)
			}
			g.GetPd().GetMem(entityIdList, battleBackup.Sce)
		}
	} else {
		isAttac = true
	}
	if isAttac {
		g.GetPd().GetMem(req.AssistMonsterEntityIdList, battleBackup.Sce)
	}

	g.SceneCastSkillProp(battleBackup.Sce) // 物品效果
	var skill *gdconf.GoppMazeSkill
	if req.SkillIndex != 0 {
		skill = gdconf.GetGoppMazeSkill(battleBackup.Sce.AvatarId, 2)
		if g.GetPd().DelMp(battleBackup.Sce.AvatarId) {
			g.GetPd().DelLineUpMp(1)
			g.Send(cmd.SceneCastSkillMpUpdateScNotify, &proto.SceneCastSkillMpUpdateScNotify{
				CastEntityId: req.CastEntityId,
				Mp:           g.GetPd().GetLineUpMp(),
			})
			g.SyncLineupNotify(g.GetPd().GetBattleLineUp())
		}
	} else {
		skill = gdconf.GetGoppMazeSkill(battleBackup.Sce.AvatarId, 1)
	}
	if skill == nil {
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	g.GetPd().SceneCastSkill(battleBackup, skill, req)
	// summonun
	if battleBackup.SummonUnitId != 0 {
		db := g.GetPd().GetSummonUnitInfo()
		db.AvatarId = battleBackup.Sce.AvatarId
		db.AttachEntityId = battleBackup.Sce.AvatarEntityId
		db.EntityId = g.GetPd().GetNextGameObjectGuid()
		db.SummonUnitId = battleBackup.SummonUnitId
		db.Pos = req.TargetMotion.Pos
		g.AddSummonUnitSceneGroupRefreshScNotify()
	}
	if len(battleBackup.Sce.EvenIdList) == 0 || !battleBackup.IsBattle { // 是否满足战斗条件
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	// 获取战斗角色
	battleBackup.BattleAvatarList = g.GetPd().GetBattleAvatarMap(g.GetPd().GetBattleLineUp())
	battleInfoPb := g.GetPd().GetSceneBattleInfo(battleBackup)
	// 记录战斗
	g.GetPd().AddBattleBackup(battleBackup)
	// 回复
	rsp.BattleInfo = battleInfoPb
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

/***********************************战斗结算***********************************/

func (g *GamePlayer) PVEBattleResultCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PVEBattleResultCsReq)
	battleBin := g.GetPd().GetBattleBackupById(req.BattleId)
	if battleBin == nil {
		return
	}
	sce := battleBin.Sce
	var teleportToAnchor = false
	rsp := &proto.PVEBattleResultScRsp{
		BattleAvatarList: make([]*proto.BattleAvatar, 0),
		BattleId:         req.BattleId,
		StageId:          req.StageId,
		EndStatus:        req.EndStatus, // 战斗结算状态
		CheckIdentical:   true,          // 反作弊验证
		BinVersion:       "",
		ResVersion:       strconv.Itoa(int(req.ClientVersion)), // 版本验证
		DropData:         &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}
	// 更新角色状态
	g.GetPd().BattleUpAvatar(req.Stt.GetBattleAvatarList(), req.GetEndStatus())
	g.SyncLineupNotify(g.GetPd().GetBattleLineUp())

	allSync := &model.AllPlayerSync{
		IsBasic:       true,
		MaterialList:  make([]uint32, 0),
		EquipmentList: nil,
		RelicList:     nil,
	}
	addPileItem := make([]*model.Material, 0)
	delPileItem := make([]*model.Material, 0)

	// 根据不同结算状态处理
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_WIN: // 胜利
		// 删除怪物实体
		if sce != nil {
			g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
				GroupRefreshList: g.GetPd().GetDelSceneGroupRefreshInfo(sce.MonsterEntityIdList),
			})
		}
		// 任务判断
		if battleBin.EventId != 0 {
			rsp.EventId = battleBin.EventId
			finishSubMission := g.GetPd().UpBattleSubMission(req.BattleId)
			if req.Stt.CustomValues != nil {
				finishSubMission = append(finishSubMission, g.GetPd().BattleCustomValues(req.Stt.CustomValues, battleBin.EventId)...)
			}
			if len(finishSubMission) != 0 {
				g.InspectMission(finishSubMission)
			}
			entity := g.GetPd().GetTriggerBattleString(battleBin.EventId)
			if entity != nil {
				blockBin := g.GetPd().GetBlock(g.GetPd().GetCurEntryId())
				g.GetPd().UpPropState(blockBin, entity.GroupId, entity.InstId, 1)    // 更新状态
				g.PropSceneGroupRefreshScNotify([]uint32{entity.EntityId}, blockBin) // 通知状态更改
			}
		}
		// 任务判断二
		if sce != nil {
			for _, entityId := range sce.MonsterEntityIdList {
				me := g.GetPd().GetMonsterEntityById(entityId)
				if me != nil {
					finishSubMission := g.GetPd().UpKillMonsterSubMission(me)
					if len(finishSubMission) != 0 {
						g.InspectMission(finishSubMission)
					}
				}
			}
		}
		// 以太战线任务判断
		if battleBin.AetherDivideId != 0 {
			finishSubMission := g.GetPd().AetherDivideCertainFinishHyperlinkDuel(battleBin.AetherDivideId)
			if len(finishSubMission) != 0 {
				g.InspectMission(finishSubMission)
			}
		}
		// 参战角色经验添加
		for _, avatar := range req.Stt.GetBattleAvatarList() {
			if _, ok := g.GetPd().AvatarAddExp(avatar.Id, battleBin.AvatarExpReward); ok {
				allSync.AvatarList = append(allSync.AvatarList, avatar.Id)
			}
		}
		// 获取奖励
		addPileItem = append(addPileItem, battleBin.DisplayItemList...)
		for _, displayItem := range battleBin.DisplayItemList {
			rsp.DropData.ItemList = append(rsp.DropData.ItemList, &proto.Item{
				ItemId: displayItem.Tid,
				Num:    displayItem.Num,
			})
		}
		if conf := gdconf.GetCocoonConfigById(battleBin.CocoonId, battleBin.WorldLevel); conf != nil { // 副本处理
			rsp.DropData.ItemList = append(rsp.DropData.ItemList,
				g.GetPd().GetBattleDropData(conf.MappingInfoID, addPileItem, battleBin.WorldLevel, allSync)...)
			finishSubMission := g.GetPd().FinishCocoon(battleBin.CocoonId)
			if len(finishSubMission) != 0 {
				g.InspectMission(finishSubMission)
			}
			delPileItem = append(delPileItem, &model.Material{
				Tid: model.Stamina,
				Num: conf.StaminaCost,
			})
		}
		if conf := gdconf.GetFarmElementConfig(req.StageId); conf != nil {
			rsp.DropData.ItemList = append(rsp.DropData.ItemList,
				g.GetPd().GetBattleDropData(conf.MappingInfoID, addPileItem, battleBin.WorldLevel, allSync)...)
			delPileItem = append(delPileItem, &model.Material{
				Tid: model.Stamina,
				Num: conf.StaminaCost,
			})
		}
	case proto.BattleEndStatus_BATTLE_END_LOSE: // 失败
		teleportToAnchor = true
	case proto.BattleEndStatus_BATTLE_END_QUIT:
		teleportToAnchor = true
	}

	switch g.GetPd().GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		g.ChallengePVEBattleResultCsReq(req)
	case spb.BattleType_Battle_CHALLENGE_Story:
		g.ChallengePVEBattleResultCsReq(req)
	case spb.BattleType_Battle_ROGUE:
		teleportToAnchor = false
		g.RoguePVEBattleResultCsReq(req, sce)
	case spb.BattleType_Battle_TrialActivity: // 角色试用
		g.TrialActivityPVEBattleResultScRsp(req)
		teleportToAnchor = true
	}

	// 是否传送到最近锚点
	if teleportToAnchor {
		// 当前坐标通知(移动到最近锚点)
		g.EnterSceneByServerScNotify(g.GetPd().GetCurEntryId(), 0, 0, 0)
	}

	g.GetPd().DelMaterial(delPileItem)
	g.GetPd().AddItem(addPileItem, allSync)
	g.StaminaInfoScNotify()
	g.AllPlayerSyncScNotify(allSync)
	// g.AllScenePlaneEventScNotify(addPileItem)

	g.GetPd().DelBattleBackupById(req.BattleId)
	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

/***********************************关卡/副本***********************************/

func (g *GamePlayer) StartCocoonStageCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartCocoonStageCsReq)
	rsp := &proto.StartCocoonStageScRsp{
		PropEntityId: req.PropEntityId,
		CocoonId:     req.CocoonId, // 关卡id
		Retcode:      0,
		Wave:         req.Wave,
	}
	battleBackup := &model.BattleBackup{
		IsBattle:   false,
		WorldLevel: req.WorldLevel,
	}
	cocoonConfig := gdconf.GetCocoonConfigById(req.CocoonId, battleBackup.WorldLevel)
	if cocoonConfig == nil {
		logger.Warn("No Cocoon like this can be found,cocoonId:%v,worldLevel:%v", req.CocoonId, battleBackup.WorldLevel)
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartCocoonStageScRsp, rsp)
		return
	}
	battleBackup.StageID = cocoonConfig.StageID
	battleBackup.StageIDList = cocoonConfig.StageIDList
	// 获取角色
	battleBackup.BattleAvatarList = g.GetPd().GetBattleAvatarMap(g.GetPd().GetBattleLineUp())
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE) // 设置战斗状态
	battleInfoPb := g.GetPd().GetSceneBattleInfo(battleBackup)
	if battleInfoPb == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartCocoonStageScRsp, rsp)
		return
	}
	rsp.BattleInfo = battleInfoPb
	// 储存战斗信息
	battleBackup.CocoonId = req.CocoonId
	battleBackup.WorldLevel = req.WorldLevel
	g.GetPd().AddBattleBackup(battleBackup)
	g.Send(cmd.StartCocoonStageScRsp, rsp)
}

func (g *GamePlayer) ActivateFarmElementCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ActivateFarmElementCsReq)

	rsp := &proto.ActivateFarmElementScRsp{
		WorldLevel: req.WorldLevel,
		EntityId:   req.EntityId,
	}
	g.Send(cmd.ActivateFarmElementScRsp, rsp)
}

func (g *GamePlayer) RefreshTriggerByClientCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.RefreshTriggerByClientCsReq)
	db := g.GetPd().GetSummonUnitInfo()
	if db.EntityId == req.TriggerEntityId {
		conf := gdconf.GetSummonUnitMazeSkillAction(db.SummonUnitId, req.TriggerName)
		if conf != nil {
			for _, action := range conf {
				if action.Type == constant.AddMazeBuff {
					if db.BuffList == nil {
						db.BuffList = make([]*model.OnBuffMap, 0)
					}
					db.BuffList = append(db.BuffList, &model.OnBuffMap{
						AvatarId:  0,
						BuffId:    action.Id,
						Level:     1,
						Count:     0,
						LifeCount: 0,
						AddTime:   uint64(time.Now().Unix()),
						LifeTime:  20,
					})
				}
			}
		}
	}

	rsp := &proto.RefreshTriggerByClientScRsp{
		RefreshTrigger:  true,
		Retcode:         0,
		TriggerName:     req.TriggerName,
		TriggerEntityId: req.TriggerEntityId,
	}
	g.Send(cmd.RefreshTriggerByClientScRsp, rsp)
}

/***********************************物品破坏处理***********************************/

func (g *GamePlayer) SceneCastSkillProp(sce *model.SceneCastEntity) {
	var addMPCost uint32 = 0
	allSync := &model.AllPlayerSync{AvatarList: make([]uint32, 0)}
	for _, propId := range sce.PropIdList {
		conf := gdconf.GetMazePropId(propId)
		if conf == nil {
			continue
		}
		if conf.RecoverMp {
			addMPCost += 2
		}
		if conf.RecoverHp {
			g.GetPd().AvatarRecoverPercent(sce.AvatarId, 0.3, 0)
			allSync.AvatarList = append(allSync.AvatarList, sce.AvatarId)
		}
	}
	if addMPCost > 0 {
		g.GetPd().AddLineUpMp(2) // 如果涉及到更新战斗中的队伍状态，这部分需要改
		g.SyncLineupNotify(g.GetPd().GetBattleLineUp())
	}
	g.AllPlayerSyncScNotify(allSync)
}
