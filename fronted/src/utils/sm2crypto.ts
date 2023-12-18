import { sm2 } from "sm-crypto";
import { Buffer } from "buffer";

export const sm2Encrypt = (data: string, publicKey: string) => {
  const publicKeyHex = Buffer.from(publicKey, "base64").toString("hex");
  const cipherMode = 1;
  return sm2.doEncrypt(data, "04" + publicKeyHex, cipherMode);
};
