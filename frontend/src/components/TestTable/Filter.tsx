import React, { useMemo } from "react";
import { Row } from "react-table";
import styled from "styled-components";
import Select from "../Select";
import Input from "../Input";

export type AllFiltersValue = {
  id: string;
  value: string[];
};

export function filterAddressAndFromAccount(
  rows: Row<{
    [key: string]: any;
    FromAccount: string;
    AddressList: string[];
  }>[],
  id: string,
  filterValue: string
) {
  if (!filterValue) return rows;

  return rows.filter((row) => {
    const value = filterValue.toLocaleLowerCase();
    const { FromAccount, AddressList } = row.original;
    return (
      FromAccount.toLocaleLowerCase().includes(value) ||
      AddressList.some((address) => address.toLocaleLowerCase().includes(value))
    );
  });
}

export function DefaultColumnFilter({ column }: any) {
  const { filterValue, setFilter, defaultFilterPlaceholder } = column;
  return (
    <Input
      style={{ height: "30px" }}
      value={filterValue || ""}
      placeholder={defaultFilterPlaceholder}
      onChange={(e: any) => {
        setFilter(e.target.value || undefined);
      }}
    />
  );
}

// export function SelectColumnFilter(
//   { column: { setFilter, preFilteredRows, id, filterValue } }: any,
//   options: string[] = []
// ) {
//   const all = useMemo(
//     () => ({
//       value: "",
//       label: "All",
//     }),
//     []
//   );

//   const defaultOptions = useMemo(() => {
//     if (options?.length > 0) return options;

//     const tmpOptions: string[] = [];
//     preFilteredRows.forEach((row: any) => {
//       if (tmpOptions.findIndex((elm) => elm === row.values[id]) === -1) {
//         tmpOptions.push(row.values[id]);
//       }
//     });

//     return [...tmpOptions];
//   }, [id, preFilteredRows, options]);

//   const selectOption = useMemo(
//     () =>
//       defaultOptions.map((option: string) => {
//         return { value: option, label: option };
//       }),
//     [defaultOptions]
//   );

//   selectOption.unshift(all);

//   return (
//     <Select
//       options={selectOption}
//       defaultValue={all}
//       onChange={(e: any) => setFilter(e.value)}
//     />
//   );
// }
export function SelectColumnFilter(
  { column: { setFilter, preFilteredRows, id, filterValue } }: any,
  options: string[] = []
) {
  const all = useMemo(
    () => ({
      value: "",
      label: "All",
    }),
    []
  );

  const defaultOptions = useMemo(() => {
    if (options?.length > 0) return options;

    const tmpOptions: string[] = [];
    preFilteredRows.forEach((row: any) => {
      if (tmpOptions.findIndex((elm) => elm === row.values[id]) === -1) {
        tmpOptions.push(row.values[id]);
      }
    });

    return [...tmpOptions];
  }, [id, preFilteredRows, options]);
  const selectOption = useMemo(
    () => [
      all,
      ...defaultOptions.map((option: string) => {
        return { value: option, label: option };
      }),
    ],
    [defaultOptions, all]
  );

  return (
    <Select
      options={selectOption}
      defaultValue={all}
      onChange={(e: any) => setFilter(e.value)}
    />
  );
}
export function filterArrayData(rows: any, id: any, filterValue: any[] = []) {
  if (filterValue.some((v) => v === "All")) return rows;
  return rows.filter((row: any) =>
    row.values[id].some((v: any) => filterValue.some((f) => f === v))
  );
}

export function statusFilter(rows: any, id: any, filterValue: string) {
  let myRows: any = [];

  if (filterValue.length === 0) {
    //選擇ALL時
    return rows;
  }
  for (let i = 0; i < rows.length; i++) {
    if (rows[i]["values"][id] === filterValue) {
      myRows.push(rows[i]);
    }
  }
  return myRows;
}

// export function MultiSelectColumnFilter({
//   column: {
//     setFilter,
//     filterValue = [], // 篩選器的值
//     filterOptions, // 自訂的篩選選項
//     id,
//     AllLabelString, // 自訂的全選標籤名稱
//   },
//   preFilteredRows,
// }: any) {
//   const options = useMemo(() => {
//     if (filterOptions?.length > 0) return filterOptions;

//     const tmpOptions: any[] = [];

//     preFilteredRows.forEach((row: any) => {
//       if (tmpOptions.findIndex((elm) => elm === row.values[id]) === -1) {
//         tmpOptions.push(row.values[id]);
//       }
//     });
//     return [...tmpOptions];
//   }, [id, preFilteredRows, filterOptions]);

