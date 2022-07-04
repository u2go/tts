package tts

import (
	"github.com/u2go/u2utils"
	"testing"
)

func TestTtsSpeak(t *testing.T) {
	u2utils.PanicOnError(TtsSpeak("Hello 中文语音测试", "tmp", "zh"))
}
