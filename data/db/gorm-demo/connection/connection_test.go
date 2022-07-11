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

func TestResultToMap(t *testing.T) {
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
	// find 到 map，将扫描结果至map[string]interface{}
	mapRecord := db.Model(&APIRequestRecord{}).Order("id desc").Limit(10)
	var mapRec map[string]interface{}
	mapRows, err := mapRecord.Rows()
	for mapRows.Next() {
		db.ScanRows(mapRows, &mapRec)
		fmt.Println("mapRec: ", mapRec)
	}

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
