package main

import (
	"errors"
	"time"

	"LifeService"
)

//ClubActivityManagerImp implement
type ClubActivityManagerImp struct {
	dataServiceProxy     *LifeService.DataService
	dataServiceObj       string
	userInfoServiceProxy *LifeService.UserInfoService
	UserInfoServiceObj   string
}

//init 初始化
func (imp *ClubActivityManagerImp) init() {
	imp.dataServiceProxy = new(LifeService.DataService)
	imp.dataServiceObj = "LifeService.DataServer.DataServiceObj"
	imp.userInfoServiceProxy = new(LifeService.UserInfoService)
	imp.UserInfoServiceObj = "LifeService.UserInfoServer.UserInfoServiceObj"

	comm.StringToProxy(imp.dataServiceObj, imp.dataServiceProxy)
	comm.StringToProxy(imp.UserInfoServiceObj, imp.userInfoServiceProxy)
}

//CreateClubManager 创建社团管理员
func (imp *ClubActivityManagerImp) CreateClubManager(wxID string, clubID string, RetCode *int32) (int32, error) {
	ret, err := imp.dataServiceProxy.CreateClubManager(wxID, clubID)
	if err != nil {
		SLOG.Error("Call Remote DataServer::CreateClubManager error: ", err.Error())
		*RetCode = 400
		return -1, err
	}
	*RetCode = ret
	if ret == 300 {
		SLOG.Error("Remote DataServer::CreateClubManager SQL error")
		return 0, nil
	}

	SLOG.Debug("CreateClubManager successfully")
	return 0, nil
}

//CreateClub 创建社团
func (imp *ClubActivityManagerImp) CreateClub(ClubInfo *LifeService.ClubInfo, RetCode *int32) (int32, error) {
	var clubID string
	// 获取当前时间
	CurrentTime := time.Now().Format("2006-01-02 15:04:05")
	ClubInfo.Create_time = CurrentTime
	// 创建社团
	_, err := imp.dataServiceProxy.CreateClub(ClubInfo, &clubID)

	if err != nil {
		SLOG.Error("Create Club Error")
		*RetCode = 400
		return -1, err
	}
	// 返回的社团id为空, 说明DB层服务SQL执行错误
	if clubID == "" {
		SLOG.Error("Create Club Error")
		*RetCode = 300
		return -1, errors.New("SQL Error")
	}
	// 创建社团管理员
	ret, err1 := imp.dataServiceProxy.CreateClubManager(ClubInfo.Chairman, clubID)
	if err1 != nil {
		SLOG.Error("Call Remote DataServer::CreateClubManager error: ", err.Error())
		*RetCode = 400
		return -1, err1
	}
	*RetCode = ret
	if ret == 300 {
		SLOG.Error("Remote DataServer::CreateClubManager SQL error")
		return 0, nil
	}

	SLOG.Debug("CreateClubManager successfully")
	return 0, nil
}

//GetClubList 获取社团列表
func (imp *ClubActivityManagerImp) GetClubList(index int32, wxID string, nextIndex *int32, clubInfoList *[]LifeService.ClubInfo, RetCode *int32) (int32, error) {
	var batch int32 = 6
	iRet, err := imp.dataServiceProxy.GetClubList(index, batch, wxID, nextIndex, clubInfoList)
	if err != nil {
		SLOG.Error("Get club list error")
		*RetCode = 400
	} else {
		if iRet == 0 {
			SLOG.Debug("GetClubList successfully")
			*RetCode = 200
		} else {
			SLOG.Debug("Cannot get Club list")
			*RetCode = 301
		}
	}
	return 0, nil
}

//GetManagerClubList 获取管理社团列表
func (imp *ClubActivityManagerImp) GetManagerClubList(index int32, wxID string, nextIndex *int32, clubInfoList *[]LifeService.ClubInfo, RetCode *int32) (int32, error) {
	var batch int32 = 6

	iRet, err := imp.dataServiceProxy.GetManagerClubList(index, batch, wxID, nextIndex, clubInfoList)
	if err != nil {
		SLOG.Error("Get club list error")
		*RetCode = 400
	} else {
		if iRet == 0 {
			SLOG.Debug("GetClubList successfully")
			*RetCode = 200
		} else {
			SLOG.Debug("Cannot get Club list")
			*RetCode = 301
		}
	}
	return 0, nil
}

