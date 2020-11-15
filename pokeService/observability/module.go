package observability

import "go.uber.org/fx"

// Module is the observability module
var Module = fx.Provide(
	NewLogger,
)
