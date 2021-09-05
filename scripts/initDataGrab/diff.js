let keebmap = require('../../src/vendorData.json');
let keebtalk = require('./output.json');

keebmap = keebmap.map((v) => v.name.toLowerCase());
keebtalk = keebtalk.map((v) => v.name.toLowerCase());

for (v of keebmap) {
  if (!keebtalk.includes(v)) {
    console.log(v);
  }
}
