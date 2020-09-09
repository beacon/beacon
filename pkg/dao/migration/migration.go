/*Package migration 提供数据库迁移, 使用文件夹而非其他方式进行迁移
基于github.com/mattes/migrate做一层封装，仅使用文件方式作为sql源
*/
package migration

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	migrate "github.com/golang-migrate/migrate/v4"
	database "github.com/golang-migrate/migrate/v4/database"
	file "github.com/golang-migrate/migrate/v4/source/file"
)

// 借助以下值预设数据库迁移所需配置
var (
	DriverName string
	DBURI      string
	SQLPath    string
)

// Init init
func Init(driverName, dbURI, sqlPath string) {
	DriverName, DBURI, SQLPath = driverName, dbURI, sqlPath
}

func must(fn func(session *migrate.Migrate) error) {
	session, err := New(DriverName, DBURI, SQLPath)
	if err != nil {
		log.Fatalf("Failed to create migration session: %v sqlPath=%s", err, SQLPath)
	}
	defer session.Close()
	if err := fn(session); err != nil {
		log.Fatalf("Error migrate: %v", err)
	}
	version, dirty, err := session.Version()
	log.Printf("Database Status: version:%d dirty:%t err:%v", version, dirty, err)
}

// MustMigrate 升级
func MustMigrate() {
	must(Migrate)
}

// MustSet 设置为指定版本
func MustSet(force bool, version string) {
	must(func(session *migrate.Migrate) error {
		return Set(session, force, version)
	})
	log.Println("database is set to version:", version)
}

// MustDown 降一级
func MustDown() {
	must(Down)
	log.Println("database is downgraded successfully")
}

// New 创建新的migrate session
func New(driverName, dbURI, sqlPath string) (*migrate.Migrate, error) {
	driver, err := database.Open(dbURI)
	if err != nil {
		return nil, err
	}

	if sqlPath[0] != '/' {
		sqlPath = "/" + sqlPath
	}

	f := &file.File{}
	srcDrv, err := f.Open("file://" + sqlPath)
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithInstance("file", srcDrv, driverName, driver)
	if err != nil {
		return nil, fmt.Errorf("migrate.NewWithInstance:%v", err)
	}

	return m, nil
}

// Migrate will ensure db version is latest
func Migrate(m *migrate.Migrate) error {
	err := m.Up() // run your migrations and handle the errors above of course
	switch err {
	case nil:
		log.Println("database migrated successfully")
		return nil
	case migrate.ErrNoChange:
		log.Println("database has migrated to latest version, nothing changed")
		return nil
	default:
		return err
	}
}

// Set is a subCommand for debug only.
func Set(m *migrate.Migrate, force bool, version string) error {
	versionParsed, err := strconv.ParseInt(version, 10, 64)
	if err != nil {
		return errors.New("invalid format, must be a number")
	}

	if force {
		return m.Force(int(versionParsed))
	}
	return m.Steps(int(versionParsed))
}

// Down is a subCommand for downgrade db for debug only.
func Down(m *migrate.Migrate) error {
	return m.Steps(-1)
}
