package snowflake

import (
	"sync"
	"time"
)

// 雪花算法默认起始时间 2020-01-01
const defaultStarTimestamp = 1579536000

const (
	MachineIdBits = uint(8)  //机器id所占的位数
	SequenceBits  = uint(12) //序列所占的位数
	//MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits)) //支持的最大机器id数量
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits)) //
	MachineIdShift = SequenceBits                     //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits     //时间戳左移位数
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
	return &snowflake{
		startTimeStamp: defaultStarTimestamp,
		machineId:      machineId,
	}
}

func (s *snowflake) GetId() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & SequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	r := (now-s.startTimeStamp)<<TimestampShift | (s.machineId << MachineIdShift) | (s.sequence)
	return r
}
