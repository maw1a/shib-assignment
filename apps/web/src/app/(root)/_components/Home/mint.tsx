import { useMutation, useQueryClient } from "@tanstack/react-query";
import { fetchTokenByOwner, mintToken } from "@/lib/api";
import { toast } from "sonner";

export function MintButton({ to }: { to?: string }) {
  const queryClient = useQueryClient();

  const mint = useMutation({
    mutationFn: async (to: string | undefined) => {
      const mint = await mintToken(to);

      if (mint?.ok) {
        toast(mint.message);
        queryClient.fetchQuery({
          queryKey: ["token"],
          queryFn: () => fetchTokenByOwner(to),
        });
        return;
      }

      toast(mint?.message || "failed to mint NFT");
    },
  });

  return (
    <button
      className="text-3xl font-medium rounded-xl px-9 py-4 text-white bg-blue-700 hover:bg-blue-600 hover:scale-105 transition-all shadow-2xl shadow-blue-500 disabled:cursor-not-allowed"
      onClick={() => mint.mutate(to)}
      disabled={mint.isPending}
    >
      Mint Phantom
    </button>
  );
}
