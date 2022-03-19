package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/btree"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"time"
)
import clientv3 "go.etcd.io/etcd/client/v3"

func main() {
	btreeDemo()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	key := "hello"

	watchDemo(cli)

	resp, err := cli.Put(ctx, key, "sample_value")

	fmt.Println(resp)

	get, err := cli.Get(ctx, key)

	for i := range get.Kvs {
		fmt.Println(string(get.Kvs[i].Value))
	}

	cli.Delete(ctx, key)

	cancel()
	defer cli.Close()
}

func watchDemo(cli *clientv3.Client) {
	key := "/watchdemo"
	kv := clientv3.NewKV(cli)
	go func() {
		for i := 0; i < 100; i++ {
			kv.Put(context.TODO(), key, fmt.Sprint("hi", i))
			kv.Put(context.TODO(), "/watchdemo/x", fmt.Sprint("hi x ", i))
			kv.Put(context.TODO(), "/watchdemo/y", fmt.Sprint("hi y ", i))
			kv.Put(context.TODO(), "/watchdemo/m", fmt.Sprint("hi z ", i))
			time.Sleep(1 * time.Second)
		}
	}()

	response, err := kv.Get(context.TODO(), key)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response != nil {
		fmt.Println(response.Kvs[0].Value)
	}

	watchRevision := response.Header.Revision

	watcher := clientv3.NewWatcher(cli)

	watchChan := watcher.Watch(context.TODO(), key, clientv3.WithRev(watchRevision), clientv3.WithPrefix())

	// 处理kv变化事件
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}
}

type keyIndex struct {
	key []byte
}

func (ki *keyIndex) Less(b btree.Item) bool {
	return bytes.Compare(ki.key, b.(*keyIndex).key) == -1
}

func btreeDemo() {
	tree := *btree.New(5)
	for i := 0; i < 19; i++ {
		key := fmt.Sprint("a", i)
		put(&tree, key)
	}
	tree.Ascend(func(i btree.Item) bool {
		fmt.Println(string(i.(*keyIndex).key))
		return true
	})
}

func put(tree *btree.BTree, key string) btree.Item {
	keyIndex := &keyIndex{key: []byte(key)}
	return tree.ReplaceOrInsert(keyIndex)
}
