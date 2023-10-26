package idgenerator

import (
	"errors"
	"math"
	"strconv"
	"time"
)

const (
	YYYYMMDDHHmmss = "20060102150405"
	YYYYMMDDHHmm   = "200601021504"
	YYYYMMDDHH     = "2006010215"
	YYYYMMDD       = "20060102"
)

type Layout int

const (
	Binary Layout = iota
	Decimal
)

type formater interface {
	MaxinalMachineCode() int64
	MaximalSequence() int64
	Format(machine, serial int64, reback bool) Id
}

type Id int64

func (i Id) String() string {
	if i == 0 {
		return ""
	}

	return strconv.FormatInt(int64(i), 10)
}

func (i Id) Int64() int64 {
	return int64(i)
}

// 标识生成器
type generator struct {
	maximalSequence    int64        // 序列段的最大值
	maxinalMachineCode int64        // 机器码的最大值
	sequenceLength     int64        // 序列段长度
	machineCodeLength  int64        // 机器码长度
	lastTimestamp      int64        // 最后回绕时间
	getTimestampFunc   func() int64 // 获取时间段的方法
}

func (g generator) MaxinalMachineCode() int64 {
	return g.maxinalMachineCode
}

func (g generator) MaximalSequence() int64 {
	return g.maximalSequence
}

func (g *generator) getTimestamp(reback bool) int64 {
	newTimestamp := g.getTimestampFunc()

	for reback && g.lastTimestamp == newTimestamp {
		time.Sleep(time.Microsecond * 10)
		newTimestamp = g.getTimestampFunc()
	}

	g.lastTimestamp = newTimestamp
	return newTimestamp
}

type binaryFormater struct {
	generator
}

func newBinaryFormatter(start int64, sequenceLength, machineCodeLength int64) (*binaryFormater, error) {
	return &binaryFormater{
		generator: generator{
			maximalSequence:    int64(-1 ^ (-1 << sequenceLength)),
			maxinalMachineCode: int64(-1 ^ (-1 << machineCodeLength)),
			sequenceLength:     sequenceLength,
			machineCodeLength:  machineCodeLength,
			getTimestampFunc:   func() int64 { return time.Now().Unix() - start },
		},
	}, nil
}

func (f *binaryFormater) Format(machine, serial int64, reback bool) Id {
	timestamp := f.getTimestamp(reback)
	return Id(timestamp<<(f.sequenceLength+f.machineCodeLength) | machine<<f.sequenceLength | serial)
}

type decimalFormater struct {
	generator
	supportReback bool
}

const (
	decimalMaxLength = 19
)

func newDecimalFormater(format string, serialLength, machineLength int) (*decimalFormater, error) {
	if len(format)+serialLength+machineLength > decimalMaxLength {
		return nil, errors.New("the data length is out of limit")
	}

	serialShift := int64(math.Pow10(serialLength))
	machineShift := int64(math.Pow10(machineLength))

	return &decimalFormater{
		generator: generator{
			sequenceLength:     serialShift,
			maximalSequence:    serialShift - 1,
			machineCodeLength:  machineShift,
			maxinalMachineCode: machineShift - 1,
			getTimestampFunc: func() int64 {
				now := time.Now()
				v := now.Format(format)
				r, _ := strconv.ParseInt(v, 10, 64)
				return r
			},
		},
		supportReback: len(format) == 14,
	}, nil
}

func (f *decimalFormater) Format(machine, serial int64, reback bool) Id {
	timestamp := f.getTimestamp(reback)
	return Id(timestamp*f.sequenceLength*f.machineCodeLength + machine*f.sequenceLength + serial)
}
