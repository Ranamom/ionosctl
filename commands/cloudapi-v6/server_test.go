package commands

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/ionos-cloud/ionosctl/internal/config"
	"github.com/ionos-cloud/ionosctl/internal/core"
	"github.com/ionos-cloud/ionosctl/internal/utils/clierror"
	cloudapiv6 "github.com/ionos-cloud/ionosctl/services/cloudapi-v6"
	"github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	// Resources
	serverCreate = resources.Server{
		Server: ionoscloud.Server{
			Properties: &ionoscloud.ServerProperties{
				Name:             &testServerVar,
				Cores:            &cores,
				Ram:              &ram,
				CpuFamily:        &testServerVar,
				AvailabilityZone: &testServerVar,
				Type:             &testServerEnterpriseType,
			},
		},
	}
	serverCubeCreate = resources.Server{
		Server: ionoscloud.Server{
			Properties: &ionoscloud.ServerProperties{
				Name:             &testServerVar,
				Type:             &testServerCubeType,
				TemplateUuid:     &testServerVar,
				CpuFamily:        &testCpuFamilyType,
				AvailabilityZone: &testServerVar,
			},
			Entities: &ionoscloud.ServerEntities{
				Volumes: &ionoscloud.AttachedVolumes{
					Items: &[]ionoscloud.Volume{
						{
							Properties: &ionoscloud.VolumeProperties{
								Name:        &testServerVar,
								Bus:         &testServerVar,
								Type:        &testVolumeType,
								LicenceType: &testLicenceType,
							},
						},
					},
				},
			},
		},
	}
	s = ionoscloud.Server{
		Id: &testServerVar,
		Metadata: &ionoscloud.DatacenterElementMetadata{
			State: &state,
		},
		Properties: &ionoscloud.ServerProperties{
			Name:             &testServerVar,
			Cores:            &cores,
			Ram:              &ram,
			CpuFamily:        &testServerVar,
			AvailabilityZone: &testServerVar,
			VmState:          &state,
		},
	}
	ss = resources.Servers{
		Servers: ionoscloud.Servers{
			Id:    &testServerVar,
			Items: &[]ionoscloud.Server{s},
		},
	}
	serverProperties = resources.ServerProperties{
		ServerProperties: ionoscloud.ServerProperties{
			Name:             &testServerNewVar,
			Cores:            &coresNew,
			Ram:              &ramNew,
			CpuFamily:        &testServerNewVar,
			AvailabilityZone: &testServerNewVar,
		},
	}
	serverNew = resources.Server{
		Server: ionoscloud.Server{
			Metadata: &ionoscloud.DatacenterElementMetadata{
				State: &state,
			},
			Id:         &testServerVar,
			Properties: &serverProperties.ServerProperties,
		},
	}
	// Resources Attributes
	cores                    = int32(2)
	coresNew                 = int32(4)
	ram                      = int32(256)
	ramNew                   = int32(256)
	state                    = "ACTIVE"
	testServerVar            = "test-server"
	testServerNewVar         = "test-new-server"
	testVolumeType           = "DAS"
	testLicenceType          = "UNKNOWN"
	testCpuFamilyType        = "INTEL_SKYLAKE"
	testServerCubeType       = serverCubeType
	testServerEnterpriseType = serverEnterpriseType
	testServerErr            = errors.New("server test: error occurred")
)

func TestServerCmd(t *testing.T) {
	var err error
	core.RootCmdTest.AddCommand(ServerCmd())
	if ok := ServerCmd().IsAvailableCommand(); !ok {
		err = errors.New("non-available cmd")
	}
	assert.NoError(t, err)
}

func TestPreRunDcServerIds(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		err := PreRunDcServerIds(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunDcServerIdsRequiredFlagErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		err := PreRunDcServerIds(cfg)
		assert.Error(t, err)
	})
}

func TestPreRunServerCreate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), testServerVar)
		err := PreRunServerCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunServerCreateCube(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerCubeType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgTemplateId), testServerVar)
		err := PreRunServerCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunServerCreateCubeErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		err := PreRunServerCreate(cfg)
		assert.Error(t, err)
	})
}

