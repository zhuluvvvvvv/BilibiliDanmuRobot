package config

import (
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	//rest.RestConf
	Log         logx.LogConf
	RoomId      int    `json:",default=4699397"`
	WsServerUrl string `json:",default=wss://broadcastlv.chat.bilibili.com:2245/sub"`
	//QrCodePath   string `json:"qr_code_path"`
	TalkRobotCmd    string            `json:",default=test"`
	RobotName       string            `json:",default=花花"`
	DanmuLen        int               `json:",default=20"`
	EntryEffect     bool              `json:",default=false"`
	EntryMsg        string            `json:",default=off"`
	ThanksGift      bool              `json:",default=false"`
	CustomizeBullet bool              `json:",default=false"`
	InteractWord    bool              `json:",default=false"`
	WelcomeSwitch   bool              `json:",default=false"`
	WelcomeString   map[string]string `json:",optional"`
	RobotMode       string            `json:",default=QingYunKe,options=QingYunKe|ChatGPT"`
	ChatGPT         struct {
		APIToken string `json:",optional"`
	}
	CronDanmu     bool `json:",default=false"`
	CronDanmuList []struct {
		Cron  string `json:",optional"`
		Danmu string `json:",optional"`
	} `json:",default=false"`
	PKNotice bool `json:",default=true"`
}
