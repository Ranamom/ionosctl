package commands

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"regexp"
	"testing"

	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/ionos-cloud/ionosctl/pkg/resources"
	"github.com/ionos-cloud/ionosctl/pkg/utils/clierror"
	ionoscloud "github.com/ionos-cloud/sdk-go/v5"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	snapshotTest = resources.Snapshot{
		Snapshot: ionoscloud.Snapshot{
			Id: &testSnapshotVar,
			Properties: &ionoscloud.SnapshotProperties{
				Name:        &testSnapshotVar,
				Location:    &testSnapshotVar,
				Description: &testSnapshotVar,
				Size:        &testSnapshotSize,
				LicenceType: &testSnapshotVar,
			},
			Metadata: &ionoscloud.DatacenterElementMetadata{
				State: &snapshotState,
			},
		},
	}
	snapshotState = "BUSY"
	snapshots     = resources.Snapshots{
		Snapshots: ionoscloud.Snapshots{
			Id:    &testSnapshotVar,
			Items: &[]ionoscloud.Snapshot{snapshotTest.Snapshot},
		},
	}
	snapshotProperties = resources.SnapshotProperties{
		SnapshotProperties: ionoscloud.SnapshotProperties{
			Name:                &testSnapshotNewVar,
			Description:         &testSnapshotNewVar,
			CpuHotPlug:          &testSnapshotBoolVar,
			CpuHotUnplug:        &testSnapshotBoolVar,
			RamHotPlug:          &testSnapshotBoolVar,
			RamHotUnplug:        &testSnapshotBoolVar,
			NicHotPlug:          &testSnapshotBoolVar,
			NicHotUnplug:        &testSnapshotBoolVar,
			DiscVirtioHotPlug:   &testSnapshotBoolVar,
			DiscVirtioHotUnplug: &testSnapshotBoolVar,
			DiscScsiHotPlug:     &testSnapshotBoolVar,
			DiscScsiHotUnplug:   &testSnapshotBoolVar,
			SecAuthProtection:   &testSnapshotBoolVar,
			LicenceType:         &testSnapshotVar,
		},
	}
	snapshotNew = resources.Snapshot{
		Snapshot: ionoscloud.Snapshot{
			Id:         &testSnapshotVar,
			Properties: &snapshotProperties.SnapshotProperties,
		},
	}
	testSnapshotBoolVar = false
	testSnapshotSize    = float32(2)
	testSnapshotVar     = "test-snapshot"
	testSnapshotNewVar  = "test-new-snapshot"
	testSnapshotErr     = errors.New("snapshot test error")
)

func TestPreRunSnapshotId(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		err := PreRunSnapshotId(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunSnapshotIdErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), "")
		err := PreRunSnapshotId(cfg)
		assert.Error(t, err)
	})
}

func TestPreRunSnapNameLicenceDcIdVolumeId(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		err := PreRunSnapNameLicenceDcIdVolumeId(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunSnapNameLicenceDcIdVolumeIdErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), "")
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), "")
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), "")
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), "")
		err := PreRunSnapNameLicenceDcIdVolumeId(cfg)
		assert.Error(t, err)
	})
}

func TestPreRunSnapshotIdDcIdVolumeId(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		err := PreRunSnapshotIdDcIdVolumeId(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		rm.Snapshot.EXPECT().List().Return(snapshots, nil, nil)
		err := RunSnapshotList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		rm.Snapshot.EXPECT().List().Return(snapshots, nil, testSnapshotErr)
		err := RunSnapshotList(cfg)
		assert.Error(t, err)
	})
}

func TestRunSnapshotListSort(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), testSnapshotVar)
		rm.Snapshot.EXPECT().List().Return(snapshots, nil, nil)
		err := RunSnapshotList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		rm.Snapshot.EXPECT().Get(testSnapshotVar).Return(&snapshotTest, nil, nil)
		err := RunSnapshotGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		rm.Snapshot.EXPECT().Get(testSnapshotVar).Return(&snapshotTest, nil, testSnapshotErr)
		err := RunSnapshotGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunSnapshotCreate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDescription), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotSecAuthProtection), false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Create(testSnapshotVar, testSnapshotVar, testSnapshotVar, testSnapshotVar, testSnapshotVar, false).Return(&snapshotTest, nil, nil)
		err := RunSnapshotCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotCreateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDescription), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotSecAuthProtection), false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Create(testSnapshotVar, testSnapshotVar, testSnapshotVar, testSnapshotVar, testSnapshotVar, false).Return(&snapshotTest, &testResponse, nil)
		err := RunSnapshotCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunSnapshotUpdate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDescription), testSnapshotNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotCpuHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotCpuHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotRamHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotRamHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotNicHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotNicHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscVirtioHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscVirtioHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscScsiHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscScsiHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotSecAuthProtection), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Update(testSnapshotVar, snapshotProperties).Return(&snapshotNew, nil, nil)
		err := RunSnapshotUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotUpdateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDescription), testSnapshotNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotName), testSnapshotNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotCpuHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotCpuHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotRamHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotRamHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotNicHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotNicHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscVirtioHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscVirtioHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscScsiHotPlug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotDiscScsiHotUnplug), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotSecAuthProtection), testSnapshotBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotLicenceType), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.Snapshot.EXPECT().Update(testSnapshotVar, snapshotProperties).Return(&snapshotNew, nil, testSnapshotErr)
		err := RunSnapshotUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunSnapshotRestore(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Restore(testSnapshotVar, testSnapshotVar, testSnapshotVar).Return(nil, nil)
		err := RunSnapshotRestore(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotRestoreErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Restore(testSnapshotVar, testSnapshotVar, testSnapshotVar).Return(nil, testSnapshotErr)
		err := RunSnapshotRestore(cfg)
		assert.Error(t, err)
	})
}

func TestRunSnapshotRestoreAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgVolumeId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Restore(testSnapshotVar, testSnapshotVar, testSnapshotVar).Return(nil, nil)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		err := RunSnapshotRestore(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotDelete(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Delete(testSnapshotVar).Return(nil, nil)
		err := RunSnapshotDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotDeleteErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.Snapshot.EXPECT().Delete(testSnapshotVar).Return(nil, testSnapshotErr)
		err := RunSnapshotDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunSnapshotDeleteAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		rm.Snapshot.EXPECT().Delete(testSnapshotVar).Return(nil, nil)
		err := RunSnapshotDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunSnapshotDeleteAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSnapshotId), testSnapshotVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		cfg.Stdin = os.Stdin
		err := RunSnapshotDelete(cfg)
		assert.Error(t, err)
	})
}

func TestGetSnapshotsCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() {}
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("snapshot", config.ArgCols), []string{"Name"})
	getSnapshotCols(core.GetGlobalFlagName("snapshot", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetSnapshotsColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() {}
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("snapshot", config.ArgCols), []string{"Unknown"})
	getSnapshotCols(core.GetGlobalFlagName("snapshot", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestGetSnapshotsIds(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() {}
	w := bufio.NewWriter(&b)
	viper.Set(config.ArgConfig, "../pkg/testdata/config.json")
	getSnapshotIds(w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`401 Unauthorized`)
	assert.True(t, re.Match(b.Bytes()))
}
