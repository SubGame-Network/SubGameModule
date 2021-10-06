import React, {
  useState,
  useEffect,
  useCallback,
  useMemo,
  useRef,
} from "react";
import { FormattedMessage } from "react-intl";
import styled from "styled-components";
import Banner from "../../components/Banner";
import { Link } from "react-router-dom";
import InfiniteScroll from "react-infinite-scroll-component";
import { useApiGetModulesData } from "../../hooks/useApi/useModule";

import NoData from "../../components/NoData";
import Loading from "../../components/Loading";
export const cardLineBg = [
  "cardLine_1",
  "cardLine_2",
  "cardLine_3",
  "cardLine_4",
];
export const randomCardColors = [
  "linear-gradient(180deg, #17262D 0%, #314C5B 100%)",
  "linear-gradient(180deg, #B1B6BB 0%, #3C4959 100%)",
  "linear-gradient(180deg, #1E9E8B 0%, #A9D167 100%)",
  "linear-gradient(180deg, #EAAB41 0%, #DF442A 100%)",
  "linear-gradient(180deg, #01C3FA 0%, #0165ED 100%)",
  "linear-gradient(180deg, #0181E7 0%, #01D5C3 100%)",
  "linear-gradient(180deg, #7C3AD9 0%, #4522D7 100%)",
  "linear-gradient(180deg, #6655A0 0%, #D9ABC4 100%)",
  "linear-gradient(180deg, #ED1486 0%, #FC5E69 100%)",
  "linear-gradient(180deg, #E15DBA 0%, #C11370 100%)",
];

export const cardStyle: { [key: number]: any } = {
  0: {
    color: randomCardColors[0],
    line: cardLineBg[0],
  },
  1: {
    color: randomCardColors[1],
    line: cardLineBg[1],
  },
  2: {
    color: randomCardColors[2],
    line: cardLineBg[2],
  },
  3: {
    color: randomCardColors[3],
    line: cardLineBg[3],
  },

  4: {
    color: randomCardColors[4],
    line: cardLineBg[0],
  },
  5: {
    color: randomCardColors[6],
    line: cardLineBg[1],
  },
  6: {
    color: randomCardColors[7],
    line: cardLineBg[3],
  },
  7: {
    color: randomCardColors[8],
    line: cardLineBg[2],
  },
  8: {
    color: randomCardColors[5],
    line: cardLineBg[0],
  },
};

function Module() {
  const [data, setData] = useState<
    { id: number; depiction: string; name: string }[]
  >([]);
  const [dataIndex, setDataIndex] = useState({ start: 0, end: 8 });
  const timeID = useRef(0);

  const getMoreData = () => {
    clearTimeout(timeID.current);

    timeID.current = window.setTimeout(() => {
      refreshData();
    }, 1000);
  };
  const { data: myModuleData, isLoading } = useApiGetModulesData();
  const originalData = useMemo(() => {
    return myModuleData?.data.data || [];
  }, [myModuleData?.data.data]);
  const refreshData = useCallback(() => {
    setData((prevState: any) => {
      let newAry = originalData.slice(dataIndex.start, dataIndex.end);

      let tmpQueryState = prevState.concat(newAry);

      return tmpQueryState;
    });
    setDataIndex((prevState: any) => {
      let tmpQueryState = { ...prevState };
      tmpQueryState["start"] = tmpQueryState["start"] + 8;
      tmpQueryState["end"] = tmpQueryState["end"] + 8;

      return tmpQueryState;
    });
  }, [dataIndex.end, dataIndex.start, originalData]);

  useEffect(() => {
    if (originalData.length > 0) refreshData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [originalData]);

  return (
    <ModuleStyle>
      <Banner text="Module" />
      <div className="wrap">
        {isLoading ? (
          <Loading />
        ) : data.length <= 0 ? (
          <NoData />
        ) : (
          <>
            {" "}
            <InfiniteScroll
              className="scrollBox"
              dataLength={data.length}
              next={getMoreData}
              hasMore={data.length < originalData.length}
              loader={
                <LoadingComponent>
                  <img src="./images/grayCard.png" alt="grayCard" />
                  <img src="./images/grayCard.png" alt="grayCard" />

                  <img src="./images/grayCard.png" alt="grayCard" />

                  <img src="./images/grayCard.png" alt="grayCard" />
                  <div className="text">
                    <FormattedMessage id="loadingModule" />
                  </div>
                </LoadingComponent>
              }
            >
              <div className="cardArea">
                {" "}
                {data.map((item, index: number) => {
                  return (
                    <Link to={`/moduledetail?id=${item.id}`}>
                      <Card
                        backgroundColor={cardStyle[(item.id - 1) % 9].color}
                        backgroundLine={cardStyle[(item.id - 1) % 9].line}
                      >
                        <div className="content">
                          <div className="h180">
                            <p className="title">{item.name}</p>
                            <div className="flex">
                              <div className="whiteDot"></div>
                              <p>SubGameScan</p>
                            </div>
                          </div>
                          <div className="desc">{item.depiction}</div>{" "}
                          <p className="detail">
                            <FormattedMessage id="Detail" />
                          </p>
                        </div>
                      </Card>
                    </Link>
                  );
                })}
              </div>
            </InfiniteScroll>
          </>
        )}
      </div>
    </ModuleStyle>
  );
}

const ModuleStyle = styled.div`
  .wrap {
    padding: 40px 0 100px;
    .scrollBox {
      overflow: unset !important;
    }
    .cardArea {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr;
      grid-column-gap: 20px;
      grid-row-gap: 20px;
    }
  }
`;
const LoadingComponent = styled.div`
  margin-top: 20px;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  grid-column-gap: 20px;
  position: relative;
  .text {
    font-size: 16px;
    line-height: 160%;

    color: #171717;
    position: absolute;
    left: 50%;
    transform: translate(-50%, 100px);
  }
`;
const Card = styled.div<{ backgroundColor: string; backgroundLine: string }>`
  padding: 20px;
  background: ${({ backgroundColor }) => {
    return backgroundColor;
  }};
  position: relative;
  .title {
    font-weight: bold;
    font-size: 24px;
    line-height: 29px;

    color: #ffffff;
  }
  .desc {
    font-size: 14px;
    line-height: 17px;
    margin-top: 30px;

    color: #ffffff;
  }

  .detail {
    display: inline-block;
    font-weight: bold;
    font-size: 14px;
    line-height: 17px;
    margin-top: 30px;
    color: #ffffff;
    :hover {
      border-bottom: 1px solid white;
    }
  }
  .content {
    position: relative;
    z-index: 10;
    top: 0;
    left: 0;
    .h180 {
      height: 180px;
    }
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
  ::before {
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
    top: 0;
    left: 0;
  }
  :hover {
    cursor: pointer;
    transform: translateY(-10px);
  }
`;

export default Module;
