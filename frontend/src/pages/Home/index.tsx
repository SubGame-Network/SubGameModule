import React, { useState } from "react";
import styled from "styled-components";
import BackgroundPlayer from "../../components/Player/BackgroundPlayer";
import PlayButton from "../../components/Button/PlayButton";
import { FormattedMessage } from "react-intl";
import Button from "../../components/Button";
import { Link } from "react-router-dom";
function Home(props: any) {
  const [isPlaying, setIsPlaying] = useState(false);

  const closeButtonStyle: any = {
    zIndex: 2,
    position: "absolute",
    top: "20px",
    right: "83px",
  };

  const end = () => setIsPlaying(false);
  const play = () => setIsPlaying(true);
  return (
    <HomeStyle>
      <div className="videoBox">
        <BackgroundPlayer
          play={isPlaying}
          controls={isPlaying}
          url="./images/module.mp4"
          coverImgPath="./images/banner.svg"
        />
        {isPlaying ? (
          <PlayButton
            onClick={end}
            isPlaying={true}
            style={isPlaying ? closeButtonStyle : undefined}
            showTips={true}
          />
        ) : (
          <div className="contentBox">
            <div className="title">
              <FormattedMessage id="SubGameModule" />
            </div>
            <div className="desc">
              <FormattedMessage id="homedesc" />
            </div>

            <PlayButton onClick={play} showTips={true} />
          </div>
        )}
      </div>
      <div className="platformTour">
        <div className="wrap">
          <div className="title">
            <FormattedMessage id="platformTour" />
          </div>
          <div className="grid_column_3">
            <div className="box"></div>
            <div className="box"></div>
            <div className="box"></div>
          </div>
          <div className="flex_center">
            <Link to="/module">
              {" "}
              <Button text="SeeModule" className="btn" />
            </Link>
          </div>
        </div>
      </div>
    </HomeStyle>
  );
}

const HomeStyle = styled.div`
  min-height: 100vh;
  width: 100%;
  color: #fff;
  position: relative;
  .videoBox {
    height: 580px;
    position: relative;
    .contentBox {
      display: flex;
      flex-direction: column;
      align-items: center;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
    }
    .title {
      font-weight: bold;
      font-size: 60px;
      line-height: 73px;

      color: #171717;
    }
    .desc {
      font-weight: 500;
      font-size: 18px;
      line-height: 22px;
      color: #171717;
      margin: 20px 0 30px;
      width: 660px;
      word-break: break-all;
      text-align: center;
    }
  }
  .platformTour {
    padding: 60px 0 132px;
    background: linear-gradient(259.92deg, #393939 8.3%, #171717 74.26%);

    .title {
      font-weight: bold;
      font-size: 36px;
      line-height: 44px;
      text-align: center;
      color: #ffffff;
    }
    .grid_column_3 {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      grid-column-gap: 20px;
      margin-top: 30px;
      .box {
        background: #3e3e3e;
        height: 250px;
        box-shadow: 0px 1px 8px -2px rgba(0, 0, 0, 0.15);
        border-radius: 5px;
      }
    }
    .flex_center {
      margin-top: 40px;

      display: flex;
      justify-content: center;
      .btn {
        background-color: #eb027d;
        width: 320px;
      }
    }
  }
`;

export default Home;
