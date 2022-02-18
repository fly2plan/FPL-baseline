// {
//   "ai": Array(14).fill('7'),
//   "fr": Array(2).fill('7'),
//   "tof": Array(2).fill('7'),
//   "noa": Array(4).fill('7'),
//   "toa": Array(8).fill('7'),
//   "wtc": Array(2).fill('7'),
//   "e_and_c": {
//       "rc_n_aae": {
//           "s": Array(4).fill('7'),
//           "a": Array(98).fill('7')
//       },
//       "surveillance": {
//           "s": Array(2).fill('7'),
//           "ssr": {
//               "mode_a_c": {
//                   "a": Array(4).fill('7')

//               },
//               "mode_s": {
//                   "a": Array(14).fill('7')
//               }
//           },
//           "ads": {
//               "c": {
//                   "a": Array(8).fill('7')
//               },
//               "b": {
//                   "a": Array(24).fill('7')
//               }
//           }
//       }
//   },
//   "t": Array(8).fill('7'),
//   "da": Array(8).fill('7'),
//   "cs": Array(10).fill('7'),
//   "cl": Array(10).fill('7'),
//   "r": Array(304).fill('7'),
//   "de": Array(8).fill('7'),
//   "teet": Array(8).fill('7'),
//   "ada": Array(8).fill('7'),
//   "adas": [Array(8).fill('7'), Array(8).fill('7')],
//   "oi": Array(152).fill('7'),
//   "e": Array(12).fill('7'),
//   "pob": Array(10).fill('7'),
//   "er": Array(8).fill('7'),
//   "se": Array(8).fill('7'),
//   "j": Array(20).fill('7'),
//   "d": Array(20).fill('7'),
//   "acm": Array(76).fill('7'),
//   "re": Array(76).fill('7'),
//   "pn": Array(38).fill('7'),
// },

function bitGroups(arrays) {
  let bitGroups = {}
  for (const [array, name] of arrays) {
    let bitGroup = ''
    for (const field of array) {
      bitGroup = bitGroup.concat(field)
    }
    if (bitGroup.length < 128) {
      bitGroup = bitGroup.concat(Array(128 - bitGroup.length).fill('0').join(''))
    }
    bitGroups[name] = bitGroup
  }
  return bitGroups
}

function castToBinary(arrays) {
  for (let [field, maxBits] of arrays) {
    field = BigInt(field).toString()
    if (field.length < maxBits) {
      field = field.concat(Array(maxBits - field.length).fill(0).join(''))
    }
  }
}

function joinPreimageDigits(arrays) {
  return arrays.reduce((previousArray, currentArray) => previousArray.concat(currentArray), []);
}

function digitsToDecimal(digits) {
  return parseInt(digits.reduce((previousValue, currentValue) => previousValue.concat(currentValue)));
}

function decimalToHexString(decimal) {
  return decimal.toString(16);
}

function prependWithZeros(hexString) {
  const numberOfZeros = 32 - hexString.length;
  if (numberOfZeros > 0) {
    return Array(numberOfZeros).fill('0').join('').concat(hexString)
  }
  return hexString;
}

function hashToHexPair(hash) {
  return [hash.slice(0, 32), hash.slice(32, 64)];
}

function hashToIntPair(hash) {
  return [BigInt('0x' + hash.slice(0, 32)).toString(), BigInt('0x' + hash.slice(32, 64)).toString()];
}


module.exports = {
  bitGroups,
  castToBinary,
  joinPreimageDigits,
  digitsToDecimal,
  decimalToHexString,
  prependWithZeros,
  hashToIntPair,
  hashToHexPair
}