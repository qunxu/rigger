module github.com/tal-tech/rigger

go 1.12

require (
	github.com/apache/calcite-avatica-go/v5 v5.0.0
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.3
	github.com/lib/pq v1.3.0
	github.com/lunny/log v0.0.0-20160921050905-7887c61bf0de
	github.com/spf13/cobra v1.0.0
	github.com/ziutek/mymysql v1.5.4
)

replace github.com/apache/calcite-avatica-go/v5 v5.0.0 => gitlab.alibaba-inc.com/amap-car-lbs/calcite-avatica-go-lindorm/v5 v5.0.29
