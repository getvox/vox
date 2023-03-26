package runtime

import (
	"github.com/iobrother/zoo/core/config"
	"github.com/iobrother/zoo/core/log"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	zdb "github.com/iobrother/zim/pkg/database/db"
	zredis "github.com/iobrother/zim/pkg/database/redis"
	znats "github.com/iobrother/zim/pkg/nats"
)

var defaultRuntime Runtime

type Runtime struct {
	db *gorm.DB
	rc *redis.Client
	nc *nats.Conn
	js nats.JetStreamContext
}

func (r *Runtime) GetDB() *gorm.DB {
	return r.db
}

func (r *Runtime) GetRedisClient() *redis.Client {
	return r.rc
}

func (r *Runtime) GetNC() *nats.Conn {
	return r.nc
}

func (r *Runtime) GetJS() nats.JetStreamContext {
	return r.js
}

func (r *Runtime) SetDB(db *gorm.DB) {
	defaultRuntime.db = db
}

func (r *Runtime) SetRedisClient(rc *redis.Client) {
	defaultRuntime.rc = rc
}

func (r *Runtime) SetNats(nc *nats.Conn) {
	r.nc = nc
	r.js, _ = r.nc.JetStream(nats.PublishAsyncMaxPending(256))
}

func GetDB() *gorm.DB {
	return defaultRuntime.GetDB()
}

func GetRedisClient() *redis.Client {
	return defaultRuntime.GetRedisClient()
}

func GetNC() *nats.Conn {
	return defaultRuntime.GetNC()
}

func GetJS() nats.JetStreamContext {
	return defaultRuntime.GetJS()
}

func Setup() {
	if config.Get("mysql") != nil {
		c := zdb.Config{}
		if err := config.Scan("mysql", &c); err != nil {
			log.Fatal(err)
		}
		db, err := zdb.Open(&c)
		if err != nil {
			log.Fatal(err)
		}
		defaultRuntime.SetDB(db)
	}

	if config.Get("redis") != nil {
		c := zredis.Config{}
		if err := config.Scan("redis", &c); err != nil {
			log.Fatal(err)
		}
		rc, err := zredis.NewClient(&c)
		if err != nil {
			log.Fatal(err)
		}
		defaultRuntime.SetRedisClient(rc)
	}

	if config.Get("nats") != nil {
		c := znats.Config{}
		if err := config.Scan("nats", &c); err != nil {
			log.Fatal(err)
		}
		nc, err := znats.Open(&c)
		if err != nil {
			log.Fatal(err)
		}

		defaultRuntime.SetNats(nc)
	}
}