//ApplyForClub 申请社团
func (imp *ClubActivityManagerImp) ApplyForClub(wxID string, clubID string, RetCode *int32) (int32, error) {
	var isInClub bool
	_, err := imp.userInfoServiceProxy.IsInClub(wxID, clubID, false, &isInClub)

	if err != nil {
		SLOG.Error("Remote Server UserInfoServer::IsInClub error")
		*RetCode = 500
		return -1, err
	}

	if isInClub {
		*RetCode = 300
		SLOG.Debug("Applied")
		return 0, nil
	}

	_, err1 := imp.dataServiceProxy.CreateApply(wxID, clubID)
	if err1 != nil {
		SLOG.Error("Call Remote DataServer::CreateApply error", err1)
		*RetCode = 400
		return -1, err1
	}
	SLOG.Debug("Apply Successfully")
	return 0, nil
}

//GetClubApply 获取社团申请列表
func (imp *ClubActivityManagerImp) GetClubApply(clubID string, index int32, applyStatus int32, nextIndex *int32, applyList *[]LifeService.ApplyInfo) (int32, error) {
	var batch int32 = 6
	iRet, err := imp.dataServiceProxy.GetApplyListByClubId(clubID, index, batch, applyStatus, nextIndex, applyList)
	if err != nil {
		SLOG.Error("Remote Server DataServer::GetApplyListByClubId error: ", err.Error())
		return -1, err
	}
	return iRet, err
}

//GetUserApply 获取用户的申请
func (imp *ClubActivityManagerImp) GetUserApply(wxID string, index int32, applyStatus int32, nextIndex *int32, applyList *[]LifeService.ApplyInfo) (int32, error) {
	var batch int32 = 6
	iRet, err := imp.dataServiceProxy.GetApplyListByUserId(wxID, index, batch, applyStatus, nextIndex, applyList)
	if err != nil {
		SLOG.Error("Remote Server DataServer::GetApplyListByUserId error: ", err.Error())
		return -1, err
	}
	return iRet, err
}

//ModifyApplyStatus 设置申请状态
func (imp *ClubActivityManagerImp) ModifyApplyStatus(wxID string, clubID string, applyStatus int32, RetCode *int32) (int32, error) {
	var ret int32
	iRet, err := imp.dataServiceProxy.SetApplyStatus(wxID, clubID, applyStatus, &ret)
	if err != nil || ret != 0 {
		SLOG.Error("Remote Server DataServer::setApplyStatus error: ", err.Error())
		return -1, err
	}
	SLOG.Debug("ModifuApplyStatus")
	*RetCode = 200
	return iRet, err
}

//DeleteApply 删除申请
func (imp *ClubActivityManagerImp) DeleteApply(wxID string, clubID string, RetCode *int32) (int32, error) {
	iRet, err := imp.dataServiceProxy.DeleteApply(wxID, clubID, RetCode)
	if err != nil {
		SLOG.Error("Remote Server DataServer::DeleteApply error: ", err.Error())
		return -1, err
	}
	SLOG.Debug("DeleteApply")
	return iRet, err
}

//CreateActivity 创建活动
func (imp *ClubActivityManagerImp) CreateActivity(wxID string, activityInfo *LifeService.ActivityInfo, RetCode *int32) (int32, error) {
	var isClubManager bool
	// 判断是否为社团管理员
	_, err := imp.userInfoServiceProxy.IsClubManager(wxID, activityInfo.Club_id, &isClubManager)

	if err != nil {
		SLOG.Error("Remote Server UserInfoServer::IsClubManager error")
		*RetCode = 500
		return -1, err
	}

	if isClubManager {
		// 创建社团
		activityInfo.Target_id = "0"
		_, err1 := imp.dataServiceProxy.CreateActivity(activityInfo)

		if err1 != nil {
			SLOG.Error("Cal Remote DataServer::CreateActivity error ", err1.Error())
			*RetCode = 400
			return -1, err1
		}
		SLOG.Debug("Create Activity successfully")
		*RetCode = 200
	} else {
		*RetCode = 400
		return -1, errors.New("Not Manager")
	}
	return 0, nil
}

