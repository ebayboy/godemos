package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//var Address = []string{"localhost:9092"}
var Address = []string{"tpaas-kafka-scdn.jdcloud.com:9092"}
var topic = []string{"deeplog-lbwaf-source-cloud"}
var groupId = "waf-cloud1610962377985"
var clientID = "C89f1f46c"
var clientPasswd = "eXktKroFrDUF9S3B"

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(2)

	//广播式消费：消费者1
	go clusterConsumer(wg, Address, topic, groupId)

	wg.Wait()
}

// 支持brokers cluster的消费者
func clusterConsumer(wg *sync.WaitGroup, brokers, topics []string, groupId string) {
	defer wg.Done()
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.ClientID = clientID
	//passwd

	// init consumer
	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupId, err)
		return
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("%s:Error: %s\n", groupId, err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("%s:Rebalanced: %+v \n", groupId, ntf)
		}
	}()

	// consume messages, watch signals
	var successes int
Loop:
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s:%s/%d/%d\t%s\t%s\n", groupId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
				successes++
			}
		case <-signals:
			break Loop
		}
	}
	fmt.Fprintf(os.Stdout, "%s consume %d messages \n", groupId, successes)
}
