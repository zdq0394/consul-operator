package handler

const (
	redisHeadlessPort     = 8500
	redisHeadlessPortName = "consul-port"
)

const (
	svcHeadlessNamePrefix = "ch"
	statefulsetNamePrefix = "cs"
	configMapNamePrefix   = "ccfg"
)

var (
	terminationGracePeriodSeconds int64 = 20
)
