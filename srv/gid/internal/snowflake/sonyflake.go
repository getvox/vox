package snowflake

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// These constants are the bit lengths of Snowflake ID parts.
const (
	BitLenTime      = 32 // bit length of time
	BitLenMachineID = 5  // bit length of machine id
	BitLenSequence  = 16 // bit length of sequence number
)

type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

// Snowflake is a distributed unique ID generator.
type Snowflake struct {
	mutex       *sync.Mutex
	startTime   int64
	elapsedTime int64
	sequence    uint16
	machineID   uint16
}

func NewSnowflake(st Settings) *Snowflake {
	sf := new(Snowflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<BitLenSequence - 1)

	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSnowflakeTime(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC))
	} else {
		sf.startTime = toSnowflakeTime(st.StartTime)
	}

	var err error
	if st.MachineID == nil {
		// sf.machineID, err = lower16BitPrivateIP()
		return nil
	} else {
		sf.machineID, err = st.MachineID()
	}
	if err != nil || (st.CheckMachineID != nil && !st.CheckMachineID(sf.machineID)) {
		return nil
	}

	return sf
}

func (sf *Snowflake) NextId() int64 {
	const maskSequence = uint16(1<<BitLenSequence - 1)

	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	current := currentElapsedTime(sf.startTime)
	if sf.elapsedTime < current {
		sf.elapsedTime = current
		// 初始序号0-9之间，增加一定的随机性
		sf.sequence = uint16(rand.Intn(10))
	} else { // sf.elapsedTime >= current
		sf.sequence = (sf.sequence + 1) & maskSequence
		if sf.sequence == 0 {
			sf.elapsedTime++
			overtime := sf.elapsedTime - current
			time.Sleep(sleepTime(overtime))
		}
	}

	return sf.toID()
}

const snowflakeTimeUnit = 1e9 // nsec, i.e. 1 sec

func toSnowflakeTime(t time.Time) int64 {
	return t.UTC().UnixNano() / snowflakeTimeUnit
}

func currentElapsedTime(startTime int64) int64 {
	return toSnowflakeTime(time.Now()) - startTime
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*snowflakeTimeUnit*time.Nanosecond -
		time.Duration(time.Now().UTC().UnixNano()%snowflakeTimeUnit)*time.Nanosecond
}

func (sf *Snowflake) toID() int64 {
	if sf.elapsedTime >= 1<<BitLenTime {
		log.Panic("over the time limit")
	}

	return sf.elapsedTime<<(BitLenSequence+BitLenMachineID) |
		int64(sf.machineID)<<BitLenSequence |
		int64(sf.sequence)
}
