# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Layer(RESTObject):
    """ Represents a Layer in the 

        Notes:
            Layer of a docker image
    """

    def __init__(self, **kwargs):
        """ Initializes a Layer instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> layer = Layer(id=u'xxxx-xxx-xxx-xxx', name=u'Layer')
              >>> layer = Layer(data=my_dict)
        """

        super(Layer, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._name = None
        self._namespacename = None
        self._parentname = None
        self._severity = None
        self._vulnerabilities = None
        
        self.expose_attribute(local_name="id", remote_name="ID")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespacename", remote_name="namespaceName")
        self.expose_attribute(local_name="parentname", remote_name="parentName")
        self.expose_attribute(local_name="severity", remote_name="severity")
        self.expose_attribute(local_name="vulnerabilities", remote_name="vulnerabilities")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return self.ID

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        self.ID = ID

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return layerIdentity

    # Properties
    @property
    def ID(self):
        """ Get ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        return self._id

    @ID.setter
    def ID(self, value):
        """ Set ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        self._id = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the entity

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the entity

              
        """
        self._name = value
    
    @property
    def namespaceName(self):
        """ Get namespaceName value.

          Notes:
              NamespaceName is the name of the namespace

              
        """
        return self._namespacename

    @namespaceName.setter
    def namespaceName(self, value):
        """ Set namespaceName value.

          Notes:
              NamespaceName is the name of the namespace

              
        """
        self._namespacename = value
    
    @property
    def parentName(self):
        """ Get parentName value.

          Notes:
              ParentName is the name of parent layer

              
        """
        return self._parentname

    @parentName.setter
    def parentName(self, value):
        """ Set parentName value.

          Notes:
              ParentName is the name of parent layer

              
        """
        self._parentname = value
    
    @property
    def severity(self):
        """ Get severity value.

          Notes:
              Severity defines the severity level of the vulnerability

              
        """
        return self._severity

    @severity.setter
    def severity(self, value):
        """ Set severity value.

          Notes:
              Severity defines the severity level of the vulnerability

              
        """
        self._severity = value
    
    @property
    def vulnerabilities(self):
        """ Get vulnerabilities value.

          Notes:
              Vulnerabilities is the list of all the vulnerabilities of a layer

              
        """
        return self._vulnerabilities

    @vulnerabilities.setter
    def vulnerabilities(self, value):
        """ Set vulnerabilities value.

          Notes:
              Vulnerabilities is the list of all the vulnerabilities of a layer

              
        """
        self._vulnerabilities = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # layerIdentity represents the Identity of the object
layerIdentity = {"name": "layer", "category": "layers", "constructor": Layer}