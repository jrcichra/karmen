# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: karmen.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='karmen.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x0ckarmen.proto\"\x17\n\x04Ping\x12\x0f\n\x07message\x18\x01 \x01(\t\"\x17\n\x04Pong\x12\x0f\n\x07message\x18\x01 \x01(\t\"\xef\x01\n\x0fRegisterRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x03\x12,\n\x06\x65vents\x18\x03 \x03(\x0b\x32\x1c.RegisterRequest.EventsEntry\x12.\n\x07\x61\x63tions\x18\x04 \x03(\x0b\x32\x1d.RegisterRequest.ActionsEntry\x1a-\n\x0b\x45ventsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a.\n\x0c\x41\x63tionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"N\n\x10RegisterResponse\x12!\n\x07request\x18\x01 \x01(\x0b\x32\x10.RegisterRequest\x12\x17\n\x06result\x18\x02 \x01(\x0b\x32\x07.Result\"v\n\x06Result\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x03\x12+\n\nparameters\x18\x02 \x03(\x0b\x32\x17.Result.ParametersEntry\x1a\x31\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x8c\x01\n\x05\x45vent\x12\x11\n\teventName\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x03\x12*\n\nparameters\x18\x03 \x03(\x0b\x32\x16.Event.ParametersEntry\x1a\x31\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x8f\x01\n\x06\x41\x63tion\x12\x12\n\nactionName\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x03\x12+\n\nparameters\x18\x03 \x03(\x0b\x32\x17.Action.ParametersEntry\x1a\x31\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"M\n\rActionRequest\x12\x17\n\x06\x61\x63tion\x18\x01 \x01(\x0b\x32\x07.Action\x12\x0c\n\x04uuid\x18\x02 \x01(\t\x12\x15\n\rrequesterName\x18\x03 \x01(\t\"\\\n\x0e\x41\x63tionResponse\x12\x1f\n\x07request\x18\x01 \x01(\x0b\x32\x0e.ActionRequest\x12\x17\n\x06result\x18\x02 \x01(\x0b\x32\x07.Result\x12\x10\n\x08hostname\x18\x03 \x01(\t\"\xb0\x01\n\x0c\x45ventRequest\x12\x15\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x06.Event\x12\x0c\n\x04uuid\x18\x02 \x01(\t\x12\x15\n\rrequesterName\x18\x03 \x01(\t\x12\x31\n\nparameters\x18\x04 \x03(\x0b\x32\x1d.EventRequest.ParametersEntry\x1a\x31\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"H\n\rEventResponse\x12\x1e\n\x07request\x18\x01 \x01(\x0b\x32\r.EventRequest\x12\x17\n\x06result\x18\x02 \x01(\x0b\x32\x07.Result2\xb8\x01\n\x06Karmen\x12/\n\x08Register\x12\x10.RegisterRequest\x1a\x11.RegisterResponse\x12*\n\tEmitEvent\x12\r.EventRequest\x1a\x0e.EventResponse\x12\x37\n\x10\x41\x63tionDispatcher\x12\x0f.ActionResponse\x1a\x0e.ActionRequest(\x01\x30\x01\x12\x18\n\x08PingPong\x12\x05.Ping\x1a\x05.PongB\x1cZ\x1agithub.com/jrcichra/karmenb\x06proto3')
)




_PING = _descriptor.Descriptor(
  name='Ping',
  full_name='Ping',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='Ping.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=16,
  serialized_end=39,
)


_PONG = _descriptor.Descriptor(
  name='Pong',
  full_name='Pong',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='Pong.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=41,
  serialized_end=64,
)


_REGISTERREQUEST_EVENTSENTRY = _descriptor.Descriptor(
  name='EventsEntry',
  full_name='RegisterRequest.EventsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='RegisterRequest.EventsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='RegisterRequest.EventsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=213,
  serialized_end=258,
)

