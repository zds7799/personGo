// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myapp/internal/service/internal/dao/internal"
)

// platnumberDao is the data access object for table platnumber.
// You can define custom methods on it to extend its functionality as you wish.
type platnumberDao struct {
	*internal.PlatnumberDao
}

var (
	// Platnumber is globally public accessible object for table platnumber operations.
	Platnumber = platnumberDao{
		internal.NewPlatnumberDao(),
	}
)

// Fill with you ideas below.
