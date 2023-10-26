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

type ReidsOption struct {
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

type Client interface {
	redis.UniversalClient
	Prefix() string
	Separator() string
	JoinKeys(keys ...string) string
	FormatKeys(keys ...string) []string
}

type redisClient struct {
	redis.UniversalClient
	prefix    string
	separator string
}

func New(opt *ReidsOption) redisClient {
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

func (rdb redisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	newKeys := rdb.FormatKeys(keys...)

	return rdb.UniversalClient.Eval(ctx, script, newKeys, args...)
}

func (rdb redisClient) Prefix() string {
	return rdb.prefix
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
