import React from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";

import { IconFilledSignWrong } from "react-icon-guanfan";
interface Props extends React.HTMLAttributes<HTMLButtonElement> {
  isPlaying?: boolean;
  showTips?: boolean;
}

const PlayButton = ({ isPlaying, showTips, ...rest }: Props) => {
  return isPlaying ? (
    <CloseButton type="button" {...rest}>
      <IconFilledSignWrong fontSize={24} color="#fff" />
      <span style={{ color: "#fff", fontSize: "18px" }}>
        <FormattedMessage id="close" />
      </span>
    </CloseButton>
  ) : (
    <Body type="button" {...rest}>
      <Triangle className="triangle" />
      {showTips && (
        <Tips className="tips">
          <FormattedMessage id="viewVideo" />
        </Tips>
      )}
    </Body>
  );
};

const Body = styled.button`
  width: 100px;
  height: 100px;
  border-radius: 50%;
  background-color: black;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 8px solid white;

  div.triangle {
    border-color: transparent transparent transparent #fff;
  }
  &:hover {
    background-color: "";

    div.tips {
      min-width: 110px;
      min-height: 30px;
      max-width: 999px;
      max-height: 999px;
      padding: 6px 16px;
    }
  }
`;

const CloseButton = styled.button`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const Triangle = styled.div`
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 11px 0 11px 16px;
  transform: translateX(3px);
  border-color: transparent transparent transparent;
`;

const Tips = styled.div`
  position: absolute;
  overflow: hidden;
  top: calc(100% + 10px);
  color: #fff;
  max-width: 0px;
  max-height: 0px;
  border-radius: 5px;
  font-weight: 500;
  font-size: 12px;
  background-color: "";
`;

export default PlayButton;
