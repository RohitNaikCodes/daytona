/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

package executor

import (
	"context"
	"fmt"

	"github.com/daytonaio/runner/pkg/runner/v2/specs"
	specsgen "github.com/daytonaio/runner/pkg/runner/v2/specs/gen"
)

func (e *Executor) createBackup(ctx context.Context, job *specsgen.Job) (any, error) {
	var p specsgen.CreateBackupPayload
	if err := specs.ParsePayload(job.Payload, &p); err != nil {
		return nil, fmt.Errorf("failed to parse payload: %w", err)
	}

	// TODO: is state cache needed?
	return nil, e.docker.CreateBackup(ctx, job.ResourceId, specs.CreateBackupPayloadToDTO(&p))
}
