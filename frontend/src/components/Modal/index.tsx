import React, { useEffect } from "react";
import styled from "styled-components";

function Modal(props: any) {
  useEffect(() => {
    if (window.screen.width < 1024) {
      const html = document.querySelector("html");
      if (html) {
        html.style.overflow = "hidden";
      }
    }
    return () => {
      if (window.screen.width < 1024) {
        const html = document.querySelector("html");
        if (html) {
          html.style.overflow = "auto";
        }
      }
    };
  }, []);
  return (
    <>
      <ModalMask
        modalShow={props.modalShow}
        customOpacity={props.customOpacity}
        customZindex={props.customZindex}
        mobile={!!props.mobile}
      >
        <div className="modal-main">{props.children}</div>
      </ModalMask>
    </>
  );
}

const ModalMask = styled.div<{
  modalShow: boolean;
  customOpacity?: string;
  customZindex?: string;
  mobile?: boolean;
}>`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: ${({ customZindex }) => {
    return !!customZindex ? parseInt(customZindex) : 6;
  }};

  background: ${({ customOpacity }) => {
    return `rgba(35, 24, 21, ${
      !!customOpacity ? parseInt(customOpacity) : 0.2
    })`;
  }};
  display: ${({ modalShow }): string => {
    return modalShow ? "block " : "none";
  }};
  .modal-main {
    position: fixed;
    background: transparent;
    height: auto;
    z-index: ${({ customZindex }) => {
      return !!customZindex ? parseInt(customZindex) : 1;
    }};
    top: 40px;
    left: 50%;

    transform: translate(-50%, 0%);
  }
`;

export default Modal;
