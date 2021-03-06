// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"

	"github.com/greenplum-db/gpupgrade/greenplum"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/upgrade"
	"github.com/greenplum-db/gpupgrade/utils"
)

func upgradeSegment(segment Segment, request *idl.UpgradePrimariesRequest, host string) error {
	err := restoreBackup(request, segment)

	if err != nil {
		return errors.Wrapf(err, "failed to restore master data directory backup on host %s for content id %d: %s",
			host, segment.Content, err)
	}

	err = RestoreTablespaces(request, segment)
	if err != nil {
		return errors.Wrapf(err, "restore tablespace on host %s for content id %d: %s",
			host, segment.Content, err)
	}

	err = performUpgrade(segment, request)

	if err != nil {
		failedAction := "upgrade"
		if request.CheckOnly {
			failedAction = "check"
		}
		return errors.Wrapf(err, "failed to %s primary on host %s with content %d", failedAction, host, segment.Content)
	}

	return nil
}

func performUpgrade(segment Segment, request *idl.UpgradePrimariesRequest) error {
	dbid := int(segment.DBID)
	segmentPair := upgrade.SegmentPair{
		Source: &upgrade.Segment{BinDir: request.SourceBinDir, DataDir: segment.SourceDataDir, DBID: dbid, Port: int(segment.SourcePort)},
		Target: &upgrade.Segment{BinDir: request.TargetBinDir, DataDir: segment.TargetDataDir, DBID: dbid, Port: int(segment.TargetPort)},
	}

	options := []upgrade.Option{
		upgrade.WithExecCommand(execCommand),
		upgrade.WithWorkDir(segment.WorkDir),
		upgrade.WithSegmentMode(),
	}

	if request.CheckOnly {
		options = append(options, upgrade.WithCheckOnly())
	} else {
		// During gpupgrade execute, tablepace mapping file is copied after
		// the master has been upgraded. So, don't pass this option during
		// --check mode. There is no test in pg_upgrade which depends on the
		// existence of this file.
		options = append(options, upgrade.WithTablespaceFile(request.TablespacesMappingFilePath))
	}

	if request.UseLinkMode {
		options = append(options, upgrade.WithLinkMode())
	}

	return upgrade.Run(segmentPair, options...)
}

func restoreBackup(request *idl.UpgradePrimariesRequest, segment Segment) error {
	if request.CheckOnly {
		return nil
	}

	return Rsync(request.MasterBackupDir, segment.TargetDataDir, []string{
		"internal.auto.conf",
		"postgresql.conf",
		"pg_hba.conf",
		"postmaster.opts",
		"gp_dbid",
		"gpssh.conf",
		"gpperfmon",
	})
}

func RestoreTablespaces(request *idl.UpgradePrimariesRequest, segment Segment) error {
	if request.CheckOnly {
		return nil
	}

	for oid, tablespace := range segment.Tablespaces {
		if !tablespace.GetUserDefined() {
			continue
		}

		targetDir := greenplum.GetTablespaceLocationForDbId(tablespace, int(segment.DBID))
		sourceDir := greenplum.GetMasterTablespaceLocation(filepath.Dir(request.TablespacesMappingFilePath), int(oid))
		if err := Rsync(sourceDir, targetDir, nil); err != nil {
			return errors.Wrap(err, "rsync master tablespace directory to segment tablespace directory")
		}

		symLinkName := fmt.Sprintf("%s/pg_tblspc/%s", segment.TargetDataDir, strconv.Itoa(int(oid)))
		if err := ReCreateSymLink(targetDir, symLinkName); err != nil {
			return errors.Wrap(err, "failed to recreate symbolic link")
		}
	}

	return nil
}

var ReCreateSymLink = func(sourceDir, symLinkName string) error {
	return reCreateSymLink(sourceDir, symLinkName)
}

func reCreateSymLink(sourceDir, symLinkName string) error {
	_, err := utils.System.Lstat(symLinkName)
	if err == nil {
		if err := utils.System.Remove(symLinkName); err != nil {
			return errors.Wrapf(err, "failed to unlink %q", symLinkName)
		}
	} else if !os.IsNotExist(err) {
		return errors.Wrapf(err, "stat symbolic link %q", symLinkName)
	}

	if err := utils.System.Symlink(sourceDir, symLinkName); err != nil {
		return errors.Wrapf(err, "create symbolic link %q to directory %q", symLinkName, sourceDir)
	}

	return nil
}
