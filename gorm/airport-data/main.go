package main

import (
	"fmt"
	"github.com/fishdemon/go-yehua/gorm/airport-data/uid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strings"
	"time"
)

type EvaOutFeedbackConfig struct {
	ConfigId   int64     `gorm:"column:ConfigId"`
	TypeId     int64     `gorm:"column:TypeId"`
	TypeName   string    `gorm:"column:TypeName"`
	SceneId    int64     `gorm:"column:SceneId"`
	SceneName  string    `gorm:"column:SceneName"`
	AreasJson  string    `gorm:"column:AreasJson"`
	UpdateTime time.Time `gorm:"column:UpdateTime"`
}

type EvaFeedbackArea struct {
	FeedbackAreaId     int64     `gorm:"column:AreaId"`
	FeedbackAreaName   string    `gorm:"column:AreaName"`
	AppIconPath string `gorm:"column:AppIconPath"`
	MiniProgIconPath string `gorm:"column:MiniProgIconPath"`
}

type Temp struct {
	PId string `json:"parent_id";gorm:"column:pid"`
	Type string `json:"type"`
	Title string `json:"title"`
	EvalueateId string `json:"evalueateId";gorm:"column:evalueateId"`
}

type EvaAppDetail struct {
	EvaId            int64     `gorm:"column:EvaId"`
	EvaTime          time.Time `gorm:"column:EvaTime"`
	LevelGrade       int       `gorm:"column:LevelGrade"`
	LevelDesc        string    `gorm:"column:LevelDesc"`
	UserId           string    `gorm:"column:UserId"`
	FlightNum        string    `gorm:"column:FlightNum"`
	EvaStageId       int64     `gorm:"column:EvaStageId"`
	EvaStageName     string    `gorm:"column:EvaStageName"`
	EvaAppSourceId   int64     `gorm:"column:EvaAppSourceId"`
	EvaAppSourceName string    `gorm:"column:EvaAppSourceName"`
	Comment          string    `gorm:"column:Comment"`
	UpdateTime       time.Time `gorm:"column:UpdateTime"`
	Cs          	 string    `gorm:"column:Cs";json:"_"`
	Id               int64     `gorm:"column:id"`
	Type             string   `gorm:"column:Type"`
	PId				 string    `gorm:"column:PId"`
}

func (*EvaAppDetail) TableName() string {
	return "eva_appdetail"
}

type EvaDetailItem struct {
	Id                 int    `gorm:"column:Id;primary_key"`
	EvaId              int64  `gorm:"column:EvaId"`
	EvaTargetTypeId    int    `gorm:"column:EvaTargetTypeId"`
	EvaDimensionItemId int64  `gorm:"column:EvaDimensionItemId"`
	EvaDimensionId     int64  `gorm:"column:EvaDimensionId"`
	EvaDimensionName   string `gorm:"column:EvaDimensionName"`
	LevelGrade         int32  `gorm:"column:LevelGrade"`
	LevelDesc          string `gorm:"column:LevelDesc"`
}

func (*EvaDetailItem) TableName() string {
	return "eva_detailitem"
}

//评价纬度表
type EvaDimension struct {
	EvaDimensionId   int64     `gorm:"column:EvaDimensionId;primary_key"`
	EvaDimensionName string    `gorm:"column:EvaDimensionName"`
	EvaLevel         int32     `gorm:"column:EvaLevel"`
	UpdateStaffId    string    `gorm:"column:UpdateStaffId"`
	UpdateStaffName  string    `gorm:"column:UpdateStaffName"`
	UpdateTime       time.Time `gorm:"column:UpdateTime"`
}

func (*EvaDimension) TableName() string {
	return "eva_dimension"
}

// 评价纬度子表
type EvaDimensionitem struct {
	EvaDimensionItemId int64     `gorm:"column:EvaDimensionItemId;primary_key"`
	EvaDimensionId     int64     `gorm:"column:EvaDimensionId"`
	LevelGrade         int32     `gorm:"column:LevelGrade"`
	LevelDesc          string    `gorm:"column:LevelDesc"`
	UpdateStaffId      string    `gorm:"column:UpdateStaffId"`
	UpdateStaffName    string    `gorm:"column:UpdateStaffName"`
	UpdateTime         time.Time `gorm:"column:UpdateTime"`
}

func (*EvaDimensionitem) TableName() string {
	return "eva_dimensionitem"
}

