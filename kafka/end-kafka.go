package kafka

import (
	"context"
	"fmt"
	"strconv"
	"time"

	e "github.com/EricKim65/abango/etc"
	g "github.com/EricKim65/abango/global"
	"github.com/Shopify/sarama"
)

//////////// Kafka EndPoint /////////////
func KafkaSyncRequest(askstr string, retTopic string) (string, string, error) {

	topic := g.XConfig["KafkaTopic"]
	conn := g.XConfig["KafkaAddr"] + ":" + g.XConfig["KafkaPort"]
	if _, _, err := KafkaSyncProducer(askstr, topic, conn); err == nil {
		TmpInt, _ := strconv.Atoi(g.XConfig["KafkaCosumerTimeout"])
		timeout := int64(TmpInt)
		if msg, err := KafkaReturnConsumer(retTopic, conn, timeout); err == nil {
			return msg, "200", nil // Normal Retrun
		} else {
			return "", "K503", err //Servuce Unavailable
		}
	} else {
		return "", "K909", err // Server or Network disconnected
	}
}

func KafkaReturnConsumer(topic string, conn string, timeout int64) (string, error) {

	kfcf := sarama.NewConfig()
	kfcf.Consumer.Return.Errors = true

	// Specify brokers address. This is default one
	brokers := []string{conn}
	master, err := sarama.NewConsumer(brokers, kfcf)
	if err != nil {
		return "", e.MyErr("DSFYODGBWSVE-Kafka Consumer Not created", err, true)
	}

	defer func() {
		if err := master.Close(); err != nil {
			return
		}
	}()

	// How to decide partition, is it fixed value...?
	// consumer, err := master.ConsumePartition(retTopic, 0, sarama.OffsetNewest)
	// 다른 놈이 읽어가고 있는 것 같다. 그래서 항상 처음 것으로 하였다.
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		e.MyErr("ConsumePartition-QERWRVZDFVNMRE", err, true)
	}

	strMsg := ""

	doneCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				strMsg = string(msg.Value)
				doneCh <- struct{}{}
			case <-ctx.Done(): // if cancel() execute
				doneCh <- struct{}{}
				return
			}
		}
	}(ctx)

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		e.OkLog("Timeout-Kafka connection after " + strconv.Itoa(int(timeout)) + " Seconds")
		cancel()
	}()

	<-doneCh

	return strMsg, nil
}
