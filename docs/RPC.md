# 社团活动服务RPC接口
以下类型在tars协议文件中定义，参照[tars数据类型定义文档](#)
```
ClubInfo          社团信息
ActivityInfo      活动信息
ApplyInfo         社团申请信息
ActivityRecord    活动报名记录信息
ErrorCode         服务错误码类型
```

## 接口列表

- [CreateClubManager](#interface-CreateClubManager)
- [CreateClub](#interface-CreateClub)
- [GetClubList](#interface-GetClubList)
- [GetManagerClubList](#interface-GetManagerClubList)
* [ApplyForClub](#interface-ApplyForClub)
* [GetClubApply](#interface-GetClubApply)
* [GetUserApply](#interface-GetUserApply)
* [ModifyApplyStatus](#interface-ModifyApplyStatus)
* [DeleteApply](#interface-DeleteApply)
- [CreateActivity](#interface-CreateActivity)
- [GetActivityList](#interface-GetActivityList)
- [UpdateActivity](#interface-UpdateActivity)
- [DeleteActivity](#interface-DeleteActivity)
- [GetActivityDetail](#interface-GetActivityDetail)
* [GetActivityParticipate](#interface-GetActivityParticipate)
* [ApplyForActivity](#interface-ApplyForActivity)
* [DeleteActivityParticipate](#interface-DeleteActivityParticipate)

## 接口详情

### <a id="interface-createclubmanager"> CreateClubManager
创建社团管理员

**定义**

```cpp
int CreateClubManager(string wxId, string clubId, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|clubId|string|社团id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-createclub"> CreateClub
创建社团

**定义**

```cpp
int CreateClub(ClubInfo clubInfo, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|clubInfo|ClubInfo|社团信息|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|错误码|

### <a id="interface-getclublist"> GetClubList
获取社团列表

**定义**

```cpp
int GetClubList(int index, string wxId, out int nextIndex, out vector<ClubInfo> clubInfoList, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|index|int|索引|
|wxId|string|用户id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|nextIndex|int|下一页索引, -1表示没有数据|
|clubInfoList|vector<ClubInfo>|社团信息列表|

### <a id="interface-getmanagerclublist"> GetManagerClubList
获取社团管理员列表

**定义**

```cpp
int GetManagerClubList(int index, string wxId, out int nextIndex, out vector<ClubInfo> clubInfoList, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|index|int|索引|
|wxId|string|用户id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|nextIndex|int|下一页索引, -1表示没有数据|
|clubInfoList|vector<ClubInfo>|社团信息列表|

### <a id="interface-applyforclub"> ApplyForClub
申请社团

**定义**

```cpp
int ApplyForClub(string wxId, string clubId, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|clubId|string|社团id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-getclubapply"> GetClubApply
获取社团申请

**定义**

```cpp
int GetClubApply(string clubId, int index, int applyStatus, out int nextIndex, out vector<ApplyInfo> applyList)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|clubInfo|string|社团id|
|index|int|索引|
|applyStatus|int|申请状态|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|nextIndex|int|下一页索引, -1表示没有数据|
|applyList|vector<ApplyInfo>|申请信息列表|

### <a id="interface-getuserapply"> GetUserApply
获取用户申请

**定义**


```cpp
int GetUserApply(string wxId, int index, int applyStatus, out int nextIndex, out vector<ApplyInfo> applyList)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|index|int|索引|
|applyStatus|int|申请状态|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|nextIndex|int|下一页索引, -1表示没有下一页数据|
|applyList|vector<ApplyInfo>|申请信息列表|

### <a id="interface-modifyapplystatus"> ModifyApplyStatus
修改申请状态

**定义**

```cpp
int ModifyApplyStatus(string wxId, string clubId, int applyStatus, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|clubId|string|社团id|
|applyStatus|int|申请状态|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-deleteapply"> DeleteApply
删除申请

**定义**

```cpp
int DeleteApply(string wxId, string clubId, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|clubId|string|社团id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-createactivity"> CreateActivity
创建活动

**定义**

```cpp
int CreateActivity(string wxId, ActivityInfo activityInfo, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|actiityInfo|ActivityInfo|活动信息|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-getactivitylist"> GetActivityList
获取活动列表

**定义**

```cpp
int GetActivityList(int index, string wxId, string clubId, out int nextIndex, out vector<map<string, string>> activityList)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|index|int|索引|
|wxId|string|用户id|
|clubId|string|社团id, 为空字符串时(即""), 表示不对社团进行筛选|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|nextIndex|int|下一页索引, -1表示没有下一页数据|
|activityList|vector<map<string, string>>|活动信息列表|

### <a id="interface-updateactivity"> UpdateActivity
更新活动信息

**定义**

```cpp
int UpdateActivity(ActivityInfo activityInfo, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|activityInfo|ActivityInfo|活动信息|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-deleteactivity"> DeleteActivity
删除活动

**定义**

```cpp
int DeleteActivity(string activityId, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|activityId|string|活动id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-getactivitydetail"> GetActivityDetail
获取活动详情

**定义**

```cpp
int GetActivityDetail(string activityId, out ActivityInfo activityInfo)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|activityId|string|活动id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|activityInfo|ActivityInfo|活动信息|

### <a id="interface-getactivityparticipate"> GetActivityParticipate
获取活动参与者

**定义**

```cpp
int GetActivityParticipate(int index, string activityId, out int nextIndex, out vector<ActivityRecord> participateList)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|index|int|索引, 为0时表示获取第一页数据|
|activityId|string|活动id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|nextIndex|int|下一页索引, -1表示没有下一页数据|
|participateList|vector<ActivityRecord>|活动参与者信息列表|

### <a id="interface-applyforactivity"> ApplyForActivity
报名活动

**定义**

```cpp
int ApplyForActivity(string wxId, string activityId, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|activityId|string|活动id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|

### <a id="interface-deleteactivityparticipate"> DeleteActivityParticipate
删除活动参与者

**定义**

```cpp
int DeleteActivityParticipate(string activityId, string wxId, out ErrorCode ErrCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|activityId|string|活动id|
|wxId|string|用户id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|ErrCode|ErrorCode|服务错误码|
