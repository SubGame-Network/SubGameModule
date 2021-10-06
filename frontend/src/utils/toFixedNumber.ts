function toFixedNumber(num: string) {
  const re = new RegExp("^-?\\d+(?:.\\d{0," + (6 || -1) + "})?");
  const result = num.match(re);

  if (result) {
    return result[0];
  }

  return "0";
}

export default toFixedNumber;
