package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"

	"LifeService"
)

var comm *tars.Communicator

//SLOG 日志
var SLOG = rogger.GetLogger("ServerLog")

func main() { //Init servant
	comm = tars.NewCommunicator()
	imp := new(ClubActivityManagerImp)                                    //New Imp
	imp.init()
	app := new(LifeService.ClubActivityManager)                                 //New init the A Tars
	cfg := tars.GetServerConfig()                                         //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".ClubActivityManagerObj") //Register Servant
	tars.Run()
}
