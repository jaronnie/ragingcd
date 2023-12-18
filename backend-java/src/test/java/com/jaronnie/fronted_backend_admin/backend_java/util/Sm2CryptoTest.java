package com.jaronnie.fronted_backend_admin.backend_java.util;

import cn.hutool.core.codec.Base64;
import cn.hutool.core.util.HexUtil;

import static org.junit.jupiter.api.Assertions.*;

class Sm2CryptoTest {
    public static void main(String[] args) {
        // base64
        String publicKeyBase64String = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEudKajERAfAePEgbn3dxnXwAVFWcoxdY53/flvIAPV+GcVKkNms95ixQgivw6sNcxzovOV/8jXamKJI1WXINozw==";
        // 解码
        byte[] publicKeyBytes = Base64.decode(publicKeyBase64String);
        // 转化成 hex 16 进制
        String publicKey = HexUtil.encodeHexStr(publicKeyBytes);

        String privateKeyBase64String = "MHcCAQEEIPdEbTstcgS4bIscJ0wz+6UhKexwUA33S0bhBfoZ/tDhoAoGCCqBHM9VAYItoUQDQgAEudKajERAfAePEgbn3dxnXwAVFWcoxdY53/flvIAPV+GcVKkNms95ixQgivw6sNcxzovOV/8jXamKJI1WXINozw==";
        byte[] privateKeyBytes = Base64.decode(privateKeyBase64String);
        String privateKey = HexUtil.encodeHexStr(privateKeyBytes);

        String str = "123456";
        String encrypt = Sm2Crypto.encryptBase64(str, publicKey);
        System.out.println("加密后结果：" + encrypt);

        String decrypt = Sm2Crypto.decryptBase64(encrypt, privateKey);
        System.out.println("解密后结果：" + decrypt);
        assertEquals(str, decrypt);
    }
}