"use client";

import { getDefaultConfig } from "@rainbow-me/rainbowkit";
import { mainnet, sepolia } from "wagmi/chains";

export const wagmiConfig = getDefaultConfig({
  appName: "Shib Assignment",
  projectId: "108e190abafd0d00c2bfe876ebd76408",
  chains: [sepolia, mainnet],
  ssr: true,
});
