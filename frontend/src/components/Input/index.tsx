import React, { useRef, useState } from "react";
import styled from "styled-components";
import { FormattedMessage, useIntl } from "react-intl";

import { IconOutlineEye, IconOutlineEyeoff } from "react-icon-guanfan";

interface InputProps extends React.HTMLProps<HTMLInputElement> {
  errorMsg?: string;
  noNeedErrorMsg?: boolean;
  unitIcon?: JSX.Element;
  unitButtonClick?: () => void;
}

const Input: React.FunctionComponent<InputProps> = ({
  errorMsg,
  noNeedErrorMsg,
  unitButtonClick,
  unitIcon,

  ...rest
}) => {
  const [eyeIsOpen, setEyeIsOpen] = useState(false);
  const passwordInput = useRef<HTMLInputElement>(null);
  const { formatMessage } = useIntl();

  const changePasswordInputType = () => {
    if (passwordInput.current?.getAttribute("type") === "password") {
      passwordInput.current?.setAttribute("type", "text");
      setEyeIsOpen(false);
    } else {
      passwordInput.current?.setAttribute("type", "password");
      setEyeIsOpen(true);
    }
  };
  return (
    <InputStyle className="custom_InputDiv">
      <input
        {...rest}
        ref={passwordInput}
        placeholder={
          rest.placeholder ? formatMessage({ id: rest.placeholder }) : undefined
        }
      ></input>
      {unitIcon && (
        <button type="button" className="unit" onClick={unitButtonClick}>
          {unitIcon}
        </button>
      )}
      {!noNeedErrorMsg && (
        <p className="errormsg">
          {errorMsg && (
            <FormattedMessage
              id={errorMsg || "error!"}
              defaultMessage={errorMsg || "error!"}
            />
          )}
        </p>
      )}
      {rest.type === "password" && (
        <>
          {" "}
          {eyeIsOpen ? (
            <IconOutlineEye
              size={20}
              className="eye"
              onClick={() => changePasswordInputType()}
            />
          ) : (
            <IconOutlineEyeoff
              className="eye"
              size={20}
              onClick={() => changePasswordInputType()}
            />
          )}
        </>
      )}
    </InputStyle>
  );
};

const InputStyle = styled.div`
  position: relative;
  transition: width 0.4s;

  input {
    position: relative;

    width: 100%;
    height: 44px;
    border-radius: 2px;
    border: 1px solid #e8e8e8;
    background-color: #fff;
    color: #231815;
    font-size: 14px;
    padding-left: 10px;
    ::placeholder {
      color: #8b8b8b;
    }

    &:focus {
    }
    &.error {
      border: solid 1px #eb027d;
    }
  }
  .unit {
    position: absolute;
    right: 5px;
    top: 7px;
    font-size: 20px;
    background-color: #eeeeee;
  }
  .errormsg {
    font-size: 12px;

    color: #eb027d;
  }
  .eye {
    position: absolute;
    width: 20px;
    height: 20px;
    padding: 0;
    background-color: transparent;
    right: 4px;
    top: 22px;
    transform: translateY(-50%);
  }
`;

export default React.memo(Input);
