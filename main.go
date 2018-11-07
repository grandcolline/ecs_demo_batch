package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("--- START Slack")

	// インスタンスIDの取得
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, _ := svc.GetInstanceIdentityDocument()
	instanceId := doc.InstanceID
	// コンテナIDの取得
	containerId, _ := os.Hostname()
	// タスクの取得
	resp, err := http.Get(os.Getenv("ECS_CONTAINER_METADATA_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var metadata interface{}
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		log.Fatal(err)
	}
	taskArn := metadata.(map[string]interface{})["Labels"].(map[string]interface{})["com.amazonaws.ecs.task-arn"].(string)
	task := strings.Split(taskArn, "/")[1]

	// slack投げる
	url := os.Getenv("WEBHOOK_URL")
	name := os.Getenv("NAME")
	text := "instanceId: " + instanceId + "\ntask: " + task + "\ncontainerId: " + containerId
	channel := os.Getenv("CHANNEL")
	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	fmt.Println("--- END Slack")
}
