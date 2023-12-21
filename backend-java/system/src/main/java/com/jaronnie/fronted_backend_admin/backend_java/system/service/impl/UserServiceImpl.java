package com.jaronnie.fronted_backend_admin.backend_java.system.service.impl;

import cn.dev33.satoken.stp.StpUtil;
import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.core.toolkit.Wrappers;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.jaronnie.fronted_backend_admin.backend_java.common.util.RandomCodeGen;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.AddUserBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.LoginBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.RegisterUserBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.po.UserPo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.LoginVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.PublicKeyVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.UserVo;
import com.jaronnie.fronted_backend_admin.backend_java.common.enumeration.errcode.UserErrorCodeEnum;
import com.jaronnie.fronted_backend_admin.backend_java.system.mapper.UserMapper;
import com.jaronnie.fronted_backend_admin.backend_java.system.service.IMailService;
import com.jaronnie.fronted_backend_admin.backend_java.system.service.IUserService;
import com.jaronnie.fronted_backend_admin.backend_java.common.util.Md5SaltUtil;
import com.jaronnie.fronted_backend_admin.backend_java.common.util.RsaCrypto;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

import static com.jaronnie.fronted_backend_admin.backend_java.common.enumeration.errcode.UserErrorCodeEnum.LogUpError;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements IUserService {
    private final UserMapper baseMapper;
    private final IMailService iMailService;

    @Value("${spring.mail.username}")
    private String From;

    @Value("${backend.encrypt.type}")
    private String Type;

    @Value("${backend.encrypt.private-key}")
    private String PrivateKey;

    @Value("${backend.encrypt.public-key}")
    private String PublicKey;

    private static final String Salt = "jaronnie";

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
    public LoginVo login(LoginBo loginBo) {
        LambdaQueryWrapper<UserPo> lqw = Wrappers.lambdaQuery();
        lqw.eq(UserPo::getUsername, loginBo.getUsername());
        UserPo userPo = this.baseMapper.selectOne(lqw);
        if (userPo == null) {
            throw UserErrorCodeEnum.LoginError.newException();
        }

        String password = RsaCrypto.decrypt(loginBo.getPassword(), PrivateKey);
        if (Md5SaltUtil.verify(password, Salt, userPo.getPassword())) {
            StpUtil.login(userPo.getId());
            return LoginVo.builder()
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
    public boolean delete(Integer id) {
        StpUtil.logout(id);
        return this.baseMapper.deleteById(id) == 1;
    }

    @Override
    public UserVo add(AddUserBo addUserBo) {
        LambdaQueryWrapper<UserPo> lqw = new LambdaQueryWrapper<>();
        lqw.eq(UserPo::getUsername, addUserBo.getUsername());
        if (this.baseMapper.exists(lqw)) {
            throw LogUpError.newException();
        }

        String password = RsaCrypto.decrypt(addUserBo.getPassword(), PrivateKey);
        UserPo userPo = UserPo.builder()
                .avatar(addUserBo.getAvatar().getUrl())
                .username(addUserBo.getUsername())
                .password(Md5SaltUtil.generateHash(password, Salt))
                .build();
        this.baseMapper.insert(userPo);

        return UserVo.builder()
                .id(userPo.getId())
                .avatar(addUserBo.getAvatar().getUrl())
                .username(addUserBo.getUsername())
                .createTime(userPo.getCreateTime())
                .updateTime(userPo.getUpdateTime())
                .build();
    }

    @Override
    public UserVo register(RegisterUserBo registerUserBo) {
        return null;
    }

    @Override
    public Boolean sendMail(String mail, String username) {
        String verificationCode = RandomCodeGen.genRandomCode(6);
        String content = String.format("你好，%s：\n  本次邮箱验证码：%s", username, verificationCode);
        return iMailService.sendSimpleEmail(From, mail, "邮箱验证码", content);
    }

    @Override
    public PublicKeyVo getPublicKey() {
        return PublicKeyVo.builder()
                .type(Type)
                .publicKey(PublicKey)
                .build();
    }
}
