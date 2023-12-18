package com.jaronnie.fronted_backend_admin.backend_java.util;

import cn.hutool.core.util.StrUtil;
import cn.hutool.crypto.asymmetric.KeyType;
import cn.hutool.crypto.asymmetric.SM2;

public class Sm2Crypto {

    /**
     * SM2公钥加密
     *
     * @param content 原文
     * @param publicKey SM2公钥
     */
    public static String encryptBase64(String content, String publicKey) {
        SM2 sm2 = new SM2(null, publicKey);
        return sm2.encryptBase64(content, KeyType.PublicKey);
    }

    /**
     * SM2私钥解密
     *
     * @param encryptStr SM2加密字符串
     * @param privateKey SM2私钥
     */
    public static String decryptBase64(String encryptStr, String privateKey) {
        SM2 sm2 = new SM2(privateKey, null);
        return StrUtil.utf8Str(sm2.decrypt(encryptStr, KeyType.PrivateKey));
    }
}