func testRead() []byte {
	fp, err := os.OpenFile("/Users/aallenma/evalueateId.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 10000)
	n, err := fp.Read(data)
	if err != nil {

	}
	fmt.Println(string(data[:n]))
	return data[:n]
}

func GetDimensionItem(db *gorm.DB, targetId int64, desc string, grade int) *EvaDimensionitem {
	sql := "SELECT A.* FROM eva_dimensionitem A JOIN eva_dimension B ON A.EvaDimensionId=B.EvaDimensionId JOIN eva_dimensionrelation C ON C.EvaDimensionId=B.EvaDimensionId WHERE C.EvaTargetId=? AND A.LevelDesc=? AND A.LevelGrade=?;"

	di := EvaDimensionitem{}
	err := db.Raw(sql, targetId, desc, grade).Find(&di).Error
	if err != nil {
		return nil
	}
	return &di
}

func GetDimension(db *gorm.DB, id int64) *EvaDimension {
	d := EvaDimension{
		EvaDimensionId: id,
	}
	err := db.Where(&d).Find(&d).Error
	if err != nil {
		return nil
	}
	return &d
}

func contains(source []string, target string) bool {
	for _, item := range source {
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	//bs := testRead()
	//var data []Temp
	//err := json.Unmarshal(bs, &data)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//user := "liuxu"
	//password :="Airport@Shenz9102!"
	//host := "dcdbt-4jcvcwra.sql.tencentcdb.com:4"
	user := "root"
	password :="123456"
	host := "localhost"
	database := "airport"
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database) +
		"?clientFoundRows=false&parseTime=true&loc=Local&timeout=5s&charset=utf8&collation=utf8_general_ci"
	db, err := gorm.Open("mysql", source)
	if err != nil {
		fmt.Println("failed to connect mysql server")
	}
	defer db.Close()

	// 禁用表名的复数形式
	db.SingularTable(true)
	db.LogMode(false)
	db = db.Set("gorm:table_options", "engine=InnoDB charset=UTF8")

	//sql := "insert into `airport`.`temp` (pid, evalueateId) values (?, ?)"
	//for _, item := range data {
	//	i, _ := strconv.Atoi(item.EvalueateId)
	//	err1 := db.Exec(sql, item.PId, i)
	//	if err1 != nil {
	//		fmt.Println(err1)
	//	}
	//}

	//area := EvaFeedbackArea{}
	//db.First(&area)
	//fmt.Println(area)
	//
	//res := &[]EvaFeedbackArea{}
	//dbRes := db.Where("TypeId=? and SceneId=?", 1, 4).Find(&res)
	//if dbRes.Error != nil {
	//	fmt.Println(dbRes.Error)
	//}
	//resByte, err := json.Marshal(*res)
	//str := string(resByte)
	//if err == nil {
	//	fmt.Println(str)
	//}

	//di := GetDimensionItem(db, "候机环境好")
	//d := GetDimension(db, di.EvaDimensionId)
	//fmt.Println(d)

	//sql := "SELECT  A.id, A.parent_id, C.EvaStageId, C.EvaStageName, A.user_id AS 'UserId', A.grade AS 'LevelGrade' ,D.LevelAppDesc AS 'LevelDesc' , A.`comment` AS 'Comment', FROM_UNIXTIME(A.created) AS 'EvaTime', FROM_UNIXTIME(A.created) AS 'UpdateTime', A.flight_number AS 'FlightNum', A.fix_comments AS 'Cs' FROM service_evaluation A JOIN temp B on A.parent_id=B.pid JOIN eva_stage C ON C.EvaStageId=B.evalueateId JOIN eva_level D ON D.LevelGrade=A.grade;"
	leftSql := "SELECT  A.id, A.parent_id AS 'PId', A.Type, C.EvaStageId, C.EvaStageName, A.user_id AS 'UserId', A.grade AS 'LevelGrade' ,D.LevelAppDesc AS 'LevelDesc' , A.`comment` AS 'Comment', FROM_UNIXTIME(A.created) AS 'EvaTime', FROM_UNIXTIME(A.created) AS 'UpdateTime', A.flight_number AS 'FlightNum', A.fix_comments AS 'Cs' FROM service_evaluation A LEFT JOIN temp B on A.parent_id=B.pid LEFT JOIN eva_stage C ON C.EvaStageId=B.evalueateId LEFT JOIN eva_level D ON D.LevelGrade=A.grade;"
	evas := []*EvaAppDetail{}
	err = db.Raw(leftSql).Find(&evas).Error
	if err != nil {
		fmt.Println(err)
	}

	unknowAves := []string{}
	unknownComments := []string{}
	unknownTotal := 0
	total := 0
	evasInsert := []*EvaAppDetail{}
	evaItems := []*EvaDetailItem{}
	for _, item := range evas {
		if item.Type == "careService" || (item.PId == "0" && item.Type != "personal") {
			t := fmt.Sprint("未识别评价：", item.Id, "   ", item.Type, "   ",item.PId)
			unknowAves = append(unknowAves, t)
			// 丢弃
			continue
		}

		if item.EvaStageName == "" {
			if  item.Type == "personal" {
				item.EvaStageId = 1255477804582047749
				item.EvaStageName = "个人"
				if item.Comment == "" && item.Cs != "" {
					item.Comment = item.Cs
				} else if item.Comment != "" && item.Cs != "" {
					item.Comment = item.Comment + "," + item.Cs
				}
			} else {
				t := fmt.Sprint("未识别评价：", item.Id, "   ", item.Type, "   ",item.PId)
				unknowAves = append(unknowAves, t)
				// 丢弃
				continue
			}
		}

		item.EvaId = uid.Generate()
		item.EvaAppSourceId = 2
		item.EvaAppSourceName = "小程序"
		evasInsert = append(evasInsert, item)

		//eva item
		cs := item.Cs
		if item.Type == "personal" || cs == "" {
			continue
		}
		cArr := strings.Split(cs, ",")

		for _, c := range cArr {

			if c == "洗手间配套全/整洁" && (item.EvaStageId ==1255480244882640900 || item.EvaStageId ==1255480244882640897){
				c = "洗手间配套全，整洁"
			}

			if c == "洗手间脏乱/配套不全" && (item.EvaStageId ==1255480244882640900 ) {
				c = "洗手间配套不全、不整洁"
			}

			if c == "室内导航清晰易用" && (item.EvaStageId ==1255479197191639044 ) {
				c = "登机口导航体验好"
			}

			if c == "交通推荐不精准" && (item.EvaStageId ==1255479197191639040 ) {
				c = "交通推荐不方便"
			}

			if c == "行李有损坏/丢失" && (item.EvaStageId ==1255479197191639042 ) {
				c = "行李损坏/丢失"
			}

			if c == "餐饮选择丰富、口味好" && (item.EvaStageId ==1255479678706126849 ) {
				c = "餐饮选择丰富/口味好"
			}



			di := GetDimensionItem(db, item.EvaStageId, c, item.LevelGrade)
			if di == nil {
				unknownTotal++
				temp := fmt.Sprintf("%d %s %d %s", item.EvaStageId, item.EvaStageName, item.LevelGrade, c)
				if !contains(unknownComments, temp) {
					unknownComments = append(unknownComments, temp)
				}
				fmt.Println("找不到：", item.EvaStageId, " ", item.EvaStageName, " ", item.LevelGrade, " ", c, " ",item.Id)
				continue
			}
			total++
			d := GetDimension(db, di.EvaDimensionId)

			evaItem := &EvaDetailItem{
				Id:                 0,
				EvaId:              item.EvaId,
				EvaTargetTypeId:    1,
				EvaDimensionItemId: di.EvaDimensionItemId,
				EvaDimensionId:     di.EvaDimensionId,
				EvaDimensionName:   d.EvaDimensionName,
				LevelGrade:         di.LevelGrade,
				LevelDesc:          di.LevelDesc,
			}

			evaItems = append(evaItems, evaItem)
		}
	}
	fmt.Println()
	fmt.Println("原总评价条数：", len(evas), " 可识别条数：", len(evasInsert) ," 可识别总纬度条数: ", len(evaItems))
	fmt.Println()
	for _, item := range unknowAves {
		fmt.Println(item)
	}
	fmt.Println()
	fmt.Println("不识别维度总数： ", unknownTotal)
	fmt.Println("不识别维度评语: ")
	for _, v := range unknownComments {
		fmt.Println(v)
	}
	fmt.Println()

	//sql := "INSERT INTO `airport`.`eva_appdetail`(`EvaId`, `EvaTime`, `LevelGrade`, `LevelDesc`, `UserId`, `FlightNum`, `EvaStageId`, `EvaStageName`, `EvaAppSourceId`, `EvaAppSourceName`, `Comment`, `UpdateTime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, 2, '小程序', ?, ?);"
	for _, eva := range evasInsert {
		err := db.Omit("id", "Cs", "Type", "PId").Create(eva).Error
		//err := db.Raw(sql, eva.EvaId, eva.EvaTime, eva.LevelGrade, eva.LevelDesc, eva.UserId, eva.FlightNum, eva.EvaStageId, eva.EvaStageName, eva.Comment, eva.UpdateTime).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, evaItem := range evaItems {
		err := db.Create(evaItem).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	//b1, _ := json.Marshal(evas)
	//evaStr := string(b1)
	//fmt.Println(evaStr)
	//
	//b2, _ := json.Marshal(evaItems)
	//evaItemStr := string(b2)
	//fmt.Println(evaItemStr)

}
