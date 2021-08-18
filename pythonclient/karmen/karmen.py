#!/usr/bin/python3 -u

import threading
import time
import queue
import socket
import grpc

import karmen.karmen_pb2 as pb
import karmen.karmen_pb2_grpc as pb_grpc


class Karmen:
    def __init__(self, name=socket.gethostname(), hostname="localhost", port=8080):
        super().__init__()
        self.name = name
        self.channel = grpc.insecure_channel(f"{hostname}:{port}")
        self.stub = pb_grpc.KarmenStub(self.channel)
        self.actions = {}

    def Pass(self) -> int:
        return 200

    def ping(self) -> str:
        result = self.stub.PingPong(pb.Ping(message="Python!"))
        return result.message

    def runEvent(self, name):
        event = pb.Event(eventName=name, timestamp=int(time.time()))
        result = self.stub.EmitEvent(pb.EventRequest(
            requesterName=self.name, event=event))
        return result

    def addAction(self, func, name):
        self.actions[name] = func

    def setupActions(self):
        send_queue = queue.SimpleQueue()

        # set up the two way connection
        recv = self.stub.ActionDispatcher(
            iter(send_queue.get, None))
        # send who we are
        send_queue.put(pb.ActionResponse(hostname=self.name))

        threading.Thread(target=self.handleActions,
                         args=(recv, send_queue)).start()

    def handleActions(self, recv, send_queue):
        while True:
            # blocking for actions
            msg = next(recv)

            # got an action
            print(f"Got an action!")
            print(msg)

            # run the action
            threading.Thread(target=self.handleAction,
                             args=(msg, send_queue)).start()

    def handleAction(self, msg, send_queue):
        # run the action
        print(f"Running action: {msg.action.actionName}")
        result = pb.ActionResponse()
        self.actions[msg.action.actionName](
            msg.action.parameters, result.result)
        print(f"Finished running action: {msg.action.actionName}")
        send_queue.put(result)

    def register(self) -> int:
        result = self.stub.Register(pb.RegisterRequest(
            name=self.name, timestamp=int(time.time())))
        self.setupActions()
        return result.result.code


if __name__ == "__main__":

    def sleep(parameters, result):
        print(f"Sleeping for {parameters['seconds']} seconds")
        time.sleep(int(parameters['seconds']))
        print(f"Done sleeping for {parameters['seconds']} seconds")
        result.code = 200

    k = Karmen(name="bob")
    print(k.ping())
    k.addAction(sleep, "sleep")
    k.register()
    print(k.runEvent("pleaseSleep"))
