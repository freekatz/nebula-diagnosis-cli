package tests

import (
	"fmt"
	"testing"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/info/physical"
)

func TestPhysicalInfoFmt(t *testing.T) {
	phyInfo := new(physical.PhyInfo)
	fmt.Printf("%+v", *phyInfo)
}
