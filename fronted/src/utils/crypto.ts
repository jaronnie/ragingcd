import { JSEncrypt } from "jsencrypt";

export const encrypt = (data: string, publicKey: string) => {
  const encryptor = new JSEncrypt();
  encryptor.setPublicKey(publicKey);
  return encryptor.encrypt(<string>data);
};
