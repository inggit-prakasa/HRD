package main

import (
	"context"
	"fmt"
	"log"

	"HRD/employee"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	//TIARA

	//LOGIN
	//PILIHAN STAFF
	//MENUSTAFF || MENUHRD
	//FUNGSISTAFF || HRD
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := employee.NewHRDClient(conn)

	variablelbuatloopmenu := true

	for variablelbuatloopmenu {
		fmt.Println("Laporan By Username")
		var username string
		fmt.Scan(&username)

		input := &employee.GetEmpByUsername{
			Username: username,
		}
		result := GetEmployeeByUsername(c, input)
		fmt.Println(result)
		variablelbuatloopmenu = false
	}
}

func GetEmployeeByUsername(c employee.HRDClient, input *employee.GetEmpByUsername) string {
	resp, err := c.LaporanByUsername(context.Background(), input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return "Id :" + resp.ID + "\n" +
		"Nama :" + resp.Nama + "\n" +
		"Total Gaji :" + resp.TotalGaji + "\n" +
		"Status :" + resp.Status
}
