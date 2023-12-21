package com.jaronnie.fronted_backend_admin.backend_java.system.service;

import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.AddUserBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.LoginBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo.RegisterUserBo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.LoginVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.PublicKeyVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.UserVo;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestParam;

public interface IUserService {
    TableDataInfo<UserVo> queryPageList(PageQuery pageQuery);
    UserVo info();

    LoginVo login(LoginBo loginBo);

    void logout();

    boolean delete(Integer id);

    UserVo add(AddUserBo addUserBo);

    UserVo register(RegisterUserBo registerUserBo);

    Boolean sendMail(String mail, String username);

    PublicKeyVo getPublicKey();
}
