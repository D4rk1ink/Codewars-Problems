const sortIntervals = (intervals) => {
  const l = intervals.length;
  let pivot = 0;

  while (pivot < l) {
    for (let i = 0; i < l; i++) {
      const interval = intervals[i];
      if (intervals[pivot][0] <= interval[0]) {
        intervals = [
          ...intervals.slice(0, i),
          intervals[pivot],
          ...intervals.slice(i, pivot),
          ...intervals.slice(pivot + 1, l),
        ];
        break;
      }
    }
    pivot++;
  }

  return intervals;
};

const sumIntervals = (intervals) => {
  const sortedIntervals = sortIntervals(intervals);

  let sum = 0;
  let start = sortedIntervals[0][0];
  let end = sortedIntervals[0][1];

  for (const interval of sortedIntervals) {
    if (end < interval[0]) {
      sum = sum + (end - start);
      start = interval[0];
      end = interval[1];
    } else if (interval[1] > end) {
      end = interval[1];
    }
  }

  sum = sum + (end - start);

  return sum;
};
