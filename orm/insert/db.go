package insert

import (
	"database/sql"
	orm_metadata "demo-golang/orm/metadata"
)

// S6OrmDB orm 框架的数据库对象
type S6OrmDB struct {
	// p7s6SqlDB 真正的数据库对象
	p7s6SqlDB *sql.DB
	// I9Registry 元数据注册中心
	I9Registry orm_metadata.I9Registry
}

// F8NewS6OrmDB 构造 S6OrmDB
func F8NewS6OrmDB(p7s6db *sql.DB) *S6OrmDB {
	return &S6OrmDB{
		p7s6SqlDB:  p7s6db,
		I9Registry: orm_metadata.F8NewI9Registry(),
	}
}
