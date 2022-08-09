package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Table struct {
	TableNo     string
	IsAvailable bool
}
type Customer struct {
	CustomerId          string
	CustomerName        string
	CustomerPhoneNumber string
}
type FnB struct {
	FnBId    string
	MenuName string
	Price    float64
}
type CustomerOrder struct {
	OrderMenu FnB
	Qty       int
}
type BillDetail struct {
	BillDetailId string
	CustomerOrder
}

type Bill struct {
	BillNo          string
	TableNo         Table
	TransactionDate time.Time
	CustomerId      Customer
	BillDetail      []BillDetail
	IsSettled       bool
}

func main() {
	// using model only
	table1 := Table{TableNo: "T01"}
	fnb1 := FnB{
		FnBId:    "F001",
		MenuName: "Nasi Goreng",
		Price:    15000,
	}
	fnb2 := FnB{
		FnBId:    "F002",
		MenuName: "Es Teh",
		Price:    3000,
	}
	customer1 := Customer{
		CustomerId:          "C001",
		CustomerName:        "Jution",
		CustomerPhoneNumber: "0891234567",
	}
	bill1 := Bill{
		BillNo:          "B001",
		TableNo:         table1,
		TransactionDate: time.Now(),
		CustomerId:      customer1,
		BillDetail: []BillDetail{
			{
				BillDetailId: "BD001", CustomerOrder: CustomerOrder{OrderMenu: fnb1, Qty: 2},
			},
			{
				BillDetailId: "BD002", CustomerOrder: CustomerOrder{OrderMenu: fnb2, Qty: 2},
			},
		},
	}
	// fmt.Println(bill1.BillDetail[1])
	fmt.Println(bill1)
	// WriteJSONFile(bill1, "Transaction")
	read := []Bill{}
	read = ReadJSONFile("/home/yurham/Documents/Golang/golang-fundamental/Data Transaction 09 Jun 22 14:42 WIB.json")
	// billRead:=[]Bill{}
	// billRead=read
	fmt.Println(read)
	for i := 0; i < len(read.CatlogNodes); i++ {
		fmt.Println("Product Id: ", read.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", read.CatlogNodes[i].Quantity)
	}
}
func WriteJSONFile(data interface{}, db string) {
	file, _ := json.MarshalIndent(data, "", " ")
	dt := time.Now()
	path := fmt.Sprintf("Data %s %v.json", db, dt.Format(time.RFC822))
	_ = ioutil.WriteFile(path, file, 0644)
}

func ReadJSONFile(filePath string) (data interface{}) {
	file, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(file), &data)
	return

}
