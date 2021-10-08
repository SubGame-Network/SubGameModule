import { Dispatch } from "react";
import { web3Accounts, web3Enable } from "@polkadot/extension-dapp";
import { WsProvider, ApiPromise } from "@polkadot/api";
import { options } from "@subgame/api";
import keyring from "@polkadot/ui-keyring";

import { I_INIT_STATE, Actions } from "./useReducer";

import { DAPP_NAME } from "./config";

export const connect = async (
  state: I_INIT_STATE,
  dispatch: Dispatch<Actions>
) => {
  const apiState = state.apiState;
  const nodeURL = state.nodeURL;

  // 連線功能只執行一次，連線失敗會自動重新連線
  if (apiState) return;

  const provider = new WsProvider(nodeURL);

  const _api = new ApiPromise(options({ provider }));

  _api.on("connected", () => dispatch({ type: "CONNECTING" }));

  _api.on("error", (err) => dispatch({ type: "CONNECT_ERROR", payload: err }));

  _api.on("ready", (_api) =>
    dispatch({ type: "CONNECT_SUCCESS", payload: _api })
  );
};

export const loadAccounts = async (
  state: I_INIT_STATE,
  dispatch: Dispatch<Actions>
) => {
  if (!!state.keyring || state.keyringState) return;
  try {
    const exts = await web3Enable(DAPP_NAME);
    if (exts.length === 0) {
      throw new Error("extension not Find");
    }
    const allAccounts = await web3Accounts();

    keyring.loadAll({}, allAccounts);
    dispatch({ type: "SET_KEYRING", payload: keyring });
  } catch (error) {
    dispatch({ type: "KEYRING_ERROR", payload: error });
  }
};
