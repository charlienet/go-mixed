package snowflake

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 雪花算法默认起始时间 2022-01-01
const defaultStarTimestamp = 1640966400000

const (
	MachineIdBits  = uint(8)                           //机器id所占的位数
	SequenceBits   = uint(12)                          //序列所占的位数
	MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits)) //支持的最大机器id数量
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits))  //
	MachineIdShift = SequenceBits                      //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits      //时间戳左移位数
)

type SnowFlake interface {
	GetId() int64
}

type snowflake struct {
	sync.Mutex
	startTimeStamp int64
	timestamp      int64
	machineId      int64
	sequence       int64
}

func CreateSnowflake(machineId int64) SnowFlake {
	timeBits := 63 - MachineIdBits - SequenceBits
	maxTime := time.UnixMilli(defaultStarTimestamp + (int64(-1 ^ (-1 << timeBits))))
	log.Println("最大可用时间:", maxTime)

	return &snowflake{
		startTimeStamp: defaultStarTimestamp,
		machineId:      machineId & MachineIdMax,
	}
}

// 组织方式   时间戳-机器码-序列号
func (s *snowflake) GetId() int64 {

	// 生成序列号规则
	// 检查当前生成时间与上次生成时间对比
	// 如等于上次生成时间，检查是否已经达到序列号的最大值，如已达到等待下一个时间点并且设置序列号为零。
	// 如不相等则序列号自增
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixMilli()
	if s.timestamp == now && s.sequence == 0 {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "下一个时间点")
		for now <= s.timestamp {
			now = time.Now().UnixMilli()
		}
	}

	s.timestamp = now
	s.sequence = (s.sequence + 1) & SequenceMask

	log.Println("时间戳:", now-s.startTimeStamp)

	log.Println("时间差:", time.Now().Sub(time.UnixMilli(defaultStarTimestamp)))

	r := (now-s.startTimeStamp)<<TimestampShift | (s.machineId << MachineIdShift) | (s.sequence)

	return r
}
