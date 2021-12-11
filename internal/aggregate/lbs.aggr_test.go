package aggregate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lbsAggr_GeoCoder(t *testing.T) {
	lbs := newLbs()
	res, err := lbs.GeoCoder(context.Background(), "浙江省杭州市西湖区")

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
