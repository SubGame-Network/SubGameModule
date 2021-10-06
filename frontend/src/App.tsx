import React, { useState, createContext, useEffect } from "react";
import { Cookies } from "react-cookie";

import IntlProvider from "./i18n";
import ThemeProvider from "./theme";
import Router from "./routers/Router";
import Feedback, {
  IFeedBackData,
  TType,
  TMessage,
  TMessageValues,
} from "./components/Feedback";

import { usePolkadotJS } from "@polkadot/api-provider";

const Cookie = new Cookies();
interface IAppContext {
  Cookie: Cookies;
  locale: string;
  showFeedBack: (
    type: TType,
    message?: TMessage,
    messageValues?: TMessageValues
  ) => void;
  setLocale: (a: string) => void;
}

export const AppContext = createContext<IAppContext>({} as IAppContext);
const cookie = new Cookies();

function App() {
  const [locale, setLocale] = useState(Cookie.get("locale") || "en");
  const [feedbackShow, setFeedbackShow] = useState(false);

  const [feedbackProps, setFeedbackProps] = useState<IFeedBackData>({
    type: "Submitted",
  });

  const {
    state: { currentAccount },
  } = usePolkadotJS();
  useEffect(() => {
    if (currentAccount) {
      cookie.set("subgame_module_address", currentAccount.address);
    }
  }, [currentAccount]);
  const showFeedBack = (
    type: TType,
    message?: TMessage,
    messageValues?: TMessageValues
  ) => {
    setFeedbackProps({
      type,
      message,
      messageValues,
    });
    setFeedbackShow(true);
  };
  // const handleRedirect = () => {
  //   Cookie.remove("subgame_stake_token");
  //   Cookie.remove("stakeB2B_uuid");
  //   Cookie.remove("stakeB2B_account");
  //   window.location.reload();
  // };

  return (
    <IntlProvider locale={locale}>
      <ThemeProvider>
        <AppContext.Provider
          value={{
            Cookie,
            locale,
            showFeedBack,
            setLocale,
          }}
        >
          {feedbackShow && (
            <Feedback setFeedbackShow={setFeedbackShow} {...feedbackProps} />
          )}
          <Router />
        </AppContext.Provider>
      </ThemeProvider>
    </IntlProvider>
  );
}

export default App;
