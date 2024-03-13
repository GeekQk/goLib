package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	IsAuth bool   `json:"is_authorized"`
}

func TestEncode(t *testing.T) {
	p := User{
		Name:   "zhangsan",
		Age:    20,
		Email:  "zhangsan@mail.com",
		IsAuth: true,
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %T ,%[1]v\n", string(b))

	userList := []User{}
	for i := 0; i < 10; i++ {
		userList = append(userList, User{
			ID:    i,
			Name:  fmt.Sprintf("%s + %d", "test", i),
			Email: "zhangsan@mail.com",
		})
	}
	b, _ = json.Marshal(userList)
	fmt.Printf("b: %T ,%[1]v\n", string(b))

}

func TestDecode(t *testing.T) {
	b := `{"id":0,"name":"zhangsan","age":20,"email":"zhangsan@mail.com","is_authorized":true}`
	var p User
	err := json.Unmarshal([]byte(b), &p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("p: %T ,%+[1]v\n", p)

	b = `[{"id":0,"name":"test + 0","age":0,"email":"zhangsan@mail.com","is_authorized":false},{"id":1,"name":"test + 1","age":0,"email":"zhangsan@mail.com","is_authorized":false},{"id":2,"name":"test + 2","age":0,"email":"zhangsan@mail.com","is_authorized":false}]`
	var list []User
	err = json.Unmarshal([]byte(b), &list)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("p: %T ,%+[1]v\n", list)

}

func TestDecodeFile(t *testing.T) {

	// 打开JSON文件
	file, err := os.Open("users.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// 创建一个新的JSON解码器
	decoder := json.NewDecoder(file)

	// 读取文件中的所有JSON数据
	var users []User
	err = decoder.Decode(&users)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// 处理JSON数据
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d, City: %d\n", user.Name, user.Age, user.ID)
	}

}

func TestEncodeFile(t *testing.T) {

	// 创建一个User实例
	user := User{
		Name: "Alice",
		Age:  30,
	}

	// 打开文件用于写入
	file, err := os.Create("user.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 创建一个新的JSON编码器
	encoder := json.NewEncoder(file)

	// 将User实例编码为JSON并写入文件
	err = encoder.Encode(user)
	if err != nil {
		fmt.Println("Error encoding user:", err)
		return
	}

	fmt.Println("User encoded and written to file successfully")
}
