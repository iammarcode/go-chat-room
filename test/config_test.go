package main

import (
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/initialize"
	"testing"
)

func TestConfig(t *testing.T)  {
	initialize.LoadConfig("../")

	expectedRuntime := "runtime/"
	loadedRuntime := global.Config.App.RuntimeRootPath

	if loadedRuntime != expectedRuntime {
		t.Fatalf(`expectedRuntime = %v want match for %v`, expectedRuntime, loadedRuntime)
	}
}
