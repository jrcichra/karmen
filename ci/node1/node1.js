const karmen = require("karmen-node-client");
let k = new karmen.Client();
(async () => {
    await k.registerContainer();
    await k.registerEvent('hi');
    await console.log("about to register an action")
    await k.registerAction('print', function () {
        console.log("I'm saying something from an action!");
    });
    await k.emitEvent('hi', { "nodejs": "iscool" });
})();