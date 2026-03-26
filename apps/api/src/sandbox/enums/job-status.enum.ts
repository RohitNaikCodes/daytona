/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 *
 * JobStatus is generated from apps/runner/specs/runner.proto.
 * Only the real application values are re-exported here so that consumers
 * (e.g. TypeORM @Column enum:) are not exposed to proto sentinel values
 * (JOB_STATUS_UNSPECIFIED, UNRECOGNIZED) that must never appear in the DB.
 */
import { JobStatus as _JobStatus } from '@daytonaio/runner-specs'

export enum JobStatus {
  PENDING = _JobStatus.PENDING,
  IN_PROGRESS = _JobStatus.IN_PROGRESS,
  COMPLETED = _JobStatus.COMPLETED,
  FAILED = _JobStatus.FAILED,
}
