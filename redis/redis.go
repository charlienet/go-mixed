package redis

import (
	"sync"
	"time"

	"github.com/charlienet/go-mixed/expr"
	"github.com/redis/go-redis/v9"
)

const (
	defaultSeparator = ":"

	blockingQueryTimeout = 5 * time.Second
	readWriteTimeout     = 2 * time.Second
	defaultSlowThreshold = time.Millisecond * 100 // 慢查询
)

var Nil = redis.Nil

var (
	once sync.Once
)

type ReidsOption struct {
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

	PoolSize        int
	PoolTimeout     time.Duration
	MinIdleConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

type Client interface {
	redis.UniversalClient
}

func New(opt *ReidsOption) Client {
	var rdb redis.UniversalClient
	once.Do(func() {
		rdb = redis.NewUniversalClient(&redis.UniversalOptions{
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
		})

		if len(opt.Prefix) > 0 {
			rdb.AddHook(renameKey{
				prefix:    opt.Prefix,
				separator: expr.Ternary(len(opt.Separator) == 0, defaultSeparator, opt.Separator),
			})
		}
	})

	return rdb
}
