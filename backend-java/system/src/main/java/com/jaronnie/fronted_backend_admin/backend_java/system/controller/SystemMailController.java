package com.jaronnie.fronted_backend_admin.backend_java.system.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.fronted_backend_admin.backend_java.common.util.R;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.SendMailBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.service.IMailService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.*;

@Api(tags = "邮件管理")
@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/system/mail")
public class SystemMailController {
    private final IMailService iMailService;

    @Value("${spring.mail.username}")
    private String From;

    @SaCheckLogin
    @ApiOperation(value = "发送邮件")
    @PostMapping("/send")
    public R<Boolean> send(@RequestBody SendMailBo sendMailBo) {
        return R.ok(iMailService.sendSimpleEmail(From, sendMailBo.getTo(), sendMailBo.getSubject(), sendMailBo.getText()));
    }
}
