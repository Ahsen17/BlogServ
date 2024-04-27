package models

const (
	SystemLevel = "system"
	CustomLevel = "custom"
)

const (
	SystemCreate   = "SC"
	SystemUpdate   = "SU"
	SystemRetrieve = "SR"
	SystemDelete   = "SD"
)

const (
	CustomCreate   = "CC"
	CustomUpdate   = "CU"
	CustomRetrieve = "CR"
	CustomDelete   = "CD"

	CustomLogin  = "login"
	CustomLogout = "logout"
	CustomRevoke = "revoke"
)

type BaseLog struct {
	logTime  int64
	logLevel string
	logType  string
	logMsg   string
}

type SystemLog struct {
	BaseLog BaseLog
}

type SystemLogMgr struct {
	SystemLog *SystemLog
}

type CustomLog struct {
	BaseLog BaseLog
}

type CustomLogMgr struct {
	CustomLog *CustomLog
}
