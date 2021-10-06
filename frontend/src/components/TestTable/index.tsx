import React, { useRef, useEffect, useMemo, useState } from "react";
import styled from "styled-components";
import {
  useTable,
  usePagination,
  useSortBy,
  useFilters,
  useGroupBy,
  useExpanded,
  useRowSelect,
  useFlexLayout,
} from "react-table";
import Loading from "../Loading";
import NoData from "../NoData";
import NotJoin from "../NotJoin";

import { FormattedMessage } from "react-intl";
import Select from "../Select";
import { DefaultColumnFilter, getAllFilterValue } from "./Filter";
import { settingOption } from "../../api/types/global";

import Pagination from "../Pagination/Index";
import { AllFiltersValue } from "./Filter";
interface Props {
  isOverFlow?: boolean;
  columns: any[];
  data: any[];
  renderDataDetails?: (row: any, page: any) => void;
  isLoading?: boolean;
  greyoutTargets?: { id: string; value: string }[];
  totalRow?: () => void;
  pageCount: number;
  turnFilterFunc?: (
    filters: { id: string; value: string }[],
    pageIndex: string,
    row: string
  ) => void;
  dataCount?: number;
  currentAddressHasJoin: boolean;
}

const Index: React.FunctionComponent<Props> = ({
  totalRow = false,
  isOverFlow = false,
  columns,
  data,
  renderDataDetails,
  isLoading,
  greyoutTargets = [],
  pageCount,
  turnFilterFunc = () => {},
  currentAddressHasJoin,
  calcTotalAmount,
  expanded,
  dataCount,
}: any) => {
  const [selectedRow, setSelectedRow] = useState<settingOption>({
    value: 10,
    label: "10 rolls",
  });

  const topFilterColumns = useMemo(() => {
    return columns
      .filter((column: any) => column.topFilter)
      .map((col: any) => col.accessor);
  }, [columns]);
  const skipPageResetRef: any = useRef(false);
  const AllFilterValue = useRef<AllFiltersValue[]>([]);

  const defaultColumn = useMemo(
    () => ({
      Filter: DefaultColumnFilter,
      Cell: (Cell: any) => <span>{Cell.value}</span>,
    }),
    []
  );

  const pageWillReset = () => {
    skipPageResetRef.current = true;
  };
  const pageWillNotReset = () => {
    skipPageResetRef.current = false;
  };
  const rowOption = [
    { value: 10, label: "10 rolls" },
    { value: 20, label: "20 rolls" },
    { value: 50, label: "50 rolls" },
    { value: 100, label: "100 rolls" },
    { value: 500, label: "500 rolls" },

    { value: 1000, label: "1000 rolls" },
  ];
  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    prepareRow,
    page,
    gotoPage,
    nextPage,
    previousPage,
    state: { pageIndex, filters },
    allColumns,
    preFilteredRows,
  } = useTable(
    {
      pageCount,
      columns,
      data,
      initialState: {
        filters: AllFilterValue.current, // 預設篩選為所有值
        hiddenColumns: topFilterColumns,
      },
      defaultColumn,
      autoResetFilters: false,
      autoResetExpanded: true,
      autoResetPage: skipPageResetRef.current,
      manualFilters: true,
      manualPagination: true,
    },
    useFilters,
    useGroupBy,
    useSortBy,
    useExpanded,
    usePagination,
    useRowSelect,
    useFlexLayout
  );

  useEffect(() => {
    AllFilterValue.current = getAllFilterValue(allColumns, preFilteredRows);
    turnFilterFunc(filters, pageIndex, selectedRow.value);
  }, [
    allColumns,
    preFilteredRows,
    filters,
    pageIndex,
    turnFilterFunc,
    selectedRow.value,
  ]);

  return (
    <BoxWrap>
      <TableDIV {...getTableProps()} isOverFlow={isOverFlow}>
        <Head isOverFlow={isOverFlow}>
          {headerGroups.map((headerGroup) => {
            return (
              <TheadTr {...headerGroup.getHeaderGroupProps()}>
                {headerGroup.headers.map((column: any) => {
                  return (
                    <TheadTh
                      {...column.getHeaderProps()}
                      textAlign={column.textAlign}
                    >
                      <p className="header">{column.render("Header")}</p>
                    </TheadTh>
                  );
                })}
              </TheadTr>
            );
          })}

          {headerGroups.map((headerGroup) => {
            return (
              <TheadTr
                {...headerGroup.getHeaderGroupProps()}
                className="filter-row"
                onClick={pageWillReset}
              >
                {headerGroup.headers.map((column) => {
                  return (
                    <TheadTh
                      {...column.getHeaderProps()}
                      onClick={pageWillReset}
                    >
                      {column.canFilter ? column.render("Filter") : null}
                    </TheadTh>
                  );
                })}
              </TheadTr>
            );
          })}
          {totalRow && totalRow()}
        </Head>
        <Body
          onClick={pageWillNotReset}
          {...getTableBodyProps()}
          isOverFlow={isOverFlow}
          pageLength={page.length}
        >
          {!currentAddressHasJoin ? (
            <NotJoin />
          ) : isLoading ? (
            <Loading />
          ) : page.length <= 0 ? (
            <NoData />
          ) : (
            <>
              {page.map((row: any) => {
                prepareRow(row);
                let isGreyout = greyoutTargets.some((item: any) => {
                  return (
                    row.original[`${item.id}`] &&
                    row.original[`${item.id}`] === item.value
                  );
                });
                return (
                  <>
                    <TbodyTr
                      {...row.getRowProps()}
                      className={`${
                        isGreyout ? "greyout" : row.isExpanded ? "expanded" : ""
                      }`}
                    >
                      {row.cells.map((cell: any) => {
                        return (
                          <TbodyTd {...cell.getCellProps()}>
                            {cell.render("Cell", { editable: false })}
                          </TbodyTd>
                        );
                      })}
                    </TbodyTr>

                    {row.isExpanded && (
                      <TbodyTrexpand>
                        {renderDataDetails(row, page)}
                      </TbodyTrexpand>
                    )}
                  </>
                );
              })}
            </>
          )}
        </Body>
      </TableDIV>
      <div className="sapcebetween">
        <div onClick={pageWillNotReset}>
          <Pagination
            pageCount={pageCount}
            pageIndex={pageIndex}
            previousPage={previousPage}
            nextPage={nextPage}
            gotoPage={gotoPage}
          />
        </div>
        <div onClick={pageWillReset} className="d_flex">
          <p>
            {dataCount}&nbsp;
            <FormattedMessage id="itemsintotal" />{" "}
          </p>
          <Select
            options={rowOption}
            value={selectedRow}
            onChange={(e: any) => {
              setSelectedRow({ value: e.value, label: e.label });
            }}
            customWidth="330px"
            customHeight="44px"
          />
        </div>
      </div>
    </BoxWrap>
  );
};
const BoxWrap = styled.div`
  .sapcebetween {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 20px;
    border: solid 1px #e8e8e8;
    border-top: none;
    .d_flex {
      display: grid;
      align-items: center;
      grid-template-columns: 1fr 330px;
      grid-column-gap: 30px;
      p {
        min-width: 100px;
        font-size: 14px;
        line-height: 160%;
        display: flex;
        align-items: center;
        text-align: right;
        color: #8b8b8b;
      }
    }
  }
`;
const TableDIV = styled.div<{ isOverFlow: boolean }>`
  background: #fff;
  padding: 20px 0 0;
  color: #fff;
  border: 1px solid #e8e8e8;
  overflow-x: ${({ isOverFlow }) => {
    return isOverFlow ? "scroll" : "unset";
  }};
`;

