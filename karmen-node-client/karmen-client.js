
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
const carrier = require('carrier');
const log4js = require('log4js');
const logger = log4js.getLogger();
logger.level = "debug";

class Message {
    makeRegisterContainer(container_name) {
        this.name = container_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = undefined;
        this.type = REGISTERCONTAINER;
        this.container_name = container_name;
        this.id = undefined;
    }
    makeRegisterEvent(container_name, event_name) {
        this.name = event_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = undefined;
        this.type = REGISTEREVENT;
        this.container_name = container_name;
        this.id = undefined;
    }
    makeRegisterAction(container_name, action_name) {
        this.name = action_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = undefined;
        this.type = REGISTERACTION;
        this.container_name = container_name;
        this.id = undefined;
    }
    makeEmitEvent(container_name, event_name, params) {
        this.name = event_name;
        this.timestamp = new Date().valueOf();
        this.response_code = OK;
        this.params = params;
        this.type = EMITEVENT;
        this.container_name = container_name;
        this.id = undefined;
    }
    makeActionResponse(container_name, action_name, rc) {
        this.name = action_name;
        this.timestamp = new Date().valueOf();
        this.response_code = rc;
        this.params = undefined;
        this.type = TRIGGERACTIONRESPONSE;
        this.container_name = container_name;
        this.id = undefined;
    }
    toJSONStr() {
        let j = {
            "type": this.type,
            "timestamp": this.timestamp,
            "container_name": this.container_name,
            "name": this.name,
            "response_code": this.response_code,
            "params": this.params,
            "id": this.id
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
        this.status = this.PASS;
    }
    Fail() {
        this.status = this.FAIL;
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
        this.carrier = carrier.carry(this.socket);      //carrier will wrap the socket and won't fire an event until a line comes through
        this.hostname = os.hostname();
        this.actions = {};
        //Connection logic
        this.socket.connect(this.port, this.host, () => {
            logger.info(`Connected to ${this.host} on port ${this.port}`);
        });
        //Handle data logic
        this.carrier.on('line', (data) => {
            logger.debug(`Received: '${data}'`);
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
                    this.handleEmitEventResponse(j);
                    break;
                case REGISTERCONTAINERRESPONSE:
                    this.handleRegisterContainerResponse(j);
                    break;
                case REGISTEREVENTRESPONSE:
                    this.handleRegisterEventResponse(j);
                    break;
                case REGISTERACTIONRESPONSE:
                    this.handleRegisterActionResponse(h);
                default:
                    logger.error(`Got unexpected message from karmen: ${j}`);
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
        let m = new Message();
        if (r.getResult()) {
            m.makeActionResponse(this.hostname, message['name'], OK);
        } else {
            m.makeActionResponse(this.hostname, message['name'], ERROR);
        }
        this.send(m.toJSONStr());
    }

    handleEmitEventResponse(j) {
        logger.info(`Event ${j['name']} finished with a return code of ${j['response_code']}`);
    }

    handleRegisterEventResponse(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j['response_code'] != OK) {
            logger.error(`While registering the event, we got a bad return code: ${j['response_code']}`);
        } else {
            logger.info(`Sucessfully registered event ${j['name']}`);
        }
    }

    handleRegisterActionResponse(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j['response_code'] != OK) {
            logger.error(`While registering the action, we got a bad return code: ${j['response_code']}`);
        } else {
            logger.info(`Sucessfully registered action ${j['name']}`);
        }
    }

    handleTriggerAction(j) {
        let r = Result();
        logger.info(`Params before processing:${j['params']}`);
        let params = this.processParams(j['params']);
        logger.info(`Params after processing:${params}`);
        //Spawn the event and send an action response after its completed
        (async () => {
            await this.actions[j['name']](params, r);
            await this.sendActionResponse(j, r);
        })();
    }

    handleDispatchedEvent(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j['response_code'] != OK) {
            logger.error(`While emitting the event ${j['name']}, we got a bad return code: ${j['response_code']}`);
        } else {
            logger.info(`Sucessfully emitted event ${j['name']}`);
        }
    }

    handleRegisterContainerResponse(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j['response_code'] != OK) {
            logger.error(`While registering the container, we got a bad return code: ${j['response_code']}`);
        } else {
            logger.info("Sucessfully registered container");
        }
    }

    send(string) {
        this.socket.write(string);
    }

    registerContainer() {
        let m = new Message();
        m.makeRegisterContainer(this.hostname);
        this.send(m.toJSONStr());
    }

    registerEvent(event_name) {
        let m = new Message();
        m.makeRegisterEvent(this.hostname, event_name);
        this.send(m.toJSONStr());
    }

    registerAction(action_name, action_function) {
        let m = new Message();
        m.makeRegisterAction(this.hostname, action_name);
        this.send(m.toJSONStr());
    }

    emitEvent(event_name, params = {}) {
        let m = new Message();
        m.makeEmitEvent(this.hostname, event_name, params);
        this.send(m.toJSONStr());
    }


}

module.exports = {
    Client: Client,
    Result: Result
}