/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// NetworkLoadBalancerForwardingRuleTarget struct for NetworkLoadBalancerForwardingRuleTarget
type NetworkLoadBalancerForwardingRuleTarget struct {
	// IP of a balanced target VM
	Ip *string `json:"ip"`
	// Port of the balanced target service. (range: 1 to 65535)
	Port *int32 `json:"port"`
	// Weight parameter is used to adjust the target VM's weight relative to other target VMs. All target VMs will receive a load proportional to their weight relative to the sum of all weights, so the higher the weight, the higher the load. The default weight is 1, and the maximal value is 256. A value of 0 means the target VM will not participate in load-balancing but will still accept persistent connections. If this parameter is used to distribute the load according to target VM's capacity, it is recommended to start with values which can both grow and shrink, for instance between 10 and 100 to leave enough room above and below for later adjustments.
	Weight *int32 `json:"weight"`
	HealthCheck *NetworkLoadBalancerForwardingRuleTargetHealthCheck `json:"healthCheck,omitempty"`
}



// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetIp() *string {
	if o == nil {
		return nil
	}


	return o.Ip

}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Ip, true
}

// SetIp sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetIp(v string) {


	o.Ip = &v

}

// HasIp returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}



// GetPort returns the Port field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetPort() *int32 {
	if o == nil {
		return nil
	}


	return o.Port

}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}


	return o.Port, true
}

// SetPort sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetPort(v int32) {


	o.Port = &v

}

// HasPort returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}



// GetWeight returns the Weight field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetWeight() *int32 {
	if o == nil {
		return nil
	}


	return o.Weight

}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}


	return o.Weight, true
}

// SetWeight sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetWeight(v int32) {


	o.Weight = &v

}

// HasWeight returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}



// GetHealthCheck returns the HealthCheck field value
// If the value is explicit nil, the zero value for NetworkLoadBalancerForwardingRuleTargetHealthCheck will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetHealthCheck() *NetworkLoadBalancerForwardingRuleTargetHealthCheck {
	if o == nil {
		return nil
	}


	return o.HealthCheck

}

// GetHealthCheckOk returns a tuple with the HealthCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleTarget) GetHealthCheckOk() (*NetworkLoadBalancerForwardingRuleTargetHealthCheck, bool) {
	if o == nil {
		return nil, false
	}


	return o.HealthCheck, true
}

// SetHealthCheck sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetHealthCheck(v NetworkLoadBalancerForwardingRuleTargetHealthCheck) {


	o.HealthCheck = &v

}

// HasHealthCheck returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}


func (o NetworkLoadBalancerForwardingRuleTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	

	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	

	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	

	if o.HealthCheck != nil {
		toSerialize["healthCheck"] = o.HealthCheck
	}
	
	return json.Marshal(toSerialize)
}

type NullableNetworkLoadBalancerForwardingRuleTarget struct {
	value *NetworkLoadBalancerForwardingRuleTarget
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRuleTarget) Get() *NetworkLoadBalancerForwardingRuleTarget {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRuleTarget) Set(val *NetworkLoadBalancerForwardingRuleTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRuleTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRuleTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRuleTarget(val *NetworkLoadBalancerForwardingRuleTarget) *NullableNetworkLoadBalancerForwardingRuleTarget {
	return &NullableNetworkLoadBalancerForwardingRuleTarget{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRuleTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRuleTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


