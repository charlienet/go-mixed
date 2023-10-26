package idgenerator

import (
	"testing"
)

func TestBinary(t *testing.T) {
	f, _ := newBinaryFormatter(DefaultStartTimeStamp, 16, 12)
	t.Log(f.maxinalMachineCode, f.maximalSequence)
}

func TestDecimal(t *testing.T) {
	f, _ := newDecimalFormater(YYYYMMDDHHmmss, 4, 1)
	t.Log(f.maxinalMachineCode, f.maxinalMachineCode)
	t.Log(f.Format(22333, 9, false))
}

func TestDecimalMonth111(t *testing.T) {
	f, _ := newDecimalFormater(YYYYMMDD, 4, 1)
	t.Log(f.maxinalMachineCode, f.maxinalMachineCode)

	t.Log(f.Format(233, 9, false))
}

func TestBinaryTimestamp(t *testing.T) {
	f, _ := newBinaryFormatter(DefaultStartTimeStamp, 10, 4)
	for i := 0; i < 100; i++ {
		if i%7 == 0 {
			t.Log(f.Format(int64(i), 0xF, true))
		} else {
			t.Log(f.Format(int64(i), 0xF, false))
		}
	}
}
