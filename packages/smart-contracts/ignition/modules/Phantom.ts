import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const PhantomModule = buildModule("PhantomModule", (m) => {
  const phantom = m.contract("Phantom", [], { id: "phantom__0_0_0" });

  return { phantom };
});

export default PhantomModule;
