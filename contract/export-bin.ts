import { task, subtask } from "hardhat/config";
import { TASK_COMPILE } from "hardhat/builtin-tasks/task-names";
import * as fs from "fs";
import * as path from "path";

subtask("export-bin", async (_, hre) => {
  const binDir = path.join(hre.config.paths.root, "out", "bin");
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }

  const contractsToExport = ["GlintDough", "GlintCardPaymaster"];

  for (const contractName of contractsToExport) {
    try {
      const artifact = await hre.artifacts.readArtifact(contractName);
      if (artifact.bytecode && artifact.bytecode !== "0x") {
        const binPath = path.join(binDir, `${contractName}.bin`);
        const cleanBytecode = artifact.bytecode.startsWith("0x")
          ? artifact.bytecode.slice(2)
          : artifact.bytecode;
        fs.writeFileSync(binPath, cleanBytecode, "utf8");
      }
    } catch (err) {
      // Contract artifact might not exist if not compiled
    }
  }
});

task(TASK_COMPILE, async (args, hre, runSuper) => {
  const result = await runSuper(args);
  await hre.run("export-bin");
  return result;
});
