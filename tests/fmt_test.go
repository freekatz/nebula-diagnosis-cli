package tests

import (
	"fmt"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/info/physical"
	"testing"
)

func TestPhysicalInfoFmt(t *testing.T) {
	phyInfo := new(physical.PhyInfo)
	fmt.Printf("%+v", *phyInfo)
}
