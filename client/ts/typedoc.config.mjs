// @ts-check
import { OptionDefaults } from "typedoc";

/** @type {Partial<import("typedoc").TypeDocOptions>} */
export default {
  entryPoints: ["src/index.ts"],
  out: "../../docs/sdk",
  // protoc-gen-es annotates every generated symbol with @generated; register
  // the tag so typedoc renders it instead of warning once per symbol.
  blockTags: [...OptionDefaults.blockTags, "@generated"],
};
