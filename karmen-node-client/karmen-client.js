
//Global constants

const REGISTERCONTAINER = "register-container"
const REGISTERCONTAINERRESPONSE = "register-container-response"
const REGISTEREVENT = "register-event"
const REGISTEREVENTRESPONSE = "register-event-response"
const REGISTERACTION = "register-action"
const REGISTERACTIONRESPONSE = "register-action-response"
const EMITEVENT = "emit-event"
const EMITEVENTRESPONSE = "emit-event-response"
const TRIGGERACTION = "trigger-action"
const TRIGGERACTIONRESPONSE = "trigger-action-response"
const DISPATCHEDEVENT = "dispatched-event"
const OK = 200
const ERROR = 503
const ONLINE = "online"
const OFFLINE = "offline"

const os = require('os');
const net = require('net');
const log4js = require("log4js");
const { finished } = require('stream')
const logger = log4js.getLogger();
logger.level = "debug";

class Message {
    makeRegisterContainer(container_name) {
        this.name = container_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = undefined;
        self.type = REGISTERCONTAINER;
        self.container_name = container_name;
        self.id = undefined;
    }
    makeRegisterEvent(container_name, event_name) {
        this.name = event_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = undefined;
        self.type = REGISTEREVENT;
        self.container_name = container_name;
        self.id = undefined;
    }
    makeRegisterAction(container_name, action_name) {
        this.name = action_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = undefined;
        self.type = REGISTERACTION;
        self.container_name = container_name;
        self.id = undefined;
    }
    makeEmitEvent(container_name, event_name, params) {
        this.name = event_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = params;
        self.type = EMITEVENT;
        self.container_name = container_name;
        self.id = undefined;
    }
    makeActionResponse(container_name, action_name, rc) {
        this.name = action_name;
        this.timestamp = new Date().valueOf();
        this.response_code = rc;
        this.params = undefined;
        self.type = TRIGGERACTIONRESPONSE;
        self.container_name = container_name;
        self.id = undefined;
    }
    toJSONStr() {
        let j = {
            "type": self.type,
            "timestamp": self.timestamp,
            "container_name": self.container_name,
            "name": self.name,
            "response_code": self.response_code,
            "params": self.params,
            "id": self.id
        }
        return JSON.stringify(j);
    }
}

class Result {
    constructor() {
        this.PASS = true;
        this.FAIL = false;
        this.status = this.PASS;
    }
    Pass() {
        this.status = self.PASS;
    }
    Fail() {
        this.status = self.FAIL;
    }
    getResult() {
        return this.status;
    }
}

class Client {
    constructor(host = "karmen", port = 8080) {
        this.host = host;
        this.port = port;
        this.socket = new net.Socket();
        this.hostname = os.hostname();
        // this.inboundQueues = {
        //     REGISTERCONTAINERRESPONSE: queue(),
        //     REGISTEREVENTRESPONSE: queue(),
        //     REGISTERACTIONRESPONSE: queue(),
        //     EMITEVENTRESPONSE: queue(),
        //     DISPATCHEDEVENT: queue()
        // };
        // this.outboundQueue = queue();
        this.actions = {};
        this.currentEvents = {};
        //Connection logic
        client.connect(this.port, this.host, () => {
            logger.info(`Connected to ${this.host} on port ${this.port}`);
        });
        //Handle data logic
        this.socket.on('data', (data) => {
            logger.debug(`Received: ${data}`);
            let j = JSON.parse(data);
            //Break out which function processes this data based on type
            switch (j['type']) {
                case TRIGGERACTION:
                    this.handleTriggerAction(j);
                    break;
                case DISPATCHEDEVENT:
                    this.handleDispatchedEvent(j);
                    break;
                case EMITEVENTRESPONSE:
                    // Tell the event that the event is done
                    this.eventQueues()
                    break;
                default:
                // Put it in the inbound queue
            }
        });
    }

    processParams(params) {
        let res = {};
        for (const [key, value] of Object.entries(params)) {
            res[key] = params[key]['value'];
        }
    }

    sendActionResponse(message, r) {
        logger.debug(`sendActionResponse got an r of ${r.getResult()}`);
        let m = Message();
        if (r.getResult()) {
            m.makeActionResponse(this.hostname, message['name'], OK);
        } else {
            m.makeActionResponse(this.hostname, message['name'], ERROR);
        }
        this.send(m.toJSONStr());
    }

    handleTriggerAction(j) {
        let r = Result();
        logger.info(`Params before processing:${j['params']}`);
        let params = this.processParams(j['params']);
        logger.info(`Params after processing:${params}`);
        // Use a promise to handle their custom action, then to get the result back
        new Promise(this.actions[j['name']](j, r)).then((j, r) => {
            this.sendActionResponse(j, r);
        });

    }

    handleDispatchedEvent(j, message) {
        let id = message['id'];
        logger.info(`Event ${j['name']} was dispatched to us!`);
        //Figure out when the event is over

        logger.info(`Event ${j['name']} finished with a return code of ${}`)
    }

    send(string) {
        this.socket.write(string);
    }

}