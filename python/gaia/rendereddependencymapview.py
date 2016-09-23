# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class RenderedDependencyMapView(RESTObject):
    """ Represents a RenderedDependencyMapView in the 

        Notes:
            Render some dependency map views from a dependency map
    """

    def __init__(self, **kwargs):
        """ Initializes a RenderedDependencyMapView instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> rendereddependencymapview = RenderedDependencyMapView(id=u'xxxx-xxx-xxx-xxx', name=u'RenderedDependencyMapView')
              >>> rendereddependencymapview = RenderedDependencyMapView(data=my_dict)
        """

        super(RenderedDependencyMapView, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._dependencymapview = None
        self._processingunittags = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="dependencyMapView", remote_name="dependencyMapView")
        self.expose_attribute(local_name="processingUnitTags", remote_name="processingUnitTags")

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
        return rendereddependencymapviewIdentity

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
    def dependencyMapView(self):
        """ Get dependencyMapView value.

          Notes:
              The dependencyMapView linked to the rendered dependency map view

              
        """
        return self._dependencymapview

    @dependencyMapView.setter
    def dependencyMapView(self, value):
        """ Set dependencyMapView value.

          Notes:
              The dependencyMapView linked to the rendered dependency map view

              
        """
        self._dependencymapview = value
    
    @property
    def processingUnitTags(self):
        """ Get processingUnitTags value.

          Notes:
              A map of the transient tags for the processing units

              
        """
        return self._processingunittags

    @processingUnitTags.setter
    def processingUnitTags(self, value):
        """ Set processingUnitTags value.

          Notes:
              A map of the transient tags for the processing units

              
        """
        self._processingunittags = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # rendereddependencymapviewIdentity represents the Identity of the object
rendereddependencymapviewIdentity = {"name": "rendereddependencymapview", "category": "rendereddependencymapviews", "constructor": RenderedDependencyMapView}