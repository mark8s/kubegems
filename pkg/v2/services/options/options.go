package options

import (
	"kubegems.io/kubegems/pkg/utils/argo"
	"kubegems.io/kubegems/pkg/utils/database"
	"kubegems.io/kubegems/pkg/utils/git"
	"kubegems.io/kubegems/pkg/utils/helm"
	"kubegems.io/kubegems/pkg/utils/jwt"
	"kubegems.io/kubegems/pkg/utils/msgbus"
	"kubegems.io/kubegems/pkg/utils/prometheus"
	"kubegems.io/kubegems/pkg/utils/redis"
	"kubegems.io/kubegems/pkg/utils/system"
)

type Options struct {
	System    *system.Options             `json:"system,omitempty"`
	Appstore  *helm.Options               `json:"appstore,omitempty"`
	Argo      *argo.Options               `json:"argo,omitempty"`
	DebugMode bool                        `json:"debugMode,omitempty"`
	Exporter  *prometheus.ExporterOptions `json:"exporter,omitempty"`
	Git       *git.Options                `json:"git,omitempty"`
	JWT       *jwt.Options                `json:"jwt,omitempty"`
	LogLevel  string                      `json:"logLevel,omitempty"`
	Msgbus    *msgbus.Options             `json:"msgbus,omitempty"`
	Mysql     *database.Options           `json:"mysql,omitempty"`
	Redis     *redis.Options              `json:"redis,omitempty"`
}

func DefaultOptions() *Options {
	defaultoptions := &Options{
		Appstore:  helm.NewDefaultOptions(),
		Argo:      argo.NewDefaultArgoOptions(),
		DebugMode: false,
		Exporter:  prometheus.DefaultExporterOptions(),
		Git:       git.NewDefaultOptions(),
		JWT:       jwt.DefaultOptions(),
		LogLevel:  "debug",
		Msgbus:    msgbus.DefaultMsgbusOptions(),
		Mysql:     database.NewDefaultOptions(),
		Redis:     redis.NewDefaultOptions(),
		System:    system.NewDefaultOptions(),
	}
	defaultoptions.System.Listen = ":8020"
	return defaultoptions
}
