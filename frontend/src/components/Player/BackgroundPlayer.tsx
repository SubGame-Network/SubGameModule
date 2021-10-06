import React, { useEffect, useRef } from "react";
import ReactPlayer from "react-player";
import styled from "styled-components";

interface Props {
  url: string;
  play: boolean;
  width?: string;
  height?: string;
  loop?: boolean;
  controls: boolean;
  coverImgPath?: string;
}

const BackgroundPlayer: React.FunctionComponent<Props> = ({
  url,
  play,
  width,
  height,
  loop,
  controls,
  coverImgPath,
}) => {
  const ref = useRef<ReactPlayer | null>(null);
  const isFristTimeRender = useRef(true);

  useEffect(() => {
    if (isFristTimeRender.current) {
      isFristTimeRender.current = false;
      return;
    }

    if (!play && ref.current?.seekTo) {
      ref.current?.seekTo(0, "seconds");
    }
  }, [play, ref]);

  if (coverImgPath && !play) {
    return (
      <Wrap>
        <figure className="cover">
          <img src={coverImgPath} alt="cover" />
        </figure>
      </Wrap>
    );
  }

  return (
    <ReactPlayer
      wrapper={Wrap}
      width={width || "100%"}
      height={height || "100%"}
      url={url}
      ref={ref}
      volume={0.2}
      controls={controls}
      playing={play}
      loop={loop}
      onEnded={() => {
        ref.current?.seekTo(0, "seconds");
      }}
    />
  );
};

const Wrap = styled.div`
  width: 100%;
  height: 100%;
  top: 0;
  z-index: 0;
  left: 0;
  video {
    object-fit: cover;
    background-color: #cfcfcf;
  }
  figure.cover {
    width: 100%;
    height: 100%;
    img {
      object-fit: cover;

      width: 100%;
      height: 100%;
    }
  }
`;

export default BackgroundPlayer;
