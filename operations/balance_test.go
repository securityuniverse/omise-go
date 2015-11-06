package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	client := testutil.NewTestClient(t)

	balance := &omise.Balance{}
	client.MustDo(balance, &RetrieveBalance{})

	a.Equal(t, balance.Object, "balance")
	testutil.LogObj(t, balance)
}
