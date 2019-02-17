package database

import (
	"database/sql"
	"fmt"

	// 引入 mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/lovestreet/mysqlDoc/config"
)

//Init 初始化
func Init(cfg config.Configuration) error {
	gMySQL.connect(cfg.ConnString())
	return nil
}

var gMySQL = new(mysql)

//mysql 连接及操作对象
type mysql struct {
	dbInstance *sql.DB
	connString string //连接字符串
}

func (m *mysql) close() {
	if m.dbInstance != nil {
		m.dbInstance.Close()
	}
}

func (m *mysql) connect(connString string) error {
	if len(connString) == 0 {
		return fmt.Errorf("connection string is empty")
	}

	const mode = "mysql"

	var err error
	if m.dbInstance, err = sql.Open(mode, connString); err != nil {
		fmt.Printf("connect database error:[%v]", err)
		m.dbInstance = nil
		return err
	}

	fmt.Println("connect database success")
	return nil
}

func (m *mysql) refreshConn() {
	if m.isAlive() {
		return
	}

	m.dbInstance = nil //clean dbInstance

	for i := 0; i < 3; i++ {
		if err := m.connect(m.connString); err != nil {
			continue
		}
		break
	}

	return
}

func (m *mysql) isAlive() bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("check mysql isAlive recover error :[%v]", err)
		}
	}()

	if m.dbInstance == nil {
		return false
	}

	if err := m.dbInstance.Ping(); err != nil {
		fmt.Printf("ping database error :[%v]", err)
		return false
	}

	return true
}

func (m *mysql) query() {

}
