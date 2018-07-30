# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetFolderResult(object):
    """
    A collection of values returned by getFolder.
    """
    def __init__(__self__, create_time=None, display_name=None, lifecycle_state=None, name=None, organization=None, parent=None, id=None):
        if create_time and not isinstance(create_time, basestring):
            raise TypeError('Expected argument create_time to be a basestring')
        __self__.create_time = create_time
        """
        Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        if display_name and not isinstance(display_name, basestring):
            raise TypeError('Expected argument display_name to be a basestring')
        __self__.display_name = display_name
        """
        The folder's display name.
        """
        if lifecycle_state and not isinstance(lifecycle_state, basestring):
            raise TypeError('Expected argument lifecycle_state to be a basestring')
        __self__.lifecycle_state = lifecycle_state
        """
        The Folder's current lifecycle state.
        """
        if name and not isinstance(name, basestring):
            raise TypeError('Expected argument name to be a basestring')
        __self__.name = name
        """
        The resource name of the Folder in the form `folders/{organization_id}`.
        """
        if organization and not isinstance(organization, basestring):
            raise TypeError('Expected argument organization to be a basestring')
        __self__.organization = organization
        """
        If `lookup_organization` is enable, the resource name of the Organization that the folder belongs.
        """
        if parent and not isinstance(parent, basestring):
            raise TypeError('Expected argument parent to be a basestring')
        __self__.parent = parent
        """
        The resource name of the parent Folder or Organization.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_folder(folder=None, lookup_organization=None):
    """
    Use this data source to get information about a Google Cloud Folder.
    
    ```hcl
    """
    __args__ = dict()

    __args__['folder'] = folder
    __args__['lookupOrganization'] = lookup_organization
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getFolder:getFolder', __args__)

    return GetFolderResult(
        create_time=__ret__.get('createTime'),
        display_name=__ret__.get('displayName'),
        lifecycle_state=__ret__.get('lifecycleState'),
        name=__ret__.get('name'),
        organization=__ret__.get('organization'),
        parent=__ret__.get('parent'),
        id=__ret__.get('id'))
