# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class DatabaseIAMPolicy(pulumi.CustomResource):
    """
    Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
    
    * `google_spanner_database_iam_policy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
    
    ~> **Warning:** It's entirely possibly to lock yourself out of your database using `google_spanner_database_iam_policy`. Any permissions granted by default will be removed unless you include them in your config.
    
    * `google_spanner_database_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
    * `google_spanner_database_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
    
    ~> **Note:** `google_spanner_database_iam_policy` **cannot** be used in conjunction with `google_spanner_database_iam_binding` and `google_spanner_database_iam_member` or they will fight over what your policy should be.
    
    ~> **Note:** `google_spanner_database_iam_binding` resources **can be** used in conjunction with `google_spanner_database_iam_member` resources **only if** they do not grant privilege to the same role.
    """
    def __init__(__self__, __name__, __opts__=None, database=None, instance=None, policy_data=None, project=None):
        """Create a DatabaseIAMPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not database:
            raise TypeError('Missing required property database')
        elif not isinstance(database, basestring):
            raise TypeError('Expected property database to be a basestring')
        __self__.database = database
        """
        The name of the Spanner database.
        """
        __props__['database'] = database

        if not instance:
            raise TypeError('Missing required property instance')
        elif not isinstance(instance, basestring):
            raise TypeError('Expected property instance to be a basestring')
        __self__.instance = instance
        """
        The name of the Spanner instance the database belongs to.
        """
        __props__['instance'] = instance

        if not policy_data:
            raise TypeError('Missing required property policy_data')
        elif not isinstance(policy_data, basestring):
            raise TypeError('Expected property policy_data to be a basestring')
        __self__.policy_data = policy_data
        """
        The policy data generated by
        a `google_iam_policy` data source.
        """
        __props__['policyData'] = policy_data

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        __self__.etag = pulumi.runtime.UNKNOWN
        """
        (Computed) The etag of the database's IAM policy.
        """

        super(DatabaseIAMPolicy, __self__).__init__(
            'gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'database' in outs:
            self.database = outs['database']
        if 'etag' in outs:
            self.etag = outs['etag']
        if 'instance' in outs:
            self.instance = outs['instance']
        if 'policyData' in outs:
            self.policy_data = outs['policyData']
        if 'project' in outs:
            self.project = outs['project']
