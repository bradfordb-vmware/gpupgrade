// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package agent

import (
	"context"

	"github.com/greenplum-db/gp-common-go-libs/gplog"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils"
)

func (s *Server) ArchiveLogDirectory(ctx context.Context, in *idl.ArchiveLogDirectoryRequest) (*idl.ArchiveLogDirectoryReply, error) {
	gplog.Info("agent starting %s", idl.Substep_ARCHIVE_LOG_DIRECTORIES)

	err := utils.System.Rename(in.GetOldDir(), in.GetNewDir())
	return &idl.ArchiveLogDirectoryReply{}, err
}
