// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

//  Creates a device registry in Google's Cloud IoT Core platform. For more information see
// [the official documentation](https://cloud.google.com/iot/docs/) and
// [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
// 
type Registry struct {
	s *pulumi.ResourceState
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOpt) (*Registry, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["credentials"] = nil
		inputs["eventNotificationConfig"] = nil
		inputs["httpConfig"] = nil
		inputs["mqttConfig"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["stateNotificationConfig"] = nil
	} else {
		inputs["credentials"] = args.Credentials
		inputs["eventNotificationConfig"] = args.EventNotificationConfig
		inputs["httpConfig"] = args.HttpConfig
		inputs["mqttConfig"] = args.MqttConfig
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["stateNotificationConfig"] = args.StateNotificationConfig
	}
	s, err := ctx.RegisterResource("gcp:kms/registry:Registry", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Registry{s: s}, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegistryState, opts ...pulumi.ResourceOpt) (*Registry, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["credentials"] = state.Credentials
		inputs["eventNotificationConfig"] = state.EventNotificationConfig
		inputs["httpConfig"] = state.HttpConfig
		inputs["mqttConfig"] = state.MqttConfig
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["stateNotificationConfig"] = state.StateNotificationConfig
	}
	s, err := ctx.ReadResource("gcp:kms/registry:Registry", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Registry{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Registry) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Registry) ID() *pulumi.IDOutput {
	return r.s.ID
}

// List of public key certificates to authenticate devices. Structure is documented below. 
func (r *Registry) Credentials() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["credentials"])
}

// A PubSub topics to publish device events. Structure is documented below.
func (r *Registry) EventNotificationConfig() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["eventNotificationConfig"])
}

// Activate or deactivate HTTP. Structure is documented below.
func (r *Registry) HttpConfig() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["httpConfig"])
}

// Activate or deactivate MQTT. Structure is documented below.
func (r *Registry) MqttConfig() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["mqttConfig"])
}

// A unique name for the resource, required by device registry.
// Changing this forces a new resource to be created.
func (r *Registry) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The project in which the resource belongs. If it is not provided, the provider project is used.
func (r *Registry) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The Region in which the created address should reside. If it is not provided, the provider region is used.
func (r *Registry) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// A PubSub topic to publish device state updates. Structure is documented below.
func (r *Registry) StateNotificationConfig() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["stateNotificationConfig"])
}

// Input properties used for looking up and filtering Registry resources.
type RegistryState struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials interface{}
	// A PubSub topics to publish device events. Structure is documented below.
	EventNotificationConfig interface{}
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig interface{}
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig interface{}
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name interface{}
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project interface{}
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region interface{}
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig interface{}
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// List of public key certificates to authenticate devices. Structure is documented below. 
	Credentials interface{}
	// A PubSub topics to publish device events. Structure is documented below.
	EventNotificationConfig interface{}
	// Activate or deactivate HTTP. Structure is documented below.
	HttpConfig interface{}
	// Activate or deactivate MQTT. Structure is documented below.
	MqttConfig interface{}
	// A unique name for the resource, required by device registry.
	// Changing this forces a new resource to be created.
	Name interface{}
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project interface{}
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	Region interface{}
	// A PubSub topic to publish device state updates. Structure is documented below.
	StateNotificationConfig interface{}
}
