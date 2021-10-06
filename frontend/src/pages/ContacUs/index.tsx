import React, { useState } from "react";
import { FormattedMessage } from "react-intl";
import styled from "styled-components";
import Banner from "../../components/Banner";
import { useFormik } from "formik";
import * as yup from "yup";
import Button from "../../components/Button";
import Select from "../../components/Select";
import { settingOption } from "../../api/types/global";
import { useApiConcat, useApiGetUserInfo } from "../../hooks/useApi/useUser";
import NotJoin from "../../components/NotJoin";
function ContacUs() {
  const { data: userInfoData } = useApiGetUserInfo();
  const currentAddressHasJoin = userInfoData?.data.code === 200;
  const { mutateAsync: cancatAsync } = useApiConcat();
  const [selectedType, setSelectedType] = useState<settingOption>({
    value: "",
    label: "Select",
  });
  const typeOption = [
    {
      value: "Module Suddestion",
      label: "Module Suddestion",
    },
    {
      value: "Others",
      label: "Others",
    },
  ];
  const {
    values,
    handleChange,
    errors,
    touched,
    handleSubmit,

    // eslint-disable-next-line react-hooks/rules-of-hooks
  } = useFormik({
    initialValues: {
      type: "",
      content: "",
    },
    validationSchema: yup.object().shape({
      content: yup.string().max(25000, "Wordlimit").required("Required"),
    }),
    onSubmit: async ({ content }) => {
      let data: any;
      data = {
        type: selectedType.value,
        contact: content,
      };
      const api = await cancatAsync(data);
      if (api.status === 200 && api.data.data) {
      }
    },
  });
  return (
    <ContacUsStyle>
      {" "}
      <Banner text="ContactUs" />
      <div className="contactusWrap">
        {currentAddressHasJoin ? (
          <form>
            {" "}
            <label>
              <FormattedMessage id="Type" />
            </label>
            <Select
              options={typeOption}
              value={selectedType}
              onChange={(e: any) => {
                setSelectedType({ value: e.value, label: e.label });
              }}
              customWidth="100%"
              customHeight="50px"
            />
            <label className="mt30">
              <FormattedMessage id="ContactDetail" />
            </label>
            <textarea
              name="content"
              id="content"
              value={values.content}
              onChange={handleChange}
              className={errors.content && touched.content ? "error" : ""}
            >
              It was a dark and stormy night...
            </textarea>
            <Button
              text="Submit"
              className="btn"
              type="button"
              onClick={() => {
                handleSubmit();
              }}
            />
          </form>
        ) : (
          <NotJoin />
        )}
      </div>
    </ContacUsStyle>
  );
}

const ContacUsStyle = styled.div`
  .contactusWrap {
    padding: 40px 0 100px;
    max-width: 520px;
    margin: auto;
    label {
      display: block;
      margin-bottom: 5px;
      &.mt30 {
        margin-top: 30px;
      }
    }
    textarea {
      width: 100%;
      height: 173px;
      &.error {
        border: 1px solid #eb027d;
      }
    }
    .btn {
      background-color: #171717;
      margin-top: 30px;
      height: 50px;
      :disabled {
        background-color: #e8e8e8;
        color: white;
      }
    }
  }
`;

export default ContacUs;
