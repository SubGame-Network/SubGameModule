import React from "react";
import styled from "styled-components";
import Tips from "../Tips";

interface Props extends React.HTMLAttributes<HTMLButtonElement> {
  icon: JSX.Element | string;
  tips?: string;
  isLoading?: boolean;
  disabled?: boolean;
}

const ActionButton: React.FunctionComponent<Props> = ({
  icon,
  isLoading,
  tips,
  disabled,
  ...rest
}) => {
  return (
    <Body className={isLoading ? "loading" : ""} disabled={disabled} {...rest}>
      {isLoading ? <i className="icon-Member-System-Icon_Loading" /> : icon}
      {tips && !disabled && <Tips message={tips} positionRelativeElm />}
    </Body>
  );
};

const Body = styled.button`
  background: #fff;
  border: 1px solid #e8e8e8;

  box-sizing: border-box;
  border-radius: 3px;
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  i,
  i::before {
    font-size: 20px;
    color: #fff;
  }
  :hover {
    border: 1px solid #171717;

    background: #f7f7f7;
  }

  &.loading {
    background-color: #313853;
  }

  &:disabled {
    border: 1px solid #e8e8e8;
    svg {
      fill: #e8e8e8;
    }
    :hover {
      border: 1px solid #e8e8e8;
      background: #fff;

      svg {
        fill: #e8e8e8;
      }
    }
  }
`;

export default ActionButton;
