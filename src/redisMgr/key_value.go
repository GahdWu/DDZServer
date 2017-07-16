package redisMgr

import (
	"fmt"
)

const (
	// 玩家前缀
	con_PlayerPrefix = "Player"

	// 玩家Key的分隔符
	con_PlayerKeyDelimiter = "_"
)

//******************************key**************************
const (
	// SocketServerCenter地址
	con_SocketServerCenterAddress_Key string = "SocketServerCenterAddress"

	// 存放超时日志的List
	con_ExpireLog_Key string = "ExpireLog"

	// 存放ServerGroupId和APIUrl的HashTable
	con_ServerGroupInfo_Key string = "ServerGroupInfo"

	//发送到gs失败的玩家数据
	con_SendFailedBattleData_key string = "SendFailedBattleData"
)

//******************************模块公有的key**************************
const (
	// 玩家名称
	con_Name_field string = "Name"

	// 玩家等级
	con_Lv_field string = "Lv"

	// VIP等级
	con_Vip_field string = "Vip"

	// 玩家战力
	con_FAP_field string = "FAP"

	// 玩家历史最高战力
	con_MaxFap_field string = "MaxFap"

	// 合作商Id
	con_PartnerId_field string = "PartnerId"

	// 服务器Id
	con_ServerId_field string = "ServerId"

	// 玩家所在服务器名称
	con_Zone_field string = "Zone"

	// 头像Id
	con_HeadImageId_field string = "HeadImageId"

	// 时装Id
	con_FashionModelId_field string = "FashionModelId"

	// 头像框
	con_PVPInterLv_field string = "PVPInterLv"

	// 公会名称
	con_GuildName_field string = "GuildName"

	// 玩家阵容数据 用于展示
	con_Formation_field string = "Formation"

	// 英雄头像展示
	con_BaseFormation_field string = "BaseFormation"

	// 玩家阵容数据 用于战斗
	con_FightInfo_field string = "FightInfo"

	// 出战宠物数据
	con_CombatPet_field string = "CombatPet"

	// 法宝模型Id
	con_MountModelId_field string = "MountModelId"

	// 法宝模型level
	con_MountModelLevel_field string = "MountLv"

	//玩家所在服务器组名称
	con_ServerGroupId_field string = "ServerGroupId"

	//posId与slotId对应关系
	con_SlotFormationInfo_field string = "SlotFormationInfo"

	// 玩家状态
	con_Status_field string = "Status"

	//socketServer地址
	con_SocketServerAddress_field string = "SocketServerAddress"

	//SocketServerIP
	con_SocketServerIP_field string = "SocketServerIP"

	//玩家当前模块
	con_CurModuleType_field string = "CurModuleType"
)

//--------------------神域------------------------
const (
	// 复活结束时间
	con_Goddomain_ReliveTime_field string = "ReliveTime_godDomain"

	//单场玩家积分
	con_Goddomain_FightScore_field string = "FightScore_godDomain"

	//赛季玩家积分
	con_Goddomain_SeasonFightScore_field string = "SeasonFightScore_godDomain"

	//单场战斗积分获得时间
	con_Goddomain_FightScoreTime_field string = "FightScoreTime_godDomain"

	//战斗开始时间
	con_Goddomain_BattleBeginTime_field string = "StartTime_godDomain"

	//战斗结束时间
	con_Goddomain_BattleEndTime_field string = "EndTime_godDomain"

	//单场击杀玩家数
	con_Goddomain_KillNum_field string = "KillNum_godDomain"

	//赛季击杀人数
	con_Goddomain_SeasonKillNum_field string = "SeasonKillNum_godDomain"

	//单场死亡次数
	con_Goddomain_DeadNum_field string = "DeadNum_godDomain"

	//赛季死亡次数
	con_Goddomain_SeasonDeadNum_field string = "SeasonDeadNum_godDomain"

	//赛季胜利次数
	con_Goddomain_SeasonWinNum_field string = "SeasonWinNum_godDomain"

	//赛季失败次数
	con_Goddomain_SeasonDefeatNum_field string = "SeasonDefeatNum_godDomain"

	//单场占领资源数
	con_Goddomain_OccupyResNum_field string = "ControlPointNum_godDomain"

	//单场荣誉值
	con_Goddomain_Glory_field string = "Glory_godDomain"

	//本队获取的资源点总数
	con_Goddomain_BattleTotalResourceNum_field string = "BattleTotalResourceNum_godDomain"

	//队伍战斗结果
	con_Goddomain_GameResult_field string = "GameResult_godDomain"

	//玩家是否挂机
	con_Goddomain_IsHangUp_field string = "IsHangUp_godDomain"

	//今日战斗次数
	con_Goddomain_FightNumDaily_field string = "FightNumDaily_godDomain"

	//刷新时间
	con_Goddomain_FreshTime_field string = "FreshTime_godDomain"

	//复活次数
	con_Goddomain_ReliveNum_field string = "ReliveNum_godDomain"

	// 飞机技能
	con_Goddomain_SkillIds_field string = "SkillIds_godDomain"
)

//--------------------------------------圣渊--------------------------------
const (
	// 复活结束时间
	con_ShengYuan_ReliveTime_field string = "ReliveTime_shengYuan"

	//赛季战斗积分
	con_ShengYuan_SeasonFightScore_field string = "SeasonFightScore_shengYuan"

	//赛季杀人数
	con_ShengYuan_SeasonKillNum_field string = "SeasonKillNum_shengYuan"

	//赛季死亡次数
	con_ShengYuan_SeasonDeadNum_field string = "SeasonDeadNum_shengYuan"

	//赛季胜利次数
	con_ShengYuan_SeasonWinNum_field string = "SeasonWinNum_shengYuan"

	//赛季失败次数
	con_ShengYuan_SeasonDefeatNum_field string = "SeasonDefeatNum_shengYuan"
)

//--------------------------------------公平对决--------------------------------
const (
	// 玩家阵容数据 用于展示
	con_Fairpvp_fairpvpFormation_field string = "FairpvpFormation"

	// 玩家阵容数据 用于战斗
	con_Fairpvp_fairpvpFightInfo_field string = "FairpvpFightInfo"

	// 赛季战斗积分
	con_Fairpvp_fairpvpSeasonScore_field string = "FairpvpSeasonScore"

	// 赛季挑战总场次
	con_Fairpvp_fairpvpSeasonFightNum_field string = "FairpvpSeasonFightNum"

	// 赛季赢的总场次
	con_Fairpvp_fairpvpSeasonWinNum_field string = "FairpvpSeasonWinNum"

	// 赛季排行榜
	con_Fairpvp_fairpvpHourRank_field string = "FairpvpHourRank"
)

//******************************value**************************
var (
	// SocketServerCenter地址
	socketServerCenterAddress string
)

// 获取玩家的Key
// playerId：玩家Id
// 返回值：
// 玩家的Key
func getPlayerKey(playerId string) string {
	return fmt.Sprintf("%s%s%s", con_PlayerPrefix, con_PlayerKeyDelimiter, playerId)
}

// 获取所有玩家Key的Pattern
// 返回值：
// 所有玩家Key的Pattern
func getAllPlayerKeysPattern() string {
	return fmt.Sprintf("%s%s*", con_PlayerPrefix, con_PlayerKeyDelimiter)
}

// 获取SocketServerCenter地址
// 返回值：
// SocketServerCenter地址
func GetSocketServerCenterAddress() string {
	return socketServerCenterAddress
}
