package master

import (
	"time"
	"go.etcd.io/etcd/clientv3"
)

// 任务管理器
type JobMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

// 初始化
func InitJobMgr() (err error)  {
	var (
		config clientv3.Client
		client clientv3.Client
		kv clientv3.KV
		lease clientv3.Lease
	)
	config = clientv3.Config{
		Endpoints: G_config.EtcdAndPoints, //集群
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond,
	}
	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	// 赋值单例
	G_jobMgr = &JobMgr{
		client: client,
		kv: kv,
		lease: lease,
	}
	return	
}


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