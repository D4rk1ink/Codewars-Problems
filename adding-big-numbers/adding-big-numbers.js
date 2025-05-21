const add = (num1, num2) => {
  const lenNum1 = num1.length;
  const lenNum2 = num2.length;
  let result = "";
  let pivot = 1;
  let extended = 0;

  while (pivot <= lenNum1 || pivot <= lenNum2) {
    n1 = +num1.charAt(lenNum1 - pivot) || 0;
    n2 = +num2.charAt(lenNum2 - pivot) || 0;
    tmp = n1 + n2 + extended;
    result = (tmp % 10).toString() + result;
    extended = Math.floor(tmp / 10);

    pivot++;
  }

  if (extended != 0) {
    result = extended + result;
  }

  return result;
};
