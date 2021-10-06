import React from "react";
import styled from "styled-components";

function Loading() {
  return (
    <>
      <Container className="loading">
        <img src="./images/Loading.svg" alt="Loading" />
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
  padding-top: 80px;

  @keyframes fadein {
    0% {
      opacity: 0.2;
    }
    50% {
      opacity: 1;
    }
    100% {
      opacity: 0.2;
    }
  }
  img {
    animation: fadein 2s ease;
    animation-iteration-count: infinite;
  }
`;

export default Loading;
