import React from "react";
import { FormattedMessage } from "react-intl";
import styled from "styled-components";
import FormattedNumber from "../../utils/FormattedNumber";

interface Props {
  text: string;
  value: string;
  backgroundColor: string;
}

function Card({ text, value, backgroundColor }: Props) {
  return (
    <CardStyle backgroundColor={backgroundColor}>
      <div className="key">
        <FormattedMessage id={text} />
      </div>

      <div className="value">
        {value === "-" ? (
          <> - &nbsp;SGB</>
        ) : (
          <>
            {" "}
            <FormattedNumber value={value} /> &nbsp;SGB
          </>
        )}
      </div>
    </CardStyle>
  );
}

const CardStyle = styled.div<{ backgroundColor: string }>`
  padding: 12px;
  background-color: ${({ backgroundColor }) => {
    return backgroundColor;
  }};
  .key {
    font-size: 14px;
    line-height: 17px;

    color: #ffffff;
  }
  .value {
    font-weight: bold;
    font-size: 24px;
    line-height: 29px;
    margin-top: 10px;
    color: #ffffff;
    text-align: right;
  }
`;

export default Card;
