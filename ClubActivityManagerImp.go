package main

import (
    "errors"
    // "strconv"
    "time"

    "LifeService"
)

//ClubActivityManagerImp implement
type ClubActivityManagerImp struct {
    dataServiceProxy       *LifeService.DataService
    dataServiceObj         string
    userInfoServiceProxy   *LifeService.UserInfoService
    UserInfoServiceObj     string
}

//init 初始化
func (imp *ClubActivityManagerImp) init() {
    imp.dataServiceProxy     = new(LifeService.DataService)
    imp.dataServiceObj       = "LifeService.DataServer.DataServiceObj"
    imp.userInfoServiceProxy = new(LifeService.UserInfoService)
    imp.UserInfoServiceObj   = "LifeService.UserInfoServer.UserInfoServiceObj"

    comm.StringToProxy(imp.dataServiceObj, imp.dataServiceProxy)
    comm.StringToProxy(imp.UserInfoServiceObj, imp.userInfoServiceProxy)
}

//CreateClubManager 创建社团管理员
func (imp *ClubActivityManagerImp) CreateClubManager(wxId string, clubId string, RetCode *int32) (int32, error) {
    var TableName = "club_managers"
    var Columns   = []LifeService.Column {
        LifeService.Column{ColumnName: "wx_id"  , DBInt: false, ColumnValue: wxId},
        LifeService.Column{ColumnName: "club_id", DBInt: true , ColumnValue: clubId},
    }
    _, err := imp.dataServiceProxy.InsertData(TableName, Columns)
    if err != nil {
        SLOG.Error("Call Remote DataServer::InsertData error: ", err)
        *RetCode = 400
        return -1, nil
    }
    *RetCode = 300
    SLOG.Debug("CreateClubManager successfully")
    return 0, nil
}

//CreateClub 创建社团
func (imp *ClubActivityManagerImp) CreateClub(ClubInfo *LifeService.ClubInfo, RetCode *int32) (int32, error) {
    var iRet int32
    
    CurrentTime := time.Now().Format("2006-01-02 15:04:05")
    ClubInfo.Create_time = CurrentTime

    _, err := imp.dataServiceProxy.CreateClub(ClubInfo, &iRet)

    if err != nil {
        SLOG.Error("Create Club Error")
        *RetCode = 400
    } else {
        *RetCode = 200
    }
    return 0, err
}

