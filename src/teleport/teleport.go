package teleport

import "github.com/ahsen17/BlogServ/src/data"

var (
	dbClient = data.DBClient()
	cache    = data.CacheClient()
)
