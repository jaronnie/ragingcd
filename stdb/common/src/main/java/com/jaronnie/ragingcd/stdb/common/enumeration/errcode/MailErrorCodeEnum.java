package com.jaronnie.ragingcd.stdb.common.enumeration.errcode;

import com.jaronnie.ragingcd.stdb.common.component.exception.assertion.BusinessExceptionAssert;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum MailErrorCodeEnum implements BusinessExceptionAssert {
    SendMailError(20001, "发送邮件失败");

    private final int code;
    private final String message;
}
