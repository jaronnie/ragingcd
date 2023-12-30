package com.jaronnie.ragingcd.stdb.system.service;

public interface IMailService {
    Boolean sendSimpleEmail(String from, String to, String subject, String text);
}
