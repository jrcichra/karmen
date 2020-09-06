import karmen
import time


def action1(params, result):
    print("In action1!")
    result.Pass()


def action2(params, result):
    print("In action2!")
    result.Pass()


k = karmen.Client()
k.registerContainer()
k.registerEvent("event1")
k.registerEvent("event2")
k.registerAction("action1", action1)
k.registerAction("action2", action2)

time.sleep(5)
k.emitEvent("event1",params={"a":"parameter", "b": 5.4})
time.sleep(5)
k.emitEvent("event2")
time.sleep(5)
