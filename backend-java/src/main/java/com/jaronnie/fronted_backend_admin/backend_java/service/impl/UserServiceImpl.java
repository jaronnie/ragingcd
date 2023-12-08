package com.jaronnie.fronted_backend_admin.backend_java.service.impl;

import cn.dev33.satoken.stp.StpUtil;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.core.toolkit.Wrappers;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LoginQuery;
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

        Page<UserPo> page = baseMapper.selectPage(pageQuery.build(), lqw);

        TableDataInfo<UserVo> userVoTableDataInfo = new TableDataInfo<>();
        userVoTableDataInfo.setTotal(page.getTotal());
        List<UserVo> result = page.getRecords().stream().map(userPo -> UserVo.builder()
                .id(userPo.getId())
                .username(userPo.getUsername())
                .createTime(userPo.getCreateTime())
                .updateTime(userPo.getUpdateTime())
                .build()).collect(Collectors.toList());
        userVoTableDataInfo.setRows(result);
        return userVoTableDataInfo;
    }

    @Override
    public UserVo info(String token) {
        Object object = StpUtil.getLoginIdByToken(token);
        Integer id = Integer.parseInt(String.valueOf(object));
        UserPo userPo = this.baseMapper.selectById(id);
        return UserVo.builder()
                .username(userPo.getUsername())
                .id(userPo.getId())
                .createTime(userPo.getCreateTime())
                .updateTime(userPo.getUpdateTime())
                .build();
    }

    @Override
    public LoginResponseVo login(LoginQuery loginQuery) {
        LambdaQueryWrapper<UserPo> lqw = Wrappers.lambdaQuery();
        lqw.eq(UserPo::getUsername, loginQuery.getUsername());
        UserPo userPo = this.baseMapper.selectOne(lqw);
        if (Objects.equals(userPo.getPassword(), loginQuery.getPassword())) {
            StpUtil.login(userPo.getId());
            return LoginResponseVo.builder()
                    .token(StpUtil.getTokenValue())
                    .build();
        }
        throw UserErrorCodeEnum.LoginError.newException();
    }
}
