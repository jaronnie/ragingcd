package com.jaronnie.fronted_backend_admin.backend_java.system.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.fronted_backend_admin.backend_java.common.util.R;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.LoginBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.LoginVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.PublicKeyVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.UserVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.service.IUserService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/system/user")
@Api(tags = "用户管理")
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

    @ApiOperation(value = "获取用户信息")
    @GetMapping("/info")
    @SaCheckLogin
    public R<UserVo> info() {
        return R.ok(iUserService.info());
    }

    @ApiOperation(value = "获取公钥")
    @GetMapping("/public-key")
    public R<PublicKeyVo> getPublicKey() {
        return R.ok(iUserService.getPublicKey());
    }
}
