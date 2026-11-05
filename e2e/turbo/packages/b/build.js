// Slow on purpose, so a cache hit is obvious.
const fs = require("fs");
const until = Date.now() + 5000;
while (Date.now() < until) {}
fs.mkdirSync("dist", { recursive: true });
fs.writeFileSync("dist/out.txt", "built b\n");
console.log("built b");
