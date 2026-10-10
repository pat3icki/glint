import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-abi-exporter";
import "./export-bin";
import "./export-go-abigen";

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.28",
        settings: {
          optimizer: { enabled: true, runs: 200 },
          evmVersion: "cancun",
        },
      },
      {
        version: "0.8.24",
        settings: {
          optimizer: { enabled: true, runs: 200 },
          evmVersion: "cancun",
        },
      },
    ],
  },
  paths: {
    sources: "./out/contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts",
  },
  abiExporter: {
    path: "./out/abi",
    runOnCompile: true,
    clear: true,
    flat: true,
    only: ["GlintDough", "GlintCardPaymaster"],
    spacing: 2,
    format: "json",
  },
  goabigen: {
    outdir: "./out/bindings-go/",
    runOnCompile: true,
    only: ["GlintDough", "GlintCardPaymaster"],
  },
};

export default config;
