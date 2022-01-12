package tests

import (
	"fmt"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/remote"
	"log"
	"testing"
)

func TestGetNebulaMetrics(t *testing.T) {
	ip := "192.168.8.49"
	var port int32 = 19669
	metrics, err := remote.GetNebulaMetrics(ip, port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(metrics)
}

func TestGetNebulaComponentStatus(t *testing.T) {
	ip := "192.168.8.49"
	var port int32 = 19669
	metrics, err := remote.GetNebulaComponentStatus(ip, port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(metrics)
}

func TestGetNebulaFlags(t *testing.T) {
	ip := "192.168.8.49"
	var port int32 = 19669
	metrics, err := remote.GetNebulaFlags(ip, port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(metrics)
}