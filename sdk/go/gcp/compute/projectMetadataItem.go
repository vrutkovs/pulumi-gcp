// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a single key/value pair on metadata common to all instances for
// a project in GCE. Using `google_compute_project_metadata_item` lets you
// manage a single key/value setting in Terraform rather than the entire
// project metadata map.
type ProjectMetadataItem struct {
	s *pulumi.ResourceState
}

// NewProjectMetadataItem registers a new resource with the given unique name, arguments, and options.
func NewProjectMetadataItem(ctx *pulumi.Context,
	name string, args *ProjectMetadataItemArgs, opts ...pulumi.ResourceOpt) (*ProjectMetadataItem, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["key"] = nil
		inputs["project"] = nil
		inputs["value"] = nil
	} else {
		inputs["key"] = args.Key
		inputs["project"] = args.Project
		inputs["value"] = args.Value
	}
	s, err := ctx.RegisterResource("gcp:compute/projectMetadataItem:ProjectMetadataItem", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectMetadataItem{s: s}, nil
}

// GetProjectMetadataItem gets an existing ProjectMetadataItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMetadataItem(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProjectMetadataItemState, opts ...pulumi.ResourceOpt) (*ProjectMetadataItem, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["key"] = state.Key
		inputs["project"] = state.Project
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("gcp:compute/projectMetadataItem:ProjectMetadataItem", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectMetadataItem{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ProjectMetadataItem) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ProjectMetadataItem) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The metadata key to set.
func (r *ProjectMetadataItem) Key() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["key"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *ProjectMetadataItem) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The value to set for the given metadata key.
func (r *ProjectMetadataItem) Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering ProjectMetadataItem resources.
type ProjectMetadataItemState struct {
	// The metadata key to set.
	Key interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The value to set for the given metadata key.
	Value interface{}
}

// The set of arguments for constructing a ProjectMetadataItem resource.
type ProjectMetadataItemArgs struct {
	// The metadata key to set.
	Key interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The value to set for the given metadata key.
	Value interface{}
}