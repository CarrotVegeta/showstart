package service

import (
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/models"
	"github.com/CarrotVegeta/showstart/pkg"
)

func GetSpectatorList() map[string]interface{} {
	s := &models.RequestSpectator{
		StFlpv:    config.Conf.StFlpv,
		Sign:      "",
		TrackPath: "",
		Terminal:  pkg.Terminal,
	}
	return RequestWithBodyParam(pkg.GetSpectatorAction, pkg.B, "POST", s)
}
