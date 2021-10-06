import isToday from "date-fns/isToday";
import formatDistanceToNow from "date-fns/formatDistanceToNow";
import format from "date-fns/format";
import isValid from "date-fns/isValid";
const formatTime = (value: any) => {
  const date = new Date(value);

  return isValid(date)
    ? isToday(date)
      ? formatDistanceToNow(date)
      : format(date, "yyyy-MM-dd HH:mm:ss")
    : value;
};

const formatNumber = (value: string | number) => {
  const number = typeof value === "string" ? parseFloat(value) : value;
  return isNaN(number) ? "-" : number;
};
const FormatUpdateTime = (time?: string | Date | null, useFullTime = true) => {
  return format(
    time ? new Date(time) : new Date(),
    useFullTime ? "yyyy-MM-dd HH:mm:ss" : "yyyy-MM-dd"
  );
};

export { formatTime, formatNumber, FormatUpdateTime };
