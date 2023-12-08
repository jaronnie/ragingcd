package com.jaronnie.fronted_backend_admin.backend_java.controller;

import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LoginQuery;
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
    public R<TableDataInfo<UserVo>> getUsers() {
        return R.ok(iUserService.queryPageList(new PageQuery()));
    }

    @PostMapping("/login")
    public R<LoginResponseVo> login(@RequestBody LoginQuery loginQuery) {
        return R.ok(iUserService.login(loginQuery));
    }

    @GetMapping("/info")
    public R<UserVo> info(@RequestParam String token) {
        return R.ok(iUserService.info(token));
    }
}
