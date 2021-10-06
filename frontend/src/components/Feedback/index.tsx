import React, { useEffect, useRef } from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";

import {
  IconOutlineSignComplete,
  IconFilledSignMention,
  IconOutlineCopy,
} from "react-icon-guanfan";

export type TType = "Submitted" | "Failed" | "Copied";
export type TMessage = string;
export type TMessageValues = { [key: string]: string | number };

export interface IFeedBackData {
  type: TType;
  message?: TMessage;
  messageValues?: TMessageValues;
}

interface IFeedBackProps extends IFeedBackData {
  delay?: number;
}

interface Props extends IFeedBackProps {
  setFeedbackShow: (a: boolean) => void;
}

const Feedback = ({
  message,
  type,
  setFeedbackShow,
  messageValues,
  delay = 3500,
}: Props) => {
  const ref = useRef<HTMLDivElement | null>(null);
  const types = {
    Submitted: <IconOutlineSignComplete fontSize={10} color="#FFF" />,
    Failed: <IconFilledSignMention fontSize={10} color="#EB027D" />,
    Copied: <IconOutlineCopy fontSize={10} color="#FFF" />,
  };
  useEffect(() => {
    let outTimeID: number, timeID: number;

    if (ref.current) {
      setTimeout(() => {
        if (ref.current) ref.current.classList.add("active");
      }, 0);
    }

    outTimeID = window.setTimeout(() => {
      if (ref.current) ref.current.classList.remove("active");
    }, delay);

    timeID = window.setTimeout(() => {
      setFeedbackShow(false);
    }, delay + 500);

    return () => {
      clearTimeout(timeID);
      clearTimeout(outTimeID);
    };
  }, [setFeedbackShow, delay]);

  return (
    <Body ref={ref} className={type}>
      <figure className={type === "Failed" ? "Failed" : ""}>
        {types[type]}
      </figure>
      <Right>
        <p className="content">
          {message ? (
            <FormattedMessage id={message} values={messageValues} />
          ) : (
            <FormattedMessage
              id={
                type === "Submitted" ? "Feedback_Submitted" : "Feedback_Failed"
              }
            />
          )}
        </p>
      </Right>
    </Body>
  );
};

const Body = styled.div`
  background: rgba(23, 23, 23, 0.8);
  box-sizing: border-box;
  border-radius: 3px;

  color: #171717;
  line-height: 18px;
  padding: 15px 15px 15px 20px;

  position: fixed;
  top: 0px;
  left: 50%;

  transform: translate(-50%, -100%);
  display: flex;
  align-items: center;
  justify-content: flex-start;
  min-width: 600px;
  z-index: 999;
  transition: 0.5s;
  font-size: 16px;
  font-weight: 400;
  color: #fff;
  height: 50px;
  figure {
    margin-right: 13px;
    display: flex;
    align-items: center;
    width: 24px;
    height: 24px;
    &.Failed {
    }
  }

  &.active {
    top: 30px;
    transform: translate(-50%, 0%);
  }

  &.Failed {
    color: ${({ theme }) => theme.Function_Error};
  }
`;

const Right = styled.div`
  .content {
    height: 30px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    color: #fff;
    font-size: 14px;
    font-weight: 400;
  }
`;

export default Feedback;
