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
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	testRuleTargetProperties = resources.NetworkLoadBalancerForwardingRuleProperties{
		NetworkLoadBalancerForwardingRuleProperties: ionoscloud.NetworkLoadBalancerForwardingRuleProperties{
			Targets: &[]ionoscloud.NetworkLoadBalancerForwardingRuleTarget{
				{
					Ip:     &testNlbRuleTargetVar,
					Port:   &testNlbRuleTargetIntVar,
					Weight: &testNlbRuleTargetIntVar,
					HealthCheck: &ionoscloud.NetworkLoadBalancerForwardingRuleTargetHealthCheck{
						Check:         &testNlbRuleTargetBoolVar,
						CheckInterval: &testNlbRuleTargetIntVar,
						Maintenance:   &testNlbRuleTargetBoolVar,
					},
				},
			},
		},
	}
	testNlbRuleTargetGet = resources.NetworkLoadBalancerForwardingRule{
		NetworkLoadBalancerForwardingRule: ionoscloud.NetworkLoadBalancerForwardingRule{
			Id:         &testNlbRuleTargetVar,
			Properties: &ionoscloud.NetworkLoadBalancerForwardingRuleProperties{},
		},
	}
	testNlbRuleTargetGetUpdated = resources.NetworkLoadBalancerForwardingRule{
		NetworkLoadBalancerForwardingRule: ionoscloud.NetworkLoadBalancerForwardingRule{
			Id:         &testNlbRuleTargetVar,
			Properties: &testRuleTargetProperties.NetworkLoadBalancerForwardingRuleProperties,
		},
	}
	testNlbRuleTargetIntVar  = int32(1)
	testNlbRuleTargetBoolVar = false
	testNlbRuleTargetVar     = "test-rule-target"
	testNlbRuleTargetErr     = errors.New("networkloadbalancer-rule-target test error")
)

func TestPreRunNetworkLoadBalancerRuleTarget(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		err := PreRunNetworkLoadBalancerRuleTarget(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunNetworkLoadBalancerRuleTargetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		err := PreRunNetworkLoadBalancerRuleTarget(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		err := RunNlbRuleTargetList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunNlbRuleTargetListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, testNlbRuleTargetErr)
		err := RunNlbRuleTargetList(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetListGetTargetsErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGet, nil, nil)
		err := RunNlbRuleTargetList(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetListGetPropertiesErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&resources.NetworkLoadBalancerForwardingRule{}, nil, nil)
		err := RunNlbRuleTargetList(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetAdd(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWeight), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheckInterval), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheck), testNlbRuleTargetBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaintenance), testNlbRuleTargetBoolVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGet, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar, &testRuleTargetProperties).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		err := RunNlbRuleTargetAdd(cfg)
		assert.NoError(t, err)
	})
}

func TestRunNlbRuleTargetAddResponseErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWeight), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheckInterval), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheck), testNlbRuleTargetBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaintenance), testNlbRuleTargetBoolVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGet, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar, &testRuleTargetProperties).Return(&testNlbRuleTargetGetUpdated, &testResponse, nil)
		err := RunNlbRuleTargetAdd(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetAddErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWeight), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheckInterval), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheck), testNlbRuleTargetBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaintenance), testNlbRuleTargetBoolVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGet, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar, &testRuleTargetProperties).Return(&testNlbRuleTargetGetUpdated, nil, testNlbRuleTargetErr)
		err := RunNlbRuleTargetAdd(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetAddGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWeight), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheckInterval), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheck), testNlbRuleTargetBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaintenance), testNlbRuleTargetBoolVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGet, nil, testNlbRuleTargetErr)
		err := RunNlbRuleTargetAdd(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetAddWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWeight), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheckInterval), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCheck), testNlbRuleTargetBoolVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaintenance), testNlbRuleTargetBoolVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGet, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar, &testRuleTargetProperties).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		err := RunNlbRuleTargetAdd(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetRemove(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar,
			&resources.NetworkLoadBalancerForwardingRuleProperties{
				NetworkLoadBalancerForwardingRuleProperties: ionoscloud.NetworkLoadBalancerForwardingRuleProperties{
					Targets: &[]ionoscloud.NetworkLoadBalancerForwardingRuleTarget{},
				},
			}).Return(&testNlbRuleTargetGet, nil, nil)
		err := RunNlbRuleTargetRemove(cfg)
		assert.NoError(t, err)
	})
}

func TestRunNlbRuleTargetRemoveErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar,
			&resources.NetworkLoadBalancerForwardingRuleProperties{
				NetworkLoadBalancerForwardingRuleProperties: ionoscloud.NetworkLoadBalancerForwardingRuleProperties{
					Targets: &[]ionoscloud.NetworkLoadBalancerForwardingRuleTarget{},
				},
			}).Return(&testNlbRuleTargetGet, nil, testNlbRuleTargetErr)
		err := RunNlbRuleTargetRemove(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetRemoveGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, testNlbRuleTargetErr)
		err := RunNlbRuleTargetRemove(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetRemoveIpErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), "x.x.x.x")
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		err := RunNlbRuleTargetRemove(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetRemovePortErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), int32(2))
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		err := RunNlbRuleTargetRemove(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetRemoveWaitErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgWaitForRequest), true)
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar,
			&resources.NetworkLoadBalancerForwardingRuleProperties{
				NetworkLoadBalancerForwardingRuleProperties: ionoscloud.NetworkLoadBalancerForwardingRuleProperties{
					Targets: &[]ionoscloud.NetworkLoadBalancerForwardingRuleTarget{},
				},
			}).Return(&testNlbRuleTargetGet, nil, nil)
		err := RunNlbRuleTargetRemove(cfg)
		assert.Error(t, err)
	})
}

func TestRunNlbRuleTargetRemoveAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		rm.NetworkLoadBalancer.EXPECT().GetForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar).Return(&testNlbRuleTargetGetUpdated, nil, nil)
		rm.NetworkLoadBalancer.EXPECT().UpdateForwardingRule(testNlbRuleTargetVar, testNlbRuleTargetVar, testNlbRuleTargetVar,
			&resources.NetworkLoadBalancerForwardingRuleProperties{
				NetworkLoadBalancerForwardingRuleProperties: ionoscloud.NetworkLoadBalancerForwardingRuleProperties{
					Targets: &[]ionoscloud.NetworkLoadBalancerForwardingRuleTarget{},
				},
			}).Return(&testNlbRuleTargetGet, nil, nil)
		err := RunNlbRuleTargetRemove(cfg)
		assert.NoError(t, err)
	})
}

func TestRunNlbRuleTargetRemoveAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgForce, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgNetworkLoadBalancerId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRuleId), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetIp), testNlbRuleTargetVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetPort), testNlbRuleTargetIntVar)
		cfg.Stdin = os.Stdin
		err := RunNlbRuleTargetRemove(cfg)
		assert.Error(t, err)
	})
}

func TestGetRuleTargetsCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() {}
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("target", config.ArgCols), []string{"TargetIp"})
	getRuleTargetsCols(core.GetGlobalFlagName("target", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetRuleColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() {}
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("target", config.ArgCols), []string{"Unknown"})
	getRuleTargetsCols(core.GetGlobalFlagName("target", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}