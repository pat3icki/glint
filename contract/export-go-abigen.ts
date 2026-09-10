import { task, subtask } from "hardhat/config";
import { TASK_COMPILE } from "hardhat/builtin-tasks/task-names";
import * as fs from "fs";
import * as path from "path";
import { execSync } from "child_process";

export interface GoAbigenUserConfig {
  outdir: string;
  runOnCompile?: boolean;
  only?: string[];
}

declare module "hardhat/types/config" {
  interface HardhatUserConfig {
    goabigen?: GoAbigenUserConfig;
  }
  interface HardhatConfig {
    goabigen?: GoAbigenUserConfig;
  }
}

subtask("export-go-abigen", async (_, hre) => {
  const config = hre.config.goabigen;
  const abiDir = path.join(hre.config.paths.root, "abi");
  const binDir = path.join(hre.config.paths.root, "bin");
  const bindingsDir = path.join(
    hre.config.paths.root,
    config?.outdir || ""
  );


  const contracts = config?.only || ["GiftCardNFT", "GiftCardPaymaster"];

  for (const name of contracts) {
    const abiFile = path.join(abiDir, `${name}.json`);
    const binFile = path.join(binDir, `${name}.bin`);
    const pkgName = name.toLowerCase();
    const directoryPath = path.join(bindingsDir, pkgName)


    const outFile = path.join(directoryPath, `${pkgName}.go`);


    if (!fs.existsSync(directoryPath)) {
      fs.mkdirSync(directoryPath, { recursive: true });
    }

    if (fs.existsSync(abiFile) && fs.existsSync(binFile)) {
      try {
        const cmd = `abigen --abi="${abiFile}" --bin="${binFile}" --pkg="${pkgName}" --out="${outFile}"`;
        execSync(cmd, { stdio: "ignore" });
      } catch (err) {
        // Handle gracefully if abigen execution fails
      }
    }
  }
});

task(TASK_COMPILE, async (args, hre, runSuper) => {
  const result = await runSuper(args);
  await hre.run("export-go-abigen");
  return result;
});
