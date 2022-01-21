package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/nebula/nebula-diagnose/pkg/remote"
)

func TestGetNebulaMetrics(t *testing.T) {
	ip := "192.168.8.49"
	var port int = 19669
	metrics, err := remote.GetNebulaMetrics(ip, port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(metrics)
}

func TestGetNebulaComponentStatus(t *testing.T) {
	ip := "192.168.8.49"
	var port int = 19669
	metrics, err := remote.GetNebulaComponentStatus(ip, port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(metrics)
}

func TestGetNebulaFlags(t *testing.T) {
	ip := "192.168.8.49"
	var port int = 19669
	metrics, err := remote.GetNebulaFlags(ip, port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(metrics)
}
