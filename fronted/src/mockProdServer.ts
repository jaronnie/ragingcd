import { createProdMockServer } from "vite-plugin-mock/es/createProdMockServer";

import user from "../mock/user";
export function setupProdMockServer() {
  createProdMockServer([...user]);
}
