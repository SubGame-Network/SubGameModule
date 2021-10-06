import React from "react";
import styled from "styled-components";

function Loading() {
  return (
    <>
      <Container className="loading">
        <img src="./images/NoData.svg" alt="Loading" />
      </Container>
    </>
  );
}

const Container = styled.div`
  width: 100% !important;
  min-height: 50vh;
  display: flex;
  justify-content: center;
  padding-top: 80px;
  img {
    width: 197px;
    height: 208px;
  }
`;

export default Loading;
