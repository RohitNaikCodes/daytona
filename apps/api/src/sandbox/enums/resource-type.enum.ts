/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 *
 * ResourceType is generated from apps/runner/specs/runner.proto.
 * Only the real application values are re-exported here so that consumers
 * (e.g. TypeORM @Column enum:) are not exposed to proto sentinel values
 * (RESOURCE_TYPE_UNSPECIFIED, UNRECOGNIZED) that must never appear in the DB.
 */
import { ResourceType as _ResourceType } from '@daytonaio/runner-specs'

export enum ResourceType {
  SANDBOX = _ResourceType.SANDBOX,
  SNAPSHOT = _ResourceType.SNAPSHOT,
  BACKUP = _ResourceType.BACKUP,
}
