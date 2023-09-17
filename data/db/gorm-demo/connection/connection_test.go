package main

import (
	"fmt"
	"gorm.io/gorm/clause"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// 测试insert
func TestNewDB(t *testing.T) {
	db := NewDB()
	record := APIRequestRecord{
		APIPath:    "/test",
		Method:     "POST",
		Params:     "xxxx",
		APIName:    "gorm测试",
		CostTime:   100,
		IP:         "10.1.111.20",
		Username:   "张三",
		StatusCode: 200,
		StatusText: "ok",
		CreateBy:   "zhangsan",
		UpdateBy:   "lisi",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
	record2 := APIRequestRecord{
		APIPath:    "test select insert",
		APIName:    "test select insert",
		Method:     "post",
		Params:     "xxx",
		CostTime:   100,
		IP:         "111.111.11.111",
		Username:   "xxx",
		StatusCode: 200,
		StatusText: "xxx",
		UpdateBy:   "xxxx",
	}
	tx := db.Create(&record)
	fmt.Println("rows: ", tx.Row())
	fmt.Println("insert pk id: ", record.ID)
	fmt.Println("error: ", tx.Error)
	fmt.Println("affect rows: ", tx.RowsAffected)
	res := db.Select("apipath", "apiname", "method",
		"params", "cost_time", "ip", "username",
		"status_code", "status_text", "update_by").Create(&record2)
	fmt.Println("res: ", res)
}

// 测试批量写入
func TestBatchInsert(t *testing.T) {
	db := NewDB()
	record := make([]APIRequestRecord, 0)
	for i := 0; i < 100; i++ {
		record = append(record, APIRequestRecord{
			APIPath:    fmt.Sprintf("/test/%s", strconv.QuoteRune(rune(i))),
			Method:     "Post",
			Params:     "xxx",
			APIName:    "test-" + time.Now().Format("2006-01-02 15:04:05"),
			CostTime:   100,
			IP:         "10.111.10.23",
			Username:   fmt.Sprintf("zhangsan-%s", strconv.QuoteRune(rune(i))),
			StatusCode: 200,
			StatusText: fmt.Sprintf("test-%s", strconv.QuoteRune(rune(i))),
			CreateBy:   fmt.Sprintf("zhangsan-%s", strconv.QuoteRune(rune(i))),
			UpdateBy:   fmt.Sprintf("zhangsan-%s", strconv.QuoteRune(rune(i))),
			CreateTime: time.Time{},
			UpdateTime: time.Time{},
		})
		fmt.Println("length: ", len(record))
	}
	fmt.Println("all length: ", len(record))
	result := db.Create(record)
	fmt.Println("affect rows: ", result.RowsAffected)
}

func TestMapInsert(t *testing.T) {
	db := NewDB()
	db.Model(&APIRequestRecord{}).Create(map[string]interface{}{
		"api_path":    "/api/v1/bill/",
		"method":      "GET",
		"params":      "xxxxxxx",
		"api_name":    "费用账单接口",
		"cost_time":   303,
		"ip":          "127.0.0.1",
		"username":    "jiayawu",
		"status_code": 200,
		"status_text": "",
		"create_by":   "jiayawu", "update_by": "jiayawu",
	})
}

func TestBatchMapInsert(t *testing.T) {
	records := make([]map[string]interface{}, 0)
	for i := 0; i < 100; i++ {
		records = append(records, map[string]interface{}{
			"api_path":    "/api/v1/bill/",
			"method":      "GET",
			"params":      fmt.Sprintf("xxxxxxx-%s", strconv.QuoteRune(rune(i))),
			"api_name":    "费用账单接口",
			"cost_time":   i,
			"ip":          "127.0.0.1",
			"username":    "jiayawu",
			"status_code": 200,
			"status_text": "",
			"create_by":   "jiayawu", "update_by": "jiayawu",
		})
	}
	fmt.Println("total records: ", len(records))
	db := NewDB()
	result := db.Model(&APIRequestRecord{}).Create(records)
	fmt.Println("affect rows :", result.RowsAffected)
}

func TestUpSertIgnore(t *testing.T) {
	db := NewDB()
	// 遇到冲突时候，什么都不做
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&APIRequestRecord{
		ID:         9591348,
		APIPath:    "测试接口",
		Method:     "get",
		Params:     "xxx",
		APIName:    "xxx",
		CostTime:   10,
		IP:         "10.90.10.23",
		Username:   "xxx",
		StatusCode: 200,
		StatusText: "xxxxxx",
		CreateBy:   "zhangsan",
		UpdateBy:   "lisi",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})
}

func TestUpsertDefault(t *testing.T) {
	/**
	  遇到冲突时候，更新值为DoUpdates中的默认值
	  类似于mysql的 INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***
	*/
	db := NewDB()
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"api_name": "冲突了。。"}),
	}).Create(&APIRequestRecord{
		ID:         9591348,
		APIPath:    "测试接口",
		Method:     "get",
		Params:     "xxx",
		APIName:    "xxx",
		CostTime:   10,
		IP:         "10.90.10.23",
		Username:   "xxx",
		StatusCode: 200,
		StatusText: "xxxxxx",
		CreateBy:   "zhangsan",
		UpdateBy:   "lisi",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})
}

