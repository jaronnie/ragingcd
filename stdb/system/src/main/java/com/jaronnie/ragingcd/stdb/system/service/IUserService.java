package com.jaronnie.ragingcd.stdb.system.service;

import com.jaronnie.ragingcd.stdb.system.domain.bo.AddUserBo;
import com.jaronnie.ragingcd.stdb.system.domain.bo.LoginBo;
import com.jaronnie.ragingcd.stdb.system.domain.query.PageQuery;
import com.jaronnie.ragingcd.stdb.system.domain.bo.RegisterUserBo;
import com.jaronnie.ragingcd.stdb.system.domain.query.SearchUserQuery;
import com.jaronnie.ragingcd.stdb.system.domain.vo.LoginVo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.PublicKeyVo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.TableDataInfo;
import com.jaronnie.ragingcd.stdb.system.domain.vo.UserVo;

public interface IUserService {
    TableDataInfo<UserVo> queryPageList(PageQuery pageQuery, SearchUserQuery searchUserQuery);
    UserVo info();

    LoginVo login(LoginBo loginBo);

    void logout();

    boolean delete(Integer id);

    UserVo add(AddUserBo addUserBo);

    UserVo register(RegisterUserBo registerUserBo);

    Boolean sendEmail(String email, String username);

    PublicKeyVo getPublicKey();
}
