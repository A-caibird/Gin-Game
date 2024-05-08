package RabbitMQ

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"testing"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Test1 send message to RabbitMQ and receive in this function
func Test1(t *testing.T) {
	// 连接 RabbitMQ 服务器
	conn, err := InitAmpq()
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 持久性标志
		false,   // 自动删除标志
		false,   // 独占标志
		false,   // 不等待服务器响应标志
		nil,     // 额外的属性
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息到队列
	body := "Hello, RabbitMQ!"
	err = ch.Publish(
		"",     // 交换机
		q.Name, // 队列名称
		false,  // 强制标志
		false,  // 立即标志
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Println("消息发送成功！")

	// 接收消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者名称
		true,   // 自动应答标志
		false,  // 独占标志
		false,  // 不等待服务器响应标志
		false,  // 额外参数
		nil,    // 额外属性
	)
	failOnError(err, "Failed to register a consumer")

	// 在 goroutine 中异步接收消息
	go func() {
		for d := range msgs {
			fmt.Printf("接收到消息: %s\n", d.Body)
		}
	}()

	// 等待接收消息
	fmt.Println("正在等待接收消息...")
	fmt.Println("按任意键退出")
	fmt.Scanln()
}

// TestSendMessage send message to rabbitmq, web RabbitMQ client(use RabbitMQ Web STOMP Plugin) receive message
func TestSendMessage(t *testing.T) {
	conn, err := InitAmpq()
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	ch2, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch2.Close()

	queueName := "tasks"
	done := make(chan int)
	go func() {
		for {
			message := "something to do"
			err := ch2.Publish("", queueName, false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
			if err != nil {
				log.Println("Failed to send message:", err)
				break
			} else {
				log.Println("Message sent:", message)
			}

			time.Sleep(1 * time.Second)
		}
		close(done)
	}()
	<-done
}
