# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import karmen_pb2 as karmen__pb2


class KarmenStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Register = channel.unary_unary(
        '/Karmen/Register',
        request_serializer=karmen__pb2.RegisterRequest.SerializeToString,
        response_deserializer=karmen__pb2.RegisterResponse.FromString,
        )
    self.EmitEvent = channel.unary_unary(
        '/Karmen/EmitEvent',
        request_serializer=karmen__pb2.EventRequest.SerializeToString,
        response_deserializer=karmen__pb2.EventResponse.FromString,
        )
    self.ActionDispatcher = channel.stream_stream(
        '/Karmen/ActionDispatcher',
        request_serializer=karmen__pb2.ActionResponse.SerializeToString,
        response_deserializer=karmen__pb2.ActionRequest.FromString,
        )
    self.PingPong = channel.unary_unary(
        '/Karmen/PingPong',
        request_serializer=karmen__pb2.Ping.SerializeToString,
        response_deserializer=karmen__pb2.Pong.FromString,
        )


class KarmenServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Register(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def EmitEvent(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ActionDispatcher(self, request_iterator, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PingPong(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_KarmenServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Register': grpc.unary_unary_rpc_method_handler(
          servicer.Register,
          request_deserializer=karmen__pb2.RegisterRequest.FromString,
          response_serializer=karmen__pb2.RegisterResponse.SerializeToString,
      ),
      'EmitEvent': grpc.unary_unary_rpc_method_handler(
          servicer.EmitEvent,
          request_deserializer=karmen__pb2.EventRequest.FromString,
          response_serializer=karmen__pb2.EventResponse.SerializeToString,
      ),
      'ActionDispatcher': grpc.stream_stream_rpc_method_handler(
          servicer.ActionDispatcher,
          request_deserializer=karmen__pb2.ActionResponse.FromString,
          response_serializer=karmen__pb2.ActionRequest.SerializeToString,
      ),
      'PingPong': grpc.unary_unary_rpc_method_handler(
          servicer.PingPong,
          request_deserializer=karmen__pb2.Ping.FromString,
          response_serializer=karmen__pb2.Pong.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Karmen', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))