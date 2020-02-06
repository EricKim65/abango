package kafka

import (
	"fmt"

	e "github.com/EricKim65/abango/etc"
	"github.com/Shopify/sarama"
)

func KafkaSyncProducer(message string, topic string, conn string) (int32, int64, error) {

	kfcf := sarama.NewConfig()
	kfcf.Producer.Retry.Max = 5
	kfcf.Producer.RequiredAcks = sarama.WaitForAll
	kfcf.Producer.Return.Successes = true

	if prd, err := sarama.NewSyncProducer([]string{conn}, kfcf); err == nil {

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}
		part, offset, err := prd.SendMessage(msg)
		if err != nil {
			fmt.Println("Error publish: ", err.Error())
		}
		return part, offset, nil
	} else {
		e.MyErr("SEWQRVCVRHBS", err, true)
		return 0, 0, err
	}
}
