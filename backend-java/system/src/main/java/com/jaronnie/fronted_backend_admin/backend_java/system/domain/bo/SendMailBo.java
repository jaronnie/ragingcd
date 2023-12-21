package com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo;

import lombok.Data;

@Data
public class SendMailBo {
    private String subject;
    private String text;
    private String to;
}
