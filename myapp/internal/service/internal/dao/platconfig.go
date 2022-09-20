// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myapp/internal/service/internal/dao/internal"
)

// platconfigDao is the data access object for table platconfig.
// You can define custom methods on it to extend its functionality as you wish.
type platconfigDao struct {
	*internal.PlatconfigDao
}

var (
	// Platconfig is globally public accessible object for table platconfig operations.
	Platconfig = platconfigDao{
		internal.NewPlatconfigDao(),
	}
)

// Fill with you ideas below.