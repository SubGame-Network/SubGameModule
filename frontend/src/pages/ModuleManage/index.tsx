import React, { useMemo, useCallback, useState, useEffect } from "react";
import styled from "styled-components";
import Banner from "../../components/Banner";

import Card from "./Card";
import { FormattedMessage, FormattedNumber } from "react-intl";
import TestTable from "../../components/TestTable";
import { SelectColumnFilter } from "../../components/TestTable/Filter";
import {
  render_status,
  renderCopyHash,
  renderModuleName,
  render_Time,
} from "../../components/TestTable/Cell";
import useAppContext from "../../hooks/useAppContext";
import {
  useApiGetUserRecord,
  useApiGetUserInfo,
} from "../../hooks/useApi/useUser";
import usePolkadotBalance from "../../hooks/usePolkadotBalance";
import { usePolkadotJS } from "@polkadot/api-provider";
const getStatus = (status: number) => {
  switch (true) {
    case status === 0:
      return "Failed";
    case status === 1:
      return "Success";
    case status === 2:
      return "Pending";
    default:
      return "";
  }
};
const statusStringToNum = (status: string) => {
  switch (true) {
    case status === "Failed":
      return "3";
    case status === "Success":
      return "1";
    case status === "Pending":
      return "2";
    default:
      return "";
  }
};
function ModuleManage() {
  const {
    state: { currentAccount },
  } = usePolkadotJS();
  const balances = usePolkadotBalance();
  const { data: userInfoData } = useApiGetUserInfo();
  const currentAddressHasJoin = userInfoData?.data.code === 200;
  const { showFeedBack } = useAppContext();
  const [currentBalance, setCurrentBalance] = useState(0);

  const [queryState, setQueryState] = useState<{ [propName: string]: string }>({
    row: "",
    page: "",
    NFTHash: "",
    statusString: "",
    PeriodOfUseMonth: "",
  });

  const turnFilterFunc = useCallback(
    (
      filters: { id: string; value: string }[],
      pageIndex: string,
      row: string
    ) => {
      setQueryState((prevState: any) => {
        let tmpQueryState = { ...prevState };
        if (filters.length > 0) {
          for (const [key] of Object.entries(tmpQueryState)) {
            const index = filters.findIndex((item) => {
              return item.id === key;
            });
            if (index !== -1) {
              tmpQueryState[key] = filters[index].value;
            } else {
              tmpQueryState[key] = "";
            }
          }
        } else if (filters.length === 0) {
          for (const [key] of Object.entries(tmpQueryState)) {
            tmpQueryState[key] = "";
          }
        }

        tmpQueryState.page = pageIndex;
        tmpQueryState.row = row;

        return tmpQueryState;
      });
    },
    []
  );

  const handleCopy = useCallback(
    (address: string) => {
      const copyinput = document.createElement("input");
      showFeedBack("Copied", "Copied");
      if (address && copyinput) {
        copyinput.value = address;
        document.body.appendChild(copyinput);
        copyinput.select();
        copyinput.setSelectionRange(0, 99999);
        document.execCommand("copy");
        document.body.removeChild(copyinput);
      }
    },
    [showFeedBack]
  );
  const { data, isLoading: recordIsLoading } = useApiGetUserRecord(
    {
      row: queryState.row,
      page: queryState.page,
      nftId: queryState.NFTHash,
      status: statusStringToNum(queryState.statusString),

      periodOfUse: queryState.PeriodOfUseMonth,
    },
    queryState.page !== "" && queryState.row !== ""
  );
  const dataCount = data?.data.data?.count || 0;

  const pageCount = Math.ceil(dataCount / parseFloat(queryState.row));
  const stakeRecord = useMemo(() => {
    if (data?.data.data?.list) {
      return data?.data.data?.list.map((item) => {
        return { ...item, statusString: getStatus(item.TxStatus) };
      });
    } else {
      return [];
    }
  }, [data?.data.data?.list]);

  const MoudleColumns = useMemo(() => {
    return [
      {
        Header: <FormattedMessage id="status" />,
        accessor: "statusString",
        Cell: render_status,
        Filter: (column: any) =>
          SelectColumnFilter(column, ["Failed", "Success", "Pending"]),
      },
      {
        Header: <FormattedMessage id="moduleName" />,
        accessor: "ModuleName",
        disableFilters: true,
        Cell: renderModuleName,
      },
      {
        Header: <FormattedMessage id="period" />,
        accessor: "PeriodOfUseMonth",
        Filter: (column: any) => SelectColumnFilter(column, ["6", "12"]),
        Cell: (Cell: any) => {
          return (
            <>
              {Cell.value} <FormattedMessage id="Day" />
            </>
          );
        },
      },
      {
        Header: <FormattedMessage id="startTime" />,
        accessor: "StartTime",
        disableFilters: true,
        Cell: render_Time,
      },
      {
        Header: <FormattedMessage id="endTime" />,
        accessor: "EndTime",
        disableFilters: true,
        Cell: render_Time,
      },
      {
        Header: <FormattedMessage id="stakeQuantity" />,
        accessor: "StakeSGB",
        disableFilters: true,
        Cell: (Cell: any) => {
          return (
            <>
              <FormattedNumber value={Cell.value} /> SGB
            </>
          );
        },
      },
      {
        Header: <FormattedMessage id="nftHash" />,
        accessor: "NFTHash",
        Cell: (Cell: any) => {
          return renderCopyHash(Cell, handleCopy);
        },
      },
    ];
  }, [handleCopy]);

  useEffect(() => {
    if (balances && currentAccount) {
      setCurrentBalance(parseFloat(balances[currentAccount?.address]));
    }
  }, [currentAccount?.address, currentAccount, balances]);
  return (
    <ModuleManageStyle>
      <Banner text="ModuleManage" />
      <div className="wrap">
        <>
          {" "}
          <div className="grid_column_3">
            <Card
              text="walletBalance"
              value={currentAddressHasJoin ? currentBalance.toString() : "-"}
              backgroundColor="black"
            />
            <Card
              text="staked"
              value={
                currentAddressHasJoin
                  ? data?.data.data?.stakedAmount || "0"
                  : "-"
              }
              backgroundColor="#EB027D"
            />
            <Card
              text="withrawn"
              value={
                currentAddressHasJoin ? data?.data.data?.withrawn || "0" : "-"
              }
              backgroundColor="#8B8B8B
          "
            />
          </div>
          <TestTable
            data={stakeRecord || []}
            columns={MoudleColumns}
            pageCount={pageCount}
            turnFilterFunc={turnFilterFunc}
            isLoading={recordIsLoading}
            currentAddressHasJoin={currentAddressHasJoin}
          />
        </>
      </div>
    </ModuleManageStyle>
  );
}

const ModuleManageStyle = styled.div`
  .wrap {
    padding: 40px 0 100px;

    .grid_column_3 {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      grid-column-gap: 20px;
      margin-bottom: 20px;
    }
  }
`;

export default ModuleManage;
