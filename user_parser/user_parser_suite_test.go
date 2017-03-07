package user_parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUserParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserParser Suite")
}
