package core

import "github.com/goverland-labs/platform-events/events"

const (
	SubjectEnsResolverRequest  = "core.ens.request"
	SubjectEnsResolverResolved = "core.ens.resolved"
)

type EnsNamePayload struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type EnsNamesPayload []EnsNamePayload

type EnsNamesHandler = events.Handler[EnsNamesPayload]
