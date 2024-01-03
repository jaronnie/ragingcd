package com.jaronnie.ragingcd.stdb.system.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.ragingcd.stdb.common.util.R;
import com.jaronnie.ragingcd.stdb.system.domain.bo.LoginBo;
import com.jaronnie.ragingcd.stdb.system.domain.bo.RegisterUserBo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.LoginVo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.PublicKeyVo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.UserVo;
import com.jaronnie.ragingcd.stdb.system.service.IUserService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/system/user")
@Api(tags = "用户登录/注册")
public class SystemUserController {
    private final IUserService iUserService;

    @ApiOperation(value = "用户登录")
    @PostMapping("/login")
    public R<LoginVo> login(@RequestBody LoginBo loginBo) {
        return R.ok(iUserService.login(loginBo));
    }

    @ApiOperation(value = "用户登出")
    @SaCheckLogin
    @GetMapping("/logout")
    public R<Void> logout() {
        iUserService.logout();
        return R.ok("ok", null);
    }

    @ApiOperation(value = "发送邮件获取验证码")
    @GetMapping("/send_email/{email}")
    public R<Boolean> sendMail(@PathVariable String email, @RequestParam String username) {
        return R.ok(iUserService.sendEmail(email, username));
    }

    @ApiOperation(value = "用户注册")
    @PostMapping("/register")
    public R<UserVo> register(@RequestBody RegisterUserBo registerUserBo) {
        return R.ok(iUserService.register(registerUserBo));
    }

    @ApiOperation(value = "获取用户信息")
    @GetMapping("/info")
    @SaCheckLogin
    public R<UserVo> info() {
        return R.ok(iUserService.info());
    }

    @ApiOperation(value = "鉴权中间件")
    @GetMapping("/auth")
    @SaCheckLogin
    public ResponseEntity<UserVo> auth() {
        UserVo userInfo = iUserService.info();
        System.out.println("=========鉴权中间件, 获取用户信息=========");
        System.out.println(userInfo);

        HttpHeaders headers = new HttpHeaders();
        headers.add("X-Forward-Auth-Header", String.valueOf(userInfo.getId()));

        return ResponseEntity.ok()
                .headers(headers)
                .body(userInfo);
    }

    @ApiOperation(value = "获取公钥")
    @GetMapping("/public-key")
    public R<PublicKeyVo> getPublicKey() {
        return R.ok(iUserService.getPublicKey());
    }
}
