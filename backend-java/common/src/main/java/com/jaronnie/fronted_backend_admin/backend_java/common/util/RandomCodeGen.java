package com.jaronnie.fronted_backend_admin.backend_java.common.util;

import java.util.concurrent.ThreadLocalRandom;

public class RandomCodeGen {
    public static String genRandomCode(int length) {
        ThreadLocalRandom random = ThreadLocalRandom.current();
        StringBuilder codeBuilder = new StringBuilder();
        for (int i = 0; i < length; i++) {
            int digit = random.nextInt(10);
            codeBuilder.append(digit);
        }

        return codeBuilder.toString();
    }
}