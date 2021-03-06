package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/streadway/amqp"
)

type RequestGetDocuments struct {
	List map[string]DocumentDAO
}

type ResponseDoc struct {
	List   map[string]DocumentDAO
	status string
}

type ResponseSaveDoc struct {
	Name string
}

type RequestSaveDoc struct {
	Name string
	File []byte
}

type RequestDeleteDoc struct {
	Id string
}

type ResponseDeleteDoc struct {
	Name string
}

/*
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}*/

func SendStorage(req RequestGetDocuments) ResponseDoc {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable  cambiar a false
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	corrId := randomString(32)

	/*user := User{ID: "1", Name: "juan", Email: "asd"}
	user2 := User{ID: "2", Name: "juan2", Email: "asd2"}
	user3 := User{ID: "3", Name: "juan3", Email: "asd3"}
	users := make([]User, 0)
	users = append(users, user, user2, user3)*/
	//-----------Initialize
	//----------
	dat := ToGOBRequestDoc(req)
	//fmt.Println("encapsulado", dat)
	err = ch.Publish(
		"",                         // exchange
		"rpc_Storage_getDocuments", // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(dat),
		})
	//log.Printf("Enviando Notificacion", dat)
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res := FromGOBResponseDoc(d.Body)
			d.Ack(true)
			//fmt.Println(res)
			return res
			// al comentar este return hace que entre en el loop forever
			//return
		}
		//return
	}

	return ResponseDoc{nil, ""}
}

func SaveStorage(name string, file []byte) string {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable  cambiar a false
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	corrId := randomString(32)

	/*user := User{ID: "1", Name: "juan", Email: "asd"}
	user2 := User{ID: "2", Name: "juan2", Email: "asd2"}
	user3 := User{ID: "3", Name: "juan3", Email: "asd3"}
	users := make([]User, 0)
	users = append(users, user, user2, user3)*/
	//-----------Initialize
	//----------
	req := RequestSaveDoc{File: file, Name: name}
	dat := ToGOBRequestSaveDoc(req)
	//fmt.Println("encapsulado saveDoc", dat)
	err = ch.Publish(
		"",                 // exchange
		"rpc_Storage_save", // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(dat),
		})
	//log.Printf("Enviando Notificacion", dat)
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res := FromGOBResponseSaveDoc(d.Body)
			d.Ack(true)
			fmt.Println("respuesta del saveDoc", res.Name)
			return res.Name
			// al comentar este return hace que entre en el loop forever
			//return
		}
		//return
	}

	return ""
}

func deleteStorage(docId string) string {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable  cambiar a false
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	corrId := randomString(32)

	/*user := User{ID: "1", Name: "juan", Email: "asd"}
	user2 := User{ID: "2", Name: "juan2", Email: "asd2"}
	user3 := User{ID: "3", Name: "juan3", Email: "asd3"}
	users := make([]User, 0)
	users = append(users, user, user2, user3)*/
	//-----------Initialize
	//----------
	req := RequestDeleteDoc{Id: docId}
	dat := ToGOBRequestDeleteDoc(req)
	//fmt.Println("encapsulado saveDoc", dat)
	err = ch.Publish(
		"",                   // exchange
		"rpc_Storage_delete", // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(dat),
		})
	//log.Printf("Enviando Notificacion", dat)
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res := FromGOBResponseDeleteDoc(d.Body)
			d.Ack(true)
			fmt.Println("respuesta del deleteDoc", res.Name)
			return res.Name
			// al comentar este return hace que entre en el loop forever
			//return
		}
		//return
	}

	return ""
}

func ToGOBRequestDoc(m RequestGetDocuments) []byte {
	//fmt.Println("convert GOB ", m)
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	gob.Register(DocumentDAO{})
	gob.Register(RequestGetDocuments{})
	err := e.Encode(m)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return b.Bytes()
}

func FromGOBResponseDoc(by []byte) ResponseDoc {
	m := ResponseDoc{}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return m
}

func ToGOBRequestSaveDoc(m RequestSaveDoc) []byte {
	//fmt.Println("convert GOB ", m)
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	gob.Register(RequestSaveDoc{})
	err := e.Encode(m)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return b.Bytes()
}

func FromGOBResponseSaveDoc(by []byte) ResponseSaveDoc {
	m := ResponseSaveDoc{}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return m
}

func ToGOBRequestDeleteDoc(m RequestDeleteDoc) []byte {
	//fmt.Println("convert GOB ", m)
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	gob.Register(RequestDeleteDoc{})
	err := e.Encode(m)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return b.Bytes()
}

func FromGOBResponseDeleteDoc(by []byte) ResponseDeleteDoc {
	m := ResponseDeleteDoc{}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return m
}

/*
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
*/
