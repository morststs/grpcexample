package main

import (
	"example/grpc_sample"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("One argument is required")
	}
	connectTarget := os.Args[1]

	var conn *grpc.ClientConn
	time.Sleep(1 * time.Second)
	conn, err := grpc.Dial(connectTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := grpc_sample.NewSampleServiceClient(conn)
	request := grpc_sample.Message{
		DoubleField:         1.2,
		FloatField:          1.2,
		Int32Field:          -1,
		Int64Field:          -1,
		Uint32Field:         1,
		Uint64Field:         1,
		Sint32Field:         -1,
		Sint64Field:         -1,
		Fixed32Field:        229,
		Fixed64Field:        257,
		Sfixed32Field:       -229,
		Sfixed64Field:       -257,
		BoolField:           true,
		StringField:         "string value",
		BytesField:          []byte("abc"),
		EnumField:           grpc_sample.Message_OK,
		RepeatedStringField: []string{"one", "two"},
		MapField:            map[string]int32{"key1": 1, "key2": 2},
		OneofField:          &grpc_sample.Message_StringOneofField{StringOneofField: "string one of value"},
	}
	response, err := c.GetData(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling GetData: %s", err)
	}
	log.Print(response)

	defer conn.Close()
}
