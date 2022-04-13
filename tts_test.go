package u2utils

import (
	"github.com/u2go/u2utils"
	"testing"
)

func TestTtsSpeak(t *testing.T) {
	u2utils.PanicOnError(TtsSpeak("1号指标下降20%", "tmp", "zh"))
}
