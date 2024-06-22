package environment

type Mode string

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

var EnvMode Mode = ReleaseMode

func SetEnvMode(s Mode) {
	EnvMode = s
}

func IsDebugMode() bool {
	return EnvMode == DebugMode
}

func GetEnvMode() Mode {
	return EnvMode
}
