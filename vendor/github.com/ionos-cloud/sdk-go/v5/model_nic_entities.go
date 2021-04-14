/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// NicEntities struct for NicEntities
type NicEntities struct {
	Firewallrules *FirewallRules `json:"firewallrules,omitempty"`
}



// GetFirewallrules returns the Firewallrules field value
// If the value is explicit nil, the zero value for FirewallRules will be returned
func (o *NicEntities) GetFirewallrules() *FirewallRules {
	if o == nil {
		return nil
	}


	return o.Firewallrules

}

// GetFirewallrulesOk returns a tuple with the Firewallrules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NicEntities) GetFirewallrulesOk() (*FirewallRules, bool) {
	if o == nil {
		return nil, false
	}


	return o.Firewallrules, true
}

// SetFirewallrules sets field value
func (o *NicEntities) SetFirewallrules(v FirewallRules) {


	o.Firewallrules = &v

}

// HasFirewallrules returns a boolean if a field has been set.
func (o *NicEntities) HasFirewallrules() bool {
	if o != nil && o.Firewallrules != nil {
		return true
	}

	return false
}


func (o NicEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Firewallrules != nil {
		toSerialize["firewallrules"] = o.Firewallrules
	}
	
	return json.Marshal(toSerialize)
}

type NullableNicEntities struct {
	value *NicEntities
	isSet bool
}

func (v NullableNicEntities) Get() *NicEntities {
	return v.value
}

func (v *NullableNicEntities) Set(val *NicEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableNicEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableNicEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicEntities(val *NicEntities) *NullableNicEntities {
	return &NullableNicEntities{value: val, isSet: true}
}

func (v NullableNicEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