//GetActivityList 获取活动列表
func (imp *ClubActivityManagerImp) GetActivityList(index int32, wxID string, clubID string, nextIndex *int32, activityList *[]map[string]string) (int32, error) {
	var batch int32 = 6

	_, err := imp.dataServiceProxy.GetActivityList(index, batch, wxID, clubID, nextIndex, activityList)

	if err != nil {
		SLOG.Error("Call Remote DataServer::GetActivityList error: ", err.Error())
		return -1, err
	}
	return 0, nil
}

//UpdateActivity 更新活动信息
func (imp *ClubActivityManagerImp) UpdateActivity(activityInfo *LifeService.ActivityInfo, RetCode *int32) (int32, error) {
	var iRet int32
	_, err := imp.dataServiceProxy.UpdateActivity(activityInfo, &iRet)

	if err != nil {
		SLOG.Error("Call Remote DataServer::UpdateActivity error: ", err.Error())
		*RetCode = 400
		return -1, err
	}

	if iRet != 0 {
		SLOG.Error("Remote DateServer::UpdateActivity Execute SQL error")
		*RetCode = 500
		return -1, errors.New("Remote DateServer::UpdateActivity Execute SQL error")
	}
	*RetCode = 200
	return 0, nil
}

//DeleteActivity 删除活动
func (imp *ClubActivityManagerImp) DeleteActivity(activityID string, RetCode *int32) (int32, error) {
	var ret int32
	_, err := imp.dataServiceProxy.DeleteActivity(activityID, &ret)
	if err != nil {
		SLOG.Error("Remote Server DataServer::deleteActivity error: ", err.Error())
		return -1, err
	}
	*RetCode = 200
	SLOG.Debug("DeleteActivity")
	return ret, err
}

//GetActivityDetail 获取活动详情
func (imp *ClubActivityManagerImp) GetActivityDetail(activityID string, activityInfo *LifeService.ActivityInfo) (int32, error) {

	_, err := imp.dataServiceProxy.GetActivityInfo(activityID, activityInfo)

	if err != nil {
		SLOG.Error("Call Remote DataServer::createActivity error: ", err.Error())
		return -1, err
	}

	return 0, nil
}

//GetActivityParticipate 获取活动参与者信息
func (imp *ClubActivityManagerImp) GetActivityParticipate(index int32, activityID string, nextIndex *int32, participateList *[]LifeService.ActivityRecord) (int32, error) {
	var batch int32 = 6
	iRet, err := imp.dataServiceProxy.GetActivityRecords(index, batch, activityID, nextIndex, participateList)

	if err != nil || iRet != 0 {
		SLOG.Error("Call Remote DataServer::getActivityRecords error")
		return -1, nil
	}
	return 0, nil
}

//ApplyForActivity 活动报名
func (imp *ClubActivityManagerImp) ApplyForActivity(WxID string, activityID string, RetCode *int32) (int32, error) {
	var isApplied bool
	// 是否已经报名活动
	_, err := imp.userInfoServiceProxy.IsAppliedActivity(WxID, activityID, &isApplied)
	if err != nil {
		SLOG.Error("Remote Server UserInfoServer::IsApplied error")
		*RetCode = 500
		return -1, err
	}

	if isApplied {
		SLOG.Debug("Applied")
		*RetCode = 300
		return 0, nil
	}

	_, err1 := imp.dataServiceProxy.CreateActivityRecord(WxID, activityID)
	if err1 != nil {
		SLOG.Error("Call Remote DataServer::createActivityRecord error", err1.Error())
		*RetCode = 400
		return -1, err1
	}

	SLOG.Debug("Apply Activity successfully")
	*RetCode = 200

	return 0, nil
}

//DeleteActivityParticipate 删除参与者信息
func (imp *ClubActivityManagerImp) DeleteActivityParticipate(activityID string, wxID string, RetCode *int32) (int32, error) {
	var iRet int32
	_, err := imp.dataServiceProxy.DeleteActivityRecord(activityID, wxID, &iRet)

	if err != nil {
		SLOG.Error("Call Remote DataServer::deleteActivityRecord error: " + err.Error())
		*RetCode = 400
		return -1, err
	}
	*RetCode = 200
	return 0, nil
}
