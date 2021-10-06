import React, { useState, useEffect } from "react";
import styled from "styled-components";
import StakeBox from "./StakeBox";
import ReactMarkdown from "react-markdown";
import gfm from "remark-gfm";
import Button from "../../components/Button";
import Loading from "../../components/Loading";

import "github-markdown-css";
import { useApiGetModuleDetail } from "../../hooks/useApi/useModule";
import { useApiGetUserInfo } from "../../hooks/useApi/useUser";
import { usePolkadotJS } from "@polkadot/api-provider";

import { useLocation } from "react-router-dom";
import { randomCardColors } from "../Module";
export const randomCardLine = [
  "lineBanner_1",
  "lineBanner_2",

  "lineBanner_3",

  "lineBanner_4",
];
const cardStyle: { [key: number]: any } = {
  0: {
    color: randomCardColors[0],
    line: randomCardLine[0],
  },
  1: {
    color: randomCardColors[1],
    line: randomCardLine[1],
  },
  2: {
    color: randomCardColors[2],
    line: randomCardLine[2],
  },
  3: {
    color: randomCardColors[3],
    line: randomCardLine[3],
  },

  4: {
    color: randomCardColors[4],
    line: randomCardLine[0],
  },
  5: {
    color: randomCardColors[6],
    line: randomCardLine[1],
  },
  6: {
    color: randomCardColors[7],
    line: randomCardLine[3],
  },
  7: {
    color: randomCardColors[8],
    line: randomCardLine[2],
  },
  8: {
    color: randomCardColors[5],
    line: randomCardLine[0],
  },
};
function ModuleDetail() {
  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const palletId = parseFloat(params.get("id") || "");

  const [isOpen, setIsOpen] = useState(false);
  const [markDownContent, setMarkDownContent] = useState("");
  const { data, isLoading } = useApiGetModuleDetail(palletId);
  const { data: userInfo } = useApiGetUserInfo();
  const moduleDetailData = data?.data.data;
  const userInfoData = userInfo?.data.data;

  const {
    state: { currentAccount, keyringState },
  } = usePolkadotJS();
  const walletIsConnect =
    keyringState === "READY" && currentAccount?.address ? true : false;
  useEffect(() => {
    fetch(moduleDetailData?.module.readmeMdUrl || "")
      .then((response) => {
        return response.text();
      })
      .then((text) => {
        setMarkDownContent(text);
      });
  }, [moduleDetailData?.module.readmeMdUrl]);
  return (
    <ModuleDetailStyle
      isOpen={isOpen}
      backgroundColor={cardStyle[(palletId - 1) % 9].color}
      backgroundLine={cardStyle[(palletId - 1) % 9].line}
    >
      {isLoading ? (
        <Loading />
      ) : (
        <>
          {" "}
          <div className="banner">
            <div className="wrap">
              <div className="title">{moduleDetailData?.module.name}</div>
              <div className="flex">
                <div className="whiteDot"></div>
                <p>SubGameScan</p>
              </div>
              <div className="desc2">{moduleDetailData?.module.depiction}</div>
              <div className="grid_column_2">
                {moduleDetailData?.program.map((item) => {
                  return (
                    <StakeBox
                      month={`${item.periodOfUse}`}
                      purchase={item.amount}
                      palletId={palletId}
                      programId={item.programID}
                      cantStake={!!userInfoData && walletIsConnect}
                    />
                  );
                })}
              </div>
            </div>
          </div>
          <div className="body">
            <div className="wrap">
              <div className="markdown-body">
                {" "}
                <ReactMarkdown
                  children={markDownContent}
                  remarkPlugins={[gfm]}
                  className="mardDownBox"
                ></ReactMarkdown>
              </div>

              {!isOpen && (
                <Button
                  text="ReadMore"
                  onClick={() => {
                    setIsOpen(true);
                  }}
                  className="readBtn"
                />
              )}
              {isOpen && <></>}
            </div>
          </div>
        </>
      )}
    </ModuleDetailStyle>
  );
}

const ModuleDetailStyle = styled.div<{
  isOpen: boolean;
  backgroundColor: string;
  backgroundLine: string;
}>`
  .banner {
    background: ${({ backgroundColor }) => {
      return backgroundColor;
    }};

    position: relative;
    .wrap {
      padding: 40px 0 54px;
      z-index: 10;
      position: relative;
      .title {
        font-weight: bold;
        font-size: 36px;
        line-height: 44px;

        color: #ffffff;
      }
      .flex {
        display: flex;
        align-items: center;
        margin: 20px 0;
        .whiteDot {
          width: 12px;
          height: 12px;
          border-radius: 50%;
          background: #ffffff;
          margin-right: 5px;
        }
        p {
          font-size: 16px;

          color: #ffffff;
        }
      }
      .desc2 {
        font-weight: normal;
        font-size: 16px;
        line-height: 160%;

        color: #ffffff;
      }
      .grid_column_2 {
        display: grid;
        grid-template-columns: 1fr 1fr;
        align-items: center;
        grid-column-gap: 20px;
        margin-top: 40px;
      }
    }
  }
  .banner::before {
    background-image: ${({ backgroundLine }) => {
      return `url("./images/${backgroundLine}.png")`;
    }};
    background-repeat: no-repeat;
    background-size: cover;
    content: "";
    display: block;
    height: 100%;
    position: absolute;
    z-index: 0;
    width: 100%;
  }

  .body {
    .wrap {
      padding: 40px 0 100px;
      .title {
        font-weight: bold;
        font-size: 24px;
        margin-bottom: 20px;
        color: #171717;
      }
      .desc {
        font-size: 14px;
        line-height: 17px;
        color: #171717;
        margin-bottom: 40px;
      }
      .mardDownBox {
        max-height: ${({ isOpen }) => {
          return isOpen ? "unset" : "270px";
        }};
        overflow-y: hidden;
        margin-top: 10px;
        margin-bottom: 40px;
        padding: 16px;

        /* grey03 */

        border: 1px solid #e8e8e8;
        box-sizing: border-box;
        border-radius: 5px;
        pre {
          margin: 0;
        }
      }
      .readBtn {
        border: 1px solid #171717;
        box-sizing: border-box;
        border-radius: 3px;
        height: 40px;
        margin-top: 20px;
        background-color: white;
        color: #171717;
      }
      .markTitle {
        font-weight: 500;
        font-size: 18px;

        color: #171717;
        &.mt40 {
          margin-top: 40px;
        }
      }
      .markDesc {
        font-size: 14px;
        line-height: 17px;
        color: #171717;
        margin-top: 10px;
      }
    }
  }
`;

export default ModuleDetail;