func TestUpsertNew(t *testing.T) {
	db := NewDB()
	// insert 冲突时候更新AssignmentColumns 中的值为新值
	res := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"api_path", "params", "api_name", "status_text"}),
	}).Create(&APIRequestRecord{
		ID:         9591346,
		APIPath:    "测试接口冲突时候更新为新值",
		Method:     "get",
		Params:     "测试接口冲突时候更新为新值",
		APIName:    "测试接口冲突时候更新为新值",
		CostTime:   10,
		IP:         "10.90.10.23",
		Username:   "xxx",
		StatusCode: 200,
		StatusText: "测试接口冲突时候更新为新值",
		CreateBy:   "zhangsan",
		UpdateBy:   "lisi",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})
	fmt.Println("res: ", res.Error, res.RowsAffected)
	fmt.Println("测试覆盖率: ", testing.Coverage())
}

func TestUpsertAll(t *testing.T) {
	db := NewDB()
	// insert 冲突时候更新除主键以外的所有列到新值
	res := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&APIRequestRecord{
		ID:         9591551,
		APIPath:    "测试接口更新除主键以外的所有列到新值",
		Method:     "get",
		Params:     "测试接口更新除主键以外的所有列到新值",
		APIName:    "测试接口更新除主键以外的所有列到新值",
		CostTime:   10,
		IP:         "10.90.10.23",
		Username:   "xxx",
		StatusCode: 200,
		StatusText: "测试接口更新除主键以外的所有列到新值",
		CreateBy:   "zhangsan",
		UpdateBy:   "lisi",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})
	fmt.Println("res: ", res.Error, res.RowsAffected)
	fmt.Println("测试覆盖率: ", testing.Coverage())
}

// 测试单个查询
func TestSelectSingle(t *testing.T) {
	db := NewDB()
	// 获取第一条记录,order by id limit 1
	firstRecordOrderByPk := db.First(&APIRequestRecord{})
	fmt.Println("firstRecordOrderByPk: ", firstRecordOrderByPk)
	// 获取一条记录，没有指定排序字段
	firstRecord := db.Take(&APIRequestRecord{})
	fmt.Println("firstRecord: ", firstRecord)
	// 获取最后一条记录 order by id desc limit 1
	lastRecord := db.Last(&APIRequestRecord{})
	fmt.Println("lastRecord: ", lastRecord, lastRecord.Error, lastRecord.RowsAffected)
	// 主键检索
	singlePkRec := db.Find(&APIRequestRecord{}, 9591551)
	fmt.Println("singlePkRec: ", singlePkRec)
	singleStrPkRec := db.Find(&APIRequestRecord{}, "9591551")
	fmt.Println("singleStrPkRec: ", singleStrPkRec)
	record := db.Find(&APIRequestRecord{}, []int{9591551, 9591550,
		9591549, 9591548,
	})
	fmt.Println("record: ", record)
	// 字符串主键检索
	strPkRec := db.Find(&APIRequestRecord{}, "id= ?", "9591550")
	fmt.Println("strPkRec: ", strPkRec)
	// 检索全部对象
	allRec := db.Find(&APIRequestRecord{})
	fmt.Println("allRec: ", allRec, allRec.Error, allRec.RowsAffected)
}

