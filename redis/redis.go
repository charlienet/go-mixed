package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charlienet/go-mixed/expr"
	"github.com/redis/go-redis/v9"
)

const (
	defaultSeparator = ":"

	blockingQueryTimeout = 5 * time.Second
	readWriteTimeout     = 2 * time.Second
	defaultSlowThreshold = "5000" // 慢查询(单位微秒)
)

var Nil = redis.Nil

type RedisOption struct {
	Addr      string
	Addrs     []string
	Password  string // 密码
	Prefix    string
	Separator string

	// Database to be selected after connecting to the server.
	// Only single-node and failover clients.
	DB int

	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration

	DialTimeout           time.Duration
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool

	// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
	PoolFIFO bool

	PoolSize        int
	PoolTimeout     time.Duration
	MinIdleConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

var _ Client = redisClient{}

type Clients []Client

func (clients Clients) LoadFunction(code string) {
	for _, c := range clients {
		c.LoadFunction(code)
	}
}

type Client interface {
	redis.UniversalClient
	LoadFunction(f string)              // 加载函数脚本
	Prefix() string                     // 统一前缀
	Separator() string                  // 分隔符
	JoinKeys(keys ...string) string     // 连接KEY
	FormatKeys(keys ...string) []string // 格式化KEY
}

type redisClient struct {
	redis.UniversalClient
	prefix    string
	separator string
}

func New(opt *RedisOption) redisClient {
	var rdb redisClient

	if len(opt.Addrs) == 0 && len(opt.Addr) > 0 {
		opt.Addrs = []string{opt.Addr}
	}

	separator := expr.Ternary(len(opt.Separator) == 0, defaultSeparator, opt.Separator)
	prefix := expr.Ternary(len(opt.Prefix) > 0, fmt.Sprintf("%s%s", opt.Prefix, separator), "")

	rdb = redisClient{
		prefix:    prefix,
		separator: separator,
		UniversalClient: redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    opt.Addrs,
			Password: opt.Password,

			DB: opt.DB,

			MaxRetries:      opt.MaxRetries,
			MinRetryBackoff: opt.MinRetryBackoff,
			MaxRetryBackoff: opt.MaxRetryBackoff,

			DialTimeout:           opt.DialTimeout,
			ReadTimeout:           opt.ReadTimeout,
			WriteTimeout:          opt.WriteTimeout,
			ContextTimeoutEnabled: opt.ContextTimeoutEnabled,

			PoolSize:        opt.PoolSize,
			PoolTimeout:     opt.PoolTimeout,
			MinIdleConns:    opt.MinIdleConns,
			MaxIdleConns:    opt.MaxIdleConns,
			ConnMaxIdleTime: opt.ConnMaxIdleTime,
			ConnMaxLifetime: opt.ConnMaxLifetime,
		})}

	rdb.ConfigSet(context.Background(), "slowlog-log-slower-than", defaultSlowThreshold)

	if len(opt.Prefix) > 0 {
		rdb.AddHook(renameKey{
			prefix: prefix,
		})
	}

	return rdb
}

func (rdb redisClient) Prefix() string {
	return rdb.prefix
}

func (rdb redisClient) LoadFunction(code string) {
	_, err := rdb.FunctionLoadReplace(context.Background(), code).Result()
	if err != nil {
		panic(err)
	}
}

func (rdb redisClient) Separator() string {
	return rdb.separator
}

func (rdb redisClient) JoinKeys(keys ...string) string {
	return strings.Join(keys, rdb.separator)
}

func (rdb redisClient) FormatKeys(keys ...string) []string {
	if len(rdb.prefix) == 0 {
		return keys
	}

	re := make([]string, 0, len(keys))
	for _, k := range keys {
		re = append(re, fmt.Sprintf("%s%s", rdb.prefix, k))
	}

	return re
}
