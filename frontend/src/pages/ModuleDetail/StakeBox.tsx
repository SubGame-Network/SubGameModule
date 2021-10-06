import React, { useState } from "react";
import { FormattedMessage, FormattedNumber } from "react-intl";
import styled from "styled-components";
import Button from "../../components/Button";
import useStake from "../../hooks/useStake";
import useAppContext from "../../hooks/useAppContext";
import { useHistory } from "react-router-dom";
interface Props {
  month: string;
  purchase: string;
  palletId: number;
  programId: number;
  cantStake: boolean;
}

function StakeBox({ month, purchase, palletId, programId, cantStake }: Props) {
  const history = useHistory();
  const { showFeedBack } = useAppContext();

  const [isDisable, setIsDisable] = useState(false);

  const { useSendStake } = useStake();

  const { mutateAsync } = useSendStake();
  return (
    <StakeBoxStyle>
      <div>
        <div className="key">
          <FormattedMessage id="PERIODOFUSD" />
        </div>
        <div className="value">
          {month}&nbsp;
          <FormattedMessage id="Day" />
        </div>
      </div>
      <div className="column2">
        <div></div>
        <div className="line">/</div>
      </div>

      <div>
        {" "}
        <div className="key">
          <FormattedMessage id="PURCHASE" />
        </div>
        <div className="value pink">
          {" "}
          <FormattedNumber value={parseFloat(purchase)} />
          &nbsp;SGB
        </div>
      </div>

      <div>
        <Button
          disabled={isDisable}
          text="Stake"
          className="btn"
          onClick={() => {
            if (!cantStake) {
              history.push("/userinfo");
              return;
            }
            setIsDisable(true);

            mutateAsync({
              program_id: programId,
              pallet_id: palletId,

              callBack: (responseErrorMessage: string) => {
                if (!responseErrorMessage) {
                  showFeedBack("Submitted", "Submitted");
                } else {
                  showFeedBack("Failed", responseErrorMessage);
                }

                setIsDisable(false);
              },
            });
          }}
        />
      </div>
    </StakeBoxStyle>
  );
}

const StakeBoxStyle = styled.div`
  z-index: 100;
  padding: 25px 30px;
  background-color: white;
  display: grid;
  grid-template-columns: 0.9fr 10px 1fr 160px;
  align-items: center;

  background: #ffffff;

  border: 1px solid #e8e8e8;
  grid-column-gap: 10px;

  box-shadow: 0px 1px 8px -2px rgba(0, 0, 0, 0.15);
  border-radius: 10px;
  .key {
    font-size: 12px;
    line-height: 14px;

    color: #8b8b8b;
  }
  .column2 {
    text-align: right;
  }
  .line {
    margin: 28px 10px 0 0;
  }

  .value {
    font-weight: bold;
    font-size: 18px;
    line-height: 22px;
    margin-top: 12px;
    color: #171717;
    &.pink {
      color: #eb027d;
    }
  }
  .btn {
    height: 44px;
  }
`;

export default StakeBox;
