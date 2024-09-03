package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint // Primary Key
	Username   string
	Password   string
	Email      string
	FirstName  string
	LastName   string
	Roles      []Role
	gorm.Model // Needed for Atlas to know this exists
}

type Org struct {
	ID      uint // Primary Key
	Name    string
	Roles   []Role
	Systems []System
	gorm.Model
}

type Role struct {
	ID          uint // Primary Key
	OrgId       uint
	UserId      uint
	Permissions []Permission
	gorm.Model
}

type Permission struct {
	ID            uint // Primary Key
	RoleId        uint
	ThingYouCanDo string
	gorm.Model
}

type System struct {
	ID       uint // Primary Key
	Name     string
	OrgId    uint
	Services []Service
	gorm.Model
}

type Service struct {
	ID           uint // Primary Key
	Name         string
	SystemId     uint
	EnvSnapshots []EnvSnapshot
	gorm.Model
}

type EnvSnapshot struct {
	ID         uint  // Primary Key
	EnvEventID *uint // LATEST STATE
	ServiceId  uint
	Version    *uint
	EnvEvents  []EnvEvent
	CRONS      []CRON
	Streams    []Stream
	PushApis   []PushApi
	Analytics  []Analytic
	Loggers    []Logger
	Policies   []Policy
	gorm.Model
}

type EnvEvent struct {
	ID            uint // Primary Key
	Name          string
	EnvSnapshotId uint
	TimeStamp     time.Time
	UserId        *uint
	ValueType     string
	Value         string
	gorm.Model
}

type CRON struct {
	ID            uint // Primary Key
	Value         string
	EnvSnapshotId uint
	gorm.Model
}

type Stream struct {
	ID            uint // Primary Key
	Value         string
	EnvSnapshotId uint
	gorm.Model
}

type PushApi struct {
	ID            uint // Primary Key
	Value         string
	EnvSnapshotId uint
	gorm.Model
}

type Analytic struct {
	ID            uint // Primary Key
	Value         string
	EnvSnapshotId uint
	gorm.Model
}

type Logger struct {
	ID            uint // Primary Key
	Value         string
	EnvSnapshotId uint
	gorm.Model
}

type Policy struct {
	ID            uint // Primary Key
	Value         string
	EnvSnapshotId uint
	gorm.Model
}
