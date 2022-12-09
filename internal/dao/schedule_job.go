// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go_blog/internal/dao/internal"
)

// internalScheduleJobDao is internal type for wrapping internal DAO implements.
type internalScheduleJobDao = *internal.ScheduleJobDao

// scheduleJobDao is the data access object for table schedule_job.
// You can define custom methods on it to extend its functionality as you wish.
type scheduleJobDao struct {
	internalScheduleJobDao
}

var (
	// ScheduleJob is globally public accessible object for table schedule_job operations.
	ScheduleJob = scheduleJobDao{
		internal.NewScheduleJobDao(),
	}
)

// Fill with you ideas below.