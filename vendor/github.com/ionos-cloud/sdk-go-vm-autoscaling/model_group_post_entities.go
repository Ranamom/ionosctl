/*
 * VM Auto Scaling API
 *
 * The VM Auto Scaling Service enables IONOS clients to horizontally scale the number of VM replicas based on configured rules. You can use VM Auto Scaling to ensure that you have a sufficient number of replicas to handle your application loads at all times.  For this purpose, create a VM Auto Scaling Group that contains the server replicas. The VM Auto Scaling Service ensures that the number of replicas in the group is always within the defined limits.   When scaling policies are set, VM Auto Scaling creates or deletes replicas according to the requirements of your applications. For each policy, specified 'scale-in' and 'scale-out' actions are performed when the corresponding thresholds are reached.
 *
 * API version: 1-SDK.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// GroupPostEntities The entities associated with this resource. The content depends on the resource type.
type GroupPostEntities struct {
	Actions *ActionsLinkResource `json:"actions,omitempty"`
	Servers *ServersLinkResource `json:"servers,omitempty"`
}

// NewGroupPostEntities instantiates a new GroupPostEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPostEntities() *GroupPostEntities {
	this := GroupPostEntities{}

	return &this
}

// NewGroupPostEntitiesWithDefaults instantiates a new GroupPostEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPostEntitiesWithDefaults() *GroupPostEntities {
	this := GroupPostEntities{}
	return &this
}

// GetActions returns the Actions field value
// If the value is explicit nil, the zero value for ActionsLinkResource will be returned
func (o *GroupPostEntities) GetActions() *ActionsLinkResource {
	if o == nil {
		return nil
	}

	return o.Actions

}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPostEntities) GetActionsOk() (*ActionsLinkResource, bool) {
	if o == nil {
		return nil, false
	}

	return o.Actions, true
}

// SetActions sets field value
func (o *GroupPostEntities) SetActions(v ActionsLinkResource) {

	o.Actions = &v

}

// HasActions returns a boolean if a field has been set.
func (o *GroupPostEntities) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// GetServers returns the Servers field value
// If the value is explicit nil, the zero value for ServersLinkResource will be returned
func (o *GroupPostEntities) GetServers() *ServersLinkResource {
	if o == nil {
		return nil
	}

	return o.Servers

}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupPostEntities) GetServersOk() (*ServersLinkResource, bool) {
	if o == nil {
		return nil, false
	}

	return o.Servers, true
}

// SetServers sets field value
func (o *GroupPostEntities) SetServers(v ServersLinkResource) {

	o.Servers = &v

}

// HasServers returns a boolean if a field has been set.
func (o *GroupPostEntities) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

func (o GroupPostEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}

	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}

	return json.Marshal(toSerialize)
}

type NullableGroupPostEntities struct {
	value *GroupPostEntities
	isSet bool
}

func (v NullableGroupPostEntities) Get() *GroupPostEntities {
	return v.value
}

func (v *NullableGroupPostEntities) Set(val *GroupPostEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPostEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPostEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPostEntities(val *GroupPostEntities) *NullableGroupPostEntities {
	return &NullableGroupPostEntities{value: val, isSet: true}
}

func (v NullableGroupPostEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPostEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
