package common

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bamboox/aceStock/domains"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/namsral/flag"
)

var (
	Engine        *xorm.Engine
	dbdriver, dsn *string
)

func init() {
	dbdriver = flag.String("dbdriver", "mysql", "Database driver name [default: mysql]")
	dsn = flag.String("dbdsn", "root:admin@/test?charset=utf8", "Data source Name [default: root:root@/test?charset=utf8]")
}

func InitEngine(profiles []string) *xorm.Engine {
	if Engine != nil {
		return Engine
	}

	log.Infof("init XORM engine")
	var err error

	log.Infoln("dbdriver", dbdriver, "dsn", dsn)

	if Engine, err = xorm.NewEngine(*dbdriver, *dsn); err != nil {
		log.Errorln("NewEngine", err)
	}

	if err = Engine.Sync2(new(domains.StockDomainStruct), new(domains.StockDayDomainStruct)); err != nil {
		log.Errorln("Sync2", err)
	}

	return Engine
}