// 测试查询结果fetch到map slice
func TestResultToMapSlice(t *testing.T) {
	db := NewDB()
	// find 到 map数组，将扫描结果至 []map[string]interface{}
	records := make([]map[string]interface{}, 0)
	lastTenRecord := db.Model(&APIRequestRecord{}).Order("id desc").Limit(10).Find(&records)
	fmt.Println("lastTenRecord: ", lastTenRecord)
	fmt.Println("records: ", records)
	for _, record := range records {
		fmt.Println("record: ", record)
		fmt.Println(reflect.TypeOf(record))
	}
}

// 查询结果fetch到map
func TestResultToMap(t *testing.T) {
	db := NewDB()
	// find 到 map，将扫描结果至map[string]interface{}
	mapRecord := db.Model(&APIRequestRecord{}).Order("id desc").Limit(10)
	var mapRec map[string]interface{}
	mapRows, err := mapRecord.Rows()

	if err != nil {
		panic(err)
	}

	for mapRows.Next() {
		db.ScanRows(mapRows, &mapRec)
		fmt.Println("mapRec: ", mapRec)
	}
}

// 查询结果fetch到struct
func TestResultToStruct(t *testing.T) {
	db := NewDB()
	// find 到 struct，将扫描结果至 struct
	structRecord := db.Model(&APIRequestRecord{}).Order("id desc").Limit(10)
	var apiRec APIRequestRecord
	rows, err := structRecord.Rows()
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		db.ScanRows(rows, &apiRec)
		fmt.Println("apiRec: ", apiRec)
		fmt.Println("api_name: ", apiRec.APIName)
		fmt.Println("api_path: ", apiRec.APIPath)
	}
}

// 查询结果fetch到 struct slice
func TestResultToStructSlice(t *testing.T) {
	db := NewDB()
	// find 到 struct array，将结果扫描到一个结构体数组中
	structArrayRec := make([]APIRequestRecord, 0)
	structArrayRecord := db.Model(&APIRequestRecord{}).Order("id desc").Limit(10).Find(&structArrayRec)
	fmt.Println("structArrayRecord: ", structArrayRecord.RowsAffected)
	fmt.Println("structArrayRec len: ", len(structArrayRec))
	for _, rec := range structArrayRec {
		fmt.Println("rec: ", rec)
		fmt.Println("rec.APIName", rec.APIName)
	}
}

func TestFindByPk(t *testing.T) {
	db := NewDB()
	var record APIRequestRecord
	// select * from api_request_record where id = 9589771;
	db.First(&record, 9589771)
	fmt.Println(record)
	var record2 APIRequestRecord
	db.First(&record2, "9589771")
	fmt.Println(record2)
	var records []APIRequestRecord
	db.Find(&records, []int{9589771, 9589772, 9589773, 9589774})
	fmt.Println("records: ")
	fmt.Println(records)
	//	如果主键是字符串，查询将被写成这样
	var rec APIRequestRecord
	db.First(&rec, "id = ?", "9589772")
	fmt.Println("rec: ", rec)
	//	如果目标对象有主键的话，主键将会用来构造条件，例如
	apiRec := APIRequestRecord{ID: 9589772}
	db.First(&apiRec)
	fmt.Println("apiRec: ", apiRec)

	var result APIRequestRecord
	db.Model(APIRequestRecord{ID: 9589772}).First(&result)
	fmt.Println("result: ", result)
}

func TestFindAll(t *testing.T) {
	db := NewDB()
	var records []APIRequestRecord
	result := db.Find(&records)
	fmt.Println(result.RowsAffected, result.Error)
}

// 字符串条件
func TestStringCondition(t *testing.T) {
	var record APIRequestRecord
	var records []APIRequestRecord
	db := NewDB()
	db.Where("method= ?", "GET").First(&record)
	//fmt.Println("record: ", record)
	// In
	db.Where("method in ?", []string{"POST", "GET"}).Find(&records)
	//fmt.Println("records: ", records)

	//	 like
	db.Where("api_path like ?", "%order%").Find(&records)
	//fmt.Println("orders: ", records)
	//	 and
	db.Where("api_path = ? and method = ?", "/api/v1/search", "GET").Find(&records)
	//fmt.Println("records: ", records)

	// Time
	now := time.Now()
	db.Where("create_time < ?", now).Find(&records)
	//fmt.Println("records: ", records)

	db.Where("create_time between ? and ?", now, now).Find(&records)
	fmt.Println("records: ", records)
}

