/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

package executor

import (
	"context"
	"errors"

	"github.com/daytonaio/runner/pkg/api/dto"
	"github.com/daytonaio/runner/pkg/runner/v2/specs"
	specsgen "github.com/daytonaio/runner/pkg/runner/v2/specs/gen"
)

func (e *Executor) buildSnapshot(ctx context.Context, job *specsgen.Job) (any, error) {
	var p specsgen.BuildSnapshotPayload
	if err := specs.ParsePayload(job.Payload, &p); err != nil {
		return nil, err
	}

	request := specs.BuildSnapshotPayloadToDTO(&p)

	if err := e.docker.BuildSnapshot(ctx, request); err != nil {
		return nil, err
	}

	info, err := e.docker.GetImageInfo(ctx, request.Snapshot)
	if err != nil {
		return nil, err
	}

	return dto.SnapshotInfoResponse{
		Name:       request.Snapshot,
		SizeGB:     float64(info.Size) / (1024 * 1024 * 1024),
		Entrypoint: info.Entrypoint,
		Cmd:        info.Cmd,
		Hash:       dto.HashWithoutPrefix(info.Hash),
	}, nil
}

func (e *Executor) pullSnapshot(ctx context.Context, job *specsgen.Job) (any, error) {
	var p specsgen.PullSnapshotPayload
	if err := specs.ParsePayload(job.Payload, &p); err != nil {
		return nil, err
	}

	request := specs.PullSnapshotPayloadToDTO(&p)

	if err := e.docker.PullSnapshot(ctx, request); err != nil {
		return nil, err
	}

	info, err := e.docker.GetImageInfo(ctx, request.Snapshot)
	if err != nil {
		return nil, err
	}

	return dto.SnapshotInfoResponse{
		Name:       request.Snapshot,
		SizeGB:     float64(info.Size) / (1024 * 1024 * 1024),
		Entrypoint: info.Entrypoint,
		Cmd:        info.Cmd,
		Hash:       dto.HashWithoutPrefix(info.Hash),
	}, nil
}

func (e *Executor) removeSnapshot(ctx context.Context, job *specsgen.Job) (any, error) {
	snapshotName := job.GetResourceId()
	if snapshotName == "" {
		return nil, errors.New("resource ID is required for REMOVE_SNAPSHOT")
	}

	return nil, e.docker.RemoveImage(ctx, snapshotName, true)
}

func (e *Executor) inspectSnapshotInRegistry(ctx context.Context, job *specsgen.Job) (any, error) {
	var p specsgen.InspectSnapshotInRegistryPayload
	if err := specs.ParsePayload(job.Payload, &p); err != nil {
		return nil, err
	}

	digest, err := e.docker.InspectImageInRegistry(ctx, p.GetSnapshot(), specs.RegistryToDTO(p.GetRegistry()))
	if err != nil {
		return nil, err
	}

	return dto.SnapshotDigestResponse{
		Hash:   dto.HashWithoutPrefix(digest.Digest),
		SizeGB: float64(digest.Size) / (1024 * 1024 * 1024),
	}, nil
}
