import { ConnectButton } from "@rainbow-me/rainbowkit";
import { WalletConnectProvider } from "./_components/WalletConnect";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <WalletConnectProvider>
      <div className="flex flex-col min-h-screen font-[family-name:var(--font-geist-sans)] select-none">
        <header className="w-full px-8 py-6 text-lg flex justify-between items-center">
          <div className="flex justify-between items-center gap-4">
            <div className="text-3xl">👻</div>
            <div className="text-xl font-bold">Phantom</div>
          </div>
          <ConnectButton />
        </header>
        <main className="w-full flex-1 flex flex-col gap-8 items-center sm:items-start">
          {children}
        </main>
      </div>
    </WalletConnectProvider>
  );
}
