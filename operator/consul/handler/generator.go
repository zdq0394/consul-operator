package handler

import (
	"fmt"
)

const (
	consulHeadlessPort     = 8500
	consulHeadlessPortName = "consul-port"
)

const (
	svcHeadlessNamePrefix = "ch"
	statefulsetNamePrefix = "cs"
	configMapNamePrefix   = "ccfg"
)

var (
	terminationGracePeriodSeconds int64 = 20
)

func generateName(prefix, name string) string {
	return fmt.Sprintf("%s-%s", prefix, name)
}
