package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

// 通知客户端进入场景
func (g *GamePlayer) EnterSceneByServerScNotify(entryId, teleportId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return
	}
	teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if teleportsMap == nil {
		return
	}

	var anchorID = mapEntrance.StartAnchorID
	var groupID = mapEntrance.StartGroupID
	var pos *proto.Vector
	var rot *proto.Vector
	if teleportsMap.Teleports[teleportId] != nil {
		anchorID = teleportsMap.Teleports[teleportId].AnchorID
		groupID = teleportsMap.Teleports[teleportId].AnchorGroupID
	}
	// 获取队伍
	curLine := g.GetCurLineUp()
	rsp.Lineup = g.GetLineUpPb(curLine)
	// 获取坐标
	serverGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, groupID)
	if serverGroup != nil && serverGroup.AnchorList != nil && serverGroup.AnchorList[anchorID] != nil {
		anchor := serverGroup.AnchorList[anchorID]
		pos = &proto.Vector{
			X: int32(anchor.PosX * 1000),
			Y: int32(anchor.PosY * 1000),
			Z: int32(anchor.PosZ * 1000),
		}
		rot = &proto.Vector{
			X: int32(anchor.RotX * 1000),
			Y: int32(anchor.RotY * 1000),
			Z: int32(anchor.RotZ * 1000),
		}
	}
	// 实在找不到就随便找一个锚点传送
	if pos == nil {
		for _, anchor := range teleportsMap.Teleports {
			pos = &proto.Vector{
				X: int32(anchor.PosX * 1000),
				Y: int32(anchor.PosY * 1000),
				Z: int32(anchor.PosZ * 1000),
			}
			rot = &proto.Vector{
				X: int32(anchor.RotX * 1000),
				Y: int32(anchor.RotY * 1000),
				Z: int32(anchor.RotZ * 1000),
			}
			break
		}
	}
	if pos == nil {
		// 这都没有那就不要传送了
		logger.Debug("entryId:%v,teleportId:%v error", entryId, teleportId)
		return
	}
	rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)
	g.SetCurEntryId(entryId)
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

// 传送到指定位置
func (g *GamePlayer) SceneByServerScNotify(entryId uint32, pos, rot *proto.Vector) {
	rsp := new(proto.EnterSceneByServerScNotify)
	// 获取队伍
	curLine := g.GetCurLineUp()
	rsp.Lineup = g.GetLineUpPb(curLine)
	rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) HandleGetEnteredSceneCsReq(payloadMsg []byte) {
	rsp := new(proto.GetEnteredSceneScRsp)
	db := g.GetScene()
	mapEntrance := gdconf.GetMapEntranceById(db.EntryId)
	if mapEntrance == nil {
		return
	}
	enteredSceneInfo := &proto.EnteredScene{
		FloorId: mapEntrance.FloorID,
		PlaneId: mapEntrance.PlaneID,
	}
	rsp.EnteredSceneList = []*proto.EnteredScene{enteredSceneInfo}

	g.Send(cmd.GetEnteredSceneScRsp, rsp)
}

// 客户端登录需要的包，不是传送的通知包
func (g *GamePlayer) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	pos := g.GetPosPb()
	rot := g.GetRotPb()
	dbScene := g.GetScene()
	curLine := g.GetCurLineUp()

	rsp := new(proto.GetCurSceneInfoScRsp)
	rsp.Scene = g.GetSceneInfo(dbScene.EntryId, pos, rot, curLine)
	g.Send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *GamePlayer) HanldeGetSceneMapInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetSceneMapInfoCsReq, payloadMsg)
	req := msg.(*proto.GetSceneMapInfoCsReq)

	// 1000101 1000001

	rsp := new(proto.GetSceneMapInfoScRsp)
	for _, entryId := range req.EntryIdList {
		mapEntrance := gdconf.GetMapEntranceById(entryId)
		if mapEntrance != nil {
			teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
			if teleportsMap != nil {
				mapList := &proto.SceneMapInfo{
					LightenSectionList: make([]uint32, 0),
					ChestList: []*proto.ChestInfo{
						{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_NORMAL},
						{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
						{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
					},
					// UnlockedTeleportList: make([]uint32, 0),
				}

				mapList.EntryId = entryId

				for i := uint32(0); i < 100; i++ {
					mapList.LightenSectionList = append(mapList.LightenSectionList, i)
				}

				for _, teleports := range teleportsMap.TeleportsByGroupId {
					mazeGroup := &proto.MazeGroup{GroupId: teleports.GroupId}
					mapList.MazeGroupList = append(mapList.MazeGroupList, mazeGroup)
				}

				for _, teleports := range teleportsMap.Teleports {
					mazeProp := &proto.MazePropState{
						State:    gdconf.GetStateValue("CheckPointEnable"), // 默认解锁
						GroupId:  teleports.AnchorGroupID,
						ConfigId: teleports.ID,
					}
					mapList.MazePropList = append(mapList.MazePropList, mazeProp)
					mapList.UnlockTeleportList = append(mapList.UnlockTeleportList, teleports.MappingInfoID)
				}
				rsp.SceneMapInfo = append(rsp.SceneMapInfo, mapList)
			}
		}
	}

	g.Send(cmd.GetSceneMapInfoScRsp, rsp)
}

func (g *GamePlayer) EnterSceneCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterSceneCsReq, payloadMsg)
	req := msg.(*proto.EnterSceneCsReq)
	rsp := &proto.GetEnteredSceneScRsp{}

	g.EnterSceneByServerScNotify(req.EntryId, req.TeleportId)

	g.Send(cmd.EnterSceneScRsp, rsp)
	g.Send(cmd.SceneUpdatePositionVersionNotify, rsp)
}

