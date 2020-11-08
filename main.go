package main

import (
	"fmt"
	"github.com/bolt-master"
)

func main() {
	db,err:=bolt.Open("datas.db",0600,nil)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	fmt.Println(db)
	//1.设置数据
	db.Update(func(tx *bolt.Tx) error {
		//Tx:transaction 交易
		tong1,err:=tx.CreateBucket([]byte("male"))
		if err != nil {
			panic(err.Error())
		}
		tong1.Put([]byte("13167582311"),[]byte("于洪伟"))
		tong1.Put([]byte("11343214134"),[]byte("隔壁老王安川"))

		tong2,err:=tx.CreateBucket([]byte("female"))
		if err != nil {
			panic(err.Error())
		}
		tong2.Put([]byte("12131231311"),[]byte("翠花"))

		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("male"))
		if bucket==nil {
			panic("读取数据发生错误，请重试")
		}
		dataBytes:=bucket.Get([]byte("11343214134"))
		fmt.Println("读取到的数据是：",string((dataBytes)))
		return nil
	})
}
