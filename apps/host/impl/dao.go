package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lxygwqf9527/demo-api/apps/host"
)

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 保存数据
	tx, err := i.db.BeginTx(ctx, nil)
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
	// 插入Resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.ExecContext(ctx, ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP)
	if err != nil {
		return err
	}

	//插入describe数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.ExecContext(ctx, ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber)
	if err != nil {
		return err
	}

	return nil
}

func (i *HostServiceImpl) update(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)
	// 开启事务 tx
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

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

	var (
		resStmt, hostStmt *sql.Stmt
	)
	// 更新Resource
	resStmt, err = tx.PrepareContext(ctx, updateResourceSQL)
	if err != nil {
		return err
	}

	// defer resStmt.Close()
	_, err = resStmt.ExecContext(ctx, ins.Vendor, ins.Region, ins.ExpireAt, ins.Description, ins.Name, ins.Id)
	if err != nil {
		return err
	}

	// 更新Host表
	hostStmt, err = tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return err
	}
	// defer hostStmt.Close()
	_, err = hostStmt.ExecContext(ctx, ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
