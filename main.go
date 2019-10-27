package main

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic  = "topic/1"
	broker = "tcp://localhost:1883"
)

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("mqtt-sample-1")
	opts.SetDefaultPublishHandler(publishHandler)

	c := MQTT.NewClient(opts)
	token := c.Connect()
	handleToken(token)
	fmt.Printf("Connected to %s\n", broker)

	token = c.Subscribe(topic, 0, nil)
	handleToken(token)

	publishMessages(c, topic) // publish test messages

	token = c.Unsubscribe(topic)
	handleToken(token)

	c.Disconnect(250)
	fmt.Printf("Disonnected from %s\n", broker)
}

func publishHandler(client MQTT.Client, msg MQTT.Message) {
	go func() {
		fmt.Printf("%s: %s\n", msg.Topic(), msg.Payload())
	}()
}

func handleToken(t MQTT.Token) {
	t.Wait()
	if t.Error() != nil {
		fmt.Println(t.Error())
		os.Exit(1)
	}
}

func publishMessages(c MQTT.Client, topic string) {
	for i := 0; i < 20; i++ {
		text := fmt.Sprintf("Message #%d!", i)
		token := c.Publish(topic, 0, false, text)
		handleToken(token)
	}
}
