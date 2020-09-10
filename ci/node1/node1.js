const log = require('why-is-node-running')
const karmen = require("karmen-node-client");
let k = new karmen.Client();
(async () => {
    await k.connect();
    console.log("Does this even run?");
    await k.registerContainer();
    await k.registerEvent('hi');
    console.log("about to register an action")
    await k.registerAction('print', function () {
        console.log("I'm saying something from an action!");
    });
    await k.emitEvent('hi', { "nodejs": "iscool" });
})();

setTimeout(function () {
    log() // logs out active handles that are keeping node running
}, 100)