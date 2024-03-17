package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	// to produce messages
	topic := "LIANJIA"
	partition := 1

	conn, err := kafka.DialLeader(context.Background(), "tcp", "tencentcloud.yawujia.cn:9092", topic, partition)
	if err != nil {
		fmt.Println("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	bytes, err := conn.WriteMessages(
		kafka.Message{Value: []byte("聪哥")},
		kafka.Message{Value: []byte("方小姐喜欢你哦")},
		kafka.Message{Value: []byte("快跟她表白吧!")},
	)
	fmt.Println("send bytes: ", bytes)
	if err != nil {
		fmt.Println("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}

}
