package version

import "fmt"

var (
	GoVersion   = ""
	GitCommitId = ""
	BuildTime   = ""
)

func FullVersion() string {
	return fmt.Sprintf("GoVesion:%s, GitCommitId:%s, BuildDate:%s", GoVersion, GitCommitId, BuildTime)
}
