// package main

// import (
// 	// "fmt"
// 	// "time"
// 	// "github.com/gorhill/cronexpr"
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/coreos/etcd/mvcc/mvccpb"
// 	"go.etcd.io/etcd/clientv3"
// 	"google.golang.org/grpc/keepalive"
// )

// func main()  {
// 	watch()
// }	


// func put()  {
// 	var (
// 		config clientv3.Config
// 		client *clientv3.Client
// 		err error
// 		kv clientv3.KV
// 		// put *clientv3.PutResponse
// 		getResp *clientv3.GetResponse
// 		// lease clientv3.Lease
// 	)
// 	config = clientv3.Config{
// 		Endpoints: []string{"127.0.0.1:9527"},
// 		DialTimeout: 5* time.Second,
// 	}
// 	if client,err = clientv3.New(config); err != nil {
// 		fmt.Println("connnect faled")
// 		return
// 	}
// 	// client = client
// 	kv = clientv3.NewKV(client)
// 	if getResp, err = kv.Get(context.TODO(), "name"); err != nil {
// 		fmt.Println("get failed")
// 		return
// 	} else  {
// 		fmt.Println(getResp.Kvs)
// 	}
// }


// func lease() {
// 	var (
// 		config clientv3.Config
// 		client *clientv3.Client
// 		err error
// 		kv clientv3.KV
// 		putResp *clientv3.PutResponse
// 		getResp *clientv3.GetResponse
// 		lease clientv3.Lease
// 		leaseGrand *clientv3.LeaseGrantResponse
// 		keepResp *clientv3.LeaseKeepAliveResponse
// 		keepRespChannel <- chan *clientv3.LeaseKeepAliveResponse
//  		leaseID clientv3.LeaseID
// 	)
// 	config = clientv3.Config{
// 		Endpoints: []string{"127.0.0.1:9527"},
// 		DialTimeout: 5* time.Second,
// 	}
// 	fmt.Println("ssssssssssssss")
// 	if client,err = clientv3.New(config); err != nil {
// 		fmt.Println("connnect faled")
// 		return
// 	}
// 	// client = client
// 	kv = clientv3.NewKV(client)
// 	if getResp, err = kv.Get(context.TODO(), "name"); err != nil {
// 		fmt.Println("get failed")
// 		return
// 	} else  {
// 		fmt.Println(getResp.Kvs)
// 	}
// 	if lease = clientv3.NewLease(client); lease == nil {
// 		fmt.Println("申请租约失败")
// 		return
// 	}
// 	if leaseGrand, err = lease.Grant(context.TODO(), 10); err  != nil {
// 		fmt.Println("申请租约失败")
// 		return
// 	}
// 	leaseID  = leaseGrand.ID
// 	if keepRespChannel, err = lease.KeepAlive(context.TODO(),leaseID); err != nil {
// 		fmt.Println("续租失败")
// 		return
// 	}
// 	go func() {
// 		for {
// 			select {
// 				case keepResp = <- keepRespChannel:
// 					if keepRespChannel == nil {
// 						fmt.Println("租约实效")
// 						goto END
// 					} else {
// 						fmt.Println("租约成功", keepResp.ID)
// 					}
// 			}
// 		}
// 		END:
// 	}()
// 	kv = clientv3.NewKV(client)
// 	if putResp, err = kv.Put(context.TODO(), "name", "wu", clientv3.WithLease(leaseID)); err != nil {
// 		fmt.Println("存入键值失败")
// 		return
// 	}
// 	fmt.Println("写入成功",putResp.Header.Revision)
// }

