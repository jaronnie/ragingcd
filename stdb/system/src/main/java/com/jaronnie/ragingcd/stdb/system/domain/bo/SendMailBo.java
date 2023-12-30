package com.jaronnie.ragingcd.stdb.system.domain.bo;

import lombok.Data;

@Data
public class SendMailBo {
    private String subject;
    private String text;
    private String to;
}
