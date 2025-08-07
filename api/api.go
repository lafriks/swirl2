package api

import (
	"time"

	"github.com/cuigh/auxo/app/ioc"
	"github.com/cuigh/auxo/net/web"
)

const defaultTimeout = 30 * time.Second

func ajax(ctx web.Context, err error) error {
	if err != nil {
		return err
	}
	return success(ctx, nil)
}

func success(ctx web.Context, data interface{}) error {
	return ctx.Result(0, "", data)
}

func init() {
	ioc.Put(NewSystem, ioc.Name("api.system"))
	ioc.Put(NewSetting, ioc.Name("api.setting"))
	ioc.Put(NewNode, ioc.Name("api.node"))
	ioc.Put(NewRegistry, ioc.Name("api.registry"))
	ioc.Put(NewNetwork, ioc.Name("api.network"))
	ioc.Put(NewService, ioc.Name("api.service"))
	ioc.Put(NewTask, ioc.Name("api.task"))
	ioc.Put(NewConfig, ioc.Name("api.config"))
	ioc.Put(NewSecret, ioc.Name("api.secret"))
	ioc.Put(NewStack, ioc.Name("api.stack"))
	ioc.Put(NewImage, ioc.Name("api.image"))
	ioc.Put(NewContainer, ioc.Name("api.container"))
	ioc.Put(NewVolume, ioc.Name("api.volume"))
	ioc.Put(NewUser, ioc.Name("api.user"))
	ioc.Put(NewRole, ioc.Name("api.role"))
	ioc.Put(NewEvent, ioc.Name("api.event"))
	ioc.Put(NewChart, ioc.Name("api.chart"))
	ioc.Put(NewDashboard, ioc.Name("api.dashboard"))
}
