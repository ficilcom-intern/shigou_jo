package main

import (
	"Users/xushao/Desktop/Apps/Gos/shigou_jo/gRPC/protobuf-lesson/pb"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"log"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"projectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	//binData, err := proto.Marshal(employee)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//if err := os.WriteFile("test.bin", binData, 0666); err != nil {
	//	log.Fatalln(err)
	//}
	//
	//in, err := os.ReadFile("test.bin")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//readEmployee := &pb.Employee{}
	//err = proto.Unmarshal(in, readEmployee)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(readEmployee)
}
