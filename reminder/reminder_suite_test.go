package reminder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestReminder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reminder Suite")
}