_REGISTERREQUEST_ACTIONSENTRY = _descriptor.Descriptor(
  name='ActionsEntry',
  full_name='RegisterRequest.ActionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='RegisterRequest.ActionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='RegisterRequest.ActionsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=260,
  serialized_end=306,
)

_REGISTERREQUEST = _descriptor.Descriptor(
  name='RegisterRequest',
  full_name='RegisterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='RegisterRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='RegisterRequest.timestamp', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='events', full_name='RegisterRequest.events', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='actions', full_name='RegisterRequest.actions', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_REGISTERREQUEST_EVENTSENTRY, _REGISTERREQUEST_ACTIONSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=67,
  serialized_end=306,
)


_REGISTERRESPONSE = _descriptor.Descriptor(
  name='RegisterResponse',
  full_name='RegisterResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='request', full_name='RegisterResponse.request', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='result', full_name='RegisterResponse.result', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=308,
  serialized_end=386,
)


_RESULT_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='Result.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='Result.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='Result.ParametersEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=457,
  serialized_end=506,
)

_RESULT = _descriptor.Descriptor(
  name='Result',
  full_name='Result',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='Result.code', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parameters', full_name='Result.parameters', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RESULT_PARAMETERSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=388,
  serialized_end=506,
)


_EVENT_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='Event.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='Event.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='Event.ParametersEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=457,
  serialized_end=506,
)

_EVENT = _descriptor.Descriptor(
  name='Event',
  full_name='Event',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='eventName', full_name='Event.eventName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='Event.timestamp', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parameters', full_name='Event.parameters', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_EVENT_PARAMETERSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=509,
  serialized_end=649,
)


_ACTION_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='Action.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='Action.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='Action.ParametersEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=457,
  serialized_end=506,
)

_ACTION = _descriptor.Descriptor(
  name='Action',
  full_name='Action',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='actionName', full_name='Action.actionName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='Action.timestamp', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parameters', full_name='Action.parameters', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ACTION_PARAMETERSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=652,
  serialized_end=795,
)


_ACTIONREQUEST = _descriptor.Descriptor(
  name='ActionRequest',
  full_name='ActionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='action', full_name='ActionRequest.action', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uuid', full_name='ActionRequest.uuid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='requesterName', full_name='ActionRequest.requesterName', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=797,
  serialized_end=874,
)


_ACTIONRESPONSE = _descriptor.Descriptor(
  name='ActionResponse',
  full_name='ActionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='request', full_name='ActionResponse.request', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='result', full_name='ActionResponse.result', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hostname', full_name='ActionResponse.hostname', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=876,
  serialized_end=968,
)


_EVENTREQUEST_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='EventRequest.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='EventRequest.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='EventRequest.ParametersEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=457,
  serialized_end=506,
)

_EVENTREQUEST = _descriptor.Descriptor(
  name='EventRequest',
  full_name='EventRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='EventRequest.event', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uuid', full_name='EventRequest.uuid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='requesterName', full_name='EventRequest.requesterName', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parameters', full_name='EventRequest.parameters', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_EVENTREQUEST_PARAMETERSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=971,
  serialized_end=1147,
)


_EVENTRESPONSE = _descriptor.Descriptor(
  name='EventResponse',
  full_name='EventResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='request', full_name='EventResponse.request', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='result', full_name='EventResponse.result', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1149,
  serialized_end=1221,
)

