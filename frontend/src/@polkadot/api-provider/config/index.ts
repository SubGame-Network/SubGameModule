import { Node } from "../../../config/config.json";
const DAPP_NAME = "Sub Game Center";

const RPC_URLS: { [netWork: string]: string } = {
  Polkadot: "wss://polkadot.api.onfinality.io/public-ws",
  Westend: "wss://westend-rpc.polkadot.io",
  Local: "ws://127.0.0.1:9944",
};

const DEFAULT_RPC_URL = Node || "ws://127.0.0.1:9944";

const DEFAULT_CONNECT_WALLET = false;

export { RPC_URLS, DEFAULT_RPC_URL, DAPP_NAME, DEFAULT_CONNECT_WALLET };
