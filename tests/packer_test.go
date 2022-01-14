package tests

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/packer"
	"testing"
)

func TestPack(t *testing.T) {
	tgzPacker := packer.NewTgzPacker()

	// pack file
	err := tgzPacker.Pack("./tmp/pack_test/pack_file.txt", "./tmp/pack_test/pack_file.tar.gz")
	if err != nil {
		t.Error(err)
	}
	// pack folder
	err = tgzPacker.Pack("./tmp/pack_test/pack_folder", "./tmp/pack_test/pack_folder.tar.gz")
	if err != nil {
		t.Error(err)
	}
}
