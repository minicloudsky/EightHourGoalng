package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

type HoueseType string

const (
	HouseTypeErshoufang HoueseType = "ershoufang"
	HouseTypeLoupan     HoueseType = "loupan"
)

type CityInfo struct {
	Name string `json:"name"`
}

var (
	ErshoufangProcessChan chan []Message
	LoupanProcessChan     chan []Message
	CommercialProcessChan chan []Message
	ZufangProcessChan     chan []Message
)

func InitLianjiaChan(ctx context.Context) {
	fmt.Println("---initializing lianjia channel...")
	ErshoufangProcessChan = make(chan []Message, 100)
	ZufangProcessChan = make(chan []Message, 100)
	LoupanProcessChan = make(chan []Message, 100)
	CommercialProcessChan = make(chan []Message, 100)
}

type Message struct {
	Content []byte
}

type poolParam struct {
	ctx       context.Context
	city      *CityInfo
	houseType HoueseType
}

func ListCityErshouFang(ctx context.Context, city *CityInfo) error {
	fmt.Println("ListCityErshouFang: ", city.Name)
	ErshoufangProcessChan <- []Message{{Content: []byte(city.Name)}}
	return nil
}

func ListCityLoupan(ctx context.Context, city *CityInfo) error {
	fmt.Println("ListCityLoupan: ", city.Name)
	LoupanProcessChan <- []Message{{Content: []byte(city.Name)}}
	return nil
}

func poolFunc(i interface{}) {
	param := i.(poolParam)
	ctx := param.ctx
	city := param.city
	houseType := param.houseType
	switch houseType {
	case HouseTypeErshoufang:
		err := ListCityErshouFang(ctx, city)
		if err != nil {
			return
		}
	case HouseTypeLoupan:
		err := ListCityLoupan(ctx, city)
		if err != nil {
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var goroutineSize = 50
	ctx := context.Background()
	InitLianjiaChan(ctx)
	var cities = []string{
		"北京", "上海", "广州", "深圳", "杭州", "成都", "重庆", "西安1", "厦门", "武汉",
		"沈阳", "长沙", "南京", "苏州", "天津", "南昌", "青岛", "大连", "合肥", "南宁",
		"福州", "长春", "贵阳", "昆明", "西宁", "海口", "石家庄", "太原", "呼和浩特", "乌鲁木齐",
		"银川", "兰州", "西双版纳", "乌兰察布", "拉萨", "包头", "南通", "哈尔滨", "西安", "西宁",
		"拉萨", "乌鲁木齐", "昆明", "西安", "西宁", "拉萨", "乌鲁木齐", "昆明", "西安", "西宁",
		"拉萨", "乌鲁木齐", "昆明", "西安", "西宁", "拉萨", "鲁木齐", "昆明", "西安2",
		"北京2", "上海2", "广州2", "深圳2", "杭州2", "成都2", "重庆2", "西安3", "厦门2", "武汉2",
		"沈阳2", "长沙2", "南京2", "苏州2", "天津2", "南昌2", "青岛2", "大连2", "合肥2", "南宁2",
		"福州2", "长春2", "贵阳2", "昆明2", "西宁2", "海口2", "石家庄2", "太原2", "呼和浩特2", "乌鲁木齐2",
		"银川2", "兰州2", "西双版纳2", "乌兰察布2", "拉萨2", "包头2", "南通2", "哈尔滨2", "西安4", "西宁3",
		"拉萨3", "乌鲁木齐3", "昆明3", "西安5", "西宁4", "拉萨4", "乌鲁木齐4", "昆明4", "西安6", "西宁5",
		"拉萨5", "乌鲁木齐5", "昆明5", "西安7", "西宁6", "拉萨6", "鲁木齐2", "昆明6", "西安8", "西宁7",
		// 继续添加城市名...
	}
	fmt.Println("cities: ", len(cities))
	p, err := ants.NewPoolWithFunc(goroutineSize, func(i interface{}) {
		poolFunc(i)
		wg.Done()
	})
	if err != nil {
		fmt.Printf("ants.NewPoolWithFunc err! err: %v", err)
	}
	for _, city := range cities {
		wg.Add(1)
		param := poolParam{
			ctx:       ctx,
			city:      &CityInfo{Name: city},
			houseType: HouseTypeErshoufang,
		}
		err = p.Invoke(param)
		if err != nil {
			fmt.Printf("invoke ershoufang task err! err: %v", err)
		}
	}
	wg.Wait()
	p.Waiting()
}
