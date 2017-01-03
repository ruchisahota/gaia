# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class APICheck(RESTObject):
    """ Represents a APICheck in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a APICheck instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> apicheck = APICheck(id=u'xxxx-xxx-xxx-xxx', name=u'APICheck')
              >>> apicheck = APICheck(data=my_dict)
        """

        super(APICheck, self).__init__()

        # Read/Write Attributes
        
        self._identity = None
        self._namespace = None
        self._token = None
        
        self.expose_attribute(local_name="identity", remote_name="identity")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="token", remote_name="token")

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
        return apicheckIdentity

    # Properties
    @property
    def identity(self):
        """ Get identity value.

          Notes:
              The identity to use to use to check the api authentication

              
        """
        return self._identity

    @identity.setter
    def identity(self, value):
        """ Set identity value.

          Notes:
              The identity to use to use to check the api authentication

              
        """
        self._identity = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              The namespace to use to check the api authentication.

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              The namespace to use to check the api authentication.

              
        """
        self._namespace = value
    
    @property
    def token(self):
        """ Get token value.

          Notes:
              The token to use to check api authentication

              
        """
        return self._token

    @token.setter
    def token(self, value):
        """ Set token value.

          Notes:
              The token to use to check api authentication

              
        """
        self._token = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("namespace", self.namespace)

        if err:
            errors.append(err)

        err = validate_required_string("token", self.token)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # apicheckIdentity represents the Identity of the object
apicheckIdentity = {"name": "apicheck", "category": "apichecks", "constructor": APICheck}