# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Image(RESTObject):
    """ Represents a Image in the 

        Notes:
            Image of a docker
    """

    def __init__(self, **kwargs):
        """ Initializes a Image instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> image = Image(id=u'xxxx-xxx-xxx-xxx', name=u'Image')
              >>> image = Image(data=my_dict)
        """

        super(Image, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._associatedtags = None
        self._name = None
        self._registry = None
        self._repository = None
        self._severity = None
        self._tag = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="registry", remote_name="registry")
        self.expose_attribute(local_name="repository", remote_name="repository")
        self.expose_attribute(local_name="severity", remote_name="severity")
        self.expose_attribute(local_name="tag", remote_name="tag")

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
        return imageIdentity

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
    def associatedTags(self):
        """ Get associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        return self._associatedtags

    @associatedTags.setter
    def associatedTags(self, value):
        """ Set associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        self._associatedtags = value
    
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
    def registry(self):
        """ Get registry value.

          Notes:
              Registry refers to the service that stores images

              
        """
        return self._registry

    @registry.setter
    def registry(self, value):
        """ Set registry value.

          Notes:
              Registry refers to the service that stores images

              
        """
        self._registry = value
    
    @property
    def repository(self):
        """ Get repository value.

          Notes:
              Repository is the name of the image repository

              
        """
        return self._repository

    @repository.setter
    def repository(self, value):
        """ Set repository value.

          Notes:
              Repository is the name of the image repository

              
        """
        self._repository = value
    
    @property
    def severity(self):
        """ Get severity value.

          Notes:
              Severity defines the severity level of the image

              
        """
        return self._severity

    @severity.setter
    def severity(self, value):
        """ Set severity value.

          Notes:
              Severity defines the severity level of the image

              
        """
        self._severity = value
    
    @property
    def tag(self):
        """ Get tag value.

          Notes:
              Tag is the tag of the image

              
        """
        return self._tag

    @tag.setter
    def tag(self, value):
        """ Set tag value.

          Notes:
              Tag is the tag of the image

              
        """
        self._tag = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        err = validate_required_string("registry", self.registry)

        if err:
            errors.append(err)

        err = validate_required_string("tag", self.tag)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # imageIdentity represents the Identity of the object
imageIdentity = {"name": "image", "category": "images", "constructor": Image}