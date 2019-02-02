package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)
//chatbot_access_token := "84RSL9DoYB7CYToRF/7n38pidk4YqaylvtyULojYQYLSRpyOf7dHIYPGK/GwtrCzXJo0FoaDznyJ+jYiygmWc0azPyO+i1QBlqj2kNobdAoVGhGWQzFmI1pZQ8/F9IxEwLw5Heoofae684hdMDobzAdB04t89/1O/w1cDnyilFU="
//func PostLineUser(userid string)

func getEvent(c *gin.Context) {
	var result map[string]interface{}
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	json.Unmarshal(buf[0:n], &result)
	events := result["events"]
	fmt.Println(events)
	s := reflect.ValueOf(events)
	v := s.Index(0).Interface().(map[string]interface{})
	fmt.Println(v["message"])
	c.String(200, "test")
}

func main() {
	r := gin.Default()
	// Ping handler
	fmt.Print("test")
	r.POST("/getEvent", getEvent)

	log.Fatal(autotls.Run(r, "www.nctusam.com"))
}
