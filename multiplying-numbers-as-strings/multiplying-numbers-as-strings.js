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

const multiply = (num1, num2) => {
  if (num1 === "0" || num2 === "0") {
    return "0";
  }
  let sum = "";

  for (let i = num2.length - 1; i >= 0; i--) {
    const n2 = +num2[i];
    if (n2 === 0) {
      continue;
    }

    let result = Array(num2.length - 1 - i)
      .fill("0")
      .join("");
    let extended = 0;
    let lastNonZeroIndex = num2.length - 1;

    for (let j = num1.length - 1; j >= 0; j--) {
      const n1 = +num1[j];

      tmp = n1 * n2 + extended;
      if (n1 === 0 && extended !== 0) {
        result = tmp.toString() + result;
        extended = 0;
        lastNonZeroIndex = j;
      } else {
        result = (tmp % 10).toString() + result;
        extended = Math.floor(tmp / 10);
      }

      if (n1 !== 0) {
        lastNonZeroIndex = j;
      }
    }
    if (extended !== 0) {
      result = extended + result;
    }

    sum = add(result.slice(lastNonZeroIndex), sum);
  }

  return sum;
};

const testcases = [
  ["2", "0", "0"],
  ["0", "30", "0"],
  ["0000001", "3", "3"],
  ["1000001", "3", "3000003"],
  ["2", "3", "6"],
  ["29", "5", "145"],
  ["30", "69", "2070"],
  ["11", "85", "935"],
  ["1009", "03", "3027"],
  [
    "01857676599633787571309758335143087859263372693536636596716263717576954",
    "9",
    "16719089396704088141787825016287790733370354241829729370446373458192586",
  ],
  
  [
    "1020303004875647366210",
    "2774537626200857473632627613",
    "2830869077153280552556547081187254342445169156730",
  ],
  [
    "58608473622772837728372827",
    "7586374672263726736374",
    "444625839871840560024489175424316205566214109298",
  ],
  [
    "01857676599633787571309758335143087859263372693536636596716263717576954",
    "60197671035752796144013026738956372679497185647413406253210790599521935",
    "111827804835570597469803982908218785104176557578866304288908024762591625515223866705731753296182036381812491317890763876470478065688973485990",
  ],
];

for (const testcase of testcases) {
  result = multiply(testcase[0], testcase[1]);
  if (result === testcase[2]) {
    console.log("Passed");
  } else {
    console.log(`Failed, expect ${testcase[2]} but got ${result}`);
  }
}
