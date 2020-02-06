package kafka

import (
	"bytes"
	"fmt"

	cf "github.com/EricKim65/abango/config"
	e "github.com/EricKim65/abango/etc"
	g "github.com/EricKim65/abango/global"
	"github.com/Shopify/sarama"
)

//////////// Kafka Service /////////////
func KafkaServiveStandBy() {

	kfcf := sarama.NewConfig()
	kfcf.Consumer.Return.Errors = true

	conn := g.XConfig["KafkaAddr"] + ":" + g.XConfig["KafkaPort"]
	brokers := []string{conn}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, kfcf)
	if err != nil {
		e.MyErr("QRVAQEAREADVSQ-Kafka Consumer Not created", err, true)
		return
	}

	// e.OkLog("Abango-Kafka->" + conn + " Service starting !")

	defer func() {
		if err := master.Close(); err != nil {
			e.MyErr("IRJDNWRTSE-Kafka Consumer Not closed", err, true)
			return
		}
	}()

	// How to decide partition, is it fixed value...?
	topic := g.XConfig["KafkaTopic"]
	if consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest); err == nil {
		doneCh := make(chan struct{})
		go func() {
			for {
				select {
				case err := <-consumer.Errors():
					fmt.Println(err)
				case msg := <-consumer.Messages():

					valarr := bytes.Split(msg.Value, []byte(g.XConfig["MsgDelimiter"])) // a[0]=frontvars a[1]=askstr
					if err := cf.GetServerVarsInSvc(valarr[0]); err == nil {
						retTopic := g.ServerVars["unique_id"]
						if _, _, err := KafkaSyncProducer(string(valarr[0]), retTopic, conn); err == nil {
							e.OkLog("Kafka-ReturnTopic: " + retTopic)
						} else {
							e.MyErr("WERRWERFAAFHQW", err, false)
						}
					} else {
						e.MyErr("QEWRQCVZVCXER-cf.ServerVars-", err, false)
					}
				}
			}
		}()
		<-doneCh

	} else {
		e.MyErr("ConsumePartition-WERDCAYQDFARW", err, true)
	}
}

func GetFrontVars(bytevar []byte) {

}
