import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const FuseNewsModule = buildModule("FuseNewsModule", (m) => {

  const News = m.contract("FuseNewsContract", []);

  return { News };
});

export default FuseNewsModule;
