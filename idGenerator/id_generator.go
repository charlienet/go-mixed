package idgenerator

import (
	"fmt"
	"math"
	"time"

	_ "unsafe"
)

// 时间段开始时间 2022-01-01
const startTimeStamp = 1640966400

const (
	MachineIdBits  = uint(8)                              //机器id所占的位数
	SequenceBits   = uint(12)                             //序列所占的位数
	MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits))    //支持的最大机器id数量
	SequenceMask   = uint64(-1 ^ (-1 << SequenceBits))    //
	MachineIdShift = uint64(SequenceBits)                 //机器id左移位数
	TimestampShift = uint64(SequenceBits + MachineIdBits) //时间戳左移位数
)

type TimePrecision int

const (
	Second TimePrecision = iota // 秒
	Minute                      // 分
	Day                         // 日
)

type Config struct {
	Machine   int
	TimeScope TimePrecision
}

type Generator struct {
	machine   uint64
	timeScope TimePrecision // 时间段精度
	timestamp uint64        // 上次生成时间
	sequence  uint64        // 上次使用序列
}

type Id struct {
	machine   uint64 // 机器标识
	Scope     int    // 时间精度
	Timestamp uint64 // 生成时间
	Sequence  uint64 // 标识序列
}

func New(cfg Config) *Generator {
	return &Generator{
		machine:   uint64(cfg.Machine),
		timeScope: cfg.TimeScope,
	}
}

func (g *Generator) Next() Id {

	now := currentTimestamp()

	if g.timestamp == now && g.sequence == 0 {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "下一个时间点")
		for now <= g.timestamp {
			// runtime.Gosched()
			now = currentTimestamp()
		}
	}

	g.timestamp = now                            // 标识生成时间
	g.sequence = (g.sequence + 1) & SequenceMask // 生成下一序列，超过最大值时返回零

	return Id{
		machine:   g.machine,
		Timestamp: g.timestamp,
		Sequence:  g.sequence,
	}
}

func (i Id) Id() uint64 {
	return i.Timestamp<<TimestampShift | i.machine<<MachineIdShift | i.Sequence
}

func (i Id) genCheck() uint64 {
	return math.MaxUint64
}

func (i Id) String() string {
	return fmt.Sprintf("Time:%d scope: %d Machine: %d Ser:%06d id:%d", i.Timestamp, i.Scope, i.machine, i.Sequence, i.Id())
}

func (g *Generator) NextId() uint64 {
	return g.Next().Id()
}

func (g *Generator) Batch(num int) []uint64 {
	ret := make([]uint64, num)

	for i := 0; i < num; i++ {
		ret[i] = g.NextId()
	}

	return ret
}

func Decode(id uint64) Id {
	return Id{}
}

func Verify(id uint64) bool {
	return false
}

func currentTimestamp() uint64 {
	return uint64(time.Now().Unix() - startTimeStamp)
}

//go:linkname runtimeNano runtime.nanotime
func runtimeNano() int64
