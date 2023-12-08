package com.jaronnie.fronted_backend_admin.backend_java.service;

import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.LoginQuery;
import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.LoginResponseVo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.UserVo;

public interface IUserService {
    TableDataInfo<UserVo> queryPageList(PageQuery pageQuery);
    UserVo info(String authorization);

    LoginResponseVo login(LoginQuery loginQuery);
}
