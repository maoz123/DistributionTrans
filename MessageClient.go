package main

import (
	"context"
	"fmt"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func GetClient() *azservicebus.Client {
	connectionString := "Endpoint=sb://ctxwsd-cloudlicenseactiveuse-eastus-jarvis.servicebus.windows.net/;SharedAccessKeyName=Manage;SharedAccessKey=uyKycRypzES4yalRHzzp1GCwofWnYq+QgMk+FcqICF4="

	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func SendMessage(message string, client *azservicebus.Client) {
	sender, err := client.NewSender("testqueue", nil)
	if err != nil {
		panic(err)
	}
	defer sender.Close(context.TODO())

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic(err)
	}
}

func GetMessage(count int, client *azservicebus.Client) {
	receiver, err := client.NewReceiverForQueue("testqueue", nil) //Change myqueue to env var
	if err != nil {
		panic(err)
	}
	defer receiver.Close(context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), count, nil)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		body := message.Body
		if err != nil {
			panic(err)
		}
		str := string(body)
		order := Order{}
		json.Unmarshal([]byte(str), &order)

		ReductStorage(order)
		InsertEvent(order)
		fmt.Println("order process succeed")
		//fmt.Printf("%s\n", string(body))

		err = receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			panic(err)
		}
	}
}