// 测试map和struct条件
func TestStructMapCondition(t *testing.T) {
	db := NewDB()
	var record APIRequestRecord
	var records []APIRequestRecord
	// 当查询是结构体时候，GORM只会查询非空的字段
	db.Where(&APIRequestRecord{APIPath: "/api/v1/myorder/", Method: "GET"}).Find(&record)
	// select * from api_request_record where api_path='/api/v1/myorder/' and method='GET'
	db.Where(&APIRequestRecord{APIPath: "/api/v1/myorder/", Method: "GET", APIName: ""}).Find(&record)

	db.Where(map[string]interface{}{"api_path": "/api/v1/myorder/", "method": "GET"}).Find(&records)
	// map会包含空值的字段查询, select * from api_request_record where api_path='' and method='GET'
	db.Where(map[string]interface{}{"api_path": "", "method": "GET"}).Find(&records)

	fmt.Println(records)

	// 指定结构体查询字段
	// select * from api_request_record where method='GET' and api_name=''
	db.Where(&APIRequestRecord{Method: "GET"}, "method", "api_name").Find(&records)
	// select * from api_request_record where  api_name=''
	db.Where(&APIRequestRecord{Method: "GET"}, "api_name").Find(&records)

	// 内联条件
	// struct
	// select * from api_request_record where method='get';
	db.Find(&records, APIRequestRecord{Method: "GET"})
	// map
	// select * from api_request_record where method='get';
	db.Find(&records, map[string]interface{}{"method": "GET"})
}

func TestNotCond(t *testing.T) {
	db := NewDB()
	var record APIRequestRecord
	// select * from api_request_record where not method='GET';
	db.Not("method = ?", "GET").Find(&record)

	// Not In
	//db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	//db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	//db.Not([]int64{1,2,3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
}

func TestOrCond(t *testing.T) {
	db := NewDB()
	var record APIRequestRecord
	db.Where("method = ?", "GET").Or("method = ?", "POST").Find(&record)
	fmt.Println(record)
}

// 选择指定字段
func TestSelectField(t *testing.T) {
	db := NewDB()
	var records []APIRequestRecord
	// select method,api_name from api_request_record limit 10;
	db.Select("method", "api_name").Find(&records).Limit(10)
	fmt.Println(records)
}

func TestOrder(t *testing.T) {
	db := NewDB()
	var records []APIRequestRecord
	db.Order("id desc,method asc").Find(&records)
	fmt.Println(records)
}

func TestLimit(t *testing.T) {
	db := NewDB()
	var records []APIRequestRecord
	// select * from api_request_record limit 10;
	db.Limit(10).Find(&records)
	fmt.Println(records)
}

func TestGroupBy(t *testing.T) {
	db := NewDB()
	type Result struct {
		ApiNum int64
		Method string
	}
	var result []Result
	// select count(*) as api_num,method from api_request_record group by method order by api_num desc;
	db.Model(&APIRequestRecord{}).Select("count(*) as api_num,method").Group("method").Order("api_num desc").Find(&result)
	// 	result:  [{728 POST} {721 GET} {332 PUT}]
	fmt.Println("result: ", result)
}

func TestDistinct(t *testing.T) {
	db := NewDB()
	type Result struct {
		ApiName string
		Method  string
	}
	var result []Result
	// select distinct api_name,method from api_request_record order by api_name desc,method;
	db.Model(&APIRequestRecord{}).Distinct("api_name", "method").Order("api_name desc,method ").Find(&result)
	fmt.Println("result: ", result)
}

func TestScan(t *testing.T) {
	db := NewDB()
	var records []APIRequestRecord
	db.Model(APIRequestRecord{}).Where("id> ?", 95000).Scan(&records)
	fmt.Println(records)
}

func TestCount(t *testing.T) {
	db := NewDB()
	var count int64
	db.Model(APIRequestRecord{}).Count(&count)
	fmt.Println("count: ", count)
}

func TestUpdate(t *testing.T) {
	db := NewDB()
	res := db.Model(APIRequestRecord{}).Where("id> ?", 9591542).Update("api_name", "test update").Commit()
	fmt.Println(res.Error, res.RowsAffected)
	var record APIRequestRecord
	db.Model(APIRequestRecord{}).First(&record)
	record.APIName = "测试更新单个列"
	fmt.Println(record.ID)
	res2 := db.Model(&record).Update("api_name", "hello").Commit()
	fmt.Println("res2: ", res2)
	deleteRes := db.Delete(&record).Commit()
	fmt.Println("deleteRes: ", deleteRes)
}
