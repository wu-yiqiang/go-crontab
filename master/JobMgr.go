package master

import (
	"go.etcd.io/etcd/clientv3"
)

// 删除
// func (jopmgr *JobMgr)DeleteJob(name string)(oldJob *common.Job, err error) {
// 	var (
// 		jobKey string
// 		delResp *clientv3.DeleteResponse
// 	)

// 	jobKey = "/cron/joob" + name
// 	type JobMgr struct {
// 		Name string
// 	}

// 	return
// }