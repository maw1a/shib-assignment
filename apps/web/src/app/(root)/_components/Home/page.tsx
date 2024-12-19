"use client";

import { useEffect } from "react";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAccount } from "wagmi";
import { MintButton } from "./mint";
import { toast } from "sonner";
import {
  CardContainer,
  CardBody,
  CardItem,
} from "@repo/ui/components/ui/3d-card";
import { fetchTokenByOwner } from "@/lib/api";

export function HomePage() {
  const account = useAccount();
  const queryClient = useQueryClient();

  const { isPending, data, error } = useQuery({
    queryKey: ["token"],
    queryFn: () => fetchTokenByOwner(account.address),
  });

  useEffect(() => {
    queryClient.prefetchQuery({
      queryKey: ["token"],
      queryFn: () => fetchTokenByOwner(account.address),
    });
  }, [account]);

  useEffect(() => {
    if (error) toast(error.message);
  }, [error]);

  if (!account.address)
    return (
      <div className="w-full h-full flex-1 flex justify-center items-center text-6xl font-semibold text-slate-500 pb-32">
        <p>Connect to continue...</p>
      </div>
    );

  if (isPending)
    return (
      <div className="w-full h-full flex-1 flex justify-center items-center text-6xl font-semibold text-slate-500 pb-32">
        <p>Loading your Phantom...</p>
      </div>
    );

  return (
    <div className="w-full h-full flex-1 flex justify-center items-center pb-32">
      {data && data.token_id ? (
        <CardContainer>
          <CardBody className="bg-gray-50 relative group/card border-black/[0.1] w-auto sm:w-[20rem] h-auto rounded-xl px-12 py-8 border space-y-4">
            <CardItem
              translateZ="40"
              className="text-base text-slate-500 w-full text-center font-medium"
            >
              Your Phantom
            </CardItem>
            <CardItem
              translateZ="60"
              className="text-[10rem] w-full text-center"
            >
              👻
            </CardItem>
            <CardItem
              translateZ="40"
              className="w-full text-slate-700 text-2xl text-center font-bold font-mono"
            >
              {data.symbol} #{data.token_id}
            </CardItem>
            <CardItem translateZ="40" className="w-full">
              <button
                className="w-full text-xl font-medium rounded-lg px-6 py-3 text-white bg-blue-700 hover:bg-blue-600 transition-all shadow-2xl shadow-blue-500 disabled:cursor-not-allowed"
                disabled
              >
                Transfer Token
              </button>
            </CardItem>
          </CardBody>
        </CardContainer>
      ) : (
        <div className="flex flex-col items-center gap-20 pb-32">
          <p className="text-8xl opacity-60 grayscale">👻</p>
          <MintButton to={account.address} />
        </div>
      )}
    </div>
  );
}
