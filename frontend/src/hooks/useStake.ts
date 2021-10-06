import { usePolkadotJS } from "../@polkadot/api-provider";
import { web3FromAddress } from "@polkadot/extension-dapp";
import "@subgame/types/interfaces/subgameStakeNft";
import { dev } from "../config/config.json";

import { useMutation } from "react-query";
import { ApiPromise } from "@polkadot/api";
import type {
  DispatchError,
  DispatchInfo,
  EventRecord,
} from "@polkadot/types/interfaces/system/types";
const useStake = () => {
  const {
    state: { api, apiState, keyring, currentAccount },
  } = usePolkadotJS();

  const parseError = (
    api: ApiPromise,
    events: EventRecord[],
    targetModule: string
  ) => {
    let errorMessage: string = "";
    events
      .filter(({ event }) => api.events.system.ExtrinsicFailed.is(event))
      .forEach(({ event: { data } }) => {
        const [error] = (data as unknown) as [DispatchError, DispatchInfo];
        if (error.isModule) {
          const decoded = api.registry.findMetaError(error.asModule);
          const { docs, method, section } = decoded;

          const message = `${section}.${method}: ${docs}`;

          errorMessage = method;

          dev && console.log("發生錯誤 ->", message);
        } else {
          dev && console.log(error.toString());
        }
      });

    return errorMessage;
  };
  const useSendStake = () => {
    return useMutation(
      async ({
        program_id,
        pallet_id,
        callBack,
      }: {
        program_id: number;
        pallet_id: number;
        callBack?: any;
      }) => {
        if (keyring && api && apiState === "READY") {
          try {
            if (currentAccount) {
              // const injector = await web3FromAddress(currentAccount?.address);
              // const result = await api.tx.crowdloan
              //   .contribute(paraID, value, null)
              //   .signAndSend(currentAccount?.address, {
              //     signer: injector.signer,
              //   });
              console.log(program_id);
              console.log(pallet_id);
              const injector = await web3FromAddress(currentAccount?.address);
              const result = await api.tx.subgameStakeNft
                .stake(program_id, pallet_id)
                .signAndSend(
                  currentAccount?.address,
                  {
                    signer: injector.signer,
                  },
                  async (result) => {
                    if (result.isInBlock) {
                      const { events } = result;
                      const responseErrorMessage = parseError(
                        api,
                        events,
                        "subgameStakeNft"
                      );
                      console.log(responseErrorMessage);
                      if (callBack) callBack(responseErrorMessage);
                    }
                  }
                );
              console.log(result);
            }
          } catch (error) {
            console.error(error);
          }
        }
      },
      {
        onError: (error: any) => {
          console.log(error);
        },
      }
    );
  };
  return { useSendStake };
};

export default useStake;
