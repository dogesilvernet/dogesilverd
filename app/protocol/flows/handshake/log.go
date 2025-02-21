package handshake

import (
	"github.com/dogesilvernet/dogesilverd/infrastructure/logger"
	"github.com/dogesilvernet/dogesilverd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
