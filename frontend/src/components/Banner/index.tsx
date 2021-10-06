import React from "react";
import styled from "styled-components";
import { FormattedMessage } from "react-intl";
interface Props {
  text: string;
}
function Banner({ text }: Props) {
  return (
    <BannerStyle>
      <div className="bannerwrap">
        <p>
          <FormattedMessage id={text} />
        </p>
      </div>
    </BannerStyle>
  );
}

const BannerStyle = styled.div`
  background-image: url("./images/Banner.png");
  background-repeat: no-repeat;
  background-size: cover;
  padding: 40px 0 52px;
  .bannerwrap {
    max-width: 1120px;
    margin: auto;
  }
  p {
    font-weight: bold;
    font-size: 36px;
    line-height: 44px;
    color: #8b8b8b;
  }
`;

export default Banner;
