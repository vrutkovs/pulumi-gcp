// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Policy struct {
	s *pulumi.ResourceState
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.DefaultAdmissionRule == nil {
		return nil, errors.New("missing required argument 'DefaultAdmissionRule'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["admissionWhitelistPatterns"] = nil
		inputs["clusterAdmissionRules"] = nil
		inputs["defaultAdmissionRule"] = nil
		inputs["description"] = nil
		inputs["project"] = nil
	} else {
		inputs["admissionWhitelistPatterns"] = args.AdmissionWhitelistPatterns
		inputs["clusterAdmissionRules"] = args.ClusterAdmissionRules
		inputs["defaultAdmissionRule"] = args.DefaultAdmissionRule
		inputs["description"] = args.Description
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:binaryauthorization/policy:Policy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["admissionWhitelistPatterns"] = state.AdmissionWhitelistPatterns
		inputs["clusterAdmissionRules"] = state.ClusterAdmissionRules
		inputs["defaultAdmissionRule"] = state.DefaultAdmissionRule
		inputs["description"] = state.Description
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:binaryauthorization/policy:Policy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Policy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Policy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Policy) AdmissionWhitelistPatterns() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["admissionWhitelistPatterns"])
}

func (r *Policy) ClusterAdmissionRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["clusterAdmissionRules"])
}

func (r *Policy) DefaultAdmissionRule() *pulumi.Output {
	return r.s.State["defaultAdmissionRule"]
}

func (r *Policy) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Policy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	AdmissionWhitelistPatterns interface{}
	ClusterAdmissionRules interface{}
	DefaultAdmissionRule interface{}
	Description interface{}
	Project interface{}
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	AdmissionWhitelistPatterns interface{}
	ClusterAdmissionRules interface{}
	DefaultAdmissionRule interface{}
	Description interface{}
	Project interface{}
}
