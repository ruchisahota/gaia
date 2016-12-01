# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class NamespaceContent(RESTObject):
    """ Represents a NamespaceContent in the 

        Notes:
            This object is used to delete the content of a namespace easily
    """

    def __init__(self, **kwargs):
        """ Initializes a NamespaceContent instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> namespacecontent = NamespaceContent(id=u'xxxx-xxx-xxx-xxx', name=u'NamespaceContent')
              >>> namespacecontent = NamespaceContent(data=my_dict)
        """

        super(NamespaceContent, self).__init__()

        # Read/Write Attributes
        
        self._contentid = None
        self._contenttype = None
        self._namespace = None
        
        self.expose_attribute(local_name="contentID", remote_name="contentID")
        self.expose_attribute(local_name="contentType", remote_name="contentType")
        self.expose_attribute(local_name="namespace", remote_name="namespace")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return ""
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        pass
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return namespacecontentIdentity

    # Properties
    @property
    def contentID(self):
        """ Get contentID value.

          Notes:
              ID of the content

              
        """
        return self._contentid

    @contentID.setter
    def contentID(self, value):
        """ Set contentID value.

          Notes:
              ID of the content

              
        """
        self._contentid = value
    
    @property
    def contentType(self):
        """ Get contentType value.

          Notes:
              Type of the content

              
        """
        return self._contenttype

    @contentType.setter
    def contentType(self, value):
        """ Set contentType value.

          Notes:
              Type of the content

              
        """
        self._contenttype = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              name of the namespace

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              name of the namespace

              
        """
        self._namespace = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("contentID", self.contentID)

        if err:
            errors.append(err)

        err = validate_required_string("contentType", self.contentType)

        if err:
            errors.append(err)

        err = validate_required_string("namespace", self.namespace)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # namespacecontentIdentity represents the Identity of the object
namespacecontentIdentity = {"name": "namespacecontent", "category": "namespacecontents", "constructor": NamespaceContent}