package global

import (
	"go.uber.org/zap"
	"server/config"
)

type GVA struct {
	GvaConfigs config.Configs
	GvaLogger  *zap.SugaredLogger
}

var Gva GVA

//var gvaInstance *GVA = &GVA{}
//
//func GetGva() *GVA {
//	return gvaInstance
//}
