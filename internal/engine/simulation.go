package engine

import (
	"context"
	"fmt"

	"github.com/nijeti/graphics/internal/simulation"
)

func (e *Engine) initSimulation(ctx context.Context) error {
	s, err := simulation.Init(textureWidth, textureHeight)
	if err != nil {
		return fmt.Errorf("failed to initialize simulation: %w", err)
	}

	e.simulation = s

	e.logger.DebugContext(ctx, "simulation initialized")

	return nil
}
