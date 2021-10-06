import { useEffect, useState } from "react";
import { usePolkadotJS } from "../@polkadot/api-provider";

import toFixedNumber from "../utils/toFixedNumber";

type TBalances = { [address: string]: string };

const usePolkadotBalance = () => {
  const {
    state: { api, keyring, apiState, chainDecimal },
  } = usePolkadotJS();
  const [balances, setBalances] = useState<TBalances | null>(null);

  useEffect(() => {
    const getBalances = async () => {
      if (keyring && api && apiState === "READY") {
        const accounts =
          keyring
            .getAccounts()
            .map((key) => keyring.encodeAddress(key.address)) || [];
        try {
          await api.query.system.account.multi(accounts, (balances) => {
            const arr = balances
              .map((balance) => balance.toJSON())
              .reduce<TBalances>((acc, balance: any, index) => {
                const value: number =
                  balance?.data?.free / Math.pow(10, chainDecimal);

                acc[accounts[index]] = toFixedNumber(value.toString());

                return acc;
              }, {});
            setBalances(arr);
          });
        } catch (error) {
          console.error(error);
        }
      }
    };
    getBalances();
  }, [api, apiState, keyring, chainDecimal]);

  return balances;
};

export default usePolkadotBalance;
