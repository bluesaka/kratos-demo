package data

import (
	"context"
	"database/sql"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	orderv1 "kratos-demo/api/order/service/v1"
	"kratos-demo/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewBunDB,
	NewUserRepo,
	NewRegistrar,
	NewDiscovery,
	NewOrderServiceClient,
)

// Data
type Data struct {
	db          *gorm.DB
	bunDB       *bun.DB
	orderClient orderv1.OrderClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, bunDB *bun.DB, orderClient orderv1.OrderClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, bunDB: bunDB, orderClient: orderClient}, cleanup, nil
}

func NewBunDB(c *conf.Data) *bun.DB {
	sqldb, err := sql.Open("mysql", c.Database.Dsn)
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("test"),
	))
	return db
}

func NewDB(c *conf.Data) *gorm.DB {
	//newLogger := logger.New(
	//	log2.New(os.Stdout, "\r\n", log2.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: time.Second,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	},
	//)

	newLogger := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("db error")
	}

	return db
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewOrderServiceClient(rd registry.Discovery) orderv1.OrderClient {
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint("discovery:///kratos-demo.order"), // 服务发现格式： discovery:///xxx 其中xxx为注册的kratos.Name
		grpc.WithDiscovery(rd),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return orderv1.NewOrderClient(conn)
}
