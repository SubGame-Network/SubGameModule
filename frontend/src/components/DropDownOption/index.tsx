/* eslint-disable jsx-a11y/anchor-is-valid */
import React from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";
interface DropDownOptionPorps {
  options: {
    label: JSX.Element | string;
    clickEvent: (e: React.MouseEvent) => void;
  }[];
  isOpen: boolean;
}
const DropDownOption: React.FunctionComponent<DropDownOptionPorps> = ({
  options,
  isOpen,
}) => {
  return (
    <DropDownOptionStyled className={isOpen ? "open" : ""}>
      {options.length > 0 ? (
        options.map((option, index) => {
          return (
            <div
              className="option"
              data-dropdown="true"
              onClick={(e) => {
                option.clickEvent(e);
              }}
              key={index}
            >
              {option.label}
            </div>
          );
        })
      ) : (
        <div className="option nodatatext" data-dropdown="true">
          <FormattedMessage id="nodatafound" />
        </div>
      )}
    </DropDownOptionStyled>
  );
};

const DropDownOptionStyled = styled.div`
  box-shadow: 0px 6px 15px rgba(0, 0, 0, 0.25);
  position: absolute;
  right: 0;
  top: 45px;
  z-index: 1000;

  width: 100%;
  max-height: 0px;
  border-radius: 4px;
  background-color: #fff;
  transform: ease max-height 0.5s;
  overflow: hidden;
  .option {
    display: block;
    width: 100%;
    min-height: 40px;
    margin: 0;
    padding: 11px 13px;
    font-family: Helvetica;
    font-size: 15px;
    letter-spacing: 1px;
    color: #231815;
    text-align: center;
    white-space: nowrap;
    box-sizing: border-box;
    border-bottom: 1px solid #eeeeee;
    &.nodatatext {
      opacity: 0.5;
    }
    &:hover {
      background-color: #231815;
      color: #fff;
    }
    &.active {
      background-color: #6d90c8;
      color: #fff;
      font-weight: bold;
    }
    &.needIcon {
      display: flex;
      i {
        :before {
          color: #918b8a;
        }
      }
      span {
        .userID {
          font-weight: normal;
          font-size: 14px;
          line-height: 22px;
          color: #231815;
        }
        .userName {
          font-weight: normal;
          font-size: 12px;
          line-height: 20px;
          letter-spacing: 0.004em;
          color: #231815;
          opacity: 0.5;
        }
      }
      :hover {
        span {
          .userID {
            font-weight: normal;
            font-size: 14px;
            line-height: 22px;
            color: #fff;
          }
          .userName {
            font-weight: normal;
            font-size: 12px;
            line-height: 20px;
            letter-spacing: 0.004em;
            color: #231815;
            opacity: 0.5;
          }
        }
      }
    }
  }
  &.open {
    max-height: 400px;
  }
`;
export default DropDownOption;
