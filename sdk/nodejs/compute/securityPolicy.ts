// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Security Policy defines an IP blacklist or whitelist that protects load balanced Google Cloud services by denying or permitting traffic from specified IP ranges. For more information
 * see the [official documentation](https://cloud.google.com/armor/docs/configure-security-policies)
 * and the [API](https://cloud.google.com/compute/docs/reference/rest/beta/securityPolicies).
 */
export class SecurityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SecurityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityPolicyState): SecurityPolicy {
        return new SecurityPolicy(name, <any>state, { id });
    }

    /**
     * An optional description of this security policy. Max size is 2048.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Fingerprint of this resource.
     */
    public /*out*/ readonly fingerprint: pulumi.Output<string>;
    /**
     * The name of the security policy.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
     * security policy, a default rule with action "allow" will be added. Structure is documented below.
     */
    public readonly rules: pulumi.Output<{ action: string, description?: string, match: { config: { srcIpRanges: string[] }, versionedExpr: string }, preview?: boolean, priority: number }[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;

    /**
     * Create a SecurityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityPolicyArgs | SecurityPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SecurityPolicyState = argsOrState as SecurityPolicyState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as SecurityPolicyArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/securityPolicy:SecurityPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityPolicy resources.
 */
export interface SecurityPolicyState {
    /**
     * An optional description of this security policy. Max size is 2048.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * The name of the security policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
     * security policy, a default rule with action "allow" will be added. Structure is documented below.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ action: pulumi.Input<string>, description?: pulumi.Input<string>, match: pulumi.Input<{ config: pulumi.Input<{ srcIpRanges: pulumi.Input<pulumi.Input<string>[]> }>, versionedExpr: pulumi.Input<string> }>, preview?: pulumi.Input<boolean>, priority: pulumi.Input<number> }>[]>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityPolicy resource.
 */
export interface SecurityPolicyArgs {
    /**
     * An optional description of this security policy. Max size is 2048.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the security policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
     * security policy, a default rule with action "allow" will be added. Structure is documented below.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ action: pulumi.Input<string>, description?: pulumi.Input<string>, match: pulumi.Input<{ config: pulumi.Input<{ srcIpRanges: pulumi.Input<pulumi.Input<string>[]> }>, versionedExpr: pulumi.Input<string> }>, preview?: pulumi.Input<boolean>, priority: pulumi.Input<number> }>[]>;
}
