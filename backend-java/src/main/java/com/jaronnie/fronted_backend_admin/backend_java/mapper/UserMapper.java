package com.jaronnie.fronted_backend_admin.backend_java.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.jaronnie.fronted_backend_admin.backend_java.domain.po.UserPo;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface UserMapper extends BaseMapper<UserPo> {

}

