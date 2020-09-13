#!/usr/bin/env node
const karmen = require("karmen-node-client");
let k = new karmen.Client();
(async () => {
    await k.connect();
    console.log("Does this even run?");
    await k.registerContainer();
    await k.registerEvent('hi');
    console.log("about to register an action")
    await k.registerAction('print', function (params, result) {
        console.log("I'm saying something from an action!");
        console.log(`I got these params: ${JSON.stringify(params)}`);
        result.Pass();
    });
    await k.emitEvent('hi', { "nodejs": "iscool" });
})();