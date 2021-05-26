package resources

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testIpBlockResourceVar = "test-ipblock-resource"
)

func TestNewIpBlockService(t *testing.T) {
	ctx := context.Background()
	t.Run("list_ipblocks_error", func(t *testing.T) {
		svc := getTestClient(t)
		backupUnitSvc := NewIpBlockService(svc.Get(), ctx)
		_, _, err := backupUnitSvc.List()
		assert.Error(t, err)
	})
	t.Run("get_ipblock_error", func(t *testing.T) {
		svc := getTestClient(t)
		backupUnitSvc := NewIpBlockService(svc.Get(), ctx)
		_, _, err := backupUnitSvc.Get(testIpBlockResourceVar)
		assert.Error(t, err)
	})
	t.Run("create_ipblock_error", func(t *testing.T) {
		svc := getTestClient(t)
		backupUnitSvc := NewIpBlockService(svc.Get(), ctx)
		_, _, err := backupUnitSvc.Create(
			testIpBlockResourceVar,
			testIpBlockResourceVar,
			int32(1),
		)
		assert.Error(t, err)
	})
	t.Run("update_ipblock_error", func(t *testing.T) {
		svc := getTestClient(t)
		backupUnitSvc := NewIpBlockService(svc.Get(), ctx)
		_, _, err := backupUnitSvc.Update(testIpBlockResourceVar, IpBlockProperties{})
		assert.Error(t, err)
	})
	t.Run("delete_ipblock_error", func(t *testing.T) {
		svc := getTestClient(t)
		backupUnitSvc := NewIpBlockService(svc.Get(), ctx)
		_, err := backupUnitSvc.Delete(testIpBlockResourceVar)
		assert.Error(t, err)
	})
}