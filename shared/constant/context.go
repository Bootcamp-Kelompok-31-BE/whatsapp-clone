package constant

type dbKey string
type redisKey string
type userIDKey string
type loggerKey string
type traceIDKey string
type userEmailKey string

const (
	DBKey        dbKey        = "db"
	RedisKey     redisKey     = "redis"
	UserIDKey    userIDKey    = "user_id"
	LoggerKey    loggerKey    = "logger"
	TraceIDKey   traceIDKey   = "trace_id"
	UserEmailKey userEmailKey = "user_email"
)
