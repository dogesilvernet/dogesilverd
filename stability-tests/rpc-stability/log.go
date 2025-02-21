package main

import (
	"github.com/dogesilvernet/dogesilverd/infrastructure/logger"
	"github.com/dogesilvernet/dogesilverd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
