package com.jaronnie.ragingcd.stdb.common.util;

import cn.hutool.core.codec.Base64;
import cn.hutool.core.util.StrUtil;
import cn.hutool.crypto.asymmetric.KeyType;
import cn.hutool.crypto.asymmetric.RSA;


public class RsaCrypto {

    /**
     * RSA 私钥解密
     *
     * @param cipherText SM2加密字符串, base64 格式
     * @param privateKey RSA PKCS8 格式私钥
     */
    public static String decrypt(String cipherText, String privateKey) {
        RSA rsa = new RSA(privateKey, null);
        byte[] cipherBytes = Base64.decode(cipherText);
        return StrUtil.utf8Str(rsa.decrypt(cipherBytes, KeyType.PrivateKey));
    }
}