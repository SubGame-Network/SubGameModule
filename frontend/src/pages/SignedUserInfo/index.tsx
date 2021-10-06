import React, { useEffect, useState } from "react";
import { FormattedMessage } from "react-intl";
import styled from "styled-components";
import { useFormik } from "formik";
import * as yup from "yup";
import { usePolkadotJS } from "@polkadot/api-provider";
import Button from "../../components/Button";
import Input from "../../components/Input";
import sliceAddress from "utils/sliceAddress";
import Loading from "../../components/Loading";

import Banner from "../../components/Banner";
import { useApiGetUserInfo, useApiEditName } from "hooks/useApi/useUser";
import { useHistory } from "react-router-dom";
import PhoneInput from "react-phone-input-2";

interface CountryData {
  name: string;
  dialCode: string;
  countryCode: string;
  format: string;
}

function SignedUserInfo() {
  const [selectedCountry, setSelectedCountry] = useState("");
  const [countryCode, setcountryCode] = useState("");

  const history = useHistory();

  const { data, refetch, isLoading } = useApiGetUserInfo();
  const { mutateAsync: editAsync } = useApiEditName();
  const userInfo = data?.data.data;
  const {
    state: { currentAccount },
  } = usePolkadotJS();
  const {
    handleChange,
    errors,
    touched,
    handleSubmit,
    values,
    // eslint-disable-next-line react-hooks/rules-of-hooks
  } = useFormik({
    initialValues: {
      nickName: userInfo?.nickName,
      phone: userInfo?.phone,
    },
    validationSchema: yup.object().shape({
      nickName: yup.string().max(20, "Wordlimit").required("Required"),
      phone: yup.string(),
    }),
    onSubmit: async ({ nickName, phone }) => {
      let data: any;
      data = {
        nickName: nickName,
        phone: phone,
        countryCode: countryCode,
        country: selectedCountry,
      };
      const api = await editAsync(data);
      if (api.status === 200 && api.data.data) {
        refetch();
      }
    },
  });
  useEffect(() => {
    if (data?.data.code !== 200) {
      history.push("/userinfo");
    }
  }, [data?.data.code, history]);
  useEffect(() => {
    if (userInfo?.country) {
      console.log(userInfo?.country);

      setSelectedCountry(userInfo?.country);
    }
    if (userInfo?.countryCode) {
      console.log(userInfo?.countryCode);

      setcountryCode(userInfo.countryCode);
    }
  }, [userInfo?.country, userInfo?.countryCode]);
  return (
    <UserInfoStyle>
      <Banner text="userInfo" />

      <div className="wrap">
        {isLoading ? (
          <Loading />
        ) : (
          <form>
            <div className="formArea">
              <div className="item">
                <label htmlFor="">
                  <FormattedMessage id="SubGameAddress" />
                </label>
                <p className="value">
                  {" "}
                  {sliceAddress(currentAccount?.address, 10, 10)}
                </p>
              </div>
              <div className="item">
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
                  value={values.nickName}
                />
              </div>

              <div className="item country">
                <label htmlFor="">
                  <FormattedMessage id="Country" />
                </label>
                {selectedCountry && (
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
                )}{" "}
              </div>

              <div className="item">
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
                    value={values.phone}

                    // className={`${touched.phone && errors.phone ? "error" : ""} `}
                    // errorMsg={errors.phone && touched.phone ? errors.phone : ""}
                  />
                </div>
              </div>

              <div className="item">
                <label htmlFor="">
                  <FormattedMessage id="Email" />
                </label>
                <p className="value">{userInfo?.email}</p>
              </div>

              <div></div>
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
        )}
      </div>
    </UserInfoStyle>
  );
}

const UserInfoStyle = styled.div`
  .wrap {
    padding: 40px 0 100px;

    .formArea {
      display: grid;
      grid-template-columns: 1fr 1fr;
      grid-template-rows: 1fr 1fr 1fr 1fr;
      grid-column-gap: 40px;
      grid-row-gap: 30px;
      .phoneColumn {
        display: grid;
        grid-template-columns: 96px 1fr;
        grid-column-gap: 30px;
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
      .item {
        padding-bottom: 20px;
        &.country {
          .react-tel-input {
            height: 50px;
            width: 100%;
          }
          .flag-dropdown {
            width: 100%;
            background-color: white;
            border: 1px solid #e8e8e8;
            .selected-flag {
              width: 100%;
            }
          }
        }
      }
      .value {
        font-size: 16px;
        line-height: 160%;
        height: 50px;
        display: flex;
        align-items: center;
        color: #171717;
        margin-top: 5px;
      }
    }
  }
`;

export default SignedUserInfo;
