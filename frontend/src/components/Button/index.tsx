import React from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";

interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  text?: string;
  textValues?: { [key: string]: string };
  iconName?: string;
  isSending?: boolean;
}

const Button = ({
  isSending,
  iconName,
  text,
  textValues,
  children,
  ...rest
}: Props) => {
  return (
    <ButtonStyle {...rest} disabled={isSending || rest.disabled}>
      {iconName && <img src={iconName} alt={iconName} className="icon" />}
      {text && (
        <FormattedMessage
          id={isSending ? "sending" : text}
          values={textValues}
        />
      )}
      {children}
    </ButtonStyle>
  );
};

const ButtonStyle = styled.button`
  height: 60px;
  width: 100%;
  color: white;
  background-color: ${({ theme }) => theme.Dark_1};
  border-radius: 3px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  font-size: 16px;
  line-height: 160%;
  &:disabled {
    background-color: #e8e8e8;
    color: #b9b9b9;
  }
  &:hover:enabled {
  }
`;

export default Button;
