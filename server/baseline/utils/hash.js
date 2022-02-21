
function bitGroups(arrays) {
  let bitGroups = {};
  for (const [array, name] of arrays) {
    let bitGroup = '';
    for (const field of array) {
      bitGroup = bitGroup.concat(field);
    };
    if (bitGroup.length < 128) {
      bitGroup = bitGroup.concat(Array(128 - bitGroup.length).fill('0').join(''));
    };
    bitGroups[name] = bitGroup;
  };
  return bitGroups;
};

function hashToIntPair(hash) {
  return [BigInt('0x' + hash.slice(0, 32)).toString(), BigInt('0x' + hash.slice(32, 64)).toString()];
};


module.exports = {
  bitGroups,
  hashToIntPair,
};