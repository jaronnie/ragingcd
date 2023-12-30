package com.jaronnie.ragingcd.stdb.system.service.impl;

import com.jaronnie.ragingcd.stdb.system.service.IMailService;
import lombok.RequiredArgsConstructor;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;

import static com.jaronnie.ragingcd.stdb.common.enumeration.errcode.MailErrorCodeEnum.SendMailError;

@Service
@RequiredArgsConstructor
public class MailServiceImpl implements IMailService {
    private final JavaMailSender javaMailSender;
    @Override
    public Boolean sendSimpleEmail(String from, String to, String subject, String text) {
        SimpleMailMessage mailMessage = new SimpleMailMessage();
        mailMessage.setFrom(from);
        mailMessage.setTo(to);

        mailMessage.setSubject(subject);
        mailMessage.setText(text);

        try {
            javaMailSender.send(mailMessage);
            return true;
        } catch (Exception e) {
            throw SendMailError.newException(e);
        }
    }
}
