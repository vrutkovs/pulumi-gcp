# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetAccountResult(object):
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, display_name=None, email=None, name=None, unique_id=None, id=None):
        if display_name and not isinstance(display_name, basestring):
            raise TypeError('Expected argument display_name to be a basestring')
        __self__.display_name = display_name
        """
        The display name for the service account.
        """
        if email and not isinstance(email, basestring):
            raise TypeError('Expected argument email to be a basestring')
        __self__.email = email
        """
        The e-mail address of the service account. This value
        should be referenced from any `google_iam_policy` data sources
        that would grant the service account privileges.
        """
        if name and not isinstance(name, basestring):
            raise TypeError('Expected argument name to be a basestring')
        __self__.name = name
        """
        The fully-qualified name of the service account.
        """
        if unique_id and not isinstance(unique_id, basestring):
            raise TypeError('Expected argument unique_id to be a basestring')
        __self__.unique_id = unique_id
        """
        The unique id of the service account.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_account(account_id=None, project=None):
    """
    Get the service account from a project. For more information see
    the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
    """
    __args__ = dict()

    __args__['accountId'] = account_id
    __args__['project'] = project
    __ret__ = pulumi.runtime.invoke('gcp:serviceAccount/getAccount:getAccount', __args__)

    return GetAccountResult(
        display_name=__ret__.get('displayName'),
        email=__ret__.get('email'),
        name=__ret__.get('name'),
        unique_id=__ret__.get('uniqueId'),
        id=__ret__.get('id'))
