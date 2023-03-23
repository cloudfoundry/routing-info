package tcp_routes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTcpRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TcpRoutes Suite")
}
