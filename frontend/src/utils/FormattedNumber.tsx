import React from "react";
import { FormattedNumber as CoreFormattedNumber } from "react-intl";

import { NumberFormatOptions } from "@formatjs/ecma402-abstract";
import { CustomFormatConfig } from "@formatjs/intl";

const FormattedNumber: React.FC<
  NumberFormatOptions &
    CustomFormatConfig & {
      value: number | bigint | string;
    }
> = ({ value, ...rest }) => {
  return (
    <CoreFormattedNumber
      maximumFractionDigits={8}
      value={typeof value === "string" ? Number(value) : value}
      {...rest}
    />
  );
};

export default FormattedNumber;