//   const selectOption = useMemo(
//     () =>
//       options.map((option: string) => {
//         return { value: option, label: option };
//       }),
//     [options]
//   );
//   const defaultValue = useMemo(
//     () =>
//       filterValue.map((v: string) => {
//         return { value: v, label: v };
//       }),
//     [filterValue]
//   );

//   const handleSelectonChange = (selectedOptions: any) => {
//     const newValueArr = selectedOptions.map((item: any) => {
//       return item.value;
//     });
//     setFilter([...newValueArr]);
//   };

//   return (
//     <MultiSelect
//       options={selectOption}
//       onChange={handleSelectonChange}
//       defaultValue={defaultValue}
//       AllLabelString={AllLabelString}
//       useCurrencyOption={id === "Currency" || id === "coinType"}
//     />
//   );
// }

export function filterColumnMultiValue(
  rows: any,
  id: any,
  filterValue: any[] = []
) {
  if (filterValue.some((v) => v === "All")) return rows;
  return rows.filter((row: any) => {
    const rowValue = row.values[id];
    return filterValue.findIndex((elm) => elm === rowValue) !== -1;
  });
}

export function filterAccountType(rows: any, id: any, filterValue: any[] = []) {
  if (filterValue.some((v) => v === "All")) return rows;
  return rows.filter((row: any) => {
    const FromAccountType = row.values["FromAccountType"];
    const ToAccountType = row.original["ToAccountType"];
    return (
      filterValue.findIndex((elm) => elm === FromAccountType) !== -1 ||
      filterValue.findIndex((elm) => elm === ToAccountType) !== -1
    );
  });
}

export function FilterMultiColumn({
  column: { filterValue, preFilteredRows, setFilter },
}: any) {
  return (
    <SearchInputContainer>
      <Input
        width="100%"
        placeholder="TxidOrAddress"
        value={filterValue || ""}
        onChange={(e: any) => {
          setFilter(e.target.value || undefined);
        }}
      />
      {filterValue ? (
        <i
          className="icon-Member-System-Icon_Failed-40"
          onClick={() => setFilter(undefined)}
        />
      ) : (
        <i className="icon-Member-System-Icon_Search" />
      )}
    </SearchInputContainer>
  );
}

export function FilterItemWithoutBalance({ column: { setFilter } }: any) {
  const all = useMemo(() => ({ value: "", label: "All" }), []);
  const options = [
    all,
    {
      value: "hide",
      label: "HideItemWithoutBalance",
    },
  ];

  return (
    <Select
      defaultValue={all}
      options={options}
      onChange={(e: any) => {
        setFilter(e.value);
      }}
    />
  );
}

export function filterItemWithoutBalance(
  rows: any[],
  id: string,
  filterValue: string
) {
  // 參數ID沒用到但不可刪

  if (filterValue === "hide") {
    return rows.filter((row) => {
      let rowValue = row.values.TotalQuantity;
      return parseFloat(rowValue) > 0;
    });
  } else {
    return rows;
  }
}

export function filterMultiColumnValue(
  rows: any[],
  id: any,
  filterValue: string,
  columnMultiId: string[]
) {
  if (!columnMultiId) return rows;

  return rows.filter((row) =>
    columnMultiId.some((id) => row.original[id]?.includes(filterValue))
  );
}

// 取得所有filter的值，用於setAllFilters
export const getAllFilterValue = (
  allColumns: any,
  preFilteredRows: any
): AllFiltersValue[] => {
  const tmpList: AllFiltersValue[] = [];
  allColumns.forEach((column: any) => {
    if (column.canFilter && column.isMultiFilter) {
      const id = column.id;
      let filterOptions: string[] = [];

      // 如果存在自訂過濾選項就使用它，不存在就用preFilteredRows自動建立
      if (column.filterOptions) {
        filterOptions = column.filterOptions;
      } else {
        filterOptions = preFilteredRows.map((row: any) => row["values"][id]);
      }

      let obj = {
        id,
        value: [""],
      };

      obj.value = [...new Set(filterOptions.map((o: string) => o))];
      tmpList.push(obj);
    }
  });
  return tmpList;
};

const SearchInputContainer = styled.div`
  position: relative;
  i {
    position: absolute;
    right: 9px;
    top: 7px;
    font-size: 18px;
    &.icon-Member-System-Icon_Search {
      :before {
        color: black;
      }
    }
    &.icon-Member-System-Icon_Failed-40,
    &.icon-Member-System-Icon_Failed-40::before {
      cursor: pointer;
      color: black;
    }
  }
`;
