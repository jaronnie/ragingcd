package com.jaronnie.fronted_backend_admin.backend_java.util;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import javax.xml.bind.DatatypeConverter;

public class Md5SaltUtil {
    public static String generateHash(String data, String salt) {
        try {
            MessageDigest messageDigest = MessageDigest.getInstance("MD5");
            messageDigest.update(salt.getBytes());
            byte[] hashedBytes = messageDigest.digest(data.getBytes());
            return DatatypeConverter.printHexBinary(hashedBytes);
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        }
        return null;
    }

    public static boolean verify(String data, String salt, String hash) {
        String generatedHash = generateHash(data, salt);
        return generatedHash != null && generatedHash.equals(hash);
    }
}
