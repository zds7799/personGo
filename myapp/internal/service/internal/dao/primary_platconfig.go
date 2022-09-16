// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myapp/internal/service/internal/dao/internal"
)

// primaryPlatconfigDao is the data access object for table platconfig.
// You can define custom methods on it to extend its functionality as you wish.
type primaryPlatconfigDao struct {
	*internal.PrimaryPlatconfigDao
}

var (
	// PrimaryPlatconfig is globally public accessible object for table platconfig operations.
	PrimaryPlatconfig = primaryPlatconfigDao{
		internal.NewPrimaryPlatconfigDao(),
	}
)

// Fill with you ideas below.