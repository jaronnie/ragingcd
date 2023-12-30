package com.jaronnie.ragingcd.stdb.system.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.ragingcd.stdb.system.domain.bo.AddUserBo;
import com.jaronnie.ragingcd.stdb.system.domain.query.PageQuery;
import com.jaronnie.ragingcd.stdb.system.domain.query.SearchUserQuery;
import com.jaronnie.ragingcd.stdb.system.domain.vo.TableDataInfo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.UserVo;
import com.jaronnie.ragingcd.stdb.system.service.IUserService;
import com.jaronnie.ragingcd.stdb.common.util.R;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/system/user_manage")
@Api(tags = "用户管理")
public class SystemUserManageController {
    private final IUserService iUserService;

    @ApiOperation(value = "获取用户列表")
    @GetMapping("/list")
    @SaCheckLogin
    public R<TableDataInfo<UserVo>> getUsers(@ModelAttribute PageQuery pageQuery, @ModelAttribute SearchUserQuery searchUserQuery) {
        return R.ok(iUserService.queryPageList(pageQuery, searchUserQuery));
    }

    @ApiOperation(value = "添加用户")
    @PostMapping("/add")
    @SaCheckLogin
    public R<UserVo> add(@RequestBody AddUserBo addUserBo) {
        return R.ok(iUserService.add(addUserBo));
    }

    @ApiOperation(value = "删除用户")
    @GetMapping("/delete/{id}")
    @SaCheckLogin
    public R<Boolean> delete(@Valid @PathVariable Integer id) {
        return R.ok(iUserService.delete(id));
    }
}