const Head = styled.div<{ isOverFlow: boolean }>`
  position: relative;
  width: ${({ isOverFlow }) => {
    return isOverFlow ? "fit-content" : "auto";
  }};
`;

const TheadTr = styled.div`
  padding: 0 6px;
  &.filter-row {
    border-top: 1px solid #e8e8e8;
  }
`;

const TheadTh = styled.div<{ textAlign?: string }>`
  padding: 12px 8px;
  height: 50px;
  display: flex;
  align-items: center;
  .header {
    font-weight: bold;
    font-size: 14px;
    line-height: 17px;

    color: #171717;
  }
`;

const Body = styled.div<{ isOverFlow: boolean; pageLength: number }>`
  width: ${({ isOverFlow, pageLength }) => {
    return isOverFlow && pageLength !== 0 ? "fit-content" : "auto";
  }};

  .empty {
    width: 100%;
    height: 500px;
    display: flex;
    justify-content: center;
    align-items: center;

    img {
      width: 116px;
      height: 115px;
    }
    p {
      text-align: center;
    }
  }
`;

const TbodyTr = styled.div`
  padding: 0 6px;
  border-top: 1px solid #e8e8e8;
  &.greyout {
    background-color: #313853;
  }
  &.expanded {
  }
  &:hover {
    background-color: #f7f7f7;
    .more {
      background-color: #f7f7f7;
    }
    &.expanded {
    }
  }
`;

const TbodyTrexpand = styled.div`
  padding: 20px;
  border-top: 1px solid #171717;
  border-bottom: 1px solid #171717;
`;

const TbodyTd = styled.div`
  padding: 12px 8px;
  min-height: 50px;
  display: flex;
  color: #171717;
  align-items: center;
  word-break: break-all;
  font-size: 14px;
`;

export default Index;