func TestPreRunServerCreateCubeImg(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerCubeType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgTemplateId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgImageId), testServerVar)
		err := PreRunServerCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().List(testServerVar).Return(ss, nil, nil)
		err := RunServerList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().List(testServerVar).Return(ss, nil, testServerErr)
		err := RunServerList(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, nil)
		err := RunServerGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerGetWait(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, nil)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, nil)
		err := RunServerGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerGetWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForState), true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, testServerErr)
		err := RunServerGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, testServerErr)
		err := RunServerGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerCreate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerEnterpriseType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), cores)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), strconv.Itoa(int(ram)))
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Create(testServerVar, serverCreate).Return(&resources.Server{Server: s}, nil, nil)
		err := RunServerCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerCreateCube(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgVolumeName), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerCubeType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgBus), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgTemplateId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgLicenceType), testLicenceType)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Create(testServerVar, serverCubeCreate).Return(&resources.Server{Server: s}, nil, nil)
		err := RunServerCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerCreateWaitState(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerEnterpriseType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), cores)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ram)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForState), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Create(testServerVar, serverCreate).Return(&resources.Server{Server: s}, nil, nil)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, nil)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&resources.Server{Server: s}, nil, nil)
		err := RunServerCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerCreateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerEnterpriseType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), cores)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ram)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Create(testServerVar, serverCreate).Return(&resources.Server{Server: s}, nil, testServerErr)
		err := RunServerCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerCreateWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgType), testServerEnterpriseType)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), cores)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ram)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Create(testServerVar, serverCreate).Return(&resources.Server{Server: s}, &testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerUpdate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), coresNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ramNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Update(testServerVar, testServerVar, serverProperties).Return(&serverNew, nil, nil)
		err := RunServerUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerUpdateWaitStateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), coresNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ramNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForState), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Update(testServerVar, testServerVar, serverProperties).Return(&serverNew, nil, nil)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&serverNew, nil, testServerErr)
		err := RunServerUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerUpdateWaitState(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), coresNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ramNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForState), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Update(testServerVar, testServerVar, serverProperties).Return(&serverNew, nil, nil)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&serverNew, nil, nil)
		rm.CloudApiV6Mocks.Server.EXPECT().Get(testServerVar, testServerVar).Return(&serverNew, nil, nil)
		err := RunServerUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerUpdateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), coresNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ramNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Update(testServerVar, testServerVar, serverProperties).Return(&serverNew, nil, testServerErr)
		err := RunServerUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerUpdateResponseErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), coresNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ramNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Update(testServerVar, testServerVar, serverProperties).Return(&serverNew, &testResponse, testServerErr)
		err := RunServerUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerUpdateWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgName), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCores), coresNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgRam), ramNew)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgAvailabilityZone), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgCPUFamily), testServerNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Update(testServerVar, testServerVar, serverProperties).Return(&serverNew, &testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerDelete(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Delete(testServerVar, testServerVar).Return(nil, nil)
		err := RunServerDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerDeleteErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Delete(testServerVar, testServerVar).Return(nil, testServerErr)
		err := RunServerDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerDeleteWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Delete(testServerVar, testServerVar).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerDeleteAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Delete(testServerVar, testServerVar).Return(nil, nil)
		err := RunServerDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerDeleteAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = os.Stdin
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerSuspend(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Suspend(testServerVar, testServerVar).Return(nil, nil)
		err := RunServerSuspend(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerSuspendErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Suspend(testServerVar, testServerVar).Return(nil, testServerErr)
		err := RunServerSuspend(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerSuspendWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Suspend(testServerVar, testServerVar).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerSuspend(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerSuspendAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = os.Stdin
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerSuspend(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerStart(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Start(testServerVar, testServerVar).Return(nil, nil)
		err := RunServerStart(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerStartErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Start(testServerVar, testServerVar).Return(nil, testServerErr)
		err := RunServerStart(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerStartWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Start(testServerVar, testServerVar).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerStart(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerStartAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = os.Stdin
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerStart(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerStop(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Stop(testServerVar, testServerVar).Return(nil, nil)
		err := RunServerStop(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerStopErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Stop(testServerVar, testServerVar).Return(nil, testServerErr)
		err := RunServerStop(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerStopWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Stop(testServerVar, testServerVar).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerStop(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerStopAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = os.Stdin
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerStop(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerReboot(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		rm.CloudApiV6Mocks.Server.EXPECT().Reboot(testServerVar, testServerVar).Return(nil, nil)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerReboot(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerRebootErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Reboot(testServerVar, testServerVar).Return(nil, testServerErr)
		err := RunServerReboot(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerRebootWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Reboot(testServerVar, testServerVar).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerReboot(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerRebootAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = os.Stdin
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerReboot(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerResume(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Resume(testServerVar, testServerVar).Return(nil, nil)
		err := RunServerResume(cfg)
		assert.NoError(t, err)
	})
}

func TestRunServerResumeErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		rm.CloudApiV6Mocks.Server.EXPECT().Resume(testServerVar, testServerVar).Return(nil, testServerErr)
		err := RunServerResume(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerResumeWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.CloudApiV6Mocks.Server.EXPECT().Resume(testServerVar, testServerVar).Return(&testResponse, nil)
		rm.CloudApiV6Mocks.Request.EXPECT().GetStatus(testRequestIdVar).Return(&testRequestStatus, nil, testRequestErr)
		err := RunServerResume(cfg)
		assert.Error(t, err)
	})
}

func TestRunServerResumeAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		cfg.Stdin = os.Stdin
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgDataCenterId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgServerId), testServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), false)
		err := RunServerResume(cfg)
		assert.Error(t, err)
	})
}

func TestGetServersCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("server", config.ArgCols), []string{"Name"})
	getServersCols(core.GetGlobalFlagName("server", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetServersColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("server", config.ArgCols), []string{"Unknown"})
	getServersCols(core.GetGlobalFlagName("server", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestGetCubeServersIds(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	err := os.Setenv(ionoscloud.IonosUsernameEnvVar, "user")
	assert.NoError(t, err)
	err = os.Setenv(ionoscloud.IonosPasswordEnvVar, "pass")
	assert.NoError(t, err)
	err = os.Setenv(ionoscloud.IonosTokenEnvVar, "tok")
	assert.NoError(t, err)
	viper.Set(config.ArgServerUrl, config.DefaultApiURL)
	viper.Set(config.ArgOutput, config.DefaultOutputFormat)
	getCubeServersIds(w, testServerVar)
	err = w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`401 Unauthorized`)
	assert.True(t, re.Match(b.Bytes()))
}