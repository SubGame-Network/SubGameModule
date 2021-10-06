import React from "react";
import { FormattedMessage } from "react-intl";
import styled from "styled-components";
import Button from "../Button";
import { Link } from "react-router-dom";
function NotJoin() {
  return (
    <>
      <Container className="loading">
        <p>
          <FormattedMessage id="notjoin" />
        </p>
        <Link to="/userinfo">
          {" "}
          <Button text="gotojoin" className="btn" />
        </Link>
      </Container>
    </>
  );
}

const Container = styled.div`
  width: 100% !important;
  min-height: 50vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  padding-top: 80px;
  font-size: 16px;
  line-height: 160%;
  /* identical to box height, or 26px */

  text-align: center;

  /* black */

  color: #171717;
  .btn {
    margin-top: 40px;
    width: 260px;
    height: 50px;
  }
`;

export default NotJoin;
