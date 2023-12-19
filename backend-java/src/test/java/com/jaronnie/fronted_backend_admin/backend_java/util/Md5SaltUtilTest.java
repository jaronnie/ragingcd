package com.jaronnie.fronted_backend_admin.backend_java.util;

import static org.junit.jupiter.api.Assertions.*;

class Md5SaltUtilTest {
    public static void main(String[] args) {
        String data = "123456";
        String salt = "jaronnie";
        String hash = Md5SaltUtil.generateHash(data, salt);
        System.out.println(hash);
        assertTrue(Md5SaltUtil.verify(data, salt, hash));
    }
}