//GetClubList 获取社团列表
func (imp *ClubActivityManagerImp) GetClubList(clubInfoList *[]LifeService.ClubInfo, RetCode *int32) (int32, error) {
    iRet, err := imp.dataServiceProxy.GetClubList(clubInfoList)
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
func (imp *ClubActivityManagerImp) ApplyForClub(wx_id string, clubId string, RetCode *int32) (int32, error) {
    var isInClub bool
    _, err := imp.userInfoServiceProxy.IsInClub(wx_id, clubId, false, &isInClub)
    
    if err != nil {
        SLOG.Error("Remote Server UserInfoServer::IsInClub error")
        *RetCode = 500
        return -1, err
    }
    
    if isInClub {
        *RetCode = 300
        SLOG.Debug("Applied")
    } else {
        var TableName = "apply_for_club"
        var Columns   = []LifeService.Column {
            LifeService.Column{ColumnName: "apply_status", DBInt: true , ColumnValue: "0"},
            LifeService.Column{ColumnName: "user_id"     , DBInt: false, ColumnValue: wx_id},
            LifeService.Column{ColumnName: "club_id"     , DBInt: true , ColumnValue: clubId},
        }
        _, err1 := imp.dataServiceProxy.InsertData(TableName, Columns)
        if err1 != nil {
            SLOG.Error("Remote Server DataServer::InsertData error", err1)
            *RetCode = 400
            return -1, err1
        }

        SLOG.Debug("Apply Successfully")
        *RetCode = 200
            
    }
    return 0, nil
}

//CreateActivity 创建活动
func (imp *ClubActivityManagerImp) CreateActivity(wx_id string, activityInfo *LifeService.ActivityInfo, RetCode *int32) (int32, error) {
    var isClubManager bool
    _, err := imp.userInfoServiceProxy.IsClubManager(wx_id, activityInfo.Club_id, &isClubManager)
    if err != nil {
        SLOG.Error("Remote Server UserInfoServer::IsClubManager error")
        *RetCode = 500
        return -1, err
    }

    if isClubManager {
        // TODO: 创建活动逻辑
        var TableName = "activities"
        var Columns   = []LifeService.Column {
            LifeService.Column{ColumnName: "name"            	, DBInt: false , ColumnValue: activityInfo.Name},
            LifeService.Column{ColumnName: "sponsor"			, DBInt: false , ColumnValue: activityInfo.Sponsor},
            LifeService.Column{ColumnName: "club_id"			, DBInt: true  , ColumnValue: activityInfo.Club_id},
            LifeService.Column{ColumnName: "start_time"         , DBInt: false , ColumnValue: activityInfo.Start_time},
            LifeService.Column{ColumnName: "stop_time"          , DBInt: false , ColumnValue: activityInfo.Stop_time},
            LifeService.Column{ColumnName: "registry_start_time", DBInt: false , ColumnValue: activityInfo.Registry_start_time},
            LifeService.Column{ColumnName: "registry_stop_time" , DBInt: false , ColumnValue: activityInfo.Registry_stop_time},
            LifeService.Column{ColumnName: "content"            , DBInt: false , ColumnValue: activityInfo.Content},
        }
        _,err1 := imp.dataServiceProxy.InsertData(TableName, Columns)
        if err1 != nil {
            SLOG.Error("Remote server DataServer::InsertData error ", err1)
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
func (imp *ClubActivityManagerImp) GetActivityList(index int32, nextIndex *int32, activityList *[]map[string]string) (int32, error) {
    // var batch     = 15
    // var TableName = "activities"
    // var Columns   = []string {"activity_id", "name", "create_time"}
    // var Condition = "`activity_id`"

    // if index == 0 {
    //     Condition += ">0"
    // } else {
    //     Condition += "<" + strconv.Itoa(int(index))
    // }
    // Condition += " and `stop_time`>'" + time.Now().Format("2006-01-02 15:04:05") + "' order by `create_time` limit " + strconv.Itoa(batch)
    _,err := imp.dataServiceProxy.GetActivityList(index, nextIndex, activityList)
    // _, err := imp.dataServiceProxy.QueryData(TableName, Columns, Condition, activityList)
    if err != nil {
        SLOG.Error("Call Remote DataServer::GetActivityList error: ", err)
        return -1, err
    }

    // var lengthL = len(*activityList)

    // tempIndex, _ := strconv.Atoi((*activityList)[lengthL - 1][Columns[0]])
    // *nextIndex = int32(tempIndex)
    return 0, nil
}

//GetActivityDetail 获取活动详情
func (imp *ClubActivityManagerImp) GetActivityDetail(activityId string, activityInfo *LifeService.ActivityInfo) (int32, error) {
    var TableName = "activities"
    var Columns   = []string {"name", "sponsor", "club_id", "target_id", "create_time", "start_time", "stop_time", "registry_start_time", "registry_stop_time", "content"}
    var Condition = "`activity_id`=" + activityId
    var Result []map[string]string

    _, err := imp.dataServiceProxy.QueryData(TableName, Columns, Condition, &Result)
    if err != nil {
        SLOG.Error("Call Remote DataServer::QueryData error: ", err)
        return -1, err
    }
    activityInfo.Activity_id         = activityId
    activityInfo.Name 		         = Result[0][Columns[1]]
    activityInfo.Sponsor             = Result[0][Columns[2]]
    activityInfo.Club_id             = Result[0][Columns[3]]
    activityInfo.Target_id           = Result[0][Columns[4]]
    activityInfo.Create_time         = Result[0][Columns[5]]
    activityInfo.Start_time          = Result[0][Columns[6]]
    activityInfo.Stop_time           = Result[0][Columns[7]]
    activityInfo.Registry_start_time = Result[0][Columns[8]]
    activityInfo.Registry_stop_time  = Result[0][Columns[9]]
    activityInfo.Content             = Result[0][Columns[10]]
    return 0, nil
}

//ApplyForActivity 活动报名
func (imp *ClubActivityManagerImp) ApplyForActivity(wx_id string, activityId string, RetCode *int32) (int32, error) {
    var isApplied bool
    
    _, err := imp.userInfoServiceProxy.IsAppliedActivity(wx_id, activityId, &isApplied)
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

    var TableName = "activity_records"
    var Columns   = []LifeService.Column {
        LifeService.Column{ColumnName: "user_id"    , DBInt: false, ColumnValue: wx_id},
        LifeService.Column{ColumnName: "activity_id", DBInt: true , ColumnValue: activityId},
    }
    
    _, err1 := imp.dataServiceProxy.InsertData(TableName, Columns)
    if err1 != nil {
        SLOG.Error("Remote Server DataServer::InsertData error")
        *RetCode = 400
        return -1, err1
    }

    SLOG.Debug("Apply Activity successfully")
    *RetCode = 200

    return 0, nil
}