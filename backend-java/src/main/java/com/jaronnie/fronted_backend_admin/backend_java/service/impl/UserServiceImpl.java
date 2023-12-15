package com.jaronnie.fronted_backend_admin.backend_java.service.impl;

import cn.dev33.satoken.stp.StpUtil;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.core.toolkit.Wrappers;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LogUpBo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LoginBo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.domain.po.UserPo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.LoginResponseVo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.UserVo;
import com.jaronnie.fronted_backend_admin.backend_java.enumeration.errcode.UserErrorCodeEnum;
import com.jaronnie.fronted_backend_admin.backend_java.mapper.UserMapper;
import com.jaronnie.fronted_backend_admin.backend_java.service.IUserService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements IUserService {
    private final UserMapper baseMapper;

    @Override
    public TableDataInfo<UserVo> queryPageList(PageQuery pageQuery) {
        LambdaQueryWrapper<UserPo> lqw = Wrappers.lambdaQuery();
        lqw.orderByDesc(UserPo::getUpdateTime);

        Page<UserPo> page = baseMapper.selectPage(pageQuery.build(), lqw);

        TableDataInfo<UserVo> userVoTableDataInfo = new TableDataInfo<>();
        userVoTableDataInfo.setTotal(page.getTotal());
        List<UserVo> result = page.getRecords().stream().map(userPo -> UserVo.builder()
                .id(userPo.getId())
                .username(userPo.getUsername())
                .createTime(userPo.getCreateTime())
                .updateTime(userPo.getUpdateTime())
                .avatar(userPo.getAvatar())
                .build()).collect(Collectors.toList());
        userVoTableDataInfo.setRows(result);
        return userVoTableDataInfo;
    }

    @Override
    public UserVo info() {
        Integer id = StpUtil.getLoginIdAsInt();
        UserPo userPo = this.baseMapper.selectById(id);
        return UserVo.builder()
                .username(userPo.getUsername())
                .id(userPo.getId())
                .createTime(userPo.getCreateTime())
                .updateTime(userPo.getUpdateTime())
                .avatar(userPo.getAvatar())
                .build();
    }

    @Override
    public LoginResponseVo login(LoginBo loginBo) {
        LambdaQueryWrapper<UserPo> lqw = Wrappers.lambdaQuery();
        lqw.eq(UserPo::getUsername, loginBo.getUsername());
        UserPo userPo = this.baseMapper.selectOne(lqw);
        if (userPo == null) {
            throw UserErrorCodeEnum.LoginError.newException();
        }

        if (Objects.equals(userPo.getPassword(), loginBo.getPassword())) {
            StpUtil.login(userPo.getId());
            return LoginResponseVo.builder()
                    .token(StpUtil.getTokenValue())
                    .build();
        }
        throw UserErrorCodeEnum.LoginError.newException();
    }

    @Override
    public void logout() {
        StpUtil.logout();
    }

    @Override
    public UserVo logUp(LogUpBo logUpBo) {
        UserPo userPo = UserPo.builder()
                .avatar(logUpBo.getAvatar().getUrl())
                .username(logUpBo.getUsername())
                .password(logUpBo.getPassword())
                .build();
        this.baseMapper.insert(userPo);

        return UserVo.builder()
                .id(userPo.getId())
                .avatar(logUpBo.getAvatar().getUrl())
                .username(logUpBo.getUsername())
                .createTime(userPo.getCreateTime())
                .updateTime(userPo.getUpdateTime())
                .build();
    }
}
