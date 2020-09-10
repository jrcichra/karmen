
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
const Channel = require(`@nodeguy/channel`);
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
        this.channels = {
            DISPATCHEDEVENT: Channel(),
            // EMITEVENTRESPONSE: Channel(),
            REGISTERCONTAINERRESPONSE: Channel(),
            REGISTEREVENTRESPONSE: Channel(),
            REGISTERACTIONRESPONSE: Channel()
        };
        // Channels to keep things in sync...because I think like a Go/Python coder
    }

    async connect() {
        //Channel hack - this is awful, but I need to block until the socket is connected...
        let c = Channel();
        //Connection logic
        this.socket.connect(this.port, this.host, () => {
            logger.info(`Connected to ${this.host} on port ${this.port}`);
            c.push(true);
        });
        //Handle data logic
        this.carrier.on('line', (data) => {
            // async so we can use await down the chain
            (async () => {
                logger.debug(`Received: '${data}'`);
                let j = JSON.parse(data);
                //Break out which function processes this data based on type
                switch (j.type) {
                    case TRIGGERACTION:
                        await this.handleTriggerAction(j);
                        break;
                    case DISPATCHEDEVENT:
                        await this.handleDispatchedEvent(j);
                        break;
                    case EMITEVENTRESPONSE:
                        await this.handleEmitEventResponse(j);
                        break;
                    case REGISTERCONTAINERRESPONSE:
                        await this.handleRegisterContainerResponse(j);
                        break;
                    case REGISTEREVENTRESPONSE:
                        await this.handleRegisterEventResponse(j);
                        break;
                    case REGISTERACTIONRESPONSE:
                        await this.handleRegisterActionResponse(h);
                        break;
                    default:
                        logger.error(`Got unexpected message from karmen: ${j}`);
                }
            });
        });
        await c.shift();
    }

    processParams(params) {
        let res = {};
        for (const [key, value] of Object.entries(params)) {
            res[key] = params[key].value;
        }
    }

    async sendActionResponse(message, r) {
        logger.debug(`sendActionResponse got an r of ${r.getResult()}`);
        let m = new Message();
        if (r.getResult()) {
            m.makeActionResponse(this.hostname, message.name, OK);
        } else {
            m.makeActionResponse(this.hostname, message.name, ERROR);
        }
        this.send(m.toJSONStr());
    }

    async handleEmitEventResponse(j) {
        logger.info(`Event ${j.name} finished with a return code of ${j.response_code}`);
    }

    async handleRegisterEventResponse(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j.response_code != OK) {
            logger.error(`While registering the event, we got a bad return code: ${j.response_code}`);
        } else {
            logger.info(`Sucessfully registered event ${j.name}`);
        }
        await this.channels.REGISTEREVENTRESPONSE.push(j);
    }

    async handleRegisterActionResponse(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j.response_code != OK) {
            logger.error(`While registering the action, we got a bad return code: ${j.response_code}`);
        } else {
            logger.info(`Sucessfully registered action ${j.name}`);
        }
        await this.channels.REGISTERACTIONRESPONSE.push(j);
    }

    async handleTriggerAction(j) {
        let r = Result();
        logger.info(`Params before processing:${j.params}`);
        let params = this.processParams(j.params);
        logger.info(`Params after processing:${params}`);
        logger.info(`Starting action: ${j.name}`);
        await this.actions[j.name](params, r);
        await this.sendActionResponse(j, r);
    }

    async handleDispatchedEvent(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j.response_code != OK) {
            logger.error(`While emitting the event ${j.name}, we got a bad return code: ${j.response_code}`);
        } else {
            logger.info(`Sucessfully emitted event ${j.name}`);
        }
        await this.channels.DISPATCHEDEVENT.push(j);
    }

    async handleRegisterContainerResponse(j) {
        logger.debug(`response=${JSON.stringify(j)}`);
        if (j.response_code != OK) {
            logger.error(`While registering the container, we got a bad return code: ${j.response_code}`);
        } else {
            logger.info("Sucessfully registered container");
        }
        await this.channels.REGISTERCONTAINERRESPONSE.push(j);
    }

    send(string) {
        logger.debug(`Sending: '${string}'`);
        this.socket.write(string);
    }

    async registerContainer() {
        let m = new Message();
        m.makeRegisterContainer(this.hostname);
        this.send(m.toJSONStr());
        await this.channels.REGISTERCONTAINERRESPONSE.shift();
    }

    async registerEvent(event_name) {
        let m = new Message();
        m.makeRegisterEvent(this.hostname, event_name);
        this.send(m.toJSONStr());
        await this.channels.REGISTEREVENTRESPONSE.shift();
    }

    async registerAction(action_name, action_function) {
        let m = new Message();
        m.makeRegisterAction(this.hostname, action_name);
        // Add action to map of actions
        this.actions[m.name] = action_function;
        this.send(m.toJSONStr());
        await this.channels.REGISTERACTIONRESPONSE.shift();
    }

    async emitEvent(event_name, params = {}) {
        let m = new Message();
        m.makeEmitEvent(this.hostname, event_name, params);
        this.send(m.toJSONStr());
        //emitEvent won't block...may need more here for JS to react, i.e a promise
    }
}

module.exports = {
    Client: Client,
    Result: Result
}