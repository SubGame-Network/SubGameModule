import React from "react";
import styled from "styled-components";
import {
  IconOutlineCopy,
  IconOutlineSignComplete,
  IconFilledPendingUgly,
  IconOutlineCrossround,
} from "react-icon-guanfan";
import { formatTime } from "utils/format";

export const render_Time = ({ value }: any) => {
  return formatTime(value);
};
export const render_status = ({ value }: any) => {
  switch (true) {
    case value === "Success":
      return <IconOutlineSignComplete size={20} color="#6BC10E" />;
    case value === "Pending":
      return <IconFilledPendingUgly size={20} color="#FFB600" />;
    case value === "Failed":
      return <IconOutlineCrossround size={20} color="#EB027D" />;

    default:
      return "-";
  }
};
export const renderCopyHash = (Cell: any, copyFunc: any) => {
  const value = Cell.value;
  return (
    <RenderCopyHashStyle>
      {value ? (
        <>
          <p> {value}</p>
          <div
            className="copyCircle"
            onClick={() => {
              copyFunc(value);
            }}
          >
            <IconOutlineCopy size={15} color="#171717" />
          </div>
        </>
      ) : (
        "-"
      )}
    </RenderCopyHashStyle>
  );
};
const RenderCopyHashStyle = styled.div`
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: space-between;
  /* padding-left: 20px; */
  .copyCircle {
    cursor: pointer;
    border-radius: 50%;
    background: #f7f7f7;
    width: 28px;
    height: 28px;
    line-height: 28px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
`;

export const renderModuleName = ({ value }: any) => {
  return <RenderModuleNameStyle>{value}</RenderModuleNameStyle>;
};
const RenderModuleNameStyle = styled.div`
  color: #004de1;
`;
