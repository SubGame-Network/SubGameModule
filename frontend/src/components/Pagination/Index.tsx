import React from "react";
import styled from "styled-components";
import PageButtonGroup from "./PageButtonGroup";
interface Props {
  pageCount: number;
  pageIndex: number;
  previousPage: () => void;
  nextPage: () => void;
  gotoPage: (targetPage: number) => void;
}

const Pagination = ({
  pageCount,
  pageIndex,
  previousPage,
  nextPage,
  gotoPage,
}: Props) => {
  return (
    <PaginationStyle>
      <div>
        <PageButtonGroup
          totalPages={pageCount}
          currentPage={pageIndex + 1}
          pageNeighbours={2}
          previousPageFunc={previousPage}
          nextPageFunc={nextPage}
          gotoPageFunc={gotoPage}
        />
      </div>
    </PaginationStyle>
  );
};

const PaginationStyle = styled.div`
  display: flex;
  width: 100%;
  background-color: #fff;
  box-sizing: border-box;
  border-radius: 0 0 4px 4px;

  div {
    display: flex;
    button {
      margin: 0 5px;
      min-width: 30px;
      padding: 0 5px;
      border: 1px solid #e8e8e8;

      border-radius: 5px;

      height: 30px;
      background-color: #ffffff;
      font-family: Helvetica;
      font-size: 13px;
      font-weight: 700;
      letter-spacing: 1px;
      text-align: center;
      color: rgb(48, 43, 60);
      cursor: pointer;
      display: flex;
      justify-content: center;
      align-items: center;

      &:focus {
        outline: none;
      }
      &:first-child {
        background: #ffffff;

        border: 1px solid #171717;
        border-radius: 5px;

        display: flex;
        justify-content: center;
        &.disabled {
          background: #e8e8e8;
          border: none;
          svg {
            fill: #b9b9b9;
          }
        }
      }
      &:last-child {
        background: #ffffff;

        border: 1px solid #171717;
        border-radius: 5px;

        display: flex;
        justify-content: center;
        &.disabled {
          background: #e8e8e8;
          border: none;
          svg {
            fill: #b9b9b9;
          }
        }
      }

      &.active {
        background-color: ${({ theme }) => theme.Pink_2};
        color: white;
      }
    }
  }
`;

export default Pagination;
