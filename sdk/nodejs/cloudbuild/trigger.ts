// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Creates a new build trigger within GCR. For more information, see
 * [the official documentation](https://cloud.google.com/container-builder/docs/running-builds/automate-builds)
 * and
 * [API](https://godoc.org/google.golang.org/api/cloudbuild/v1#BuildTrigger).
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState): Trigger {
        return new Trigger(name, <any>state, { id });
    }

    /**
     * A build resource in the Container Builder API.
     * Structure is documented below. At a high
     * level, a `build` describes where to find source code, how to build it (for
     * example, the builder image to run on the source), and where to store
     * the built artifacts. Fields can include the following variables, which
     * will be expanded when the build is created:
     * * `$PROJECT_ID`: the project ID of the build.
     * * `$BUILD_ID`: the autogenerated ID of the build.
     * * `$REPO_NAME`: the source repository name specified by RepoSource.
     * * `$BRANCH_NAME`: the branch name specified by RepoSource.
     * * `$TAG_NAME`: the tag name specified by RepoSource.
     * * `$REVISION_ID` or `$COMMIT_SHA`: the commit SHA specified by RepoSource
     * or resolved from the specified branch or tag.
     * * `$SHORT_SHA`: first 7 characters of `$REVISION_ID` or `$COMMIT_SHA`.
     */
    public readonly build: pulumi.Output<{ images?: string[], steps?: { args?: string, name?: string }[], tags?: string[] } | undefined>;
    /**
     * A brief description of this resource.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Specify the path to a Cloud Build configuration file
     * in the Git repo. This is mutually exclusive with `build`. This is typically
     * `cloudbuild.yaml` however it can be specified by the user.
     */
    public readonly filename: pulumi.Output<string | undefined>;
    /**
     * ID of the project that owns the Cloud Source Repository.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * Location of the source in a Google
     * Cloud Source Repository. Structure is documented below.
     */
    public readonly triggerTemplate: pulumi.Output<{ branchName?: string, commitSha?: string, dir?: string, project?: string, repoName?: string, tagName?: string } | undefined>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TriggerArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TriggerState = argsOrState as TriggerState | undefined;
            inputs["build"] = state ? state.build : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["triggerTemplate"] = state ? state.triggerTemplate : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            inputs["build"] = args ? args.build : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["triggerTemplate"] = args ? args.triggerTemplate : undefined;
        }
        super("gcp:cloudbuild/trigger:Trigger", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * A build resource in the Container Builder API.
     * Structure is documented below. At a high
     * level, a `build` describes where to find source code, how to build it (for
     * example, the builder image to run on the source), and where to store
     * the built artifacts. Fields can include the following variables, which
     * will be expanded when the build is created:
     * * `$PROJECT_ID`: the project ID of the build.
     * * `$BUILD_ID`: the autogenerated ID of the build.
     * * `$REPO_NAME`: the source repository name specified by RepoSource.
     * * `$BRANCH_NAME`: the branch name specified by RepoSource.
     * * `$TAG_NAME`: the tag name specified by RepoSource.
     * * `$REVISION_ID` or `$COMMIT_SHA`: the commit SHA specified by RepoSource
     * or resolved from the specified branch or tag.
     * * `$SHORT_SHA`: first 7 characters of `$REVISION_ID` or `$COMMIT_SHA`.
     */
    readonly build?: pulumi.Input<{ images?: pulumi.Input<pulumi.Input<string>[]>, steps?: pulumi.Input<{ args?: pulumi.Input<string>, name?: pulumi.Input<string> }[]>, tags?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * A brief description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specify the path to a Cloud Build configuration file
     * in the Git repo. This is mutually exclusive with `build`. This is typically
     * `cloudbuild.yaml` however it can be specified by the user.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * ID of the project that owns the Cloud Source Repository.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Location of the source in a Google
     * Cloud Source Repository. Structure is documented below.
     */
    readonly triggerTemplate?: pulumi.Input<{ branchName?: pulumi.Input<string>, commitSha?: pulumi.Input<string>, dir?: pulumi.Input<string>, project?: pulumi.Input<string>, repoName?: pulumi.Input<string>, tagName?: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * A build resource in the Container Builder API.
     * Structure is documented below. At a high
     * level, a `build` describes where to find source code, how to build it (for
     * example, the builder image to run on the source), and where to store
     * the built artifacts. Fields can include the following variables, which
     * will be expanded when the build is created:
     * * `$PROJECT_ID`: the project ID of the build.
     * * `$BUILD_ID`: the autogenerated ID of the build.
     * * `$REPO_NAME`: the source repository name specified by RepoSource.
     * * `$BRANCH_NAME`: the branch name specified by RepoSource.
     * * `$TAG_NAME`: the tag name specified by RepoSource.
     * * `$REVISION_ID` or `$COMMIT_SHA`: the commit SHA specified by RepoSource
     * or resolved from the specified branch or tag.
     * * `$SHORT_SHA`: first 7 characters of `$REVISION_ID` or `$COMMIT_SHA`.
     */
    readonly build?: pulumi.Input<{ images?: pulumi.Input<pulumi.Input<string>[]>, steps?: pulumi.Input<{ args?: pulumi.Input<string>, name?: pulumi.Input<string> }[]>, tags?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * A brief description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specify the path to a Cloud Build configuration file
     * in the Git repo. This is mutually exclusive with `build`. This is typically
     * `cloudbuild.yaml` however it can be specified by the user.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * ID of the project that owns the Cloud Source Repository.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Location of the source in a Google
     * Cloud Source Repository. Structure is documented below.
     */
    readonly triggerTemplate?: pulumi.Input<{ branchName?: pulumi.Input<string>, commitSha?: pulumi.Input<string>, dir?: pulumi.Input<string>, project?: pulumi.Input<string>, repoName?: pulumi.Input<string>, tagName?: pulumi.Input<string> }>;
}
