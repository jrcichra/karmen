const karmen = require("karmen-node-client");
let k = new karmen.Client();
k.registerContainer();
k.registerEvent('hi');
k.registerAction('print', function () {
    console.log("I'm saying something from an action!");
})
k.emitEvent('hi', { "nodejs": "iscool" });