package customesort

import "go-admin-example/common/model/authorize"

type UserMenuSlice []authorize.UserMenu

func (x UserMenuSlice) Len() int           { return len(x) }
func (x UserMenuSlice) Less(i, j int) bool { return x[i].Sort < x[j].Sort }
func (x UserMenuSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
