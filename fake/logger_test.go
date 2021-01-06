package fake

import (
	"testing"

	"github.com/Sholanki/logger"
)

func Test_Fake_Interface(t *testing.T) {
	var _ logger.Interface = New()
}
