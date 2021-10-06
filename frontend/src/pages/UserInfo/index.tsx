import React, { useState, useEffect } from "react";
import { FormattedMessage } from "react-intl";
import styled from "styled-components";
import { useFormik } from "formik";
import * as yup from "yup";
import { usePolkadotJS } from "@polkadot/api-provider";
import Button from "../../components/Button";
import Input from "../../components/Input";
import sliceAddress from "utils/sliceAddress";
import ChangeWalletBtn from "./ChangeWalletBtn";
import { useHistory } from "react-router-dom";
import PhoneInput from "react-phone-input-2";
import "react-phone-input-2/lib/style.css";
import {
  useApiSendVerifyCode,
  useApiUserJoin,
  useApiGetUserInfo,
} from "hooks/useApi/useUser";

interface CountryData {
  name: string;
  dialCode: string;
  countryCode: string;
  format: string;
}

function UserInfo() {
  const history = useHistory();

  const { data, refetch: infoRefetch } = useApiGetUserInfo();

  const [selectedCountry, setSelectedCountry] = useState("us");
  const [countryCode, setcountryCode] = useState("1");

  const [waitingSecond, setWaitingSecond] = useState(5);
  const [isDisable, setIsDisable] = useState(false);

  const { mutateAsync: verifyAsync } = useApiSendVerifyCode();
  const { mutateAsync: joinAsync } = useApiUserJoin();
  const startTimer = () => {
    let myInterval = setInterval(() => {
      setWaitingSecond((prevSec) => {
        if (prevSec > 0) {
          return prevSec - 1;
        } else {
          clearInterval(myInterval);
          setIsDisable(false);

          return 5;
        }
      });
    }, 1000);
  };
  const {
    state: { currentAccount, keyringState },
    connectToWallet,
  } = usePolkadotJS();
  useEffect(() => {
    if (data?.status === 200 && currentAccount?.address) {
      history.push(`/signeduserinfo`);
    }
  }, [data?.data.data, data?.status, history, currentAccount?.address]);
  const sendVerifyCode = async () => {
    startTimer();
    setIsDisable(true);
    const api = await verifyAsync({
      email: values.email,
    });
    if (api.status === 200 && api.data.data) {
    }
  };

  const {
    handleChange,
    errors,
    touched,
    handleSubmit,
    values,

    // eslint-disable-next-line react-hooks/rules-of-hooks
  } = useFormik({
    initialValues: {
      nickName: "",
      phone: "",
      email: "",
      verifyCode: "",
    },
    validationSchema: yup.object().shape({
      nickName: yup.string().max(20, "Wordlimit").required("Required"),
      email: yup.string().required("Required").email("EmailError"),
      verifyCode: yup.string().required("Required"),

      phone: yup.string(),
    }),
    onSubmit: async ({ nickName, phone, email, verifyCode }) => {
      let data: any;
      data = {
        nickName,
        phone,
        email,
        verifyCode,
        address: currentAccount?.address,
        country: selectedCountry,
        countryCode: countryCode,
      };
      const api = await joinAsync(data);
      if (api.status === 200 && api.data.data) {
        infoRefetch();
      }
    },
  });

  return (
    <UserInfoStyle>
      <div className="wrap">
        <div className="title">
          <p>
            <FormattedMessage id="hello" />,
          </p>
          <span>
            {" "}
            <FormattedMessage id="hello2" />
          </span>
          <span className="hand">👉</span>
        </div>
        <p className="infoDesc">
          <FormattedMessage id="infoDesc" />
        </p>
        <form>
          <div className="formArea">
            <div>
              <div className="flex_between">
                <label htmlFor="">
                  <FormattedMessage id="SubGameAddress" />
                </label>
                {keyringState === "READY" && currentAccount && (
                  <ChangeWalletBtn infoRefetch={infoRefetch} />
                )}
              </div>
              {keyringState === "READY" && currentAccount ? (
                <div className="grayBox">
                  {sliceAddress(currentAccount?.address, 10, 10)}
                </div>
              ) : (
                <Button
                  text="ConnectaWallet"
                  className="connectBtn"
                  type="button"
                  onClick={() => {
                    connectToWallet();
                  }}
                />
              )}
            </div>
            <div>
              <label htmlFor="">
                <FormattedMessage id="nickName" />
              </label>
              <Input
                name="nickName"
                id="nickName"
                onChange={handleChange}
                className={`${
                  touched.nickName && errors.nickName ? "error" : ""
                } `}
                errorMsg={
                  errors.nickName && touched.nickName ? errors.nickName : ""
                }
              />
            </div>

            <div className="country">
              <label htmlFor="">
                <FormattedMessage id="Country" />
              </label>
              <PhoneInput
                country={selectedCountry}
                value={countryCode}
                inputStyle={{
                  display: "none",
                }}
                onChange={(value, country: CountryData) => {
                  setSelectedCountry(country.name);
                  setcountryCode(country.dialCode);
                }}
              />
            </div>

            <div>
              <label htmlFor="">
                <FormattedMessage id="Phone" />
              </label>
              <div className="phoneColumn">
                <Input
                  value={countryCode}
                  onChange={(e) => {
                    setcountryCode(e.currentTarget.value);
                  }}
                  disabled
                  // className={`${touched.phone && errors.phone ? "error" : ""} `}
                  // errorMsg={errors.phone && touched.phone ? errors.phone : ""}
                />
                <Input
                  name="phone"
                  id="phone"
                  onChange={handleChange}
                  // className={`${touched.phone && errors.phone ? "error" : ""} `}
                  // errorMsg={errors.phone && touched.phone ? errors.phone : ""}
                />
              </div>
            </div>

            <div>
              <div className="flex_between">
                <label htmlFor="">
                  <FormattedMessage id="Email" />
                </label>
                <span
                  className={isDisable ? "blueBtn disable" : "blueBtn"}
                  onClick={() => {
                    if (!isDisable) {
                      sendVerifyCode();
                    }
                  }}
                >
                  <FormattedMessage id="sendCode" />
                  {isDisable && <>({waitingSecond})</>}
                </span>
              </div>
              <Input
                name="email"
                id="email"
                onChange={handleChange}
                className={`${touched.email && errors.email ? "error" : ""} `}
                errorMsg={errors.email && touched.email ? errors.email : ""}
              />
            </div>

            <div>
              <label htmlFor="">
                <FormattedMessage id="Verificationcode" />
              </label>
              <Input
                width="100"
                name="verifyCode"
                id="verifyCode"
                onChange={handleChange}
                className={`${
                  touched.verifyCode && errors.verifyCode
                    ? "error w160"
                    : "w160"
                } `}
                errorMsg={
                  errors.verifyCode && touched.verifyCode
                    ? errors.verifyCode
                    : ""
                }
              />
            </div>
            <div>
              <Button
                text="Save"
                type="button"
                onClick={() => {
                  handleSubmit();
                }}
              />
            </div>
            <div></div>
          </div>
        </form>
      </div>
    </UserInfoStyle>
  );
}

