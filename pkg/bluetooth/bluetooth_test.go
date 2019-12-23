package bluetooth

import (
	"github.com/andyinabox/go-dumbphone/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBluetoothSend(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	f, err := utils.CreateTempFile("")
	require.Nil(err)

	if testing.Verbose() {
		err = Send(f)
		assert.Nil(err, "File should send without error")
	}
}
