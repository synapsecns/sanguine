import { task } from "hardhat/config";
import * as fsWalk from '@nodelib/fs.walk';


task("merge", "Prints the list of accounts", async (_, hre) => {
  const qualifiedNames = await hre.artifacts.getAllFullyQualifiedNames();
  console.log(hre.config.paths.sources)

  const entries = fsWalk.walkSync(hre.config.paths.sources, new fsWalk.Settings());
  for (const entry of entries){
      let res = await hre.run("flatten", {
          files: [entry.path],
      })

      console.log(res)
  }
});