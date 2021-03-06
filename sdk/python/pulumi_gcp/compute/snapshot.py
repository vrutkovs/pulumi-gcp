# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Snapshot(pulumi.CustomResource):
    """
    Creates a new snapshot of a disk within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/snapshots).
    """
    def __init__(__self__, __name__, __opts__=None, labels=None, name=None, project=None, snapshot_encryption_key_raw=None, source_disk=None, source_disk_encryption_key_raw=None, zone=None):
        """Create a Snapshot resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['project'] = project

        __props__['snapshot_encryption_key_raw'] = snapshot_encryption_key_raw

        if not source_disk:
            raise TypeError('Missing required property source_disk')
        __props__['source_disk'] = source_disk

        __props__['source_disk_encryption_key_raw'] = source_disk_encryption_key_raw

        __props__['zone'] = zone

        __props__['label_fingerprint'] = None
        __props__['self_link'] = None
        __props__['snapshot_encryption_key_sha256'] = None
        __props__['source_disk_encryption_key_sha256'] = None
        __props__['source_disk_link'] = None

        super(Snapshot, __self__).__init__(
            'gcp:compute/snapshot:Snapshot',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

