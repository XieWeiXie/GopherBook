package chapter4

import "testing"

func TestClientUsage(t *testing.T) {
	ClientUsage()
	UserClientUsage()
}

func TestServerUsage(t *testing.T) {
	ServerUsage()
}
