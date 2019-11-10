// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/groundcheck"
	"github.com/sandykarunia/fudge/groundcheck/checkers"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/server/handler"
	"github.com/sandykarunia/fudge/utils"
)

// Injectors from server_wire.go:

func Instance() Server {
	osFunctions := sdk.ProvideOSFunctions()
	execFunctions := sdk.ProvideExecFunctions()
	system := utils.ProvideSystem(osFunctions, execFunctions)
	ioFunctions := sdk.ProvideIOFunctions()
	file := utils.ProvideFile(ioFunctions, osFunctions)
	path := utils.ProvidePath(system)
	checkersCheckers := checkers.Provider(system, file, path)
	groundCheck := groundcheck.Provider(checkersCheckers)
	graderGrader := grader.Provider()
	handlerHandler := handler.Provider(graderGrader)
	loggerLogger := logger.Provider(osFunctions, system)
	server := Provider(groundCheck, handlerHandler, loggerLogger)
	return server
}
