import React from "react";
import { FormattedRelativeTime } from "react-intl";
import { isValid, isToday, format } from "date-fns";

interface Props {
  value: number | Date | string;
}

const FormattedCustomTime: React.FunctionComponent<Props> = ({ value }) => {
  const date = new Date(value);

  if (isValid(date)) {
    if (isToday(date)) {
      return (
        <FormattedRelativeTime
          numeric="auto"
          updateIntervalInSeconds={1}
          value={date.getTime() / 1000 - new Date().getTime() / 1000}
        />
      );
    } else {
      return <>{format(date, "yyyy-MM-dd HH:mm:ss")}</>;
    }
  } else {
    return <>{JSON.stringify(date)}</>;
  }
};

export default FormattedCustomTime;
