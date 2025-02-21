package connmanager

import (
	"github.com/dogesilvernet/dogesilverd/infrastructure/logger"
	"github.com/dogesilvernet/dogesilverd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
