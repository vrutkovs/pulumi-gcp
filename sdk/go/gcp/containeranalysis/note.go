// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Note struct {
	s *pulumi.ResourceState
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOpt) (*Note, error) {
	if args == nil || args.AttestationAuthority == nil {
		return nil, errors.New("missing required argument 'AttestationAuthority'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attestationAuthority"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
	} else {
		inputs["attestationAuthority"] = args.AttestationAuthority
		inputs["name"] = args.Name
		inputs["project"] = args.Project
	}
	s, err := ctx.RegisterResource("gcp:containeranalysis/note:Note", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Note{s: s}, nil
}

// GetNote gets an existing Note resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNote(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NoteState, opts ...pulumi.ResourceOpt) (*Note, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attestationAuthority"] = state.AttestationAuthority
		inputs["name"] = state.Name
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:containeranalysis/note:Note", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Note{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Note) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Note) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Note) AttestationAuthority() *pulumi.Output {
	return r.s.State["attestationAuthority"]
}

func (r *Note) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Note) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering Note resources.
type NoteState struct {
	AttestationAuthority interface{}
	Name interface{}
	Project interface{}
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
	AttestationAuthority interface{}
	Name interface{}
	Project interface{}
}