func (g *GamePlayer) InteractPropCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.InteractPropCsReq, payloadMsg)
	req := msg.(*proto.InteractPropCsReq)
	rsp := &proto.InteractPropScRsp{
		Retcode:      0,
		PropState:    0,
		PropEntityId: req.PropEntityId,
	}
	var propEntityIdList []uint32

	pe := g.GetPropEntityById(req.PropEntityId)
	if pe == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	blockBin := g.GetBlock(pe.EntryId)
	mapEntrance := gdconf.GetMapEntranceById(pe.EntryId)
	if mapEntrance == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	confInteract := gdconf.GetInteractConfigById(req.InteractId)
	if confInteract == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	mazeProp := gdconf.GetMazePropId(pe.PropId)
	if mazeProp == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	switch mazeProp.PropType {
	case gdconf.PROP_ORDINARY:
		confProp := gdconf.GetServerPropById(mapEntrance.PlaneID, mapEntrance.FloorID, pe.GroupId, pe.InstId)
		if confProp == nil {
			return
		}
		// 处理交互的物体
		if confProp.GoppValue != nil {
			for _, goppValue := range confProp.GoppValue {
				g.UpPropState(blockBin, goppValue.GroupId, goppValue.InstId, 1) // 更新地图
				enep := g.GetPropEntity(goppValue.GroupId, goppValue.InstId)
				if enep == nil {
					continue
				}
				propEntityIdList = append(propEntityIdList, enep.EntityId)
			}
		}
	}
	// 更新本体
	propEntityIdList = append(propEntityIdList, req.PropEntityId)
	rsp.PropState = gdconf.GetStateValue(confInteract.TargetState) // 获取新状态
	g.UpPropState(blockBin, pe.GroupId, pe.InstId, rsp.PropState)  // 更新地图
	// 统一通知
	g.PropSceneGroupRefreshScNotify(propEntityIdList, blockBin) // 通知状态更改
	g.UpInteractSubMission(pe, rsp.PropState)                   // 检查交互任务
	g.UpdateBlock(blockBin)                                     // 保存地图
	g.Send(cmd.InteractPropScRsp, rsp)
}

func (g *GamePlayer) PropSceneGroupRefreshScNotify(propEntityIdList []uint32, db *spb.BlockBin) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	for _, propId := range propEntityIdList {
		pe := g.GetPropEntityById(propId)
		if pe == nil {
			continue
		}
		isAddend := true
		info := new(proto.GroupRefreshInfo)
		for _, ninfo := range notify.GroupRefreshList {
			if ninfo.GroupId == pe.GroupId {
				info = ninfo
				isAddend = false
				break
			}
		}
		info.RefreshEntity = append(info.RefreshEntity, &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				EntityId: pe.EntityId,
				GroupId:  pe.GroupId,
				InstId:   pe.InstId,
				Motion: &proto.MotionInfo{
					Pos: pe.Pos,
					Rot: pe.Rot,
				},
				Prop: &proto.ScenePropInfo{
					PropId:    pe.PropId, // PropID
					PropState: g.GetPropState(db, pe.GroupId, pe.InstId, ""),
				},
			},
		})
		if isAddend {
			info.GroupId = pe.GroupId
			notify.GroupRefreshList = append(notify.GroupRefreshList, info)
		}
	}

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *GamePlayer) AddAvatarSceneGroupRefreshScNotify(avatarId uint32, isTrial bool, pos, rot *proto.Vector) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	info := new(proto.GroupRefreshInfo)
	actor := &proto.SceneActorInfo{
		AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
		BaseAvatarId: avatarId,
		MapLayer:     0,
		Uid:          0,
	}
	entityId := g.GetNextGameObjectGuid()
	if isTrial {
		conf := gdconf.GetSpecialAvatarById(avatarId)
		if conf == nil {
			return
		}
		actor = &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_TRIAL_TYPE,
			BaseAvatarId: conf.AvatarID,
			MapLayer:     0,
			Uid:          0,
		}
	}
	info.RefreshEntity = append(info.RefreshEntity, &proto.SceneEntityRefreshInfo{
		AddEntity: &proto.SceneEntityInfo{
			EntityId: entityId,
			Actor:    actor,
			Motion: &proto.MotionInfo{
				Rot: pos,
				Pos: rot,
			},
		},
	})
	g.AddEntity(&AvatarEntity{
		Entity: Entity{
			EntityId: entityId,
			GroupId:  0,
			Pos:      pos,
			Rot:      rot,
		},
		AvatarId: actor.BaseAvatarId,
	})
	notify.GroupRefreshList = append(notify.GroupRefreshList, info)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}
