// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// Code generated by enterprise/einterfaces_gen.go; DO NOT EDIT.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type ComplianceInterface interface {
	RunComplianceJob(job *model.Compliance) *model.AppError
	StartComplianceDailyJob()
}
