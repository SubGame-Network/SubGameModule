import React, { useMemo, useState } from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";

import { IconOutlineCheck } from "@subgame/react-icon-subgame";

import { usePolkadotJS } from "@polkadot/api-provider";
import sliceAddress from "utils/sliceAddress";

interface Props {
  infoRefetch: () => void;
}

const ChangeWalletBtn = ({ infoRefetch }: Props) => {
  const [walletMenuShow, setWalletMenuShow] = useState(false);
  const {
    state: { keyring, keyringState, currentAccount },
    changeAccount,
  } = usePolkadotJS();

  const walletList = useMemo(() => {
    if (keyringState === "READY" && keyring) {
      return keyring.getAccounts();
    }
    return [];
  }, [keyring, keyringState]);

  return (
    <WalletListBtnStyle
      onClick={() => {
        setWalletMenuShow(!walletMenuShow);
      }}
    >
      <span className="blueBtn">
        <FormattedMessage id="Changewallet" />
      </span>
      {walletMenuShow && (
        <ListMenu className="walletListMenu">
          {walletList.map((item, index) => {
            return (
              <li
                key={item.address}
                onClick={(e) => {
                  e.stopPropagation();
                  changeAccount(item);
                  setWalletMenuShow(false);
                  infoRefetch();
                }}
                className={
                  currentAccount?.address === item.address ? "active" : ""
                }
              >
                <img src="./images/Account.svg" alt="Account" />
                <div>
                  <p className="name">{item?.meta?.name}</p>

                  <p className="address">{sliceAddress(item?.address)}</p>
                </div>

                {currentAccount?.address === item.address ? (
                  <IconOutlineCheck size={20} color="#EB027D" />
                ) : (
                  <div></div>
                )}
              </li>
            );
          })}
        </ListMenu>
      )}
    </WalletListBtnStyle>
  );
};

const WalletListBtnStyle = styled.div`
  position: relative;

  .blueBtn {
    font-size: 12px;
    line-height: 14px;
    text-align: right;
    color: #004de1;
    cursor: pointer;
  }
  .address {
    font-size: 12px;
    line-height: 14px;

    /* grey01 */
    margin-top: 4px;
    color: #8b8b8b;
    &.fw400 {
      font-weight: 400;
      color: black;
      .font16 {
        color: black;
        font-size: 14px;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        word-break: break-all;
        -webkit-line-clamp: 1;
        white-space: pre;
      }
    }
  }
`;
const ListMenu = styled.ul`
  position: absolute;
  top: 20px;
  right: 0;
  padding: 0px;
  z-index: 100;

  box-shadow: 0px 10px 25px -5px rgba(0, 0, 0, 0.26);
  border-radius: 3px;

  background-color: #fff;
  display: grid;
  li {
    padding: 15px;
    box-sizing: border-box;
    display: grid;
    grid-template-columns: 30px 1fr 20px;
    align-items: center;
    grid-column-gap: 10px;
    .name {
      text-align: left;
      font-size: 14px;
      line-height: 17px;

      color: #171717;
    }
  }
`;

export default ChangeWalletBtn;
