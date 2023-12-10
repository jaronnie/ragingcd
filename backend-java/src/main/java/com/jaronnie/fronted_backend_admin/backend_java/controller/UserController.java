package com.jaronnie.fronted_backend_admin.backend_java.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LogUpBo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LoginBo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.LoginResponseVo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.UserVo;
import com.jaronnie.fronted_backend_admin.backend_java.service.IUserService;
import com.jaronnie.fronted_backend_admin.backend_java.util.R;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/user")
public class UserController {
    private final IUserService iUserService;

    @GetMapping("/list")
    @SaCheckLogin
    public R<TableDataInfo<UserVo>> getUsers(@ModelAttribute PageQuery pageQuery) {
        return R.ok(iUserService.queryPageList(pageQuery));
    }

    @PostMapping("/add")
    @SaCheckLogin
    public R<UserVo> add(@RequestBody LogUpBo logUpBo) {
        return R.ok(iUserService.logUp(logUpBo));
    }

    @PostMapping("/login")
    public R<LoginResponseVo> login(@RequestBody LoginBo loginBo) {
        return R.ok(iUserService.login(loginBo));
    }

    @GetMapping("/info")
    @SaCheckLogin
    public R<UserVo> info(@RequestHeader("Authorization") String authorization) {
        return R.ok(iUserService.info(authorization));
    }
}
