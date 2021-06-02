package red_envelope

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 红包类
type RedEnvelope struct {
	RedId 		int64
	TotalMoney 	float64
	TotalSize  	int32
	RemainMoney float64
	RemainSize  int32
	SponsorId   int64
	SponsorName string
	Comment 	string
	StartTime   time.Time
	EndTime     time.Time
	ExpireTime  time.Time
	Records 	[]*LuckyRecord
	Status      int32			// 0 进行中；2 已抢完；3 已过期；4 取消
}

// 幸运人
type LuckyRecord struct {
	UserId     int64
	UserName   string
	luckyMoney float64
	luckyTime  time.Time
}

// 发布红包
func (redEnvelope *RedEnvelope) Start() *RedEnvelope {
	return redEnvelope
}

// 抢红包
func (redEnvelope *RedEnvelope) Grab() *RedEnvelope {
	return redEnvelope
}

// 查询记录
func (redEnvelope *RedEnvelope) Scan() []*LuckyRecord {
	return redEnvelope.Records
}

// 随机红包
func (redEnvelope *RedEnvelope) getLuckyMoney(remainNum int32, remainMoney float64) float64 {
	if remainMoney < 1 {
		return 0
	}
	if remainNum == 1 {
		return remainMoney
	}

	min := 0.01
	max := remainMoney / float64(remainNum) * 2.0
	lucky := rand.Float64() * max
	if lucky <= min {
		return min
	}

	lucky, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", lucky), 64)
	return lucky
}

// 审计红包（所有抢到的金额加起来了是否等于总红包金额）
func (redEnvelope *RedEnvelope) audit() bool {
	return true
}

func (redEnvelope *RedEnvelope) expire()  {

}
