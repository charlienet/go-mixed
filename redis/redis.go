package redis

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/redis/go-redis/v9"
)

const (
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
	LoadFunction(f string)                  // 加载函数脚本
	Prefix() string                         // 统一前缀
	Separator() string                      // 分隔符
	AddPrefix(prefix ...string) redisClient // 添加前缀
	JoinKeys(keys ...string) string         // 连接KEY
	FormatKeys(keys ...string) []string     // 格式化KEY
	ServerVersion() string
}

type redisClient struct {
	redis.UniversalClient
	prefix redisPrefix
	conf   *redis.UniversalOptions
}

type constraintFunc func(redisClient) error

func Ping() constraintFunc {
	return func(rc redisClient) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return rc.Ping(ctx).Err()
	}
}

func VersionConstraint(expended string) constraintFunc {
	return func(rc redisClient) error {
		v := rc.ServerVersion()
		if len(v) == 0 {
			return errors.New("version not obtained")
		}
		current, err := version.NewVersion(v)
		if err != nil {
			return err
		}

		constraint, err := version.NewConstraint(expended)
		if err != nil {
			return err
		}

		if !constraint.Check(current) {
			return fmt.Errorf("the desired version is %v, which does not match the expected version %v", current, expended)
		}

		return nil
	}
}

func New(opt *RedisOption, constraints ...constraintFunc) redisClient {
	if len(opt.Addrs) == 0 && len(opt.Addr) > 0 {
		opt.Addrs = []string{opt.Addr}
	}

	conf := &redis.UniversalOptions{
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
	}

	rdb := new(conf, newPrefix(opt.Separator, opt.Prefix))

	if len(constraints) > 0 {
		for _, f := range constraints {
			if err := f(rdb); err != nil {
				panic(err)
			}
		}
	}

	return rdb
}

func NewEnforceConstraints(opt *RedisOption, constraints ...constraintFunc) redisClient {
	rdb := New(opt)
	for _, f := range constraints {
		if err := f(rdb); err != nil {
			panic(err)
		}
	}

	return rdb
}

func (rdb redisClient) Prefix() string {
	return rdb.prefix.Prefix()
}

func (rdb redisClient) Separator() string {
	return rdb.prefix.Separator()
}

func (rdb redisClient) AddPrefix(prefixes ...string) redisClient {
	old := rdb.prefix
	p := newPrefix(old.separator, old.join(prefixes...))

	return new(rdb.conf, p)
}

func (rdb redisClient) LoadFunction(code string) {
	_, err := rdb.FunctionLoadReplace(context.Background(), code).Result()
	if err != nil {
		panic(err)
	}
}

func (rdb redisClient) JoinKeys(keys ...string) string {
	return rdb.prefix.join(keys...)
}

func (rdb redisClient) FormatKeys(keys ...string) []string {
	if !rdb.prefix.hasPrefix() {
		return keys
	}

	re := make([]string, 0, len(keys))
	for _, k := range keys {
		re = append(re, fmt.Sprintf("%s%s", rdb.prefix.Prefix(), k))
	}

	return re
}

func (rdb redisClient) ServerVersion() string {
	info, err := rdb.Info(context.Background(), "server").Result()
	if err != nil {
		return ""
	}

	for _, line := range strings.Split(info, "\r\n") {
		after, found := strings.CutPrefix(line, "redis_version:")
		if found {
			return after
		}
	}

	return ""
}

func new(conf *redis.UniversalOptions, prefix redisPrefix) redisClient {
	c := redis.NewUniversalClient(conf)
	c.ConfigSet(context.Background(), "slowlog-log-slower-than", defaultSlowThreshold)
	c.AddHook(renameHook{prefix: prefix})

	return redisClient{
		UniversalClient: c,
		prefix:          prefix,
		conf:            conf,
	}
}