// func watch() {
// 	var (
// 		config clientv3.Config
// 		client *clientv3.Client
// 		err error
// 		kv clientv3.KV
// 		// putResp *clientv3.PutResponse
// 		getResp *clientv3.GetResponse
// 		watcher clientv3.Watcher
// 		watchRespChan <- chan clientv3.WatchResponse
// 		watchResp clientv3.WatchResponse
// 		watchRevision int64
// 		// lease clientv3.Lease
// 		// leaseGrand *clientv3.LeaseGrantResponse
// 		// keepResp *clientv3.LeaseKeepAliveResponse
// 		// keepRespChannel <- chan *clientv3.LeaseKeepAliveResponse
//  		// leaseID clientv3.LeaseID
// 	)
// 	config = clientv3.Config{
// 		Endpoints: []string{"127.0.0.1:2379"},
// 		DialTimeout: 5* time.Second,
// 	}
// 	if client,err = clientv3.New(config); err != nil {
// 		fmt.Println("connnect faled")
// 		return
// 	}
// 	kv = clientv3.NewKV(client)
// 	go func() {
// 		for {
// 			kv.Put(context.TODO(), "/cron/jobs/job7", "job7")
// 			kv.Delete(context.TODO(), "/cron/jobs/job7")
// 			time.Sleep( 5 * time.Second)
// 		}
// 	}()

// 	if  getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err  != nil {
// 		fmt.Println("获取值失败")
// 		return
// 	}

// 	if len(getResp.Kvs) != 0 {
// 		fmt.Println("当前值:", getResp.Kvs[0])
// 	}
// 	watchRevision = getResp.Header.Revision + 1
// 	watcher = clientv3.NewWatcher(client)
// 	watchRespChan = watcher.Watch(context.TODO(),"/cron/jobs/job7", clientv3.WithRev(watchRevision))
// 	for watchResp = range watchRespChan {
// 		for _, event := range watchResp.Events {
// 			switch event.Type {
// 			case mvccpb.PUT:
// 				fmt.Println("修改成功", event.Kv.Value)
// 			case mvccpb.DELETE:
// 				fmt.Println("删除", event.Kv.Key)
// 			}
			
// 		}
// 	}
// }

// func do() {
// 	var (
// 		config clientv3.Config
// 		client *clientv3.Client
// 		err error
// 		kv clientv3.KV
// 		// putResp *clientv3.PutResponse
// 		getResp *clientv3.GetResponse
// 		watcher clientv3.Watcher
// 		watchRespChan <- chan clientv3.WatchResponse
// 		watchResp clientv3.WatchResponse
// 		watchRevision int64
// 		putOp clientv3.O
// 		opResp clientv3.OpResponse
// 		lease clientv3.Lease
// 		leaseGrand *clientv3.LeaseGrantResponse
// 		keepResp *clientv3.LeaseKeepAliveResponse
// 		keepRespChannel <- chan *clientv3.LeaseKeepAliveResponse
//  		leaseID clientv3.LeaseID
// 		ctx context.Context
// 		cancelFunc context.CancelFunc
// 		txn clientv3.Txn
// 		kv clientv3.KV
// 	)
// 	config = clientv3.Config{
// 		Endpoints: []string{"127.0.0.1:2379"},
// 		DialTimeout: 5* time.Second,
// 	}
// 	if client,err = clientv3.New(config); err != nil {
// 		fmt.Println("connnect faled")
// 		return
// 	}
// 	lease = clientv3.NewLease(client)
// 	// 申请租约
// 	if leaseGrand, err = lease.Grant(context.TODO(), 5);err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	leaseID = leaseGrand.ID
// 	ctx, cancelFunc = context.WithCancel(context.TODO())
// 	defer cancelFunc()
// 	lease.Revoke(context.TODO(),leaseID)
// 	if keepRespChannel, err = lease.KeepAlive(context.TODO(), leaseID); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	go func() {
// 		for {
// 			select {
// 				case keepResp = <- keepRespChannel:
// 					if keepRespChannel == nil {
// 						fmt.Println("租约实效")
// 						goto END
// 					} else {
// 						fmt.Println("租约成功", keepResp.ID)
// 					}
// 			}
// 		}
// 		END:
// 	}()

// 	kv = clientv3.NewKV(client)
// 	if putResp, err = kv.Put(context.TODO(), "name", "wu", clientv3.WithLease(leaseID)); err != nil {
// 		fmt.Println("存入键值失败")
// 		return
// 	}
// 	fmt.Println("写入成功",putResp.Header.Revision)

// 	kv = clientv3.NewKV(client)
// 	putOp  = clientv3.OpPut("/cron/job8", "1212")
// 	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }