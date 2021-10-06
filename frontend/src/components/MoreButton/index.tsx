/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { useState } from "react";
import styled from "styled-components";
interface MoreButton2 {
  FillIcon: {
    icon: JSX.Element;
    openColor: string;
    closeColor: string;
  };
  options: {
    label: JSX.Element | string;
    clickEvent: (e: React.MouseEvent) => void;
    userID?: string;
    disable?: boolean;
  }[];
  pageIndex?: number;
  disabled?: boolean;
  messageNoReadCount?: number;
  userActionMenu?: string;
}

const MoreButton: React.FunctionComponent<MoreButton2> = ({
  children,
  FillIcon,
  options,
  pageIndex,
  messageNoReadCount,
  disabled = false,
  userActionMenu = "",
}) => {
  const [isEnabled, setIsEnabled] = useState(false);
  const clickFunc = (e: any) => {
    const isDropdown = e.toElement?.getAttribute("data-dropdown");
    if (!isDropdown || isDropdown === "false") {
      const rootDiv: Element | null = document.querySelector("#root");
      if (rootDiv) {
        setIsEnabled(false);
        rootDiv.removeEventListener("click", clickFunc, true);
      }
    }
  };
  const clickShowMenu = () => {
    if (!disabled) {
      const rootDiv: Element | null = document.querySelector("#root");
      if (!isEnabled) {
        setIsEnabled(true);
        if (rootDiv) {
          rootDiv.addEventListener("click", clickFunc, true);
        }
      } else {
        setIsEnabled(false);
      }
    }
  };

  return (
    <>
      <Button
        type="button"
        className={isEnabled ? "moreBtn open" : "moreBtn"}
        onClick={clickShowMenu}
        pageIndex={pageIndex}
        FillIcon={FillIcon}
        isEnabled={isEnabled}
        data-dropdown={isEnabled}
        isDisabled={disabled}
        id={`${userActionMenu}`}
      >
        {messageNoReadCount && messageNoReadCount > 0 ? (
          <span
            className="dot"
            data-dropdown={isEnabled}
            onClick={clickShowMenu}
          ></span>
        ) : (
          ""
        )}
        <div className="titleBox">
          {FillIcon.icon}
          <div className="children">{children}</div>
        </div>

        <div
          data-dropdown="true"
          className="dropdownlist"
          style={{ overflow: "hidden" }}
        >
          {options.map((option, index) => {
            return (
              <a
                data-dropdown="true"
                onClick={(e) => {
                  option.clickEvent(e);
                }}
                className={option.disable ? "disable" : ""}
                key={index}
              >
                {option.label}
              </a>
            );
          })}
        </div>
      </Button>
    </>
  );
};

const Button = styled.button<{
  FillIcon: { openColor: string; closeColor: string };
  isEnabled: boolean;
  pageIndex?: number;
  isDisabled: boolean;
}>`
  background-color: #fff;
  border: none;
  position: relative;
  outline: none;

  cursor: ${(props) => (props.isDisabled ? "not-allowed" : "pointer")};
  svg {
    &,
    &::before {
      color: ${(props) =>
        props.isEnabled ? props.FillIcon.openColor : props.FillIcon.closeColor};
    }
  }
  .titleBox {
    display: flex;
    align-items: center;
    i {
      margin-left: 12px;
    }
    .children {
      span {
        &.accountName {
          color: ${(props) =>
            props.isEnabled
              ? props.FillIcon.openColor
              : props.FillIcon.closeColor};
        }
      }
    }
  }

  .dropdownlist {
    box-shadow: 0px 6px 15px rgba(0, 0, 0, 0.25);
    position: absolute;
    right: 0;
    top: 30px;
    z-index: 1000;

    min-width: 150px;
    max-height: 0px;
    border-radius: 4px;
    background-color: #fff;
    transform: ease max-height 0.5s;
    overflow: hidden;
    a {
      display: block;
      width: 100%;
      min-height: 40px;
      margin: 0;
      padding: 12px 20px;

      font-size: 14px;
      line-height: 17px;

      color: #171717;

      text-align: left;
      white-space: nowrap;
      box-sizing: border-box;
      &:hover {
        background: #e8e8e8;
      }
      &.active {
        background-color: #6d90c8;
        color: #fff;
        font-weight: bold;
      }
      &.disable {
        color: rgba(35, 24, 21, 0.2);
        pointer-events: none;
        cursor: default;
        :hover {
          background-color: #fff;
        }
      }
    }
  }
  &.open {
    div {
      max-height: 400px;
    }
  }
  &:hover {
    svg,
    svg::before {
      fill: ${(props) => {
        return props.FillIcon.openColor ? props.FillIcon.openColor : "#264c86";
      }};
    }
  }
`;

export default MoreButton;