const UserInfoStyle = styled.div`
  padding: 40px 0 100px;
  .wrap {
    .title {
      font-style: normal;
      font-weight: bold;
      font-size: 60px;

      color: #171717;
      .hand {
        font-size: 40px;
      }
    }
    .infoDesc {
      font-size: 16px;
      line-height: 160%;

      color: #171717;
      margin: 20px 0 40px;
    }
    .formArea {
      display: grid;
      grid-template-columns: 1fr 1fr;
      grid-template-rows: 1fr 1fr 1fr 1fr;
      grid-column-gap: 40px;
      grid-row-gap: 30px;
      .w160 {
        width: 160px;
      }
      .phoneColumn {
        display: grid;
        grid-template-columns: 96px 1fr;
        grid-column-gap: 30px;
      }
      .blueBtn {
        font-size: 12px;
        line-height: 14px;
        text-align: right;
        color: #004de1;
        cursor: pointer;
        &.disable {
          color: #b9b9b9;
        }
      }
      .country {
        .react-tel-input {
          height: 50px;
          width: 100%;
        }
        .flag-dropdown {
          width: 100%;
          .selected-flag {
            width: 100%;
          }
        }
      }
      .grayBox {
        padding: 0 6px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        height: 50px;
        background: #f7f7f7;
        border: 1px solid #e8e8e8;
        box-sizing: border-box;
        border-radius: 5px;
      }
      input {
        height: 50px;
      }
      label {
        display: block;
        font-weight: bold;
        font-size: 12px;
        line-height: 15px;
        margin-bottom: 5px;
        color: #171717;
      }
      .flex_between {
        display: flex;
        justify-content: space-between;
      }
      .connectBtn {
        width: 145px;
        height: 50px;

        border: 1px solid #171717;
        box-sizing: border-box;
        border-radius: 5px;
        color: #171717;
        background-color: white;
      }
    }
  }
`;

export default UserInfo;
