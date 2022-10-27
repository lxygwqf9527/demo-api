package impl

import (
	"context"
	"fmt"

	"github.com/lxygwqf9527/demo-api/apps/host"
)

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 保存数据
	fmt.Println(ctx, ins, "AAAAAAAAAAAAAAAAAAA")
	tx, err := i.db.BeginTx(ctx, nil)
	fmt.Println(ctx, "AAAAAAAAAAAAAAAAAAA")
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	// 通过defer处理事务提交方式
	// 1.无错误，则commit事务
	// 2.有报错，则rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error,%s", err)
			}
		}
	}()
	fmt.Println(ins, "AAAAAAAAAAAAAAAAAAA")
	// 插入Resource数据
	rstmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.Exec(ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP)
	if err != nil {
		return err
	}

	//插入describe数据
	dstmt, err := tx.Prepare(InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber)
	if err != nil {
		return err
	}

	return nil
}
