export const getTime = (): string => {
  const hours = new Date().getHours();
  console.log(hours);
  if (hours >= 0 && hours <= 6) {
    return "凌晨";
  }
  if (hours > 6 && hours <= 9) {
    return "早上";
  }
  if (hours > 9 && hours <= 12) {
    return "上午";
  }
  if (hours > 12 && hours <= 14) {
    return "中午";
  }
  if (hours > 14 && hours <= 18) {
    return "下午";
  }
  if (hours > 18 && hours <= 23) {
    return "晚上";
  }
};