_REGISTERREQUEST_EVENTSENTRY.containing_type = _REGISTERREQUEST
_REGISTERREQUEST_ACTIONSENTRY.containing_type = _REGISTERREQUEST
_REGISTERREQUEST.fields_by_name['events'].message_type = _REGISTERREQUEST_EVENTSENTRY
_REGISTERREQUEST.fields_by_name['actions'].message_type = _REGISTERREQUEST_ACTIONSENTRY
_REGISTERRESPONSE.fields_by_name['request'].message_type = _REGISTERREQUEST
_REGISTERRESPONSE.fields_by_name['result'].message_type = _RESULT
_RESULT_PARAMETERSENTRY.containing_type = _RESULT
_RESULT.fields_by_name['parameters'].message_type = _RESULT_PARAMETERSENTRY
_EVENT_PARAMETERSENTRY.containing_type = _EVENT
_EVENT.fields_by_name['parameters'].message_type = _EVENT_PARAMETERSENTRY
_ACTION_PARAMETERSENTRY.containing_type = _ACTION
_ACTION.fields_by_name['parameters'].message_type = _ACTION_PARAMETERSENTRY
_ACTIONREQUEST.fields_by_name['action'].message_type = _ACTION
_ACTIONRESPONSE.fields_by_name['request'].message_type = _ACTIONREQUEST
_ACTIONRESPONSE.fields_by_name['result'].message_type = _RESULT
_EVENTREQUEST_PARAMETERSENTRY.containing_type = _EVENTREQUEST
_EVENTREQUEST.fields_by_name['event'].message_type = _EVENT
_EVENTREQUEST.fields_by_name['parameters'].message_type = _EVENTREQUEST_PARAMETERSENTRY
_EVENTRESPONSE.fields_by_name['request'].message_type = _EVENTREQUEST
_EVENTRESPONSE.fields_by_name['result'].message_type = _RESULT
DESCRIPTOR.message_types_by_name['Ping'] = _PING
DESCRIPTOR.message_types_by_name['Pong'] = _PONG
DESCRIPTOR.message_types_by_name['RegisterRequest'] = _REGISTERREQUEST
DESCRIPTOR.message_types_by_name['RegisterResponse'] = _REGISTERRESPONSE
DESCRIPTOR.message_types_by_name['Result'] = _RESULT
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
DESCRIPTOR.message_types_by_name['Action'] = _ACTION
DESCRIPTOR.message_types_by_name['ActionRequest'] = _ACTIONREQUEST
DESCRIPTOR.message_types_by_name['ActionResponse'] = _ACTIONRESPONSE
DESCRIPTOR.message_types_by_name['EventRequest'] = _EVENTREQUEST
DESCRIPTOR.message_types_by_name['EventResponse'] = _EVENTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Ping = _reflection.GeneratedProtocolMessageType('Ping', (_message.Message,), dict(
  DESCRIPTOR = _PING,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:Ping)
  ))
_sym_db.RegisterMessage(Ping)

Pong = _reflection.GeneratedProtocolMessageType('Pong', (_message.Message,), dict(
  DESCRIPTOR = _PONG,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:Pong)
  ))
_sym_db.RegisterMessage(Pong)

RegisterRequest = _reflection.GeneratedProtocolMessageType('RegisterRequest', (_message.Message,), dict(

  EventsEntry = _reflection.GeneratedProtocolMessageType('EventsEntry', (_message.Message,), dict(
    DESCRIPTOR = _REGISTERREQUEST_EVENTSENTRY,
    __module__ = 'karmen_pb2'
    # @@protoc_insertion_point(class_scope:RegisterRequest.EventsEntry)
    ))
  ,

  ActionsEntry = _reflection.GeneratedProtocolMessageType('ActionsEntry', (_message.Message,), dict(
    DESCRIPTOR = _REGISTERREQUEST_ACTIONSENTRY,
    __module__ = 'karmen_pb2'
    # @@protoc_insertion_point(class_scope:RegisterRequest.ActionsEntry)
    ))
  ,
  DESCRIPTOR = _REGISTERREQUEST,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:RegisterRequest)
  ))
_sym_db.RegisterMessage(RegisterRequest)
_sym_db.RegisterMessage(RegisterRequest.EventsEntry)
_sym_db.RegisterMessage(RegisterRequest.ActionsEntry)

RegisterResponse = _reflection.GeneratedProtocolMessageType('RegisterResponse', (_message.Message,), dict(
  DESCRIPTOR = _REGISTERRESPONSE,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:RegisterResponse)
  ))
_sym_db.RegisterMessage(RegisterResponse)

