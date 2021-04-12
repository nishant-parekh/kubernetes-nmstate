package state

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
)

func TestUnit(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.state-state_suite_test.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Helper Test Suite", []Reporter{junitReporter})
}
