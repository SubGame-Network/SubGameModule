import React, { useRef, useEffect, useCallback } from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";
interface Props {
  message?: string;
  positionRelativeElm?: boolean;
  width?: string;
  copiedMode?: boolean;
}

const Tips = ({
  message = "click To Do something",
  positionRelativeElm,
  width,
  copiedMode,
}: Props) => {
  const TipsRef = useRef<HTMLDivElement>(null);

  const _onmousemove = useCallback(
    (e: any) => {
      if (TipsRef.current) {
        if (positionRelativeElm) {
          const rect = e.currentTarget.getBoundingClientRect();
          const left = e.clientX - rect.left - 200;
          const top = e.clientY - rect.top + 17;
          TipsRef.current.style.left = left.toString() + "px";
          TipsRef.current.style.top = top.toString() + "px";
        } else {
          TipsRef.current.style.left = (e.clientX + 15).toString() + "px";
          TipsRef.current.style.top = (e.clientY + 17).toString() + "px";
        }
      }
    },
    [positionRelativeElm]
  );

  const _onmouseenter = useCallback(() => {
    if (TipsRef.current) {
      TipsRef.current.style.display = "flex";
    }
  }, []);

  const _onmouseleave = useCallback(() => {
    if (TipsRef.current) {
      TipsRef.current.style.display = "none";
    }
  }, []);

  useEffect(() => {
    if (TipsRef.current) {
      const parent = TipsRef.current.parentElement;
      if (parent) {
        if (positionRelativeElm) {
          parent.style.position = "relative";
        }
        if (!copiedMode) {
          parent.onmouseenter = _onmouseenter;
          parent.onmouseleave = _onmouseleave;
          parent.onmousemove = _onmousemove;
        }
      }
    }
  }, [
    _onmousemove,
    _onmouseenter,
    _onmouseleave,
    copiedMode,
    positionRelativeElm,
  ]);
  return (
    <TipsStyle
      ref={TipsRef}
      copiedMode={copiedMode}
      style={{ width, ...(copiedMode ? { display: "flex" } : {}) }}
    >
      {copiedMode ? (
        <>
          <Copied>
            <i className="icon-Member-System-Icon_Check2" />
          </Copied>
          <span>
            <FormattedMessage id="Copied" />
          </span>
        </>
      ) : (
        <FormattedMessage
          id="updateAt"
          values={{
            span: <span style={{ marginLeft: "3px" }}>{message}</span>,
          }}
        />
      )}
    </TipsStyle>
  );
};

export const TipsStyle = styled.p<{ copiedMode?: boolean }>`
  display: none;
  position: absolute;
  width: auto;
  min-height: 32px;
  color: #fff !important;
  z-index: 999;
  font-size: 14px;
  font-weight: normal;
  padding: 5px 10px;
  background: rgba(23, 23, 23, 0.6);
  box-shadow: 0px 4px 20px -3px rgba(0, 0, 0, 0.5);
  border-radius: 4px;
  justify-content: center;
  align-items: center;

  white-space: nowrap;
  left: -100px;
  ${({ copiedMode }) =>
    copiedMode
      ? `
    top : calc(100% + 5px);
    right:0;
  `
      : ``}
`;

const Copied = styled.div`
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: rgba(23, 23, 23, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  margin-right: 5px;
  i {
    font-size: 18px;
    &,
    &::before {
      color: #242740;
    }
  }
`;

export default Tips;