Result = _reflection.GeneratedProtocolMessageType('Result', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _RESULT_PARAMETERSENTRY,
    __module__ = 'karmen_pb2'
    # @@protoc_insertion_point(class_scope:Result.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _RESULT,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:Result)
  ))
_sym_db.RegisterMessage(Result)
_sym_db.RegisterMessage(Result.ParametersEntry)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _EVENT_PARAMETERSENTRY,
    __module__ = 'karmen_pb2'
    # @@protoc_insertion_point(class_scope:Event.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _EVENT,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:Event)
  ))
_sym_db.RegisterMessage(Event)
_sym_db.RegisterMessage(Event.ParametersEntry)

Action = _reflection.GeneratedProtocolMessageType('Action', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _ACTION_PARAMETERSENTRY,
    __module__ = 'karmen_pb2'
    # @@protoc_insertion_point(class_scope:Action.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _ACTION,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:Action)
  ))
_sym_db.RegisterMessage(Action)
_sym_db.RegisterMessage(Action.ParametersEntry)

ActionRequest = _reflection.GeneratedProtocolMessageType('ActionRequest', (_message.Message,), dict(
  DESCRIPTOR = _ACTIONREQUEST,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:ActionRequest)
  ))
_sym_db.RegisterMessage(ActionRequest)

ActionResponse = _reflection.GeneratedProtocolMessageType('ActionResponse', (_message.Message,), dict(
  DESCRIPTOR = _ACTIONRESPONSE,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:ActionResponse)
  ))
_sym_db.RegisterMessage(ActionResponse)

EventRequest = _reflection.GeneratedProtocolMessageType('EventRequest', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _EVENTREQUEST_PARAMETERSENTRY,
    __module__ = 'karmen_pb2'
    # @@protoc_insertion_point(class_scope:EventRequest.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _EVENTREQUEST,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:EventRequest)
  ))
_sym_db.RegisterMessage(EventRequest)
_sym_db.RegisterMessage(EventRequest.ParametersEntry)

EventResponse = _reflection.GeneratedProtocolMessageType('EventResponse', (_message.Message,), dict(
  DESCRIPTOR = _EVENTRESPONSE,
  __module__ = 'karmen_pb2'
  # @@protoc_insertion_point(class_scope:EventResponse)
  ))
_sym_db.RegisterMessage(EventResponse)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\032github.com/jrcichra/karmen'))
_REGISTERREQUEST_EVENTSENTRY.has_options = True
_REGISTERREQUEST_EVENTSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_REGISTERREQUEST_ACTIONSENTRY.has_options = True
_REGISTERREQUEST_ACTIONSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_RESULT_PARAMETERSENTRY.has_options = True
_RESULT_PARAMETERSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_EVENT_PARAMETERSENTRY.has_options = True
_EVENT_PARAMETERSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_ACTION_PARAMETERSENTRY.has_options = True
_ACTION_PARAMETERSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_EVENTREQUEST_PARAMETERSENTRY.has_options = True
_EVENTREQUEST_PARAMETERSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))

_KARMEN = _descriptor.ServiceDescriptor(
  name='Karmen',
  full_name='Karmen',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=1224,
  serialized_end=1408,
  methods=[
  _descriptor.MethodDescriptor(
    name='Register',
    full_name='Karmen.Register',
    index=0,
    containing_service=None,
    input_type=_REGISTERREQUEST,
    output_type=_REGISTERRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='EmitEvent',
    full_name='Karmen.EmitEvent',
    index=1,
    containing_service=None,
    input_type=_EVENTREQUEST,
    output_type=_EVENTRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ActionDispatcher',
    full_name='Karmen.ActionDispatcher',
    index=2,
    containing_service=None,
    input_type=_ACTIONRESPONSE,
    output_type=_ACTIONREQUEST,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PingPong',
    full_name='Karmen.PingPong',
    index=3,
    containing_service=None,
    input_type=_PING,
    output_type=_PONG,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_KARMEN)

DESCRIPTOR.services_by_name['Karmen'] = _KARMEN

# @@protoc_insertion_point(module_scope)
