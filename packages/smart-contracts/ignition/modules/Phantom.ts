import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const PhantomModule = buildModule("PhantomModule", (m) => {
  const phantom003 = m.contract("Phantom", [], { id: "phantom__0_0_3" });

  return { phantom003 };
});

export default PhantomModule;
