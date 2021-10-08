import React from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";
import useAppContext from "../../hooks/useAppContext";

import WalletListBtn from "./WalletListBtn";
import { useHistory, Link } from "react-router-dom";
import { usePolkadotJS } from "@polkadot/api-provider";
import MoreButton from "../MoreButton";
import { IconOutlineLanguage } from "@subgame/react-icon-subgame";
const Nav = () => {
  const { Cookie, setLocale } = useAppContext();

  const {
    state: { currentAccount, keyringState },
  } = usePolkadotJS();
  const history = useHistory();
  const languageList = [
    {
      label: "EN",
      clickEvent: () => {
        setLocale("en");
        Cookie.set("locale", "en");
      },
    },
    {
      label: "簡中",
      clickEvent: () => {
        setLocale("cn");
        Cookie.set("locale", "cn");
      },
    },
  ];

  const manageList = [
    {
      label: <FormattedMessage id="UserInfo" />,
      clickEvent: () => {
        history.push("/userinfo");
      },
    },
    {
      label: <FormattedMessage id="ModuleManage" />,
      clickEvent: () => {
        history.push("/modulemanage");
      },
    },
    {
      label: <FormattedMessage id="ContactUs" />,
      clickEvent: () => {
        history.push("/contactus");
      },
    },
  ];

  const clickOpenLanguageMenu = () => {
    const btn: HTMLButtonElement | null =
      document.querySelector("#languageRef");
    if (btn) btn.click();
  };
  const clickOpenManage = () => {
    const btn: HTMLButtonElement | null = document.querySelector("#Manage");
    if (btn) btn.click();
  };
  const walletIsConnect = keyringState === "READY" && currentAccount?.address;
  return (
    <NavStyle walletIsConnect={walletIsConnect}>
      <div className="wrap">
        <div className="navLeft">
          <img
            src="./images/Logo.svg"
            alt="Logo"
            onClick={() => {
              history.push("/");
            }}
          />
        </div>
        <div className="navRight">
          <Link to="/">
            {" "}
            <p className="pageBtn">
              <FormattedMessage id="Home" />
            </p>
          </Link>
          <Link to="/module">
            <p className="pageBtn">
              <FormattedMessage id="Module" />
            </p>
          </Link>

          {keyringState === "READY" && currentAccount ? (
            <MoreButton
              FillIcon={{
                icon: <p></p>,
                openColor: "",
                closeColor: "",
              }}
              options={manageList}
              userActionMenu={"Manage"}
            >
              <p
                className="pageBtn"
                onClick={clickOpenManage}
                data-dropdown={true}
              >
                {" "}
                <FormattedMessage id="Manage" />
              </p>
            </MoreButton>
          ) : (
            ""
          )}

          <WalletListBtn />

          <div className="language">
            <MoreButton
              FillIcon={{
                icon: <IconOutlineLanguage color="#171717" />,
                openColor: "",
                closeColor: "",
              }}
              options={languageList}
              userActionMenu={"languageRef"}
            >
              <span
                className="accountName "
                onClick={clickOpenLanguageMenu}
                data-dropdown={true}
              ></span>
            </MoreButton>
          </div>
        </div>
      </div>
    </NavStyle>
  );
};

const NavStyle = styled.div<{ walletIsConnect: string | false | undefined }>`
  background-color: #fff;
  border-bottom: 1px solid rgba(23, 23, 23, 0.1);

  .wrap {
    padding: 12px 0 18px 0;
    display: flex;
    justify-content: space-between;
    align-items: center;
    ${({ theme }) => theme.fontStyles[16]}

    .navLeft {
      display: grid;
      grid-template-columns: 172px 122px;
      align-items: center;
      grid-column-gap: 30px;

      img {
        cursor: pointer;
      }
    }
    .navRight {
      display: grid;
      grid-template-columns: ${({ walletIsConnect }) => {
        return walletIsConnect
          ? "70px 70px 70px 148px 30px"
          : "70px 70px  148px 30px";
      }};
      grid-column-gap: 20px;
      align-items: center;
      .pageBtn {
        color: #171717;
        font-weight: normal;
        font-size: 14px;
        line-height: 17px;

        :hover {
          color: #eb027d;
        }
      }
      .language {
        cursor: pointer;
        box-sizing: border-box;
        display: flex;
        align-items: center;
        justify-content: center;
        height: 30px;
        border-radius: 50%;
        :hover {
          background-color: #ffd8ed;

          svg {
            background-color: #ffd8ed;
            fill: #eb027d;
          }
        }
      }
    }
  }
`;

export default Nav;
