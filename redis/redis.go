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

var (
	once sync.Once
)

type ReidsOption struct {
	Addrs     []string
	Password  string // 密码
	Prefix    string
	Separator string

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

type Client struct {
	redis.UniversalClient
}

func New(opt *ReidsOption) Client {
	var rdb redis.UniversalClient
	once.Do(func() {
		rdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    opt.Addrs,
			Password: opt.Password,

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

	return Client{UniversalClient: rdb}
}
