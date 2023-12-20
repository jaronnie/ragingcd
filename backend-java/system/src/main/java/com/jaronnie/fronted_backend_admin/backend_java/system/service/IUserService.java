package com.jaronnie.fronted_backend_admin.backend_java.system.service;

import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.LogUpBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.LoginBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.LoginResponseVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.PublicKeyVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.UserVo;

public interface IUserService {
    TableDataInfo<UserVo> queryPageList(PageQuery pageQuery);
    UserVo info();

    LoginResponseVo login(LoginBo loginBo);

    void logout();

    boolean delete(Integer id);

    UserVo logUp(LogUpBo logUpBo);

    PublicKeyVo getPublicKey();
}
