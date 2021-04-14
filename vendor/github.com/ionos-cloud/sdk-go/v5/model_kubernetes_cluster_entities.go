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

// KubernetesClusterEntities struct for KubernetesClusterEntities
type KubernetesClusterEntities struct {
	Nodepools *KubernetesNodePools `json:"nodepools,omitempty"`
}



// GetNodepools returns the Nodepools field value
// If the value is explicit nil, the zero value for KubernetesNodePools will be returned
func (o *KubernetesClusterEntities) GetNodepools() *KubernetesNodePools {
	if o == nil {
		return nil
	}


	return o.Nodepools

}

// GetNodepoolsOk returns a tuple with the Nodepools field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterEntities) GetNodepoolsOk() (*KubernetesNodePools, bool) {
	if o == nil {
		return nil, false
	}


	return o.Nodepools, true
}

// SetNodepools sets field value
func (o *KubernetesClusterEntities) SetNodepools(v KubernetesNodePools) {


	o.Nodepools = &v

}

// HasNodepools returns a boolean if a field has been set.
func (o *KubernetesClusterEntities) HasNodepools() bool {
	if o != nil && o.Nodepools != nil {
		return true
	}

	return false
}


func (o KubernetesClusterEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Nodepools != nil {
		toSerialize["nodepools"] = o.Nodepools
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesClusterEntities struct {
	value *KubernetesClusterEntities
	isSet bool
}

func (v NullableKubernetesClusterEntities) Get() *KubernetesClusterEntities {
	return v.value
}

func (v *NullableKubernetesClusterEntities) Set(val *KubernetesClusterEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterEntities(val *KubernetesClusterEntities) *NullableKubernetesClusterEntities {
	return &NullableKubernetesClusterEntities{value: val, isSet: true}
}

func (v NullableKubernetesClusterEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

