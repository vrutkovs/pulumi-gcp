# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Variavble(pulumi.CustomResource):
    """
    Manages a RuntimeConfig variable in Google Cloud. For more information, see the
    [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
    or the
    [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
    """
    def __init__(__self__, __name__, __opts__=None, name=None, parent=None, project=None, text=None, value=None):
        """Create a Variavble resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        if not parent:
            raise TypeError('Missing required property parent')
        __props__['parent'] = parent

        __props__['project'] = project

        __props__['text'] = text

        __props__['value'] = value

        __props__['update_time'] = None

        super(Variavble, __self__).__init__(
            'gcp:runtimeconfig/variavble:Variavble',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

