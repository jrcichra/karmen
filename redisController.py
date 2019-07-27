import redis
from rejson import Client, Path
import logging
import json


class redisController:
    def __init__(self, hostname, port):
        # Set class variables
        self.hostname = hostname
        self.port = port
        # Create the basic redis connection
        logging.debug("Connecting to the database with hostname: " +
                      hostname + " on port: " + port + ".")
        self.db = Client(host=hostname, port=port, decode_responses=True)

    def registerContainer(self, obj):
        # Internal error if we somehow don't go through the if or else
        msg = "Internal registerContainer error"
        code = 503
        # Grab the timestamp in the packet
        timestamp = obj['timestamp']
        # Go the container being registered
        container = obj['data']['container']
        # Pull out the valuable attributes from this layer
        name = container['name']
        container_id = container['container_id']
        # Check if this container already exists in redis
        existing_container = json.loads(self.db.jsonget(container_id))
        logging.debug(
            "Checking for an existing container returned: " + existing_container)
        if existing_container != None:
            msg = "Container " + container + \
                " was already registered in redis at " + \
                existing_container['timestamp']
            code = 1
            logging.warning(msg)

        else:
            # Build a redis object
            robj = {
                container_id: {
                    'state': "online",
                    'name': name,
                    'timestamp': timestamp,
                    'events': [],
                    'actions': []
                }
            }
            self.db.jsonset(container_id, Path.rootPath(), robj)
            msg = "OK"
            code = 0
        return code, msg

    def registerEvent(self, obj):
        # Internal error if we somehow don't go through the if or else
        msg = "Internal registerEvent error"
        code = 504
        # Grab the timestamp in the packet
        timestamp = obj['timestamp']
        # Go the event being registered
        event = obj['data']['event']
        # Pull out the valuable attributes from this layer
        name = event['name']
        container_id = obj['container_id']
        # Check if this already exists in redis by first pulling all events for this container
        existing_events = json.loads(self.db.jsonget(id, Path('.events')))
        logging.debug(
            "Checking for existing events returned: " + existing_events)
        # See if the event name we are trying to register already exists
        if name in existing_events:
            msg = "Event " + name + \
                " is already registered in redis."
            code = 1
            logging.warning(msg)

        else:
            # Insert the event into the proper container's object
            robj = {
                'name': name
            }
            self.db.jsonarrappend(id, Path('.events'), robj)
            msg = "OK"
            code = 0
        return code, msg
