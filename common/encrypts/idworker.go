package encrypts

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"m-sec/common/errs"
	"math/big"
	"sync"
	"time"
)

const (
	epoch              int64 = 1679892966                  // 开始时间戳，自定义
	workerIDBits       uint8 = 10                          // 机器ID所占位数
	sequenceBits       uint8 = 12                          // 序列号所占位数
	workerIDMax        int64 = -1 ^ (-1 << workerIDBits)   // 机器ID最大值
	sequenceMax        int64 = -1 ^ (-1 << sequenceBits)   // 序列号最大值
	workerIDShift      uint8 = sequenceBits                // 机器ID左移位数
	timestampLeftShift uint8 = sequenceBits + workerIDBits // 时间戳左移位数
)

// Worker节点

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	sequence  int64
}

// 快乐的Worker节点，使用随机数生成伪装worker节点

func FunnyWorkerId() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(1024))
	if err != nil {
		return 1
	}
	return n.Int64() + 1
}

// 新建一个Worker节点

func NewWorker() (*Worker, *errs.BError) {
	workerID := FunnyWorkerId()
	if workerID < 0 || workerID > workerIDMax {
		return nil, errs.WorkerIdExcessOfQuantity
	}
	return &Worker{
		timestamp: 0,
		workerID:  workerID,
		sequence:  0,
	}, nil
}

// ID生成器

func (w *Worker) NextID() (int64, *errs.BError) {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.sequence = (w.sequence + 1) & sequenceMax
		if w.sequence == 0 {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.sequence = 0
	}
	if now < w.timestamp {
		return 0, errs.ClockMovedBackwards
	}
	w.timestamp = now
	id := (now-epoch)<<timestampLeftShift | w.workerID<<workerIDShift | w.sequence
	return id, nil
}

// ID - Hash 生成方法，M-SEC混合加密

func LayIDHash() string {
	node, err := NewWorker()
	if err != nil {
		return GetRandHex()
	}

	// 获取伪装Worker节点生成的雪花❄ID
	snowflakeId, _ := node.NextID()

	// 获取混淆加密串
	var binaryInt uint64
	_ = binary.Read(rand.Reader, binary.LittleEndian, &binaryInt)

	// 异或运算
	hashId := int64(binaryInt) ^ snowflakeId

	// MD5生成加密字符串
	s := fmt.Sprintf("%d", hashId)
	hash := md5.New()
	hash.Write([]byte(s))
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}

// 防止加密失败的最终策略

func GetRandHex() string {
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)
	hashString := hex.EncodeToString(randomBytes)
	return hashString
}
