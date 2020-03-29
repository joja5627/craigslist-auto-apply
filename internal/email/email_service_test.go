package email

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShouldInitEmailService(t *testing.T) {
	emailServiceImpl := NewEmailService("joja5627")
	fmt.Print(emailServiceImpl.Srv.BasePath)
	require.NotNil(t,emailServiceImpl)
}
func TestShouldFindSentEmail(t *testing.T) {

}
