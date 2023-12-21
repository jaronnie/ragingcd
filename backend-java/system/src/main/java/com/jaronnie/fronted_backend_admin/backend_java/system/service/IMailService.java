package com.jaronnie.fronted_backend_admin.backend_java.system.service;

import java.util.Base64;

public interface IMailService {
    Boolean sendSimpleEmail(String from, String to, String subject, String text);
}
